package tasks

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"task-solver/server/src/model"
	"task-solver/server/src/services/tasks/functional"
)

type ITaskSolver interface {
	Solve(task model.Task) (float64, error)
	solveFirstTask(context [][]string) float64
	solveSecondTask(context [][]string) float64
}

const (
	SectionKey = "tasks"
)

type Settings struct {
	ChannelSize int `json:"channelSize"`
}

type MapReduceSolver struct {
	IndexMapper map[int64]func([][]string) float64
	Settings    *Settings
}

func CreateMapReduceSolver(settings map[string]interface{}) (ITaskSolver, error) {
	mrs := &MapReduceSolver{}

	var exists bool
	var config interface{}
	if config, exists = settings[SectionKey]; !exists {
		return nil, fmt.Errorf("the %s section does not exist", SectionKey)
	}

	var taskSolverSettings Settings
	if err := mapstructure.Decode(config, &taskSolverSettings); err != nil {
		return nil, fmt.Errorf("could not deserialize %s settings: %w", SectionKey, err)
	}

	mrs.Settings = &taskSolverSettings
	mrs.IndexMapper = map[int64]func([][]string) float64{
		1: mrs.solveFirstTask,
		2: mrs.solveSecondTask,
	}

	return mrs, nil
}

func (mrs *MapReduceSolver) Solve(task model.Task) (float64, error) {
	var ok bool
	var solver func([][]string) float64

	if solver, ok = mrs.IndexMapper[task.Index]; !ok {
		return 0.0, fmt.Errorf("invalid solution index")
	}

	return solver(task.Context), nil
}

//curl --request POST \
//-H "Content-Type: application/json" \
//--data "{\"context\":[[\"string\"]],\"index\":2}" \
//http://127.0.0.1:8080/api/v1/tasks/solve

func (mrs *MapReduceSolver) solveFirstTask(context [][]string) float64 {
	// Use a communication channel to process data
	input := make(chan interface{}, mrs.Settings.ChannelSize)

	// Send data through the channel on a separate thread, to avoid deadlocks,
	// which will be processed by a pipe
	go func() {
		for _, row := range context {
			for _, entry := range row {
				input <- entry
			}
		}

		close(input)
	}()

	// Create a custom pipe
	pipe := functional.CreatePipeChannel(
		input,
		functional.CreateFilterer(func(x interface{}) bool {
			var entry string

			_ = mapstructure.Decode(x, &entry)

			vowels, consonants, _ := CountCharTypes(entry)

			return vowels%2 == 0 && consonants%3 == 0 && len(entry) != 0
		}),
		functional.CreateMapper(func(x interface{}) interface{} {
			return 1.0
		}),
		functional.CreateReducer(func(prev interface{}, curr interface{}) interface{} {
			var x, y float64

			_ = mapstructure.Decode(prev, &x)
			_ = mapstructure.Decode(curr, &y)

			return x + y/float64(len(context))
		}, 0.0),
	)

	// Start processing data on a separate thread
	output, _, _ := pipe.Execute()

	// Wait for the pipe to finish processing the data
	pipe.Wait()

	// Fetch result from pipeline
	res := <-output

	// Check for no result
	if res == nil {
		return 0.0
	}

	// Return the data as float
	return res.(float64)
}

func (mrs *MapReduceSolver) solveSecondTask(context [][]string) float64 {
	return 2.0
}
