DO $$
DECLARE
    i INTEGER;
BEGIN
    FOR i IN 1..512 LOOP
            EXECUTE format('
                CREATE TABLE IF NOT EXISTS schema_%s.routing_rules (
                    id BIGSERIAL PRIMARY KEY,
                    name VARCHAR(255) NOT NULL,
                    source_chat_id BIGINT NOT NULL,
                    receiver_id BIGINT NOT NULL,
                    keywords JSONB DEFAULT ''[]'',
                    is_active BOOLEAN DEFAULT true
                )', LPAD(i::text, 3, '0'));
    END LOOP;
END $$;