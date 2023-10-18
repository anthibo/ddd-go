package app

import "chat-service/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}


type Commands struct {
	CreateMessage   command.CreateMessageHandler
	CreateUser command.CreateUserHandler

}

type Queries struct {
	
}