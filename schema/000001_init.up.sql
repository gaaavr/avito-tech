CREATE TABLE IF NOT EXISTS users
(
    user_id serial PRIMARY KEY,
    balance real CHECK (balance >= 0)
    );