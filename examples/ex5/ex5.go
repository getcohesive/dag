package main

import (
	"context"
	"github.com/getcohesive/dag"
)

func main() {
	d := dag.New()
	d.Spawns(f1, f2, f3).
		Join().
		Pipeline(f4, f5).
		Then().
		Spawns(f6, f7, f8)
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

func f6(ctx context.Context) error {
	println("f6")
	return nil
}

func f7(ctx context.Context) error {
	println("f7")
	return nil
}

func f8(ctx context.Context) error {
	println("f8")
	return nil
}
