CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    nickname TEXT NOT NULL UNIQUE,
    avatar TEXT NOT NULL DEFAULT '',
    hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS announcements (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_announcements_user_id ON announcements(user_id);
