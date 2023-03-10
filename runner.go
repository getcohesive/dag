package dag

import "context"

func run(ctx context.Context, job *Job) error {
	if job.sequential {
		return runSync(ctx, job)
	} else {
		return runAsync(ctx, job)
	}
}
