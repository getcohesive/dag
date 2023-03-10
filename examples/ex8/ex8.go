package main

import (
	"sync"

	"github.com/getcohesive/dag"
)

var wg = &sync.WaitGroup{}

func main() {

	wg.Add(1)

	d := dag.New()
	d.Pipeline(f1, f2).Then().Spawns(f3, f4)
	d.RunAsync(onComplete)

	wg.Wait()
}

func f1() error {
	println("f1")
	return nil
}

func f2() error {
	println("f2")
	return nil
}

func f3() error {
	println("f3")
	return nil
}

func f4() error {
	println("f4")
	return nil
}

func onComplete(err error) {
	println("All functions have completed")
	wg.Done()
}
