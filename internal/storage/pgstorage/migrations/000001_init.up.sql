CREATE TABLE message_audit (
    id BIGINT PRIMARY KEY,
    source_platform BIGINT NOT NULL,  
    source_chat_id BIGINT NOT NULL,
    sender VARCHAR(50) NOT NULL,
    message_text TEXT,
    received_at TIMESTAMPTZ NOT NULL
);