package dag

import "sync"

func runAsync(job *Job) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(job.tasks))
	errs := make([]error, len(job.tasks), len(job.tasks))

	for i, task := range job.tasks {
		go func(i int, task func() error) {
			err := task()
			errs[i] = err
			wg.Done()
		}(i, task)
	}

	wg.Wait()

	for _, e := range errs {
		if e != nil {
			return e
		}
	}

	if job.onComplete != nil {
		return job.onComplete()
	}

	return nil
}
