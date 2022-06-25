package main

import (
	"fmt" // we will create this later

	"github.com/decendgame/bot/bot"
	"github.com/decendgame/bot/config"
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
