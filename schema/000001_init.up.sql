CREATE TABLE IF NOT EXISTS users
(
    user_id serial PRIMARY KEY,
    balance numeric(100,2) CHECK (balance >= 0)
);
CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id serial PRIMARY KEY,
    user_id  integer,
    amount numeric(100,2) NOT NULL,
    date_time timestamp NOT NULL,
    message varchar(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);
CREATE TABLE IF NOT EXISTS orders
(
    order_id serial PRIMARY KEY,
    user_id  integer,
    service_id integer,
    amount numeric(100,2) NOT NULL,
    date_time timestamp NOT NULL,
    block boolean DEFAULT true,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);