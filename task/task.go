package task

import (
	"context"
	"github.com/getcohesive/dag"
)

func Of(d *dag.Dag) dag.TaskFunc {
	return func(ctx context.Context) error {
		return d.Run(ctx)
	}
}
