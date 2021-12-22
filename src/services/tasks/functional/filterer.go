package functional

import "task-solver/server/src/utils"

type FilterOperator struct {
	filter func(interface{}, int) bool
}

func CreateFilterer(filter func(interface{}) bool) IOperator {
	return createFiltererWithIndices(func(o interface{}, i int) bool {
		return filter(o)
	})
}

func CreateFiltererWithIndices(filter func(interface{}, int) bool) IOperator {
	return createFiltererWithIndices(filter)
}

func createFiltererWithIndices(filter func(interface{}, int) bool) *FilterOperator {
	filterOperator := &FilterOperator{
		filter: filter,
	}
	return filterOperator
}

func (filterOperator *FilterOperator) execute(input <-chan interface{}, output chan<- interface{}) {
	index := 0

	for entry := range input {
		if filterOperator.filter(entry, index) {
			output <- utils.Clone(entry)
		}
		index++
	}

	close(output)
}
