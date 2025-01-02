-- +goose Up
-- +goose StatementBegin
INSERT
  INTO project
       (title)
VALUES ('Best app ever'),
       ('School project'),
       ('Android game');
-- +goose StatementEnd
