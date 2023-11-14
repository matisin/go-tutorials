CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE users
(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name            VARCHAR(100) NOT NULL,
    lastname       VARCHAR(100) NOT NULL,
    mail            VARCHAR(100) UNIQUE NOT NULL,
    state           CHAR(3) DEFAULT 'PEN',
    phone           VARCHAR(16) UNIQUE NOT NULL,
    identification  VARCHAR(16) UNIQUE NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
