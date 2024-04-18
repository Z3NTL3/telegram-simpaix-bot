package types

type Message struct {
	ID          int    `json:"update_id"`
	MsgThreadID int    `json:"message_thread_id"`
	From        *User  `json:"from"`
	Date        int    `json:"date"`
	Chat        *Chat  `json:"chat"`
	Text        string `json:"text"`
}
