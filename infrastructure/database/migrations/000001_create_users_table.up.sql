CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id uuid DEFAULT uuid_generate_v4() NOT NULL,
  username VARCHAR(20) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO users (username) VALUES ('hoge');
INSERT INTO users (username) VALUES ('fuga');
INSERT INTO users (username) VALUES ('foo');
