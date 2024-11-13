DROP TABLE IF EXISTS "branch";
DROP TABLE IF EXISTS "branch_log";
DROP TRIGGER IF EXISTS "branch_log_insert";
DROP TRIGGER IF EXISTS "branch_log_update";
DROP TRIGGER IF EXISTS "branch_log_delete";


CREATE TABLE "branch" (
    "id"   INTEGER NOT NULL,
    "name" TEXT    NOT NULL,
    "url"  TEXT    NOT NULL UNIQUE,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TABLE "branch_log" (
    "id"             INTEGER NOT NULL,
    "effective_from" TEXT    NOT NULL,
    "effective_to"   TEXT,
    "branch_id"      INTEGER NOT NULL,
    "branch_name"    TEXT    NOT NULL,
    "branch_url"     TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "branch_log_insert" AFTER INSERT ON "branch" FOR EACH ROW BEGIN
    INSERT
      INTO "branch_log"
           ("effective_from", "branch_id", "branch_name", "branch_url")
    VALUES (datetime(),       NEW.id,      NEW.name,      NEW.url);
END;


CREATE TRIGGER "branch_log_update" AFTER UPDATE ON "branch" FOR EACH ROW BEGIN
    UPDATE "branch_log"
       SET "effective_to" = datetime()
     WHERE "branch_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT
      INTO "branch_log"
           ("effective_from", "branch_id", "branch_name", "branch_url")
    VALUES (datetime(),       NEW.id,      NEW.name,      NEW.url);
END;


CREATE TRIGGER "branch_log_delete" AFTER DELETE ON "branch" FOR EACH ROW BEGIN
    UPDATE "branch_log"
       SET "effective_to" = datetime()
     WHERE "branch_id" = OLD.id
       AND "effective_to" IS NULL;
END;
