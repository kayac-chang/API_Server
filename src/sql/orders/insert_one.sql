
-- Insert One
INSERT INTO orders (
    order_id, state, bet, game_id, user_id
) VALUES (
    :order_id, :state, :bet, :game_id, :user_id
)