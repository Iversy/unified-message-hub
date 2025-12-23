CREATE TABLE platform_connections (
    id BIGSERIAL PRIMARY KEY,
    platform VARCHAR(20) NOT NULL,           
    connection_name VARCHAR(255) NOT NULL,   
    credentials JSONB NOT NULL,              
    is_active BOOLEAN DEFAULT true,
    last_sync_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

SELECT create_distributed_table(
    'platform_connections',
     'id'
    );