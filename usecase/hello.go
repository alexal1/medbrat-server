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
	var msgStart1, msgStart2, msgStart3, msgStart4, msgStart5, msgNoBlood1, msgNoBlood2, msgVision1, msgManual1, msgManual2 Message

	msgStart1 = Message{
		Id:          NextMessageId(),
		Text:        "Добрый день!",
		NextMessage: &msgStart2,
	}

	msgStart2 = Message{
		Id:          NextMessageId(),
		Text:        "Я постараюсь помочь Вам провести диагностику Вашего здоровья.",
		NextMessage: &msgStart3,
	}

	msgStart3 = Message{
		Id:          NextMessageId(),
		Text:        "Для начала работы мне необходимо изучить состав Вашей крови.",
		NextMessage: &msgStart4,
	}

	msgStart4 = Message{
		Id:              NextMessageId(),
		Text:            "У Вас есть общий анализ крови, полученный в любой клинике или поликлинике?",
		PossibleAnswers: []Answer{Yes, No},
		NextMessageByValue: func(value interface{}) (nextMessage *Message) {
			switch value {
			case Yes:
				nextMessage = &msgStart5
				break
			case No:
				nextMessage = &msgNoBlood1
				break
			}
			return
		},
	}

	msgNoBlood1 = Message{
		Id:              NextMessageId(),
		Text:            "Тогда мне очень жаль, но в текущей версии я ничем не могу Вам помочь. Без анализа крови мне не хватит информации для исследования. Вы можете сдать анализы в ближайшем медицинском центре.",
		PossibleAnswers: []Answer{StartAgain},
		NextMessageByValue: func(value interface{}) (nextMessage *Message) {
			switch value {
			case StartAgain:
				nextMessage = &msgNoBlood2
				break
			}
			return
		},
	}

	msgNoBlood2 = Message{
		Id:          NextMessageId(),
		Text:        "Добрый день ещё раз!",
		NextMessage: &msgStart2,
	}

	msgStart5 = Message{
		Id:              NextMessageId(),
		Text:            "Хорошо. Сможете сфотографировать и отправить мне этот анализ, чтобы я его распознал?",
		PossibleAnswers: []Answer{Yes, No},
		NextMessageByValue: func(value interface{}) (nextMessage *Message) {
			switch value {
			case Yes:
				nextMessage = &msgVision1
				break
			case No:
				nextMessage = &msgManual1
				break
			}
			return
		},
	}

	msgManual1 = Message{
		Id:              NextMessageId(),
		Text:            "Ладно, тогда давайте я пройдусь по всем пунктам анализа крови, и вы назовете соответствующие значения.",
		PossibleAnswers: []Answer{AllRight, No},
		NextMessageByValue: func(value interface{}) (nextMessage *Message) {
			switch value {
			case AllRight:
				nextMessage = &msgManual2
				break
			case No:
				nextMessage = &msgNoBlood1
				break
			}
			return
		},
	}

	msgManual2 = Message{
		Id:          NextMessageId(),
		Text:        "Я буду называть строки анализа крови по одной. Вписывайте соответствующие значения, если они есть, или нажимайте \"Пропустить\".",
		NextMessage: startManual(),
	}

	msgVision1 = Message{
		Id:              NextMessageId(),
		Text:            "Ок. Пришлите одной фотографией, постарайтесь чтобы слова и числа были хорошо различимы. Неправильно распознанные значения можно будет подкорректировать. Имейте в виду, что распознавание рукописного текста в текущей версии не поддерживается.",
		PossibleAnswers: []Answer{ChangedMyMind},
		NextMessageByValue: func(value interface{}) (nextMessage *Message) {
			switch value {
			case ChangedMyMind:
				nextMessage = &msgManual1
				break
			}
			return
		},
	}

	return &hello{
		&msgStart1,
	}
}

func (h *hello) GetFirstMessage() *Message {
	return h.firstMessage
}
