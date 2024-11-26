CREATE TABLE IF NOT EXISTS role (
    id          SERIAL PRIMARY KEY,
    name        TEXT   NOT NULL UNIQUE,
    description TEXT
);


COMMENT ON TABLE  role             IS 'User role information';
COMMENT ON COLUMN role.id          IS 'User role id';
COMMENT ON COLUMN role.name        IS 'User role name';
COMMENT ON COLUMN role.description IS 'User role description';


INSERT
  INTO role
       (name,     description)
VALUES ('Admin',  'Administrator with full access'),
       ('Editor', 'User with permissions to create and edit issues'),
       ('Viewer', 'User with read-only access');
