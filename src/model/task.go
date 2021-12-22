package model

type Task struct {
	Context [][]string `json:"context" binding:"required"`
	Index   int64      `json:"index" binding:"required"`
}

type TaskResult struct {
	Id     string  `json:"id" binding:"required"`
	Task   Task    `json:"task" binding:"required"`
	Answer float64 `json:"answer" binding:"required"`
}

type BadRequestError struct {
	Error string
}
