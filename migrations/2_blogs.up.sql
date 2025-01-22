CREATE TABLE IF NOT EXISTS blogs (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR,
    author_id BIGINT REFERENCES authors(id)
)