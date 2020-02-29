


CREATE TABLE orders (

    -- pk
    order_id VARCHAR(32) PRIMARY KEY,

    -- fk
    game_id,
    account_id,

    -- properties
    state VARCHAR NOT NULL,
    total_bet VARCHAR NOT NULL,
    total_win VARCHAR NOT NULL,

    -- times
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    complete_at TIMESTAMPTZ,
);