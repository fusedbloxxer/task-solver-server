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
	//data := [][]string{
	//	{"aabbb", "ebep", "blablablaa", "hijk", "wsww"},
	//	{"abba", "eeeppp", "cocor", "ppppppaa", "qwerty", "acasq"},
	//	{"lalala", "lalal", "papapa", "papap"},
	//}
	//
	//input := make(chan interface{}, 3)
	//
	//go func() {
	//	for _, row := range data {
	//		for _, entry := range row {
	//			input <- entry
	//		}
	//	}
	//
	//	close(input)
	//}()
	//
	//pipe := functional.CreatePipeChannel(
	//	input,
	//	functional.CreateFilterer(func(x interface{}) bool {
	//		var entry string
	//
	//		_ = mapstructure.Decode(x, &entry)
	//
	//		vowels, consonants, _ := CountCharTypes(entry)
	//
	//		if vowels%2 == 0 && consonants%3 == 0 {
	//			fmt.Println(entry, vowels, consonants)
	//		}
	//
	//		return vowels%2 == 0 && consonants%3 == 0
	//	}),
	//	functional.CreateMapper(func(x interface{}) interface{} {
	//		return 1.0
	//	}),
	//	functional.CreateReducer(func(prev interface{}, curr interface{}) interface{} {
	//		var x, y float64
	//
	//		_ = mapstructure.Decode(prev, &x)
	//		_ = mapstructure.Decode(curr, &y)
	//
	//		return x + y/float64(len(data))
	//	}, 0.0),
	//)

	//work := sync.WaitGroup{}
	//
	//work.Add(3)
	//
	//go func() {
	//	output, _, e := pipe.Execute()
	//	fmt.Printf("p1: %v - %v\n", output, e)
	//	pipe.Wait()
	//	if output != nil {
	//		for o := range output {
	//			fmt.Println(o)
	//		}
	//	}
	//	work.Done()
	//}()
	//
	//go func() {
	//	output, _, e := pipe.Execute()
	//	fmt.Printf("p2: %v - %v\n", output, e)
	//	pipe.Wait()
	//	if output != nil {
	//		for o := range output {
	//			fmt.Println(o)
	//		}
	//	}
	//	work.Done()
	//}()
	//
	//go func() {
	//	output, _, e := pipe.Execute()
	//	fmt.Printf("p3: %v - %v\n", output, e)
	//	pipe.Wait()
	//	if output != nil {
	//		for o := range output {
	//			fmt.Println(o)
	//		}
	//	}
	//	work.Done()
	//}()
	//
	//work.Wait()

	s := server.GetInstance()

	if err := s.Initialize(); err != nil {
		log.Fatalln(err)
	}

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
