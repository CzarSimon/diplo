-- Schema changes

CREATE TABLE channel(
  id VARCHAR(64) PRIMARY KEY,
  name VARCHAR(100),
  created_at TIMESTAMP,
  game_id VARCHAR(64)
);

CREATE TABLE chat_user(
  id VARCHAR(64) PRIMARY KEY
);

CREATE TABLE channel_members(
  user_id VARCHAR(64) REFERENCES chat_user(id),
  channel_id VARCHAR(64) REFERENCES channel(id),
  PRIMARY KEY (user_id, channel_id)
);

CREATE TABLE message(
  id VARCHAR(64) PRIMARY KEY,
  message_text TEXT,
  channel_id VARCHAR(64) REFERENCES channel(id),
  author VARCHAR(64) REFERENCES chat_user(id),
  created_at TIMESTAMP
);

-- Data changes

INSERT INTO channel(id, name, created_at, game_id)
VALUES
  ('channel-1', 'Inital channel', NOW()::TIMESTAMP, 'game-1');

INSERT INTO chat_user(id)
VALUES
  ('user-1'),
  ('user-2');

INSERT INTO channel_members(user_id, channel_id)
VALUES
  ('user-1', 'channel-1'),
  ('user-2', 'channel-1');
