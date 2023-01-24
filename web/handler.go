package web

import (
	"customers/db"
	"net/http"
)

type Handler struct {
	*http.ServeMux
	store db.Store
}

func NewHandler(store db.Store) *Handler {
	h := &Handler{
		Mux:   http.NewServeMux(),
		store: store,
	}
	customers := CustomerHandler{store: store}
	h.Handle("/customer", customers.Show())
	return h
}
