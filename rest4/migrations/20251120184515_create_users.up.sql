CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    email varchar NOT NULL UNIQUE,
    encrypted_password varchar NOT NULL
);