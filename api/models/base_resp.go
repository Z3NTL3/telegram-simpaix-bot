/*
Work for: Simpaix.net
Author: Z3NTL3
License: GNU
*/
package models

import "z3ntl3/telegram-simpaix-bot/api/models/types"

type BaseResp struct {
	Success     bool   `json:"ok"`                    // always present
	Description string `json:"description,omitempty"` // optional
}

/*
Telegram's API will always satisfy this model in any response.
Which is good to know for robust error handling.
*/
type API_Resp[R ResultContext] struct {
	BaseResp
	Result R
}

type ResultContext interface {
	~*types.CallbackQuery | ~*types.Chat | ~*types.Message | ~*types.Update | ~*types.User
	// todo: more
	// IsEmpy(min int) bool -> if n fields are empty retuns true, otherwise false
}
