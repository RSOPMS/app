DROP TABLE IF EXISTS "issue";
DROP TABLE IF EXISTS "issue_log";
DROP TRIGGER IF EXISTS "issue_log_insert";
DROP TRIGGER IF EXISTS "issue_log_update";
DROP TRIGGER IF EXISTS "issue_log_delete";


CREATE TABLE "issue" (
    "id"          INTEGER NOT NULL,
    "title"       TEXT    NOT NULL,
    "description" TEXT,
    "project_id"  INTEGER NOT NULL,
    "status_id"   INTEGER NOT NULL,
    "priority_id" INTEGER NOT NULL,
    "branch_id"   INTEGER NOT NULL,
    "created_at"  TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT),
    FOREIGN KEY("project_id")  REFERENCES project("id"),
    FOREIGN KEY("status_id")   REFERENCES status("id"),
    FOREIGN KEY("priority_id") REFERENCES priority("id"),
    FOREIGN KEY("branch_id")   REFERENCES branch("id")
) STRICT;


CREATE TABLE "issue_log" (
    "id"                INTEGER NOT NULL,
    "effective_from"    TEXT    NOT NULL,
    "effective_to"      TEXT,
    "issue_id"          INTEGER NOT NULL,
    "issue_title"       TEXT    NOT NULL,
    "issue_description" TEXT,
    "issue_project_id"  INTEGER NOT NULL,
    "issue_status_id"   INTEGER NOT NULL,
    "issue_priority_id" INTEGER NOT NULL,
    "issue_branch_id"   INTEGER NOT NULL,
    "issue_created_at"  TEXT    NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
) STRICT;


CREATE TRIGGER "issue_log_insert" AFTER INSERT ON "issue" FOR EACH ROW BEGIN
    INSERT
      INTO "issue_log"
           ("effective_from",
            "issue_id",
            "issue_title",
            "issue_description",
            "issue_project_id",
            "issue_status_id",
            "issue_priority_id",
            "issue_branch_id",
            "issue_created_at")
    VALUES (datetime(),
            NEW.id,
            NEW.title,
            NEW.description,
            NEW.project_id,
            NEW.status_id,
            NEW.priority_id,
            NEW.branch_id,
            NEW.created_at);
END;


CREATE TRIGGER "issue_log_update" AFTER UPDATE ON "issue" FOR EACH ROW BEGIN
    UPDATE "issue_log"
       SET "effective_to" = datetime()
     WHERE "issue_id" = OLD.id
       AND "effective_to" IS NULL;

    INSERT
      INTO "issue_log"
           ("effective_from",
            "issue_id",
            "issue_title",
            "issue_description",
            "issue_project_id",
            "issue_status_id",
            "issue_priority_id",
            "issue_branch_id",
            "issue_created_at")
    VALUES (datetime(),
            NEW.id,
            NEW.title,
            NEW.description,
            NEW.project_id,
            NEW.status_id,
            NEW.priority_id,
            NEW.branch_id,
            NEW.created_at);
END;


CREATE TRIGGER "issue_log_delete" AFTER DELETE ON "issue" FOR EACH ROW BEGIN
    UPDATE "issue_log"
       SET "effective_to" = datetime()
     WHERE "issue_id" = OLD.id
       AND "effective_to" IS NULL;
END;
