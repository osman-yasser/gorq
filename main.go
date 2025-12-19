package main

import (
	"fmt"
	"gorq/jobqueue"
	"log"
	"math/rand"
	"time"
)

func SendEmailJob(payload string) error {
	log.Printf("Sending mail to %s", payload)
	time.Sleep(1 * time.Second)
	return nil
}

func ResizeImageJob(payload string) error {
	log.Printf("Resizing image %s", payload)
	time.Sleep(2 * time.Second)
	return nil
}

func FailingJob(payload string) error {
	if rand.Intn(2) == 0 {
		return fmt.Errorf("random error")
	}
	log.Printf("Job Succeeded: %s", payload)
	return nil
}

func main() {
	queue := jobqueue.NewQueue(10, 3)

	queue.Start()

	defer queue.Shutdown()

	log.Println("workers started")

	queue.Enqueue(
		jobqueue.Job{
			Name:       "send_mail",
			Payload:    "osmanyasseradel77@gmail.com",
			Execute:    SendEmailJob,
			MaxRetries: 0,
		})

	queue.Enqueue(jobqueue.Job{
		Name:       "resize_image",
		Payload:    "image1.png",
		Execute:    ResizeImageJob,
		MaxRetries: 3,
	})

	queue.Enqueue(jobqueue.Job{
		Name:       "send_main",
		Payload:    "test@example.com",
		Execute:    SendEmailJob,
		MaxRetries: 0,
	})

	queue.Enqueue(jobqueue.Job{
		Name:       "test_failed",
		Payload:    "test failed",
		Execute:    FailingJob,
		MaxRetries: 3,
	})

	time.Sleep(5 * time.Second)
}
