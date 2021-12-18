package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-solver/server/src/settings"
)

type Router struct {
	Engine *gin.Engine
}

func Create(s *settings.AppSettings) (router *Router, err error) {
	// Create an empty router
	router = &Router{}

	// Get the current environment mode
	var mode string
	if mode, err = getAppMode(s.Environment); err != nil {
		return nil, err
	}

	// Set the Gin Engine to the current environment
	gin.SetMode(mode)

	// Create the engine
	router.Engine = gin.Default()

	// Add the routes
	if err = router.addRoutes(); err != nil {
		return nil, fmt.Errorf("could not add routes: %w", err)
	}

	// Return the new router
	return
}

func (router *Router) addRoutes() error {
	api := router.Engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/hello-world", func(c *gin.Context) {
				c.String(http.StatusOK, "Hello, world!")
			})
		}
	}

	return nil
}

func getAppMode(env string) (mode string, err error) {
	switch env {
	case "dev":
		return gin.DebugMode, nil
	case "pro":
		return gin.ReleaseMode, nil
	default:
		return "", fmt.Errorf("invalid environment type: %s", env)
	}
}
