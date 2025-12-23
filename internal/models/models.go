package models

type Source int

const (
	VK Source = iota
	WhatsApp
	Telegram
	Max
)

type Timestamp struct {
	Second int
	Minute int
	Hour   int
	Day    int
	Month  int
	Year   int
}

type Message struct {
	// ID        int    `json:"id"` //?optional
	Client    Source `json:"source"`
	Sender    string `json:"sender"`
	ChatId    int    `json:"chat_id"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}
