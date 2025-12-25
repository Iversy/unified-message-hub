DO $$
DECLARE
    i INTEGER;
BEGIN
    FOR i IN 1..512 LOOP
            EXECUTE format('
                CREATE TABLE IF NOT EXISTS schema_%s.message_audit (
                    id BIGSERIAL NOT NULL,
                    source_platform BIGINT NOT NULL,
                    source_chat_id BIGINT NOT NULL,
                    sender VARCHAR(50) NOT NULL,
                    message_text TEXT,
                    received_at TIMESTAMPTZ NOT NULL,
                    PRIMARY KEY (id)
                )', LPAD(i::text, 3, '0'));
    END LOOP;
END $$;