
-- Insert One
INSERT INTO orders (
    order_id, 
    state, 
    bet, 
    game_id, 
    user_id,
    completed_at
) VALUES (
    :order_id, 
    :state, 
    :bet, 
    :game_id, 
    :user_id,
    :completed_at
)
