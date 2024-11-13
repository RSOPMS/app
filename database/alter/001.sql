DROP TABLE IF EXISTS "project";
DROP TABLE IF EXISTS "project_log";
DROP TRIGGER IF EXISTS "project_log_insert";
DROP TRIGGER IF EXISTS "project_log_update";
DROP TRIGGER IF EXISTS "project_log_delete";


CREATE TABLE "project" (
    "id"    INTEGER NOT NULL,
    "title" TEXT    NOT NULL UNIQUE,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TABLE "project_log" (
    "id"             INTEGER NOT NULL,
    "effective_from" TEXT    NOT NULL,
    "effective_to"   TEXT,
    "project_id"     INTEGER NOT NULL,
    "project_title"  TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "project_log_insert" AFTER INSERT ON "project" FOR EACH ROW BEGIN
    INSERT
      INTO "project_log"
           ("effective_from", "project_id", "project_title")
    VALUES (datetime(),       NEW.id,       NEW.title);
END;


CREATE TRIGGER "project_log_update" AFTER UPDATE ON "project" FOR EACH ROW BEGIN
    UPDATE "project_log"
       SET "effective_to" = datetime()
     WHERE "project_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT
      INTO "project_log"
           ("effective_from", "project_id", "project_title")
    VALUES (datetime(),       NEW.id,       NEW.title);
END;


CREATE TRIGGER "project_log_delete" AFTER DELETE ON "project" FOR EACH ROW BEGIN
    UPDATE "project_log"
       SET "effective_to" = datetime()
     WHERE "project_id" = OLD.id
       AND "effective_to" IS NULL;
END;
