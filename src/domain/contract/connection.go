package contract

import "context"

type Connection interface {
	SetNext(connection Connection)
	CountByUsername(ctx context.Context, username string) (int, error)
	Count(ctx context.Context) (int, error)
}
