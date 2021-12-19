package tasks

import "task-solver/server/src/model"

type ITaskSolver interface {
	Solve(task model.Task) float64
	solveFirstTask(context [][]string) float64
	solveSecondTask(context [][]string) float64
}

type MapReduceSolver struct {
}

func (mrs *MapReduceSolver) Solve(task model.Task) float64 {
	return 0.0
}

func (mrs *MapReduceSolver) solveFirstTask(context [][]string) float64 {
	return 0.0
}

func (mrs *MapReduceSolver) solveSecondTask(context [][]string) float64 {
	return 0.0
}
