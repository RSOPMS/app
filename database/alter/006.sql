CREATE TABLE IF NOT EXISTS "user" (
    id            SERIAL PRIMARY KEY,
    email         TEXT   NOT NULL UNIQUE,
    password_hash TEXT   NOT NULL,
    name          TEXT   NOT NULL,
    surname       TEXT   NOT NULL,
    avatar        TEXT,
    role_id       SERIAL REFERENCES role(id)
);


COMMENT ON TABLE  "user"               IS 'User information';
COMMENT ON COLUMN "user".id            IS 'User id';
COMMENT ON COLUMN "user".email         IS 'User email address';
COMMENT ON COLUMN "user".password_hash IS 'User password hash';
COMMENT ON COLUMN "user".name          IS 'User name';
COMMENT ON COLUMN "user".surname       IS 'User surname';
COMMENT ON COLUMN "user".avatar        IS 'User avatar url';
COMMENT ON COLUMN "user".role_id       IS 'User role';
