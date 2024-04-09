CREATE TABLE users (
    id serial PRIMARY KEY UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(100) NOT NULL
);
CREATE TABLE tokens
(
    id serial PRIMARY KEY UNIQUE ,
    refresh_token varchar NOT NULL ,
    user_id  serial UNIQUE,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES "users" (id)
);
