package functional

import (
	"fmt"
	"sync"
)

const (
	PipeAvailable = iota
	PipeStarted
)

type Pipe struct {
	operators []IOperator
	status    int
	data      chan interface{}
	mutex     sync.Mutex
	wait      *sync.WaitGroup
	output    chan interface{}
}

// CreatePipeChannel
// The default way to create a pipe.
// Data is received and passed through the pipe until the channel is closed.
func CreatePipeChannel(data chan interface{}, operators ...IOperator) *Pipe {
	pipe := &Pipe{
		operators: operators,
		status:    PipeAvailable,
		data:      data,
		wait:      &sync.WaitGroup{},
		output:    make(chan interface{}, cap(data)),
	}
	return pipe
}

// CreatePipe
// A simpler way to create a pipe.
// All data is received from the start and is passed from another thread through the pipe to avoid starvation.
func CreatePipe(data []interface{}, channelSize int, operators ...IOperator) *Pipe {
	dataChannel := make(chan interface{}, channelSize)

	pipe := CreatePipeChannel(dataChannel, operators...)

	go func() {
		for _, entry := range data {
			dataChannel <- entry
		}
		close(dataChannel)
	}()

	return pipe
}

func (pipe *Pipe) Pipe(operators ...IOperator) *Pipe {
	// Lock this function because it has potential side effects
	pipe.mutex.Lock()
	defer pipe.mutex.Unlock()

	// Check the status of the pipe before proceeding to modify its internals
	// In case the pipeline was already started, return nil to make the error more obvious
	if pipe.status != PipeAvailable {
		return nil
	}

	// Modify the pipeline and allow access
	pipe.operators = append(pipe.operators, operators...)
	return pipe
}

func (pipe *Pipe) Execute() (<-chan interface{}, *sync.WaitGroup, error) {
	// A pipe may be executed only once
	pipe.mutex.Lock()
	defer pipe.mutex.Unlock()

	// Check the current status
	if pipe.status != PipeAvailable {
		return nil, pipe.wait, fmt.Errorf("pipeline cannot be executed again")
	}

	// Mark the pipe as being already started
	pipe.status = PipeStarted

	// Allow others to wait for full execution of the pipe
	pipe.wait.Add(len(pipe.operators) + 1)

	// Start consuming data and processing it
	go pipe.runTasks()

	// Return the waiting group
	return pipe.output, pipe.wait, nil
}

func (pipe *Pipe) Wait() {
	pipe.wait.Wait()
}

func (pipe *Pipe) runTasks() {
	defer pipe.wait.Done()

	// No work to be done
	if len(pipe.operators) == 0 {
		return
	}

	//
	var output chan interface{}

	// Apply each operator in parallel
	for i, operator := range pipe.operators {
		// Declare input & output temp variables
		// To be used in closure
		var input chan interface{}

		// The input is the output from the previous operator
		if i >= 1 {
			input = output
		} else {
			input = pipe.data
		}

		// The output is either the final output channel
		// Or a new intermediary channel
		if i == len(pipe.operators)-1 {
			output = pipe.output
		} else {
			output = make(chan interface{}, cap(pipe.data))
		}

		// Hide outer variable to be used
		// Safely in closure
		output := output

		// Avoid loop var usage problem in closures, by creating
		// a temporary variable
		operator := operator

		// Run the operators in parallel
		go func() {
			// Apply the operator asynchronously using channels
			operator.execute(input, output)
			pipe.wait.Done()
		}()
	}
}
