-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS project (
    id    SERIAL PRIMARY KEY,
    title TEXT   NOT NULL UNIQUE
);

COMMENT ON TABLE  project       IS 'Project information';
COMMENT ON COLUMN project.id    IS 'Project id';
COMMENT ON COLUMN project.title IS 'Project title';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS project;
-- +goose StatementEnd
