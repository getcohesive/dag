package dag

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type AsyncRunner struct {
	wg map[string]*sync.WaitGroup
}

func (a *AsyncRunner) New() AsyncRunner {
	return AsyncRunner{
		wg: make(map[string]*sync.WaitGroup),
	}
}

func (a *AsyncRunner) Run(ctx context.Context, job *Job) error {
	fmt.Println("AsyncRunner.Run: starting", job.Id)
	a.wg[job.Id] = &sync.WaitGroup{}
	defer delete(a.wg, job.Id)

	a.wg[job.Id].Add(len(job.tasks))

	errs := make([]error, len(job.tasks), len(job.tasks))

	for i, task := range job.tasks {
		go func(i int, task TaskFunc) {
			err := task(ctx)
			errs[i] = err
			a.wg[job.Id].Done()
		}(i, task)
	}

	a.wg[job.Id].Wait()

	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	delete(a.wg, job.Id)
	if job.onComplete != nil {
		return job.onComplete(ctx)
	}

	return nil
}

func (a *AsyncRunner) Stop(_ context.Context, job *Job) error {
	fmt.Println("AsyncRunner.Stop: stopping", job.Id)
	if a.wg[job.Id] == nil {
		return errors.New("job wg not found")
	}
	for _, _ = range job.tasks {
		a.wg[job.Id].Done()
	}
	return nil
}
