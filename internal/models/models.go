package models

type Source int

const (
	VK Source = iota
	WhatsApp
	Telegram
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
	Client    Source `json:"source"`
	Sender    string `json:"sender"`
	ChatId    int    `json:"chat_id"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}
