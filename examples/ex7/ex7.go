package main

import (
	"context"
	"github.com/getcohesive/dag"
)

func main() {
	d := dag.New()
	d.Pipeline(f1, f2).OnComplete(f3).
		Then().
		Spawns(f1, f2).OnComplete(f4)
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
	println("complete")
	return nil
}

func f4(ctx context.Context) error {
	println("finish")
	return nil
}
