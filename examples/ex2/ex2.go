package main

import (
	"context"
	"github.com/getcohesive/dag"
)

func main() {
	d := dag.New()
	d.Spawns(f1, f2, f3)
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
