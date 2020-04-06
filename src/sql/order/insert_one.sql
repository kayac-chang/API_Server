
-- Insert One
INSERT INTO orders (
    order_id, 
    state, 
    bet, 
    win, 
    game_id, 
    user_id,
    created_at,
    completed_at
) VALUES (
    :order_id, 
    :state, 
    :bet, 
    :win, 
    :game_id, 
    :user_id,
    :created_at,
    :completed_at
)
