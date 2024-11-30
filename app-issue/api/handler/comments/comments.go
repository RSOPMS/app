package comments

import (
	"app-issue/pkg"
	"app-issue/template"
	"database/sql"
	"net/http"
)

type CommentHandler struct {
	Db *sql.DB
}

func NewCommentHandler(db *sql.DB) *CommentHandler {
	return &CommentHandler{
		Db: db,
	}
}

func (h *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) error {
	return template.RenderLayout(w, "comments", nil)
}

func (h *CommentHandler) GetCommentList(w http.ResponseWriter, r *http.Request) error {
	comments, err := pkg.ReadComments(h.Db)
	if err != nil {
		return err
	}

	return template.RenderComments(w, "table", comments)
}
