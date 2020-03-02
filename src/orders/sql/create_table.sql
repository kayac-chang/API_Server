
-- Create
CREATE TABLE IF NOT EXISTS orders (
    -- pk
    order_id CHAR(64) PRIMARY KEY,

    -- properties
    state CHAR(1) NOT NULL,
    bet NUMERIC NOT NULL,

    -- fk
    game_id SMALLINT NOT NULL REFERENCES games,
    user_id CHAR(32) NOT NULL REFERENCES users,

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

-- Drop
DROP TABLE IF EXISTS orders;

-- Trigger --
CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 
$$
BEGIN
    NEW.updated_at = NOW();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp 
    BEFORE UPDATE ON orders 
    FOR EACH ROW 
    EXECUTE PROCEDURE trigger_set_timestamp();