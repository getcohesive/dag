package dag

import "context"

func runSync(ctx context.Context, job *Job) error {
	for _, task := range job.tasks {
		err := task(ctx)
		if err != nil {
			return err
		}
	}

	if job.onComplete != nil {
		return job.onComplete(ctx)
	}

	return nil
}
