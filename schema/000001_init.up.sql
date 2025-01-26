CREATE TABLE users (
                       id serial NOT NULL UNIQUE PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE posts (
                       id serial NOT NULL UNIQUE PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT,
                       user_id INT REFERENCES users (id) ON DELETE cascade NOT NULL
);

CREATE TABLE likes (
                       id serial NOT NULL UNIQUE PRIMARY KEY,
                       user_id INT REFERENCES users (id) ON DELETE cascade NOT NULL,
                       post_id INT REFERENCES posts (id) ON DELETE cascade NOT NULL
);