DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "role_log";
DROP TRIGGER IF EXISTS "role_log_insert";
DROP TRIGGER IF EXISTS "role_log_update";
DROP TRIGGER IF EXISTS "role_log_delete";


CREATE TABLE "role" (
    "id"          INTEGER NOT NULL,
    "name"        TEXT    NOT NULL UNIQUE,
    "description" TEXT,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TABLE "role_log" (
    "id"               INTEGER NOT NULL,
    "effective_from"   TEXT    NOT NULL,
    "effective_to"     TEXT,
    "role_id"          INTEGER NOT NULL,
    "role_name"        TEXT    NOT NULL,
    "role_description" TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "role_log_insert" AFTER INSERT ON "role" FOR EACH ROW BEGIN
    INSERT
      INTO "role_log"
           ("effective_from", "role_id", "role_name", "role_description")
    VALUES (datetime(),       NEW.id,    NEW.name,    NEW.description);
END;


CREATE TRIGGER "role_log_update" AFTER UPDATE ON "role" FOR EACH ROW BEGIN
    UPDATE "role_log"
       SET "effective_to" = datetime()
     WHERE "role_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT
      INTO "role_log"
           ("effective_from", "role_id", "role_name", "role_description")
    VALUES (datetime(),       NEW.id,    NEW.name,    NEW.description);
END;


CREATE TRIGGER "role_log_delete" AFTER DELETE ON "role" FOR EACH ROW BEGIN
    UPDATE "role_log"
       SET "effective_to" = datetime()
     WHERE "role_id" = OLD.id
       AND "effective_to" IS NULL;
END;


INSERT
  INTO "role"
       ("name",   "description")
VALUES ("Admin",  "Administrator with full access"),
       ("Editor", "User with permissions to create and edit issues"),
       ("Viewer", "User with read-only access");
