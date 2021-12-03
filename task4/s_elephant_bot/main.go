package main

import (
	//	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	//"gopkg.in/telegram-bot-api.v4"
)

//type Config struct {
//	tgbotapikey string
//}
const tgbotapikey = "2137759225:AAF_bU-jiHY0-TCPwXZcLQfWFGeQabemLaU"

func main() {
	var (
		//file, _ := os.Open("config.json")
		//decoder := json.NewDecoder(file)
		//configuration := Config{}
		//err := decoder.Decode(&configuration)
		//if err != nil {
		//		log.Panic(err)
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

	//	btn := tgbotapi.KeyboardButton{Text: "My Git"}

	for {

		update = <-updChannel
		if update.Message != nil {

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if strings.ToLower(cmdText) == "git" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"https://github.com/79donaldson/andersen_devops_2021")
					bot.Send(msgConfig)
				} else if strings.ToLower(cmdText) == "tasks" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"https://github.com/79donaldson/andersen_devops_2021/tree/main/task1\n"+
							"https://github.com/79donaldson/andersen_devops_2021/tree/main/task2\n"+
							"https://github.com/79donaldson/andersen_devops_2021/tree/main/task3\n"+
							"https://github.com/79donaldson/andersen_devops_2021/tree/main/task4\n")
					bot.Send(msgConfig)
				} else if strings.ToLower(cmdText) == "task1" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"https://github.com/79donaldson/andersen_devops_2021/tree/main/task1")
					bot.Send(msgConfig)
				} else if strings.ToLower(cmdText) == "task2" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"https://github.com/79donaldson/andersen_devops_2021/tree/main/task2")
					bot.Send(msgConfig)
				} else if strings.ToLower(cmdText) == "task3" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"https://github.com/79donaldson/andersen_devops_2021/tree/main/task3")
					bot.Send(msgConfig)
				} else if strings.ToLower(cmdText) == "task4" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"https://github.com/79donaldson/andersen_devops_2021/tree/main/task4")
					bot.Send(msgConfig)
				}
			} else {
				fmt.Printf(
					"from: %s; chatID: %v; messege: %s\n",
					update.Message.From.FirstName,
					update.Message.Chat.ID,
					update.Message.Text)

				msgConfig := tgbotapi.NewMessage(
					update.Message.Chat.ID,
					update.Message.Text)
				bot.Send(msgConfig)
			}
		} else {
			fmt.Printf("not messege: %+v\n", update)
		}

	}

	bot.StopReceivingUpdates()
}
