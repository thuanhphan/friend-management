<<<<<<< HEAD
<<<<<<< HEAD
=======
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

>>>>>>> d88719a (Arrange the layered architecture)
=======
>>>>>>> 324c60e (FM-4)
CREATE TABLE users (
    email VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE friendships (
    id SERIAL PRIMARY KEY,
<<<<<<< HEAD
<<<<<<< HEAD
    user_email VARCHAR(255) REFERENCES users(email) NOT NULL,
    friend_email VARCHAR(255) REFERENCES users(email) NOT NULL,
    status VARCHAR(50) NOT NULL -- 'friends', 'subscribed', 'blocked'
=======
    user_email VARCHAR(255) REFERENCES users(email),
    friend_email VARCHAR(255) REFERENCES users(email),
    status VARCHAR(50) -- 'friends', 'subscribed', 'blocked'
>>>>>>> d88719a (Arrange the layered architecture)
=======
    user_email VARCHAR(255) REFERENCES users(email) NOT NULL,
    friend_email VARCHAR(255) REFERENCES users(email) NOT NULL,
    status VARCHAR(50) NOT NULL -- 'friends', 'subscribed', 'blocked'
>>>>>>> a78f4b5 (update script create table)
);