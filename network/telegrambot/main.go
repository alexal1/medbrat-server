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

		switch message.Text {
		case "/start":
			onNewChatStarted(chat)
			break

		case "Да":
			onAnswerReceived(chat, usecase.Yes)
			break

		case "Нет":
			onAnswerReceived(chat, usecase.No)
			break

		case "Начать сначала":
			onAnswerReceived(chat, usecase.StartAgain)
			break

		case "Я передумал":
			onAnswerReceived(chat, usecase.ChangedMyMind)
			break

		case "Хорошо":
			onAnswerReceived(chat, usecase.AllRight)
			break

		default:
			onAnswerReceived(chat, message.Text)
		}

		// TODO: handle image answers
	}

	return nil
}

func (telegramBot *telegrambot) SendMessage(chat network.Chat, message string, possibleAnswers []usecase.Answer) {
	if chat.Source != network.Telegram {
		log.Println("Attempt to send a non-Telegram message by Telegram Bot API")
		return
	}

	bot := telegramBot.bot

	msg := tgbotapi.NewMessage(chat.Id, message)

	if replyKeyboard := getReplyKeyboard(possibleAnswers); replyKeyboard != nil {
		msg.ReplyMarkup = replyKeyboard
	}

	if _, err := bot.Send(msg); err != nil {
		log.Println("Error when sending message by Telegram Bot API")
		return
	}
}

func getReplyKeyboard(possibleAnswers []usecase.Answer) interface{} {
	if len(possibleAnswers) == 0 {
		return tgbotapi.NewRemoveKeyboard(false)
	}

	var buttons []tgbotapi.KeyboardButton
	for _, answer := range possibleAnswers {
		switch answer {
		case usecase.Yes:
			buttons = append(buttons, tgbotapi.NewKeyboardButton("Да"))
			break

		case usecase.No:
			buttons = append(buttons, tgbotapi.NewKeyboardButton("Нет"))
			break

		case usecase.StartAgain:
			buttons = append(buttons, tgbotapi.NewKeyboardButton("Начать сначала"))
			break

		case usecase.ChangedMyMind:
			buttons = append(buttons, tgbotapi.NewKeyboardButton("Я передумал"))
			break

		case usecase.AllRight:
			buttons = append(buttons, tgbotapi.NewKeyboardButton("Хорошо"))
			break
		}
	}

	keyboard := tgbotapi.NewReplyKeyboard(buttons)
	keyboard.OneTimeKeyboard = true

	return keyboard
}
