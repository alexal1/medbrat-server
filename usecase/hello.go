package usecase

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type HelloUseCase interface {
	GetFirstMessage() *Message
}

// ---------------------------------------------------------------------------------------------------------------------

type hello struct {
	firstMessage *Message
}

func NewHello() HelloUseCase {
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
		Text:         "Пидора ответ",
		AnswerFormat: None,
	}

	msgHello1.NextMessage = &msgHello2
	msgHello2.NextMessage = &msgHello3
	msgHello3.NextMessage = &msgHello4
	msgHello4.NextMessageByCondition = map[interface{}]*Message{
		AnswerYes: &msgHello5,
		AnswerNo:  &msgHello6,
	}

	return &hello{
		&msgHello1,
	}
}

func (h *hello) GetFirstMessage() *Message {
	return h.firstMessage
}
