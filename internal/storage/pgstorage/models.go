package pgstorage

type MessageAudit struct {
	ID             uint64 `db:"id"`
	SourcePlatform uint64 `db:"source_platform"`
	SourceChatID   uint64 `db:"source_chat_id"`
	SenderID       string `db:"sender"`
	MessageText    string `db:"message_text"`
	Received_at    string `db:"received_at"`
}
