-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bugbase_user (
    id            SERIAL PRIMARY KEY,
    email         TEXT   NOT NULL UNIQUE,
    password_hash TEXT   NOT NULL,
    name          TEXT   NOT NULL,
    surname       TEXT   NOT NULL,
    avatar        TEXT,
    role_id       SERIAL REFERENCES role(id)
);


COMMENT ON TABLE  bugbase_user               IS 'User information';
COMMENT ON COLUMN bugbase_user.id            IS 'User id';
COMMENT ON COLUMN bugbase_user.email         IS 'User email address';
COMMENT ON COLUMN bugbase_user.password_hash IS 'User password hash';
COMMENT ON COLUMN bugbase_user.name          IS 'User name';
COMMENT ON COLUMN bugbase_user.surname       IS 'User surname';
COMMENT ON COLUMN bugbase_user.avatar        IS 'User avatar url';
COMMENT ON COLUMN bugbase_user.role_id       IS 'User role';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bugbase_user;
-- +goose StatementEnd
