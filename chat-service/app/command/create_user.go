package command

import (
	"chat-service/message"
	"common/decorator"
	"context"

	"github.com/sirupsen/logrus"
)

type CreateUser struct {
	username string
}

type CreateUserHandler decorator.CommandHandler[CreateUser]

type createUserHandler struct {
	// TODO: replace with user repository
	userRepo message.Repository
}

// TODO: replace with user repository
func NewCreateUserHandler(
	messageRepo message.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CreateUserHandler {
	if messageRepo == nil {
		panic("hourRepo is nil")
	}

	// TODO: replace with user repository
	return decorator.ApplyCommandDecorators[CreateUser](
		createUserHandler{userRepo: messageRepo},
		logger,
		metricsClient,
	)
}

func (c createUserHandler) Handle(ctx context.Context, cmd CreateUser) error {
	// create message logic here

	return nil
}
