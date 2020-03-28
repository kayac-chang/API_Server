
-- Update One
UPDATE games
SET name = :name,
    href = :href,
    category = :category
WHERE
    game_id = :game_id;
