package main

import (
	"dc-bot/bot"
	"dc-bot/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}