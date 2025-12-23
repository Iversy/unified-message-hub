package models

func (s Source) String() string {
	return [...]string{"VK", "WhatsApp", "Telegram"}[s]
}
