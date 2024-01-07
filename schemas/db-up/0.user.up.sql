CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50),
    emailid VARCHAR(50),
    password TEXT
);
