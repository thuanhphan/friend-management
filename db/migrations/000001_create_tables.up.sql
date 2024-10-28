-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     email VARCHAR(255) NOT NULL UNIQUE
-- );

-- CREATE TABLE connections (
--     id SERIAL PRIMARY KEY,
--     requestor_id INT NOT NULL REFERENCES users(id),
--     target_id INT NOT NULL REFERENCES users(id),
--     status VARCHAR(50) NOT NULL
-- );

CREATE TABLE users (
    email VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE friendships (
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(255) REFERENCES users(email),
    friend_email VARCHAR(255) REFERENCES users(email),
    status VARCHAR(50) -- 'friends', 'subscribed', 'blocked'
);