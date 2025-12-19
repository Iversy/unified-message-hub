CREATE TABLE message_audit (
    id BIGSERIAL PRIMARY KEY,
    event_id UUID NOT NULL UNIQUE,          
    source_platform VARCHAR(20) NOT NULL,  
    source_chat_id BIGINT NOT NULL,
    sender_id BIGINT NOT NULL,
    message_text TEXT,
    message_type VARCHAR(50) NOT NULL,      
    received_at TIMESTAMPTZ NOT NULL,
    
    destination_platform VARCHAR(20),        
    destination_chat_id VARCHAR(100),
    delivery_status VARCHAR(20),             
    delivered_at TIMESTAMPTZ,
    error_message TEXT
);