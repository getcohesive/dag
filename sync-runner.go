package dag

func runSync(job *Job) error {
	for _, task := range job.tasks {
		err := task()
		if err != nil {
			return err
		}
	}

	if job.onComplete != nil {
		return job.onComplete()
	}

	return nil
}
