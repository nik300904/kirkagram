ALTER TABLE "user" RENAME TO users;

ALTER TABLE "users"
    DROP COLUMN IF EXISTS followers,
    DROP COLUMN IF EXISTS following;

CREATE TABLE IF NOT EXISTS "follow" (
    id SERIAL PRIMARY KEY,
    follower_id INTEGER NOT NULL,
    following_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (following_id) REFERENCES "user"(id) ON DELETE CASCADE,
    UNIQUE (follower_id, following_id)
);

CREATE INDEX IF NOT EXISTS follow_follower_idx ON follow(follower_id);
CREATE INDEX IF NOT EXISTS follow_following_idx ON follow(following_id);