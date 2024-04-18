package types

type Update struct {
	ID         int            `json:"update_id"`
	Message    *Message       `json:"message"`
	CallBQuery *CallbackQuery `json:"callback_query"`
}
