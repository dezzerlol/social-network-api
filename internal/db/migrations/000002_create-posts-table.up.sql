CREATE TABLE IF NOT EXISTS posts (
    id text PRIMARY KEY,
    userId text NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    body text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS post_images (
    id text PRIMARY KEY,
    post_id text NOT NULL REFERENCES posts (id),
    url text NOT NULL
);

CREATE INDEX ON post_images (post_id);