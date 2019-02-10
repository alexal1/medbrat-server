package telegrambot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"medbrat-server/configs"
	"medbrat-server/network"
	"medbrat-server/usecase"
)

type telegrambot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot() network.Network {
	return &telegrambot{}
}

func (telegramBot *telegrambot) Run(onNewChatStarted func(chat network.Chat), onAnswerReceived func(chat network.Chat, answer interface{})) (err error) {
	bot, err := tgbotapi.NewBotAPI(configs.TelegramBotToken)
	if err != nil {
		return
	}

	telegramBot.bot = bot

	//noinspection GoBoolExpressions
	bot.Debug = configs.DebugMode

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updatesChan, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		return
	}

	for update := range updatesChan { // receive values from the channel repeatedly until it is closed
		message := update.Message

		if message == nil { // ignore any non-Message Updates
			continue
		}

		chat := network.Chat{
			Id:     message.Chat.ID,
			Source: network.Telegram,
		}

		if message.Text == "/start" {
			onNewChatStarted(chat)
			continue
		}

		onAnswerReceived(chat, message.Text)

		// TODO: handle image answers
	}

	return nil
}

func (telegramBot *telegrambot) SendMessage(chat network.Chat, message string, answerFormat usecase.AnswerFormat) {
	if chat.Source != network.Telegram {
		log.Println("Attempt to send a non-Telegram message by Telegram Bot API")
		return
	}

	bot := telegramBot.bot

	msg := tgbotapi.NewMessage(chat.Id, message)
	if _, err := bot.Send(msg); err != nil {
		log.Println("Error when sending message by Telegram Bot API")
		return
	}

	// TODO: implement answer format
}
