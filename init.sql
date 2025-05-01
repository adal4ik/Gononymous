CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    user_id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    avatar_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE posts(
    post_id UUID PRIMARY KEY,
    /*user_id UUID  REFERENCES users(user_id),*/
    created_at TIMESTAMP NOT NULL  DEFAULT NOW(),
    title VARCHAR(50) NOT NULL,
    subject VARCHAR(50),
    content TEXT,
    image_url TEXT
);

CREATE TABLE comments(
    comment_id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts(post_id),
    parent_id UUID REFERENCES comments(comment_id) ON DELETE CASCADE,
    user_id  UUID REFERENCES users(user_id),
    content TEXT,
    image_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);