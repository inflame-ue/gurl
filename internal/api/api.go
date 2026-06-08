package api

import (
	"net/http"

	"github.com/inflame-ue/gurl/internal/database"
)

type API struct {
	DB *database.DB
	Mux *http.ServeMux
}

func NewServer(db *database.DB) *API {
	svc := API{
		DB: db,
		Mux: http.NewServeMux(),
	}

	svc.Mux.HandleFunc("POST /shorten", svc.HandleCreateURL)

	return &svc
}

