package usecase

import "sync/atomic"

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type GlobalUseCase interface {
	Start() (messages []*Message)
}

// ---------------------------------------------------------------------------------------------------------------------

type global struct {
	currentMessage *Message
	hello          *HelloUseCase
}

var lastMessageId uint64 = 0

func NextMessageId() uint64 {
	return atomic.AddUint64(&lastMessageId, 1)
}

func NewGlobal() GlobalUseCase {
	hello := NewHello()
	return &global{
		hello.GetFirstMessage(),
		&hello,
	}
}

func (g *global) Start() (messages []*Message) {
	messages, g.currentMessage = zipMessages((*g.hello).GetFirstMessage())
	return
}

func zipMessages(startMessage *Message) (messages []*Message, lastMessage *Message) {
	lastMessage = startMessage
	for {
		messages = append(messages, lastMessage)
		if lastMessage.NextMessage != nil {
			lastMessage = lastMessage.NextMessage
		} else {
			break
		}
	}
	return
}
