package _go

type Task[args any, res any] struct {
	handler func(args) res
	args    args
	result  chan res
	stop    bool
}

func NewTask[args any, res any](handler func(args) res, arg args) *Task[args, res] {
	return &Task[args, res]{
		handler: handler,
		args:    arg,
		result:  make(chan res),
	}
}
