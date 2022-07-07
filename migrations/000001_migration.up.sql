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
-----------------------------------------------------------------------
---- Rentals
-----------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS rentals (
  user_id uuid NOT NULL,
  car_id uuid NOT NULL,
  start_date date NOT NULL,
  end_date date NOT NULL,
  price int NOT NULL,
  paid boolean NOT NULL,
  PRIMARY KEY (user_id, car_id)
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  FOREIGN KEY (car_id) REFERENCES cars (id) ON DELETE CASCADE
);