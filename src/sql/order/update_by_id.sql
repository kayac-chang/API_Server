
UPDATE orders 
SET 
    state=:state, 
    bet=:bet,
    game_id=:game_id,
    user_id=:user_id,
    completed_at=:completed_at
WHERE 
    order_id = :order_id
