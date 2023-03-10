package main

import "github.com/getcohesive/dag"

func main() {
	d := dag.New()
	d.Spawns(f1, f2, f3).
		Join().
		Pipeline(f4, f5).
		Then().
		Spawns(f6, f7, f8)
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

func f6() error {
	println("f6")
	return nil
}

func f7() error {
	println("f7")
	return nil
}

func f8() error {
	println("f8")
	return nil
}
