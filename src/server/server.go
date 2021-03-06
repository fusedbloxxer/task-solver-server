package server

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sync"
	"task-solver/server/docs"
	endpoints "task-solver/server/src/api"
	"task-solver/server/src/dal"
	"task-solver/server/src/dal/repository"
	"task-solver/server/src/router"
	"task-solver/server/src/services/tasks"
	"task-solver/server/src/settings"
)

// Use Mutex to prevent race conditions
var lock = &sync.Mutex{}

// The Server which holds the functionality of the app
type Server struct {
	Settings *settings.AppSettings
	Router   *router.Router
}

// The singleton instance
var instance *Server

// GetInstance
// Create only an instance of the Server
// Using Singleton Pattern with Double-Locking
func GetInstance() *Server {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Server{}
			fmt.Println("created server instance")
		}
	}

	return instance
}

// Initialize all components for the server
func (s *Server) Initialize() (err error) {
	// Load the application settings for the current environment
	if s.Settings, err = settings.LoadConfig(); err != nil {
		panic(fmt.Errorf("could not init server: %w", err))
	}

	// Create a router with no routes added
	if s.Router, err = router.Create(s.Settings); err != nil {
		panic(fmt.Errorf("coult not create router: %w", err))
	}

	// Add the routes
	if err = s.addRoutes(); err != nil {
		panic(fmt.Errorf("could not add routes: %w", err))
	}

	fmt.Println("initialized server instance")
	return
}

func (s *Server) addRoutes() error {

	api := s.Router.Engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// Add basic test endpoint
			testHandler := endpoints.TestAPI{}
			v1.GET("/test", testHandler.Test)

			// Add configuration endpoint
			configHandler := endpoints.ConfigAPI{}
			v1.GET("/config", configHandler.GetConfig)

			// Create services and inject dependencies
			taskSolver, err := tasks.CreateMapReduceSolver(s.Settings.Services)
			if err != nil {
				return fmt.Errorf("could not create taskSolver service: %w", err)
			}

			taskRepository, err := repository.Create(s.Settings.Services)
			if err != nil {
				return fmt.Errorf("could not create repository: %w", err)
			}

			tasksHandler := endpoints.TasksAPI{
				UnitOfWork: dal.Create(taskRepository, taskSolver),
			}

			taskGroup := v1.Group("/tasks")
			{
				taskGroup.POST("/solve", tasksHandler.SolveTask)

				taskGroup.GET("/", tasksHandler.GetAllTaskResults)
				taskGroup.GET("/indexes", tasksHandler.GetAllTaskIndexes)
				taskGroup.GET("/:taskId", tasksHandler.GetTaskResultById)

				taskGroup.DELETE("/", tasksHandler.DeleteAllTaskResults)
				taskGroup.DELETE("/:taskId", tasksHandler.DeleteTaskResult)
			}
		}
	}

	// Add Swagger / OpenAPI
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.Router.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return nil
}

// Start the server
// Listen to incoming requests and respond to them
func (s *Server) Start() (err error) {
	if err = s.Router.Engine.Run(s.Settings.Server.Host.GetHost()); err != nil {
		return fmt.Errorf("could not start the server: %w", err)
	}

	return
}
