package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/inflame-ue/gurl/internal/response"
	"github.com/inflame-ue/gurl/internal/shorten"
)

const shortCodeLength = 6

type postBody struct {
	URL string `json:"url"`
}

func (api *API) HandleCreateURL(w http.ResponseWriter, r *http.Request) {
	var requestBody postBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	if err := validateURL(requestBody.URL); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	shortCode := shorten.ShortenURL(shortCodeLength)
	id, err := api.DB.CreateShortURL(requestBody.URL, shortCode)
	if err != nil {
		response.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	urlResp, err := api.DB.GetShortURLByID(id)
	if err != nil {
		response.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	response.WriteJSON(w, urlResp, http.StatusCreated)
}

func (api *API) HandleRetrieveOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")
	if shortCode == "" || len(shortCode) != shortCodeLength {
		response.WriteErrorAndLog(w, errors.New("invalid shortcode length or format"), http.StatusBadRequest)
		return
	}

	urlResp, err := api.DB.GetOriginalURLByShortURL(shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteErrorAndLog(w, err, http.StatusNotFound)
			return
		}
		response.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	response.WriteJSON(w, urlResp, http.StatusOK)
}
