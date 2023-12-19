package main

import (
	"fmt"
	"log"
	"time"

	"z3ntl3/telegram-simpaix-bot/api"
	"z3ntl3/telegram-simpaix-bot/api/models"
	"z3ntl3/telegram-simpaix-bot/driver"

	"github.com/spf13/viper"
)

func main() {
	// under development see '/api'
	driver.ReadEnv_Variables()

	c := api.NewClient(viper.GetString("botToken"), time.Second*10)
	bot_, err := c.Authorize()
	/*
		returns interface{}, bcs it is a composition of a generic struct that links child
		with parent
	*/if err != nil {
		log.Fatal(err)
	}

	bot, err := (api.Transform[models.API_Resp[models.User]](bot_))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bot)
}
