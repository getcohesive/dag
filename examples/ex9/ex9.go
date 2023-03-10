package main

import (
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

	d.Run()
}

func g(n string) func() error {
	return func() error {
		println(n)
		return nil
	}
}
