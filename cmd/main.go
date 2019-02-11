package main

import (
	"log"
	"medbrat-server/network"
	"medbrat-server/network/ocr"
	"medbrat-server/network/telegrambot"
	"medbrat-server/usecase"
)

var storage map[network.Chat]*usecase.GlobalUseCase
var ocrInstance = ocr.NewOCR()
var telegramBot = telegrambot.NewTelegramBot()

func main() {
	storage = make(map[network.Chat]*usecase.GlobalUseCase)

	if err := telegramBot.Run(onNewChatStarted, onAnswerReceived); err != nil {
		log.Fatal("Cannot start Telegram Bot: ", err)
	} else {
		log.Println("Telegram Bot is running!")
	}
}

func onNewChatStarted(chat network.Chat) {
	globalUsecase := usecase.NewGlobal(&ocrInstance)
	storage[chat] = &globalUsecase
	startMessages := globalUsecase.Start()
	sendMessages(chat, startMessages)
}

func onAnswerReceived(chat network.Chat, answer interface{}) {
	globalUsecase := storage[chat]
	if globalUsecase == nil {
		log.Printf("We've received an answer from an unknown chat: Source=%s, Id=%d", chat.Source, chat.Id)
		return
	}

	resonseMessages := (*globalUsecase).Answer(answer)
	sendMessages(chat, resonseMessages)
}

func sendMessages(chat network.Chat, messages []*usecase.Message) {
	for _, message := range messages {
		switch chat.Source {
		case network.Telegram:
			telegramBot.SendMessage(chat, message.Text, message.PossibleAnswers)
			break

		case network.Alice:
			log.Println("Alice is not implemented yet!")
			break
		}
	}
}
