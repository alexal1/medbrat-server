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
		Id:           NextMessageId(),
		Text:         "Добрый день!",
		AnswerFormat: None,
	}

	msgHello2 := Message{
		Id:           NextMessageId(),
		Text:         "Меня зовут МедБрат, я помогу вам поставить диагноз",
		AnswerFormat: None,
	}

	msgHello3 := Message{
		Id:           NextMessageId(),
		Text:         "Мне надо посмотреть на ваш общий анализ крови, а потом я задам вам несколько вопросов о самочувствии. Прямо как настоящий врач!",
		AnswerFormat: None,
	}

	msgHello4 := Message{
		Id:           NextMessageId(),
		Text:         "Вы можете сфотографировать общий анализ крови?",
		AnswerFormat: YesNo,
	}

	msgHello5 := Message{
		Id:           NextMessageId(),
		Text:         "Ок, постарайтесь держать камеру ровно над листом и сделайте фото",
		AnswerFormat: Image,
	}

	msgHello6 := Message{
		Id:           NextMessageId(),
		Text:         "Тогда придется внести все значения вручную",
		AnswerFormat: None,
	}

	msgHello1.NextMessage = &msgHello2
	msgHello2.NextMessage = &msgHello3
	msgHello3.NextMessage = &msgHello4

	msgHello4.NextMessageByValue = func(value interface{}) (nextMessage *Message) {
		switch value {
		case AnswerYes:
			nextMessage = &msgHello5
		case AnswerNo:
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
