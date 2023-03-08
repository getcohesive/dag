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

func f1() {
	println("f1")
}
func f2() {
	println("f2")
}
func f3() {
	println("f3")
}
func f4() {
	println("f4")
}
func f5() {
	println("f5")
}
