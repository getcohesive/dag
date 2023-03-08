package task

import "github.com/getcohesive/dag"

func Of(d *dag.Dag) func() {
	return func() {
		d.Run()
	}
}
