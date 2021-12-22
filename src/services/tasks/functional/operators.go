package functional

type IOperator interface {
	execute(input <-chan interface{}, output chan<- interface{})
}
