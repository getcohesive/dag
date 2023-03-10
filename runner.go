package dag

func run(job *Job) error {
	if job.sequential {
		return runSync(job)
	} else {
		return runAsync(job)
	}
}
