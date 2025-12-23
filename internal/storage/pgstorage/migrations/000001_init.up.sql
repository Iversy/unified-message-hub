
CREATE TABLE message_audit (
    id BIGSERIAL NOT NULL,
    source_platform BIGINT NOT NULL,  
    source_chat_id BIGINT NOT NULL,
    sender VARCHAR(50) NOT NULL,
    message_text TEXT,
    received_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id)  
);

SELECT create_distributed_table(
    'message_audit',
     'id'
    );