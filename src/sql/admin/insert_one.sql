
-- Insert One
INSERT INTO admins (
    admin_id, 
    email, 
    username, 
    password, 
    organization, 
    created_at
) VALUES (
    :admin_id,
    :email,
    :username,
    :password,
    :organization,
    :created_at
)
