package model

type Task struct {
	Context [][]string
	Index   uint8
}

type TaskResult struct {
	Context [][]string
	Index   uint8
	Answer  float32
}
