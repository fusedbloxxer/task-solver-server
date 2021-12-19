package dto

type TaskRequest struct {
	Context [][]string `json:"context" binding:"required"`
	Index   uint8      `json:"index" binding:"required,min=1,max=2"`
}

type TaskResponse struct {
	Context [][]string `json:"context" binding:"required"`
	Answer  float32    `json:"answer" binding:"required"`
	Index   uint8      `json:"index" binding:"required,min=1,max=2"`
}
