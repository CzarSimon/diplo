-- Schema changes
CREATE TABLE account(
  id VARCHAR(64) PRIMARY KEY,
  email VARCHAR(100) UNIQUE,
  username VARCHAR(50),
  password VARCHAR(256),
  given_name VARCHAR(100),
  surname VARCHAR(100),
  joined_at TIMESTAMP,
  ranking NUMERIC(8,4)
);

-- Data changes
