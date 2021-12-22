package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ITestAPI interface {
	Test(*gin.Context)
}

type TestAPI struct {
}

// Test
// @Summary Test that the API is responding
// @Description Tests if the API is working. A "Hello, World!" message should always be returned.
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} string "The message "Hello, World!" is returned"
// @Router /test [get]
func (testApi *TestAPI) Test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello, world!")
}
