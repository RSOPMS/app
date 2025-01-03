INSERT
  INTO bugbase_user
       (email,             password_hash,                                                  name,       surname,      role_id)
VALUES ('a.a@bugbase.com', '$2a$10$9aXsOicMNDly4rzkfDXZgeIa6irsxPtILZoTV5us2uqUnHEKSRjau', 'anze',     'arhar',      1),
       ('k.k@bugbase.com', '$2a$10$9aXsOicMNDly4rzkfDXZgeIa6irsxPtILZoTV5us2uqUnHEKSRjau', 'kristjan', 'kostanjsek', 2),
       ('n.l@bugbase.com', '$2a$10$9aXsOicMNDly4rzkfDXZgeIa6irsxPtILZoTV5us2uqUnHEKSRjau', 'nejc',     'locicnik',   3);

-- NOTE: all users have a default password: "password"
