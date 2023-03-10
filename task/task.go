package task

import "github.com/getcohesive/dag"

func Of(d *dag.Dag) func() error {
	return func() error {
		return d.Run()
	}
}
