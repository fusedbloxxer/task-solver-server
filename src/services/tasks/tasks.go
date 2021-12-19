package tasks

type ITaskSolver interface {
	Solve(context [][]string, index uint8) float32
	solveFirstTask(context [][]string) float32
	solveSecondTask(context [][]string) float32
}

type MapReduceSolver struct {
}

func (mrs *MapReduceSolver) Solve(context [][]string, index uint8) float32 {
	return 0.0
}

func (mrs *MapReduceSolver) solveFirstTask(context [][]string) float32 {
	return 0.0
}

func (mrs *MapReduceSolver) solveSecondTask(context [][]string) float32 {
	return 0.0
}
