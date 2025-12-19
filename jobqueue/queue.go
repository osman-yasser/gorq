package jobqueue

import "context"

type Queue struct {
	jobs    chan Job
	workers int
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewQueue(bufferSize int, workers int) *Queue {
	return &Queue{
		jobs:    make(chan Job, bufferSize),
		workers: workers,
	}
}

func (q *Queue) Enqueue(job Job) {
	q.jobs <- job
}

func (q *Queue) Start() {
	q.ctx, q.cancel = context.WithCancel(context.Background())

	for i := 0; i < q.workers; i++ {
		w := NewWorker(i, q)
		w.Start(q.ctx)
	}
}

func (q *Queue) Shutdown() {
	if q.cancel != nil {
		q.cancel()
	}
}
