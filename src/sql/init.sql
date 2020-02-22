
-- Function --
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
        BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;


-- Games Start --

-- Table --
CREATE TABLE IF NOT EXISTS games (
    id       SERIAL NOT NULL PRIMARY KEY,
	name     TEXT,
	href     TEXT,
	create_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Trigger --
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON games
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- Games End --