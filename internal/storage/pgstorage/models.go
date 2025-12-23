package pgstorage

type MessageAudit struct {
	SourcePlatform uint64 `db:"source_platform"`
	SourceChatID   uint64 `db:"source_chat_id"`
	SenderID       string `db:"sender"`
	MessageText    string `db:"message_text"`
	Received_at    string `db:"received_at"`
}
