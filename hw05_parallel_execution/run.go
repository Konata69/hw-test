package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type ErrorLimiter struct {
	count int32
	limit int32
}

func (l *ErrorLimiter) Inc() {
	atomic.AddInt32(&l.count, 1)
}

func (l *ErrorLimiter) IsExceeded() bool {
	return l.limit > 0 && atomic.LoadInt32(&l.count) >= l.limit
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 || m <= 0 {
		return ErrErrorsLimitExceeded
	}

	limiter := &ErrorLimiter{
		count: 0,
		limit: int32(m),
	}

	var wg sync.WaitGroup
	taskChannel := make(chan Task)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for task := range taskChannel {
				err := task()
				if err != nil {
					limiter.Inc()
				}
			}
		}()
	}

	for _, task := range tasks {
		taskChannel <- task
		if limiter.IsExceeded() {
			break
		}
	}

	close(taskChannel)
	wg.Wait()

	if limiter.IsExceeded() {
		return ErrErrorsLimitExceeded
	}

	return nil
}
