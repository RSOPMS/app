-- +goose Up
-- +goose StatementBegin
INSERT
  INTO branch
       (name,            url)
VALUES ('main',          'https://github.com/example/repo/tree/main'),
       ('feature-login', 'https://github.com/example/repo/tree/feature-login'),
       ('bugfix-header', 'https://github.com/example/repo/tree/bugfix-header'),
       ('release-v1.0',  'https://github.com/example/repo/tree/release-v1.0'),
       ('hotfix-crash',  'https://github.com/example/repo/tree/hotfix-crash');
-- +goose StatementEnd
