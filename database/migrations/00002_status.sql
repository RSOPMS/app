-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS status (
    id            SERIAL  PRIMARY KEY,
    name          TEXT    NOT NULL UNIQUE,
    display_order INTEGER NOT NULL UNIQUE
);


COMMENT ON TABLE  status               IS 'Status information';
COMMENT ON COLUMN status.id            IS 'Status id';
COMMENT ON COLUMN status.name          IS 'Status name';
COMMENT ON COLUMN status.display_order IS 'Status display order';


INSERT
  INTO status
       (name,          display_order)
VALUES ('Open',        0),
       ('Planning',    1),
       ('In progress', 2),
       ('Testing',     3),
       ('Closed',      4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS status;
-- +goose StatementEnd
