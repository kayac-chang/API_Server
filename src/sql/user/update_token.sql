
UPDATE users
SET 
    access_token = :access_token
WHERE
    user_id = :user_id