CREATE TABLE IF NOT EXISTS priority (
    id      SERIAL  PRIMARY KEY,
    name    TEXT    NOT NULL UNIQUE,
    "order" INTEGER NOT NULL UNIQUE
);


COMMENT ON TABLE  priority       IS 'Priority information';
COMMENT ON COLUMN priority.id    IS 'Priority id';
COMMENT ON COLUMN priority.name  IS 'Priority name';
COMMENT ON COLUMN priority.order IS 'Priority display order';


INSERT
  INTO priority
       (name,       "order")
VALUES ('Critical', 0),
       ('High',     1),
       ('Medium',   2),
       ('Low',      3),
       ('Backlog',  4);
