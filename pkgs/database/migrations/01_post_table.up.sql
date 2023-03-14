CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE posts (
  post_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  post_name TEXT NOT NULL,
  post_like BOOLEAN DEFAULT false,
  post_star BOOLEAN DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);