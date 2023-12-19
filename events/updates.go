package events

import (
	"z3ntl3/telegram-simpaix-bot/api/models"
)

type Event struct {
	models.API_Resp[[]models.Update]
}

type EventChan = chan Event

func UpdateRetriever_Worker() {
}
