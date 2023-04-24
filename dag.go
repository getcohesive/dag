package dag

import (
	"context"
)

// Dag represents directed acyclic graph
type Dag struct {
	jobs []*Job
}

// New creates new DAG
func New() *Dag {
	return &Dag{
		jobs: make([]*Job, 0),
	}
}

func (dag *Dag) lastJob() *Job {
	jobsCount := len(dag.jobs)
	if jobsCount == 0 {
		return nil
	}

	return dag.jobs[jobsCount-1]
}

// Run starts the tasks
// It will block until all functions are done
func (dag *Dag) Run(ctx context.Context) error {
	for _, job := range dag.jobs {
		err := run(ctx, job)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dag *Dag) Stop(ctx context.Context) error {
	for _, job := range dag.jobs {
		err := stop(ctx, job)
		if err != nil {
			return err
		}
	}
	return nil
}

// RunAsync executes Run on another goroutine
func (dag *Dag) RunAsync(ctx context.Context, onComplete func(error)) chan struct{} {
	exitChannel := make(chan struct{})
	go func() {
		select {
		case <-exitChannel:

		default:
			err := dag.Run(ctx)
			if onComplete != nil {
				onComplete(err)
			}
		}
	}()
	return exitChannel
}

// Pipeline executes tasks sequentially
func (dag *Dag) Pipeline(tasks ...TaskFunc) *PipelineResult {
	job := &Job{
		tasks:      make([]TaskFunc, len(tasks)),
		sequential: true,
	}

	for i, task := range tasks {
		job.tasks[i] = task
	}

	dag.jobs = append(dag.jobs, job)

	return &PipelineResult{
		dag,
	}
}

// Spawns executes tasks concurrently
func (dag *Dag) Spawns(tasks ...TaskFunc) *SpawnsResult {

	job := &Job{
		tasks:      make([]TaskFunc, len(tasks)),
		sequential: false,
	}

	for i, task := range tasks {
		job.tasks[i] = task
	}

	dag.jobs = append(dag.jobs, job)

	return &SpawnsResult{
		dag,
	}
}

//
// func (dag *Dag) Stop(ctx context.Context) error {
// 	var errs []error
// 	for _, job := range dag.jobs {
// 		if job.Stop() != nil {
// 			err := job.Stop()
// 			errs = append(errs, err)
// 		}
// 	}
// 	accumulatedErrString := ""
// 	for _, err := range errs {
// 		accumulatedErrString += err.Error() + "|"
// 	}
//
// 	return errors.New(accumulatedErrString)
// }
