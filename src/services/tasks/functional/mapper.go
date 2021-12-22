package functional

import "task-solver/server/src/utils"

type MapOperator struct {
	mapper func(interface{}, int) interface{}
}

func CreateMapper(mapper func(interface{}) interface{}) IOperator {
	return createMapperWithIndices(func(o interface{}, i int) interface{} {
		return mapper(o)
	})
}

func CreateMapperWithIndices(mapper func(interface{}, int) interface{}) IOperator {
	return createMapperWithIndices(mapper)
}

func createMapperWithIndices(mapper func(interface{}, int) interface{}) *MapOperator {
	mapOperator := &MapOperator{
		mapper: mapper,
	}
	return mapOperator
}

func (mapOperator *MapOperator) execute(input <-chan interface{}, output chan<- interface{}) {
	index := 0

	for entry := range input {
		output <- mapOperator.mapper(utils.Clone(entry), index)
		index++
	}

	close(output)
}
