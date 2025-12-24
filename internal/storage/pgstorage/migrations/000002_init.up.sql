CREATE TABLE routing_rules (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    source_chat_id BIGINT NOT NULL,           
    receiver_id BIGINT NOT NULL, 
    keywords JSONB DEFAULT '[]',
    is_active BOOLEAN DEFAULT true
);

SELECT create_distributed_table(
    'routing_rules',
     'id'
    );