CREATE TABLE IF NOT EXISTS users
(
    user_id serial PRIMARY KEY,
    balance real CHECK (balance >= 0)
);
CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id serial PRIMARY KEY,
    user_id  integer
    amount real NOT NULL,
    date_time timestamp NOT NULL,
    message varchar(200) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id) NOT NULL
);