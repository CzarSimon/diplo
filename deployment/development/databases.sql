CREATE DATABASE directory;
CREATE ROLE directory WITH LOGIN PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE directory TO directory;

CREATE DATABASE chat;
CREATE ROLE chat WITH LOGIN PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE chat TO chat;

CREATE DATABASE game;
CREATE ROLE game WITH LOGIN  PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE game TO game;
