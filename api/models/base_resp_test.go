package models_test

import (
	"fmt"
	"testing"

	"z3ntl3/telegram-simpaix-bot/api/models"
)

func TestBasectx(t *testing.T) {
	var q models.API_Resp[*models.CallbackQuery] = models.API_Resp[*models.CallbackQuery]{
		Result: &models.CallbackQuery{
			Data: 1,
		},
	}

	fmt.Printf("%+v", q.Result)
}
