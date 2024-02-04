package contract

import "context"

type Connection interface {
	SetNext(connection Connection)
	CountByUsername(ctx context.Context, username string) int
	Count(ctx context.Context) int
}
