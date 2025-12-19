package jobqueue

type JobFunc func(payload string) error

type Job struct {
	Name       string
	Payload    string
	Execute    JobFunc
	retries    int
	MaxRetries int
}
