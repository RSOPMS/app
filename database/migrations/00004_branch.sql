-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS branch (
    id   SERIAL PRIMARY KEY,
    name TEXT   NOT NULL,
    url  TEXT   NOT NULL UNIQUE
);


COMMENT ON TABLE  branch      IS 'Branch information';
COMMENT ON COLUMN branch.id   IS 'Branch id';
COMMENT ON COLUMN branch.name IS 'Branch name';
COMMENT ON COLUMN branch.url  IS 'Branch url';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS branch;
-- +goose StatementEnd
