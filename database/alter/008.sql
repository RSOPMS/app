CREATE TABLE IF NOT EXISTS comment (
    id         SERIAL    PRIMARY KEY,
    issue_id   SERIAL    REFERENCES issue(id),
    content    TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


COMMENT ON TABLE  comment            IS 'Comment information';
COMMENT ON COLUMN comment.id         IS 'Comment id';
COMMENT ON COLUMN comment.issue_id   IS 'Issue to which this comment belongs to';
COMMENT ON COLUMN comment.content    IS 'Comment content';
COMMENT ON COLUMN comment.created_at IS 'Comment creation time';
