package main

import (
	"context"
	"github.com/getcohesive/dag"
	"github.com/getcohesive/dag/task"
)

func main() {
	d1 := dag.New()
	d1.Pipeline(f1, f3)

	d2 := dag.New()
	d2.Pipeline(f2, f4)

	d := dag.New()
	d.Spawns(task.Of(d1), task.Of(d2)).
		Join().
		Pipeline(f5)
	_ = d.Run(context.Background())
}

func f1(ctx context.Context) error {
	println("f1")
	return nil
}

func f2(ctx context.Context) error {
	println("f2")
	return nil
}

func f3(ctx context.Context) error {
	println("f3")
	return nil
}

func f4(ctx context.Context) error {
	println("f4")
	return nil
}

func f5(ctx context.Context) error {
	println("f5")
	return nil
}
