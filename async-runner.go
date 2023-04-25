package dag

import (
	"context"
	"fmt"
	"sync"
)

type AsyncRunner struct {
	wg          map[string]*sync.WaitGroup
	exitChannel map[string]chan struct{}
}

func (a *AsyncRunner) New() AsyncRunner {
	return AsyncRunner{
		wg:          make(map[string]*sync.WaitGroup),
		exitChannel: make(map[string]chan struct{}),
	}
}

func (a *AsyncRunner) Run(ctx context.Context, job *Job) error {
	a.wg[job.Id] = &sync.WaitGroup{}
	a.exitChannel[job.Id] = make(chan struct{})

	a.wg[job.Id].Add(len(job.tasks))

	errs := make([]error, len(job.tasks), len(job.tasks))

	for i, task := range job.tasks {
		go func(i int, task TaskFunc) {
			select {
			case <-a.exitChannel[job.Id]:
				a.wg[job.Id].Done()
				close(a.exitChannel[job.Id])
				delete(a.exitChannel, job.Id)
				return
			default:
				err := task(ctx)
				errs[i] = err
				a.wg[job.Id].Done()
			}
		}(i, task)
	}

	a.wg[job.Id].Wait()
	delete(a.wg, job.Id)

	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	delete(a.exitChannel, job.Id)
	delete(a.wg, job.Id)
	if job.onComplete != nil {
		return job.onComplete(ctx)
	}

	return nil
}

func (a *AsyncRunner) Stop(_ context.Context, job *Job) error {
	for {
		if _, ok := a.exitChannel[job.Id]; ok {
			break
		}
	}
	a.exitChannel[job.Id] <- struct{}{}
	fmt.Println("signal sent")
	return nil
}
