package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"task-solver/server/src/model"
	"task-solver/server/src/services/firebase"
)

type ITaskRepository interface {
	GetTaskResultById(id string) (model.TaskResult, error)

	GetAllTaskResults() ([]model.TaskResult, error)

	SaveTaskResult(task *model.TaskResult) error

	DeleteTaskResult(id string) error

	DeleteAllTaskResults() error
}

type FirebaseTaskRepository struct {
	App *firebase.Firebase
}

type TaskFirebase struct {
	Context string
	Index   int64
}

type TaskResultFirebase struct {
	Answer float64
	Task   TaskFirebase
}

func Create(settings map[string]interface{}) (app *FirebaseTaskRepository, err error) {
	app = &FirebaseTaskRepository{}

	if app.App, err = firebase.Create(settings); err != nil {
		return nil, fmt.Errorf("could not create firebase repository: %w", err)
	}

	return
}

func (ftr *FirebaseTaskRepository) GetTaskResultById(id string) (model.TaskResult, error) {
	doc, err := ftr.App.Firestore.Collection("tasks").Doc(id).Get(*ftr.App.Context)
	if err != nil {
		return model.TaskResult{}, fmt.Errorf("could not retrieve task result with id %v: %w", id, err)
	}

	var taskResult model.TaskResult
	if taskResult, err = toTaskResult(doc); err != nil {
		return model.TaskResult{}, fmt.Errorf("could not retrieve task result with id %v: %w", id, err)
	}

	return taskResult, nil
}

func (ftr *FirebaseTaskRepository) GetAllTaskResults() ([]model.TaskResult, error) {
	tasks := make([]model.TaskResult, 0)

	iter := ftr.App.Firestore.Collection("tasks").Documents(*ftr.App.Context)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("could not retrieve all tasks: %w", err)
		}

		var taskResult model.TaskResult
		if taskResult, err = toTaskResult(doc); err != nil {
			return nil, fmt.Errorf("could convert from firebase model: %w", err)
		}

		tasks = append(tasks, taskResult)
	}

	return tasks, nil
}

func (ftr *FirebaseTaskRepository) SaveTaskResult(task *model.TaskResult) (err error) {
	var taskResultFirebase TaskResultFirebase
	if taskResultFirebase, err = fromTaskResult(*task); err != nil {
		return fmt.Errorf("could not convert from task result to firebase model: %w", err)
	}

	var doc *firestore.DocumentRef
	doc, _, err = ftr.App.Firestore.Collection("tasks").Add(*ftr.App.Context, taskResultFirebase)

	// Return the id of the newly created task
	task.Id = doc.ID

	if err != nil {
		return fmt.Errorf("could not save the task to firestore: %w", err)
	}

	return nil
}

func (ftr *FirebaseTaskRepository) DeleteTaskResult(id string) (err error) {
	_, err = ftr.App.Firestore.Collection("tasks").Doc(id).Delete(*ftr.App.Context)

	if err != nil {
		return fmt.Errorf("could not delete task with id %v: %w", id, err)
	}

	return
}

func (ftr *FirebaseTaskRepository) DeleteAllTaskResults() (err error) {
	err = deleteCollection(*ftr.App.Context, ftr.App.Firestore, ftr.App.Firestore.Collection("tasks"), 10)

	if err != nil {
		return fmt.Errorf("could not delete all tasks: %w", err)
	}

	return
}

func toTaskResult(doc *firestore.DocumentSnapshot) (taskResult model.TaskResult, err error) {
	taskResultFirebase := TaskResultFirebase{}
	if err = mapstructure.Decode(doc.Data(), &taskResultFirebase); err != nil {
		return model.TaskResult{}, fmt.Errorf("invalid firebase model conversion: %w", err)
	}

	taskContext := make([][]string, 0)
	err = json.Unmarshal([]byte(taskResultFirebase.Task.Context), &taskContext)
	if err != nil {
		return model.TaskResult{}, fmt.Errorf("invalid firebase task model: %w", err)
	}

	// TODO: replace with some kind of automapper
	taskResult = model.TaskResult{
		Id: doc.Ref.ID,
		Task: model.Task{
			Index:   taskResultFirebase.Task.Index,
			Context: taskContext,
		},
		Answer: taskResultFirebase.Answer,
	}

	return
}

func fromTaskResult(task model.TaskResult) (taskResultFirebase TaskResultFirebase, err error) {
	bytes, err := json.Marshal(task.Task.Context)

	if err != nil {
		return TaskResultFirebase{}, fmt.Errorf("could not serialize the task: %w", err)
	}

	// TODO: replace with some kind of automapper
	taskResultFirebase = TaskResultFirebase{
		Task: TaskFirebase{
			Index:   task.Task.Index,
			Context: string(bytes),
		},
		Answer: task.Answer,
	}

	return
}

func deleteCollection(ctx context.Context, client *firestore.Client, ref *firestore.CollectionRef, batchSize int) error {
	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}
