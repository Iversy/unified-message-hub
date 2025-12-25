DO $$
DECLARE
    i INTEGER;
BEGIN
    FOR i IN 1..512 LOOP
            EXECUTE format('DROP TABLE IF EXISTS schema_%s.routing_rules CASCADE', LPAD(i::text, 3, '0'));
    END LOOP;
END $$;