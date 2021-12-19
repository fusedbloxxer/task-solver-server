package repository

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"task-solver/server/src/model"
	"task-solver/server/src/services/firebase"
)

type ITaskRepository interface {
	GetTasks() ([]model.TaskResult, error)
	SaveTask(task model.TaskResult) error
}

type FirebaseTaskRepository struct {
	App *firebase.Firebase
}

func Create(settings map[string]interface{}) (app *FirebaseTaskRepository, err error) {
	app = &FirebaseTaskRepository{}

	if app.App, err = firebase.Create(settings); err != nil {
		return nil, fmt.Errorf("could not create firebase repository: %w", err)
	}

	return
}

func (ftr *FirebaseTaskRepository) GetTasks() ([]model.TaskResult, error) {
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

		task := model.TaskResult{
			Index:  uint8(doc.Data()["Index"].(int64)),
			Answer: float32(doc.Data()["Answer"].(float64)),
		}

		err = json.Unmarshal([]byte(doc.Data()["Context"].(string)), &task.Context)

		if err != nil {
			return nil, fmt.Errorf("invalid firebase task model: %w", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (ftr *FirebaseTaskRepository) SaveTask(task model.TaskResult) error {
	bytes, err := json.Marshal(task.Context)

	if err != nil {
		return fmt.Errorf("could not serialize the task: %w", err)
	}

	_, _, err = ftr.App.Firestore.Collection("tasks").Add(*ftr.App.Context, map[string]interface{}{
		"Context": string(bytes),
		"Index":   task.Index,
		"Answer":  task.Answer,
	})

	if err != nil {
		return fmt.Errorf("could not save the task to firestore: %w", err)
	}

	return nil
}
