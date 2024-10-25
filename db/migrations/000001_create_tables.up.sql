CREATE TABLE users (
    email VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE friendships (
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(255) REFERENCES users(email) NOT NULL,
    friend_email VARCHAR(255) REFERENCES users(email) NOT NULL,
    status VARCHAR(50) NOT NULL -- 'friends', 'subscribed', 'blocked'
);