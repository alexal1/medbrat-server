package usecase

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type HelloUseCase interface {
	GetFirstMessage() *Message
}

// ---------------------------------------------------------------------------------------------------------------------

type hello struct {
	firstMessage *Message
}

func NewHello(startVision func(imageBase64 *string) (nextMessage *Message), startManual func() (nextMessage *Message)) HelloUseCase {
	msgHello1 := Message{
		Id:   NextMessageId(),
		Text: "Добрый день!",
	}

	msgHello2 := Message{
		Id:   NextMessageId(),
		Text: "Я постараюсь помочь Вам провести диагностику Вашего здоровья",
	}

	msgHello3 := Message{
		Id:   NextMessageId(),
		Text: "Для начала работы мне необходимо получить данные Вашего общего анализа крови",
	}

	msgHello4 := Message{
		Id:              NextMessageId(),
		Text:            "Вы можете сфотографировать такой анализ и отправить мне фото?",
		PossibleAnswers: []Answer{Yes, No},
	}

	msgHello5 := Message{
		Id:   NextMessageId(),
		Text: "Ок, постарайтесь держать камеру ровно над листом и сделайте фото",
	}

	msgHello6 := Message{
		Id:   NextMessageId(),
		Text: "Тогда придется внести все значения вручную",
	}

	msgHello1.NextMessage = &msgHello2
	msgHello2.NextMessage = &msgHello3
	msgHello3.NextMessage = &msgHello4

	msgHello4.NextMessageByValue = func(value interface{}) (nextMessage *Message) {
		switch value {
		case Yes:
			nextMessage = &msgHello5
		case No:
			nextMessage = &msgHello6
		}
		return
	}

	msgHello5.NextMessageByValue = func(value interface{}) (nextMessage *Message) {
		stringValue := value.(string)
		return startVision(&stringValue)
	}

	msgHello6.NextMessage = startManual()

	return &hello{
		&msgHello1,
	}
}

func (h *hello) GetFirstMessage() *Message {
	return h.firstMessage
}
