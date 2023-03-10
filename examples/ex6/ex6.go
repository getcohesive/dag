package main

import "github.com/getcohesive/dag"
import "github.com/getcohesive/dag/pipeline"

func main() {
	d := dag.New()
	d.Spawns(pipeline.Of(f1, f3), pipeline.Of(f2, f4)).
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
