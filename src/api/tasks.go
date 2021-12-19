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

	DeleteTaskResult(ctx *gin.Context)

	SolveTask(ctx *gin.Context)
}

type TasksAPI struct {
	UnitOfWork *dal.UnitOfWork
}

func (tasksApi *TasksAPI) DeleteAllTaskResults(c *gin.Context) {
	err := tasksApi.UnitOfWork.TaskRepository.DeleteAllTaskResults()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("could not delete tasks: %w", err).Error(),
		})
		return
	}
}

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

func (tasksApi *TasksAPI) SolveTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("invalid model: %w", err).Error(),
		})
		return
	}

	// Solve the problem and return the answer
	answer := tasksApi.UnitOfWork.TaskSolver.Solve(task)

	// Create task result object
	taskResult := model.TaskResult{
		Id:     "",
		Task:   task,
		Answer: answer,
	}

	// Save the solved task
	err := tasksApi.UnitOfWork.TaskRepository.SaveTaskResult(&taskResult)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("could not save the solved task: %w", err).Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, taskResult)
}
