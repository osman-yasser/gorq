package jobqueue

type Queue struct {
	jobs chan Job
}

func NewQueue(bufferSize int) *Queue {
	return &Queue{
		jobs: make(chan Job, bufferSize),
	}
}

func (q *Queue) Push(job Job) {
	q.jobs <- job
}
