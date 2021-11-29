package main

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const tgbotapikey = "2137759225:AAF_bU-jiHY0-TCPwXZcLQfWFGeQabemLaU"

func main() {
	var (
		bot        *tgbotapi.BotAPI
		err        error
		updChannel tgbotapi.UpdatesChannel
		update     tgbotapi.Update
		updConfig  tgbotapi.UpdateConfig
	)
	bot, err = tgbotapi.NewBotAPI(tgbotapikey)
	if err != nil {
		panic("bot init error: " + err.Error())
	}

	updConfig.Timeout = 60
	updConfig.Limit = 1
	updConfig.Offset = 0

	updChannel, err = bot.GetUpdatesChan(updConfig)
	if err != nil {
		panic("update channel error: " + err.Error())
	}

	for {

		update = <-updChannel
		if update.Message != nil {
			fmt.Printf("messege: %s\n", update.Message.Text)
		}
	}

	bot.StopReceivingUpdates()
}
