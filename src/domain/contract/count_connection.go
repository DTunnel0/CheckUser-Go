package contract

import "context"

type CountConnection interface {
	SetNext(connection CountConnection)
	ByUsername(ctx context.Context, username string) (int, error)
	All(ctx context.Context) (int, error)
}
