CREATE TABLE Users (
    id serial PRIMARY KEY NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(100) NOT NULL
);