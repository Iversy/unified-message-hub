DO $$
DECLARE
i INTEGER;
BEGIN
    FOR i IN REVERSE 512..1 LOOP
        EXECUTE format('DROP TABLE IF EXISTS schema_%s.routing_rules CASCADE', LPAD(i::text, 3, '0'));
    END LOOP;
END $$;
