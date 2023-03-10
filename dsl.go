package dag

type PipelineResult struct {
	dag *Dag
}

func (result *PipelineResult) Then() *PipelineDSL {
	return &PipelineDSL{
		result.dag,
	}
}

func (result *PipelineResult) OnComplete(action func() error) *PipelineResult {
	job := result.dag.lastJob()
	if job != nil {
		job.onComplete = action
	}
	return result
}

type PipelineDSL struct {
	dag *Dag
}

func (dsl *PipelineDSL) Spawns(tasks ...func() error) *SpawnsResult {
	dsl.dag.Spawns(tasks...)
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

func (result *SpawnsResult) OnComplete(action func() error) *SpawnsResult {
	job := result.dag.lastJob()
	if job != nil {
		job.onComplete = action
	}
	return result
}

type SpawnsDSL struct {
	dag *Dag
}

func (dsl *SpawnsDSL) Pipeline(tasks ...func() error) *PipelineResult {
	dsl.dag.Pipeline(tasks...)
	return &PipelineResult{
		dsl.dag,
	}
}
