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

type Route struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	SourceChatID int      `json:"source_chat_id"`
	ReceiverID   int      `json:"receiver_id"`
	Keywords     []string `json:"keywords"`
	IsActive     bool     `json:"is_active"`
}

func NewMessage() *Message {
	return &Message{}
}

func NewRoute() *Route {
	return &Route{}
}
