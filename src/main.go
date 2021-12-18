package main

import (
	"log"
	"task-solver/server/src/server"
)

func main() {
	s := server.GetInstance()

	if err := s.Initialize(); err != nil {
		log.Fatalln(err)
	}

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
