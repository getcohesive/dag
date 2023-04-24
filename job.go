package dag

import "context"

type TaskFunc func(ctx context.Context) error

// Job - Each job consists of one or more tasks
// Each Job can runs tasks in order(Sequential) or unordered
type Job struct {
	Id         string
	tasks      []TaskFunc
	sequential bool
	onComplete TaskFunc
}
