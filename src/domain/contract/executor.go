package contract

import "context"

type Executor interface {
	Execute(ctx context.Context, command string) (string, error)
}
