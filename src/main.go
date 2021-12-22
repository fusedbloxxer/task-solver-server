package main

import (
	"log"
	"task-solver/server/src/server"
)

// TasksAPI
// @title Tasks API
// @version 1.0
// @description This API can be used to solve tasks and save the results to firebase

// @host 127.0.0.1:8080
// @BasePath /api/v1

// @contact.url https://github.com/fusedbloxxer

func main() {
	s := server.GetInstance()

	if err := s.Initialize(); err != nil {
		log.Fatalln(err)
	}

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
