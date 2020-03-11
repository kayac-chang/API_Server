
-- DROPS ALL
DROP TABLE IF EXISTS games CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS orders CASCADE;

-- Tables

-- Games
CREATE TABLE IF NOT EXISTS games (
    -- pk
    game_id CHAR(32) PRIMARY KEY,

    -- properties
    name VARCHAR(256) NOT NULL UNIQUE,
    href VARCHAR NOT NULL UNIQUE,
    category VARCHAR(32) NOT NULL,

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Users
CREATE TABLE IF NOT EXISTS users (
    -- pk
    user_id CHAR(32) PRIMARY KEY,

    -- properties
    username VARCHAR(256) NOT NULL UNIQUE,

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Orders
CREATE TABLE IF NOT EXISTS orders (
    -- pk
    order_id CHAR(32) PRIMARY KEY,

    -- properties
    state CHAR(1) NOT NULL DEFAULT 'P',
    bet NUMERIC NOT NULL DEFAULT 0,
    
    -- fk
    game_id CHAR(32) NOT NULL REFERENCES games,
    user_id CHAR(32) NOT NULL REFERENCES users,

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

-- Triggers
CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 
$$
BEGIN
    NEW.updated_at = NOW();

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Users
DROP TRIGGER IF EXISTS set_timestamp ON users;

CREATE TRIGGER set_timestamp 
    BEFORE UPDATE ON users 
    FOR EACH ROW 
    EXECUTE PROCEDURE trigger_set_timestamp();

-- Orders
DROP TRIGGER IF EXISTS set_timestamp ON orders;

CREATE TRIGGER set_timestamp 
    BEFORE UPDATE ON orders
    FOR EACH ROW 
    EXECUTE PROCEDURE trigger_set_timestamp();