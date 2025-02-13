-- +goose Up
CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title VARCHAR NOT NULL,
  url VARCHAR UNIQUE NOT NULL,
  description VARCHAR NULL,
  published_at TIMESTAMP NOT NULL,
  feed_id UUID references feeds(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE posts;