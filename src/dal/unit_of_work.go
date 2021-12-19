package dal

import (
	"task-solver/server/src/dal/repository"
	"task-solver/server/src/services/tasks"
)

type UnitOfWork struct {
	TaskRepository repository.ITaskRepository
	TaskSolver     tasks.ITaskSolver
}

func Create(taskRepository repository.ITaskRepository, taskSolver tasks.ITaskSolver) (uow *UnitOfWork) {
	uow = &UnitOfWork{}
	uow.TaskSolver = taskSolver
	uow.TaskRepository = taskRepository
	return uow
}
