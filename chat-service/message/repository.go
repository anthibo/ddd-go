package message

import "context"

type Repository interface {
	CreateMessage(ctx context.Context, msg Message) (*Message, error)
}
