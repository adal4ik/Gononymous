CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users{
    user_id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    avatar_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL 
};

CREATE TABLE posts{
    post_id UUID PRIMARY KEY,
    user_id UUID FOREIGN KEY,
    created_at TIMESTAMP NOT NULL,
    title VARCHAR(50) NOT NULL,
    content TEXT,
    image_url TEXT
};

CREATE TABLE comments{
    comment_id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts(post_id),
    parent_id UUID REFERENCES comments(comment_id) ON DELETE CASCADE,
    user_id  UUID REFERENCES users(user_id),
    content TEXT,
    image_url TEXT,
    created_at TIMESTAMP NOT NULL
};