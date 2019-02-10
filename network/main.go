package network

import "medbrat-server/usecase"

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type Network interface {
	Run(
		onNewChatStarted func(chat Chat),
		onAnswerReceived func(chat Chat, answer interface{}),
	) (err error)
	SendMessage(chat Chat, message string, possibleAnswers []usecase.Answer)
}

type Chat struct {
	Id     int64
	Source ChatSource
}

type ChatSource string

const (
	Telegram ChatSource = "Telegram"
	Alice    ChatSource = "Alice"
)

// ---------------------------------------------------------------------------------------------------------------------
