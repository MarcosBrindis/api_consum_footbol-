package ports

import "context"

type EventConsumer interface {
	GetMessage(ctx context.Context) (string, error)
}
