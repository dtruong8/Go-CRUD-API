package web

import (
	"customers/db"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	store db.Store
}

func (h *CustomerHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := h.store.CustomerStore.Customers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tx)
	}

}
