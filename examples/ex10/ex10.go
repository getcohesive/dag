package main

import (
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
	println("f3")
	return nil
}

func f4() error {
	println("f4")
	return nil
}

func f5() error {
	println("f5")
	return nil
}
