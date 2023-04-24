package dag

import "context"

var asyncRunner = (&AsyncRunner{}).New()

func run(ctx context.Context, job *Job) error {
	if job.sequential {
		return runSync(ctx, job)
	} else {
		return asyncRunner.Run(ctx, job)
	}
}

func stop(ctx context.Context, job *Job) error {
	if job.sequential {
		return asyncRunner.Stop(ctx, job)
	}
	return nil
}
