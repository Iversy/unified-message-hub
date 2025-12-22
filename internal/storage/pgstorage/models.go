package pgstorage

type MessageAudit struct {
	SourcePlatform uint64 `db:"source_platform"`
	SourceChatID   uint64 `db:"source_chat_id"`
	SenderID       string `db:"sender_id"`
	MessageText    string `db:"message_text"`
	MessageType    string `db:"message_type"`
	Received_at    string `db:"received_at"`
}
