package main

import "github.com/getcohesive/dag"

func main() {
	d := dag.New()
	d.Pipeline(f1, f2, f3)
	_ = d.Run()
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
