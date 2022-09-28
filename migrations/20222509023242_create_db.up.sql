CREATE TABLE IF NOT EXISTS users(
    id TEXT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    nickname TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    country TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
