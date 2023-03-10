package main

import (
	"context"
	"sync"

	"github.com/getcohesive/dag"
)

var wg = &sync.WaitGroup{}

func main() {

	wg.Add(1)

	ctx := context.Background()
	d := dag.New()
	d.Pipeline(f1, f2).Then().Spawns(f3, f4)
	d.RunAsync(ctx, onComplete)

	wg.Wait()
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

func onComplete(err error) {
	println("All functions have completed")
	wg.Done()
}
