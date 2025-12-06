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
	// ID         int	//?optional
	Client     Source `json:"source"`
	Sender     string `json:"sender"`
	ChatId     int    `json:"chat_id"`
	Text       string `json:"text"`
	Attachment string `json:"attachment"`
	Timestamp  string `json:"timestamp"` // Timestamp

}

func NewMessage() *Message {
	return &Message{}
}
