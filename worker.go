package main

import (
	"context"
	"log"
)

type Worker struct {
	id    int
	queue *Queue
}

func NewWorker(id int, queue *Queue) *Worker {
	return &Worker{
		id:    id,
		queue: queue,
	}
}

func (w *Worker) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("[Worker %d]: Shuttind down...", w.id)
				return

			case job := <-w.queue.jobs:
				log.Printf("[Worker %d]: Starting job %s", w.id, job.Name)
				err := job.Execute(job.Payload)
				if err != nil {
					log.Printf("[Worker %d]: Error running job %s: %v (retry %d / %d)",
						w.id, job.Name, err, job.retries, job.MaxRetries)
					if job.retries < job.MaxRetries {
						job.retries++
						go func() { w.queue.jobs <- job }()
						continue
					}
				} else {
					log.Printf("[Worker %d]: Finished job %s", w.id, job.Name)
				}
			}
		}
	}()
}
