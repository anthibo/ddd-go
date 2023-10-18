package command

import (
	"chat-service/message"
	"common/decorator"
	"context"

	"github.com/sirupsen/logrus"
)

type CreateMessage struct {
	message string
}

type CreateMessageHandler decorator.CommandHandler[CreateMessage]

type createMessageHandler struct {
	messageRepo message.Repository
}

func NewMakeHoursAvailableHandler(
	messageRepo message.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CreateMessageHandler {
	if messageRepo == nil {
		panic("hourRepo is nil")
	}

	return decorator.ApplyCommandDecorators[CreateMessage](
		createMessageHandler{messageRepo: messageRepo},
		logger,
		metricsClient,
	)
}

func (c createMessageHandler) Handle(ctx context.Context, cmd CreateMessage) error {
	// create message logic here

	return nil
}
