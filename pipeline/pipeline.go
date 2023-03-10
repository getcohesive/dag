package pipeline

import (
	"context"
	"github.com/getcohesive/dag"
)

// Of wraps tasks as a single function
func Of(tasks ...dag.TaskFunc) dag.TaskFunc {
	return func(ctx context.Context) error {
		for _, task := range tasks {
			err := task(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
