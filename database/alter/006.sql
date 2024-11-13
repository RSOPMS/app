DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "user_log";
DROP TRIGGER IF EXISTS "user_log_insert";
DROP TRIGGER IF EXISTS "user_log_update";
DROP TRIGGER IF EXISTS "user_log_delete";


CREATE TABLE "user" (
    "id"            INTEGER NOT NULL,
    "email"         TEXT    NOT NULL UNIQUE,
    "password_hash" TEXT    NOT NULL,
    "name"          TEXT    NOT NULL,
    "surname"       TEXT    NOT NULL,
    "avatar"        TEXT,
    "role_id"       INTEGER NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT),
    FOREIGN KEY("role_id") REFERENCES role("id")
) STRICT;


CREATE TABLE "user_log" (
    "id"                 INTEGER NOT NULL,
    "effective_from"     TEXT    NOT NULL,
    "effective_to"       TEXT,
    "user_id"            INTEGER NOT NULL,
    "user_email"         TEXT    NOT NULL UNIQUE,
    "user_password_hash" TEXT    NOT NULL,
    "user_name"          TEXT    NOT NULL,
    "user_surname"       TEXT    NOT NULL,
    "user_avatar"        TEXT,
    "user_role"          TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "user_log_insert" AFTER INSERT ON "user" FOR EACH ROW BEGIN
    INSERT INTO "user_log"
           ("effective_from",
            "user_id",
            "user_email",
            "user_password_hash",
            "user_name",
            "user_surname",
            "user_avatar",
            "user_role")
    VALUES (datetime(),
            NEW.id,
            NEW.email,
            NEW.password_hash,
            NEW.name,
            NEW.surname,
            NEW.avatar,
            NEW.role_id);
END;


CREATE TRIGGER "user_log_update" AFTER UPDATE ON "user" FOR EACH ROW BEGIN
    UPDATE "user_log"
       SET "effective_to" = datetime()
     WHERE "user_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT INTO "user_log"
           ("effective_from",
            "user_id",
            "user_email",
            "user_password_hash",
            "user_name",
            "user_surname",
            "user_avatar",
            "user_role")
    VALUES (datetime(),
            NEW.id,
            NEW.email,
            NEW.password_hash,
            NEW.name,
            NEW.surname,
            NEW.avatar,
            NEW.role_id);
END;


CREATE TRIGGER "user_log_delete" AFTER DELETE ON "user" FOR EACH ROW BEGIN
    UPDATE "user_log"
       SET "effective_to" = datetime()
     WHERE "user_id" = OLD.id
       AND "effective_to" IS NULL;
END;
