CREATE TABLE gamers (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(id),
    task TEXT NOT NULL,
    done BOOLEAN NOT NULL
);