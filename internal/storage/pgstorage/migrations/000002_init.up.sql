CREATE TABLE routing_rules (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    source_chat_id BIGINT NOT NULL,           
    destination_chat_id VARCHAR(100) NOT NULL, 
    keywords JSONB DEFAULT '[]',             
    senders JSONB DEFAULT '[]',               
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

SELECT create_distributed_table(
    'routing_rules',
     'id'
    );