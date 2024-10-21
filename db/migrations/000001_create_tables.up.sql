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