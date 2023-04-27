package dag

type PipelineResult struct {
	dag *Dag
}

func (result *PipelineResult) Then() *PipelineDSL {
	return &PipelineDSL{
		result.dag,
	}
}

func (result *PipelineResult) OnComplete(action TaskFunc) *PipelineResult {
	job := result.dag.lastJob()
	if job != nil {
		job.onComplete = action
	}
	return result
}

type PipelineDSL struct {
	dag *Dag
}

func (dsl *PipelineDSL) Spawns(jobId string, tasks ...TaskFunc) *SpawnsResult {
	dsl.dag.Spawns(jobId, tasks...)
	return &SpawnsResult{
		dsl.dag,
	}
}

type SpawnsResult struct {
	dag *Dag
}

func (result *SpawnsResult) Join() *SpawnsDSL {
	return &SpawnsDSL{
		result.dag,
	}
}

func (result *SpawnsResult) OnComplete(action TaskFunc) *SpawnsResult {
	job := result.dag.lastJob()
	if job != nil {
		job.onComplete = action
	}
	return result
}

type SpawnsDSL struct {
	dag *Dag
}

func (dsl *SpawnsDSL) Pipeline(jobId string, tasks ...TaskFunc) *PipelineResult {
	dsl.dag.Pipeline(jobId, tasks...)
	return &PipelineResult{
		dsl.dag,
	}
}
