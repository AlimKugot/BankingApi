CREATE TABLE user_account
(
    id      SERIAL PRIMARY KEY,
    name    varchar(250) NOT NULL,
    balance numeric DEFAULT 0.00
);
