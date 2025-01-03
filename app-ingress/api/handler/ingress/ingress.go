package ingress

import (
	"database/sql"
)

type IngressHandler struct {
	Db *sql.DB
}

func NewIngressHandler(db *sql.DB) *IngressHandler {
	return &IngressHandler{
		Db: db,
	}
}
