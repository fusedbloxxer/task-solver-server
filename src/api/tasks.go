package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-solver/server/src/dal"
	"task-solver/server/src/model"
)

type ITasksAPI interface {
	DeleteAllTaskResults(ctx *gin.Context)

	GetTaskResultById(ctx *gin.Context)

	GetAllTaskResults(ctx *gin.Context)

	GetAllTaskIndexes(ctx *gin.Context)

	DeleteTaskResult(ctx *gin.Context)

	SolveTask(ctx *gin.Context)
}

type TasksAPI struct {
	UnitOfWork *dal.UnitOfWork
}

// DeleteAllTaskResults
// @Summary Delete all stored task results
// @Description Deletes the task results from the server.
// @Tags task
// @Accept json
// @Produce json
// @Success 200 "All tasks are deleted"
// @Failure 400 {object} model.BadRequestError The tasks could not be deleted
// @Router /tasks [delete]
func (tasksApi *TasksAPI) DeleteAllTaskResults(c *gin.Context) {
	err := tasksApi.UnitOfWork.TaskRepository.DeleteAllTaskResults()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.BadRequestError{
			Error: fmt.Errorf("could not delete tasks: %w", err).Error(),
		})
		return
	}
}

// GetTaskResultById
// @Summary Get a saved task result by its document id
// @Description Fetch the saved result from the server.
// @Tags task
// @Accept json
// @Produce json
// @Param taskId query string true "Used to identify the task"
// @Success 200 {object} model.TaskResult "The task result is returned"
// @Failure 400 {object} model.BadRequestError "The task does not exist"
// @Router /tasks/:taskId [get]
func (tasksApi *TasksAPI) GetTaskResultById(c *gin.Context) {
	taskId := c.Param("taskId")

	var err error
	var taskResult model.TaskResult
	if taskResult, err = tasksApi.UnitOfWork.TaskRepository.GetTaskResultById(taskId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("could not retrieve task: %w", err).Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, taskResult)
}

// GetAllTaskResults
// @Summary Retrieve all stored task results
// @Description Fetch all the stored task results from the server. They are unordered and unfiltered.
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {object} []model.TaskResult "All task results are returned as an array"
// @Failure 500 {object} model.BadRequestError "Failed to fetch the tasks"
// @Router /tasks [get]
func (tasksApi *TasksAPI) GetAllTaskResults(c *gin.Context) {
	tasks, err := tasksApi.UnitOfWork.TaskRepository.GetAllTaskResults()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("could retrieve the tasks: %w", err).Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetAllTaskIndexes
// @Summary Retrieve the possible problem types or indexes
// @Description Fetch from the server the possible problem types implemented.
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {object} []int64 "The problem indexes are returned as an array. It is unordered."
// @Failure 500 {object} model.BadRequestError "Failed to fetch the problem indexes."
// @Router /tasks/indexes [get]
func (tasksApi *TasksAPI) GetAllTaskIndexes(c *gin.Context) {
	indexes, err := tasksApi.UnitOfWork.TaskRepository.GetAllTaskIndexes()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("could retrieve the indexes: %w", err).Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, indexes)
}

// DeleteTaskResult
// @Summary Delete a task result using its id
// @Description Delete a saved task result using the id.
// @Tags task
// @Accept json
// @Produce json
// @Param taskId query string true "Used to identify the task"
// @Success 200 "The task is removed"
// @Failure 400 {object} model.BadRequestError "The taskId does not exist"
// @Router /tasks/:taskId [delete]
func (tasksApi *TasksAPI) DeleteTaskResult(c *gin.Context) {
	taskId := c.Param("taskId")
	err := tasksApi.UnitOfWork.TaskRepository.DeleteTaskResult(taskId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("could not delete task: %w", err).Error(),
		})
		return
	}
}

// Using CURL
/**
curl --request POST \
	 -H "Content-Type: application/json" \
	 --data "{\"context\":[[\"string\"]],\"index\":2}" \
	 http://127.0.0.1:8080/api/v1/tasks/solve
*/

// SolveTask
// @Summary Solve a task and save the result
// @Description Solve a task by using the context and the index of the problem. Save the results.
// @Tags task
// @Accept json
// @Produce json
// @Param Task body model.Task true "The task to be solved. Its index must be obtained from /tasks/indexes."
// @Success 200 {object} model.TaskResult "The task result containing an id for the saved value and the answer"
// @Failure 400 {object} model.BadRequestError "The task model is invalid"
// @Router /tasks/solve [post]
func (tasksApi *TasksAPI) SolveTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("invalid model: %w", err).Error(),
		})
		return
	}

	// TODO: uncomment
	// Fetch current available problems
	indexes, err := tasksApi.UnitOfWork.TaskRepository.GetAllTaskIndexes()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("could not retrieve all task indexes: %w", err).Error(),
		})
		return
	}

	// Validate task index
	exists := false
	for _, index := range indexes {
		if index == task.Index {
			exists = true
			break
		}
	}

	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid index: %v", task.Index),
		})
		return
	}

	// Solve the problem and return the answer
	answer, err := tasksApi.UnitOfWork.TaskSolver.Solve(task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("could not solve the problem: %w", err).Error(),
		})
		return
	}

	// Create task result object
	taskResult := model.TaskResult{
		Id:     "",
		Task:   task,
		Answer: answer,
	}

	// TODO: uncomment
	// Save the solved task
	err = tasksApi.UnitOfWork.TaskRepository.SaveTaskResult(&taskResult)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("could not save the solved task: %w", err).Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, taskResult)
}
