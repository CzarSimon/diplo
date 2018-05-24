-- Schema changes

CREATE TABLE channel(
  id VARCHAR(64) PRIMARY KEY,
  name VARCHAR(100),
  created_at DATETIME DEFAULT,
  game_id VARCHAR(64)
)

CREATE TABLE chat_user(
  id VARCHAR(64)
)

CREATE TABLE channel_members(
  user_id REFERENCES chat_user(id),
  channel_id REFERENCES channel(id),
  PRIMARY KEY (user_id, channel_id)
);

CREATE TABLE message(
  id VARCHAR(64) PRIMARY KEY,
  message_text TEXT,
  channel_id VARCHAR(64) REFERENCES channel(id),
  author VARCHAR(64) REFERENCES chat_user(id),
  created_at DATETIME
)

-- Data changes
