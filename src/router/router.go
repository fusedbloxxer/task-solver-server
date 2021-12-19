package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"task-solver/server/src/settings"
)

type Router struct {
	Engine *gin.Engine
}

func Create(config *settings.AppSettings) (router *Router, err error) {
	// Create an empty router
	router = &Router{}

	// Get the current environment mode
	var mode string
	if mode, err = getAppMode(config.Environment); err != nil {
		return nil, fmt.Errorf("could not initialize the router: %w", err)
	}

	// Set the Gin Engine to the current environment
	gin.SetMode(mode)

	// Create the engine
	router.Engine = gin.Default()
	return
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
