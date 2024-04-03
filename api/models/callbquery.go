package models

type CallbackQuery struct {
	ID      int      `json:"id"`
	From    *User    `json:"user"`
	Message *Message `json:"message"`
	Data    int      `json:"data"`
}
