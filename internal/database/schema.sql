CREATE TABLE IF NOT EXISTS users
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    name     TEXT NOT NULL,
    email    TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    bio      TEXT
);

CREATE TABLE IF NOT EXISTS tasks
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    title    TEXT NOT NULL,
    user_id  INTEGER NOT NULL,
    due_date DATE,
    status      TEXT NOT NULL CHECK(status IN ('open', 'in progress', 'closed')) DEFAULT 'open',
    description TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);