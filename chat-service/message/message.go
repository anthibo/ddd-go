package message

import (
	"errors"
	"time"
)

type Message struct {
	ID       int
	Username string
	RoomID   string
	Message  string
	time     time.Time
}

func NewMessage(username string, roomID string, message string, time time.Time) (*Message, error) {
	if roomID == "" {
		return nil, errors.New("empty room id")
	}

	if username == "" {
		return nil, errors.New("empty username")
	}

	if message == "" {
		return nil, errors.New("empty username")
	}
	return &Message{
		time:     time,
		ID:       1,
		Username: username,
		RoomID:   roomID,
		Message:  message,
	}, nil
}
