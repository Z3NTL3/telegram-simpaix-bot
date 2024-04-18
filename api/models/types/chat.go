package types

type ChatType = string

const (
	Private    ChatType = "private"
	Group      ChatType = "group"
	SuperGroup ChatType = "supergroup"
	Channel    ChatType = "channel"
)

type Chat struct {
	ID      int      `json:"id"`
	Type    ChatType `json:"type"`
	Title   string   `json:"title"`
	IsForum bool     `json:"is_forum"`
	Link    string   `json:"invite_link"`
}
