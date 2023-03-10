package main

import (
	"github.com/getcohesive/dag"
)

func main() {
	d := dag.New()
	d.Pipeline(f1, f2).OnComplete(f3).
		Then().
		Spawns(f1, f2).OnComplete(f4)
	d.Run()
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
	println("complete")
	return nil
}

func f4() error {
	println("finish")
	return nil
}
