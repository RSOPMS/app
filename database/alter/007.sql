DROP TABLE IF EXISTS "comment";
DROP TABLE IF EXISTS "comment_log";
DROP TRIGGER IF EXISTS "comment_log_insert";
DROP TRIGGER IF EXISTS "comment_log_update";
DROP TRIGGER IF EXISTS "comment_log_delete";


CREATE TABLE "comment" (
    "id"         INTEGER NOT NULL,
    "issue_id"   INTEGER NOT NULL,
    "content"    TEXT    NOT NULL,
    "created_at" TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT),
    FOREIGN KEY("issue_id") REFERENCES issue("id")
) STRICT;


CREATE TABLE "comment_log" (
    "id"                 INTEGER NOT NULL,
    "effective_from"     TEXT    NOT NULL,
    "effective_to"       TEXT,
    "comment_id"         INTEGER NOT NULL,
    "comment_issue_id"   INTEGER NOT NULL,
    "comment_content"    TEXT    NOT NULL,
    "comment_created_at" TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "comment_log_insert" AFTER INSERT ON "comment" FOR EACH ROW BEGIN
    INSERT
      INTO "comment_log"
           ("effective_from", "comment_id", "comment_issue_id", "comment_content", "comment_created_at")
    VALUES (datetime(),       NEW.id,        NEW.issue_id,      NEW.content,       NEW.created_at);
END;


CREATE TRIGGER "comment_log_update" AFTER UPDATE ON "comment" FOR EACH ROW BEGIN
    UPDATE "comment_log"
       SET "effective_to" = datetime()
     WHERE "comment_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT
      INTO "comment_log"
           ("effective_from", "comment_id", "comment_issue_id", "comment_content", "comment_created_at")
    VALUES (datetime(),       NEW.id,        NEW.issue_id,      NEW.content,       NEW.created_at);
END;


CREATE TRIGGER "comment_log_delete" AFTER DELETE ON "comment" FOR EACH ROW BEGIN
    UPDATE "comment_log"
       SET "effective_to" = datetime()
     WHERE "comment_id" = OLD.id
       AND "effective_to" IS NULL;
END;
