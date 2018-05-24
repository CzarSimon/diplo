-- Schema changes

CREATE TABLE province_type(
  name VARCHAR(20) PRIMARY KEY
);

CREATE TABLE country(
  name VARCHAR(20) PRIMARY KEY
);

CREATE TABLE province(
  short_name VARCHAR(3) PRIMARY KEY,
  name VARCHAR(100),
  has_supply_center BOOLEAN,
  type VARCHAR(20) REFERENCES province_type(name)
);

CREATE TABLE round_time_unit(
  name VARCHAR(1) PRIMARY KEY
);

CREATE TABLE game(
  id INTEGER SERIAL PRIMARY KEY,
  round_time INTEGER,
  round_time_unit VARCHAR(1) REFERENCES round_time_unit(name),
  has_started BOOLEAN DEFAULT FALSE,
  started_at DATETIME,
  has_ended BOOLEAN DEFAULT FALSE,
  ended_at DATETIME,
  message_visibilty FLOAT
)

CREATE TABLE user_game(
  user_id VARCHAR(256)
  game_id INTEGER REFERENCES game(id),
  country VARCHAR(20) REFERENCES country(name),
  PRIMARY KEY (user_id, game_id)
)

CREATE TABLE phase(
  name VARCHAR(20) PRIMARY KEY,
  order INTEGER UNIQUE
)

CREATE TABLE round(
  id INTEGER SERIAL PRIMARY KEY,
  year INTEGER,
  phase VARCHAR(20),
  game_id INTEGER REFERENCES game(id)
)

-- Data changes

INSERT INTO province_type(name)
VALUES
  ('Land'),
  ('Costal'),
  ('Water');

INSERT INTO country(name)
VALUES
  ('Austria'),
  ('England'),
  ('France'),
  ('Germany'),
  ('Italy'),
  ('Russia'),
  ('Turkey');

INSERT INTO province(short_name, name, has_supply_center, type)
VALUES
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, ''),
  ('', '', TRUE, '');

INSERT INTO round_time_unit(name)
VALUES
  ('minute'),
  ('hour'),
  ('day');

INSERT INTO phase(name, order)
VALUES
  ('Spring', 1),
  ('Fall', 2),
  ('Build', 3);
