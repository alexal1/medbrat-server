package usecase

import (
	"sync/atomic"
)

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type GlobalUseCase interface {
	Start() (messages []*Message)
	Answer(value interface{}) (messages []*Message)
}

type OCR interface {
	RecognizeGeneralBloodTest(blood *BloodGeneral, image *string)
}

// ---------------------------------------------------------------------------------------------------------------------

type global struct {
	currentMessage *Message
	hello          *HelloUseCase
	vision         *VisionUseCase
	manual         *ManualUseCase
	blood          *BloodGeneral
}

var lastMessageId uint64 = 0

func NextMessageId() uint64 {
	return atomic.AddUint64(&lastMessageId, 1)
}

func NewGlobal(ocr *OCR) GlobalUseCase {
	blood := NewBloodGeneral()
	vision := NewVision(&blood, ocr)
	manual := NewManual(&blood)
	hello := NewHello(vision.Start, manual.Start)
	return &global{
		hello.GetFirstMessage(),
		&hello,
		&vision,
		&manual,
		&blood,
	}
}

func (g *global) Start() (messages []*Message) {
	messages, g.currentMessage = zipMessages((*g.hello).GetFirstMessage())
	return
}

func (g *global) Answer(value interface{}) (messages []*Message) {
	nextMessage := g.currentMessage.NextMessageByValue(value)
	messages, g.currentMessage = zipMessages(nextMessage)
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
