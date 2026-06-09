package api

import (
	"net/http"

	"github.com/inflame-ue/gurl/internal/database"
)

type API struct {
	DB  *database.DB
	Mux *http.ServeMux
}

func NewServer(db *database.DB) *API {
	svc := API{
		DB:  db,
		Mux: http.NewServeMux(),
	}

	svc.Mux.HandleFunc("POST /shorten", svc.HandleCreateURL)
	svc.Mux.HandleFunc("GET /shorten/{shortCode}", svc.HandleRetrieveURL)
	svc.Mux.HandleFunc("PUT /shorten/{shortCode}", svc.HandleUpdateURL)
	svc.Mux.HandleFunc("DELETE /shorten/{shortCode}", svc.HandleDeleteURL)
	svc.Mux.HandleFunc("GET /shorten/{shortCode}/stats", svc.HandleStatisticsURL)

	return &svc
}
