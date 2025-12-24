package pgstorage

type MessageAudit struct {
	SourcePlatform uint64 `db:"source_platform"`
	SourceChatID   uint64 `db:"source_chat_id"`
	SenderID       string `db:"sender"`
	MessageText    string `db:"message_text"`
	Received_at    string `db:"received_at"`
}

type RoutingRules struct {
	ID           int      `db:"id"`
	Name         string   `db:"name"`
	SourceChatID int      `db:"source_chat_id"`
	ReceiverID   int      `db:"receiver_id"`
	Keywords     []string `db:"keywords"`
	IsActive     bool     `db:"is_active"`
}
