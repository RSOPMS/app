DROP TABLE IF EXISTS "priority";
DROP TABLE IF EXISTS "priority_log";
DROP TRIGGER IF EXISTS "priority_log_insert";
DROP TRIGGER IF EXISTS "priority_log_update";
DROP TRIGGER IF EXISTS "priority_log_delete";


CREATE TABLE "priority" (
    "id"    INTEGER NOT NULL,
    "name"  TEXT    NOT NULL UNIQUE,
    "order" INTEGER NOT NULL UNIQUE,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TABLE "priority_log" (
    "id"             INTEGER NOT NULL,
    "effective_from" TEXT    NOT NULL,
    "effective_to"   TEXT,
    "priority_id"    INTEGER NOT NULL,
    "priority_name"  TEXT    NOT NULL,
    "priority_order" INTEGER NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "priority_log_insert" AFTER INSERT ON "priority" FOR EACH ROW BEGIN
    INSERT
      INTO "priority_log"
           ("effective_from", "priority_id", "priority_name", "priority_order")
    VALUES (datetime(),       NEW.id,        NEW.name,        NEW."order");
END;


CREATE TRIGGER "priority_log_update" AFTER UPDATE ON "priority" FOR EACH ROW BEGIN
    UPDATE "priority_log"
       SET "effective_to" = datetime()
     WHERE "priority_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT
      INTO "priority_log"
           ("effective_from", "priority_id", "priority_name", "priority_order")
    VALUES (datetime(),       NEW.id,        NEW.name,        NEW."order");
END;


CREATE TRIGGER "priority_log_delete" AFTER DELETE ON "priority" FOR EACH ROW BEGIN
    UPDATE "priority_log"
       SET "effective_to" = datetime()
     WHERE "priority_id" = OLD.id
       AND "effective_to" IS NULL;
END;


INSERT
  INTO "priority"
       ("name",     "order")
VALUES ("Critical", 0),
       ("High",     1),
       ("Medium",   2),
       ("Low",      3),
       ("Backlog",  4);
