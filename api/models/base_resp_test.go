package models_test

import (
	"fmt"
	"testing"
	"z3ntl3/telegram-simpaix-bot/api/models"
	 "z3ntl3/telegram-simpaix-bot/api/models/types"
)

func TestBasectx(t *testing.T) {
	
	var q models.API_Resp[*types.CallbackQuery] = models.API_Resp[*types.CallbackQuery]{
		Result: &types.CallbackQuery{
			Data: 1,
		},
	}

	fmt.Printf("%+v", q.Result)
}
