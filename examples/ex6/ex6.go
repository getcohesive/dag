package main

import (
	"context"
	"github.com/getcohesive/dag"
)
import "github.com/getcohesive/dag/pipeline"

func main() {
	d := dag.New()
	d.Spawns(pipeline.Of(f1, f3), pipeline.Of(f2, f4)).
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
