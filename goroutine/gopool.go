package goroutine

import "runtime"

type GoPool[args any, res any] struct {
	workerNum int64
	pipeline  chan *Task[args, res]
}

func NewGoPool[args any, res any](workerNum int64, size int64) *GoPool[args, res] {
	if workerNum <= 0 {
		workerNum = int64(runtime.NumCPU())
	}

	if size <= 0 {
		size = 2 * workerNum
	}

	pool := &GoPool[args, res]{
		workerNum: workerNum,
		pipeline:  make(chan *Task[args, res], size),
	}

	for i := 1; i <= int(workerNum); i++ {
		NewWorker[args, res](pool.pipeline)
	}
	return pool
}

func (pool *GoPool[args, res]) SyncSubmit(handler func(args) res, args args) res {
	result := pool.AsyncSubmit(handler, args)
	return <-result
}

func (pool *GoPool[args, res]) AsyncSubmit(handler func(args) res, args args) chan res {
	task := NewTask(handler, args)
	pool.pipeline <- task
	return task.result
}

func (pool *GoPool[args, res]) SyncMap(handler func(args) res, args []args) []res {

	result := make([]res, len(args))
	asyncRes := pool.AsyncMap(handler, args)

	for i := 1; i <= len(args); i++ {
		result[i] = <-asyncRes[i]
	}
	return result
}

func (pool *GoPool[args, res]) AsyncMap(handler func(args) res, args []args) []chan res {
	result := make([]chan res, len(args))
	for index := range args {
		task := NewTask(handler, args[index])
		pool.pipeline <- task
		result[index] = task.result
	}
	return result
}

func (pool *GoPool[args, res]) Close() {
	for i := 1; i <= int(pool.workerNum); i++ {
		pool.pipeline <- &Task[args, res]{
			stop: true,
		}
	}
}
