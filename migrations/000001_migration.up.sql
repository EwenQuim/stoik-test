-----------------------------------------------------------------------
---- Users
-----------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY,
  name text NOT NULL,
  password text,
  email text NOT NULL
);
CREATE INDEX IF NOT EXISTS users_name ON users (name);
-----------------------------------------------------------------------
---- Cars
-----------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS cars (
  id uuid PRIMARY KEY,
  name text NOT NULL,
  price int,
  year int,
  color text
);
