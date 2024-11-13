DROP TABLE IF EXISTS "status";
DROP TABLE IF EXISTS "status_log";
DROP TRIGGER IF EXISTS "status_log_insert";
DROP TRIGGER IF EXISTS "status_log_update";
DROP TRIGGER IF EXISTS "status_log_delete";


CREATE TABLE "status" (
    "id"    INTEGER NOT NULL,
    "name"  TEXT    NOT NULL UNIQUE,
    "order" INTEGER NOT NULL UNIQUE,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TABLE "status_log" (
    "id"             INTEGER NOT NULL,
    "effective_from" TEXT    NOT NULL,
    "effective_to"   TEXT,
    "status_id"      INTEGER NOT NULL,
    "status_name"    TEXT    NOT NULL,
    "status_order"   INTEGER NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "status_log_insert" AFTER INSERT ON "status" FOR EACH ROW BEGIN
    INSERT
      INTO "status_log"
           ("effective_from", "status_id", "status_name", "status_order")
    VALUES (datetime(),       NEW.id,       NEW.name,     NEW."order");
END;


CREATE TRIGGER "status_log_update" AFTER UPDATE ON "status" FOR EACH ROW BEGIN
    UPDATE "status_log"
       SET "effective_to" = datetime()
     WHERE "status_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT
      INTO "status_log"
           ("effective_from", "status_id", "status_name", "status_order")
    VALUES (datetime(),       NEW.id,       NEW.name,     NEW."order");
END;


CREATE TRIGGER "status_log_delete" AFTER DELETE ON "status" FOR EACH ROW BEGIN
    UPDATE "status_log"
       SET "effective_to" = datetime()
     WHERE "status_id" = OLD.id
       AND "effective_to" IS NULL;
END;


INSERT
  INTO "status"
       ("name",        "order")
VALUES ("Open",        0),
       ("Planning",    1),
       ("In progress", 2),
       ("Testing",     3),
       ("Closed",      4);
