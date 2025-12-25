DO $$
DECLARE
    i INTEGER;
BEGIN
    FOR i IN 1..512 LOOP
            EXECUTE format('CREATE SCHEMA IF NOT EXISTS schema_%s', LPAD(i::text, 3, '0'));
    END LOOP;
END $$;