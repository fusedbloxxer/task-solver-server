package functional

import "task-solver/server/src/utils"

type ReduceOperator struct {
	reducer     func(interface{}, interface{}, int) interface{}
	accumulator interface{}
	identity    interface{}
}

func CreateReducer(reducer func(interface{}, interface{}) interface{}, identity interface{}) IOperator {
	return createReducerWithIndices(func(prev interface{}, curr interface{}, i int) interface{} {
		return reducer(prev, curr)
	}, identity)
}

func CreateReducerWithIndices(reducer func(interface{}, interface{}, int) interface{}, identity interface{}) IOperator {
	return createReducerWithIndices(reducer, identity)
}

func createReducerWithIndices(reducer func(interface{}, interface{}, int) interface{}, identity interface{}) *ReduceOperator {
	reduceOperator := &ReduceOperator{
		identity: utils.Clone(identity),
		reducer:  reducer,
	}
	return reduceOperator
}

func (reduceOperator *ReduceOperator) execute(input <-chan interface{}, output chan<- interface{}) {
	index := 0

	for entry := range input {
		if index == 0 {
			reduceOperator.accumulator = reduceOperator.reducer(
				reduceOperator.identity,
				entry,
				index,
			)
		} else {
			reduceOperator.accumulator = reduceOperator.reducer(
				reduceOperator.accumulator,
				entry,
				index,
			)
		}
		index++
	}

	output <- reduceOperator.accumulator
	close(output)
}
