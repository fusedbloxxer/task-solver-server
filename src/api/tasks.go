package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-solver/server/src/dal"
	"task-solver/server/src/model"
	"task-solver/server/src/model/dto"
)

type ITasksAPI interface {
	SolveTask(ctx *gin.Context)
	SaveTask(ctx *gin.Context)
	GetTasks(ctx *gin.Context)
}

type TasksAPI struct {
	UnitOfWork *dal.UnitOfWork
}

func (tasksApi *TasksAPI) SolveTask(c *gin.Context) {
	var task dto.TaskRequest

	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("invalid model: %w", err).Error(),
		})
		return
	}

	// Solve the problem and return the answer
	answer := tasksApi.UnitOfWork.TaskSolver.Solve(task.Context, task.Index)
	res := dto.TaskResponse{
		Context: task.Context,
		Index:   task.Index,
		Answer:  answer,
	}

	c.IndentedJSON(http.StatusOK, res)
}

func (tasksApi *TasksAPI) SaveTask(c *gin.Context) {
	var task dto.TaskResponse

	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("invalid model: %w", err).Error(),
		})
		return
	}

	err := tasksApi.UnitOfWork.TaskRepository.SaveTask(model.TaskResult{
		Context: task.Context,
		Index:   task.Index,
		Answer:  task.Answer,
	})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("could not save the task.go: %w", err).Error(),
		})
		return
	}

	tasks, err := tasksApi.UnitOfWork.TaskRepository.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("could retrieve the tasks: %w", err).Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func (tasksApi *TasksAPI) GetTasks(c *gin.Context) {
	tasks, err := tasksApi.UnitOfWork.TaskRepository.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("could retrieve the tasks: %w", err).Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}
