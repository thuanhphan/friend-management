CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE connections (
    id SERIAL PRIMARY KEY,
    requestor_id INT NOT NULL REFERENCES users(id),
    target_id INT NOT NULL REFERENCES users(id),
    status VARCHAR(50) NOT NULL
);