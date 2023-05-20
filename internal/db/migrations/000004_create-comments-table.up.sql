CREATE TABLE IF NOT EXISTS comments (
    id text PRIMARY KEY,
    post_id text NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    user_id text NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    body text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS comment_images (
    id text PRIMARY KEY,
    comment_id text NOT NULL REFERENCES comments (id),
    url text NOT NULL
);

CREATE INDEX ON comment_images (comment_id);