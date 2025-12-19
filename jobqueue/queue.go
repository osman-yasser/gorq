package jobqueue

import "context"

const (
	defaultBufferSize = 100
	defaultWorker     = 1
)

type Queue struct {
	jobs    chan Job
	workers int
	ctx     context.Context
	cancel  context.CancelFunc
}

type Option func(*Queue)

func NewQueue(opts ...Option) *Queue {
	q := &Queue{
		workers: defaultWorker,
	}

	for _, opt := range opts {
		opt(q)
	}

	if q.jobs == nil {
		q.jobs = make(chan Job, defaultBufferSize)
	}

	return q
}

func QueueBufferSize(size int) Option {
	return func(q *Queue) {
		q.jobs = make(chan Job, size)
	}
}

func QueueWorkers(num int) Option {
	return func(q *Queue) {
		q.workers = num
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
