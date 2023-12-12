DO $$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_database WHERE datname = 'database_test'
    )
    THEN
        CREATE DATABASE "database_test";
    END IF;
END
$$;
