CREATE TABLE IF NOT EXISTS project (
    id    SERIAL PRIMARY KEY,
    title TEXT   NOT NULL UNIQUE,
    CHECK (title <> '' AND title ~* '[^\s]')
);

COMMENT ON TABLE  project       IS 'Project information';
COMMENT ON COLUMN project.id    IS 'Project id';
COMMENT ON COLUMN project.title IS 'Project title';
