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
	svc.Mux.HandleFunc("GET /shorten/{shortCode}", svc.HandleRetrieveOriginalURL)
	svc.Mux.HandleFunc("PUT /shorten/{shortCode}", svc.HandleUpdateOriginalURL)

	return &svc
}

