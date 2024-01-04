package _go

type Worker[args any, res any] struct {
	taskPipeline chan *Task[args, res]
}

func NewWorker[args any, res any](pipeline chan *Task[args, res]) *Worker[args, res] {
	w := &Worker[args, res]{
		taskPipeline: pipeline,
	}
	go w.run()
	return w
}

func (w *Worker[args, res]) run() {
	for task := range w.taskPipeline {
		if task.stop { // 停止信号，直接worker执行结束
			return
		}
		task.result <- task.handler(task.args)
	}
}
