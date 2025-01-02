-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS issue (
    id          SERIAL    PRIMARY KEY,
    title       TEXT      NOT NULL,
    description TEXT,
    project_id  SERIAL    REFERENCES project(id),
    status_id   SERIAL    REFERENCES status(id),
    priority_id SERIAL    REFERENCES priority(id),
    branch_id   SERIAL    REFERENCES branch(id),
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (title, project_id, branch_id)
);


COMMENT ON TABLE  issue             IS 'Issue information';
COMMENT ON COLUMN issue.id          IS 'Issue id';
COMMENT ON COLUMN issue.title       IS 'Issue title';
COMMENT ON COLUMN issue.description IS 'Issue description';
COMMENT ON COLUMN issue.project_id  IS 'Project to which this issue belongs to';
COMMENT ON COLUMN issue.status_id   IS 'Issue status';
COMMENT ON COLUMN issue.priority_id IS 'Issue priority';
COMMENT ON COLUMN issue.branch_id   IS 'Issue branch';
COMMENT ON COLUMN issue.created_at  IS 'Issue creation time';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS issue;
-- +goose StatementEnd
