package main

import (
	"context"
	"github.com/getcohesive/dag"
	"github.com/getcohesive/dag/pipeline"
	"github.com/getcohesive/dag/task"
)

func main() {
	d1 := dag.New()
	d1.Pipeline(g("f1"), g("f2")).Then().Spawns(pipeline.Of(g("f3"), g("f5")), pipeline.Of(g("f4"), g("f6"))).Join().Pipeline(g("f7"))

	d2 := dag.New()
	d2.Pipeline(g("f8"), g("f9")).Then().Spawns(g("f10"), g("f11")).Join().Pipeline(g("f12"))

	d := dag.New()
	d.Spawns(task.Of(d1), task.Of(d2)).Join().Pipeline(g("f13"))

	_ = d.Run(context.Background())
}

func g(n string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		println(n)
		return nil
	}
}
