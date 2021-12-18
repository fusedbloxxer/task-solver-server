package server

import (
	"fmt"
	"sync"
	"task-solver/server/src/router"
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
	if s.Settings, err = settings.LoadConfig(); err != nil {
		panic(fmt.Errorf("could not init server: %w", err))
	}

	if s.Router, err = router.Create(s.Settings); err != nil {
		panic(fmt.Errorf("coult not create router: %w", err))
	}

	fmt.Println("initialized server instance")
	return
}

// Start the server
// Listen to incoming requests and respond to them
func (s *Server) Start() (err error) {
	if err = s.Router.Engine.Run(s.Settings.HostSettings.GetHost()); err != nil {
		return fmt.Errorf("could not start the server: %w", err)
	}

	return
}
