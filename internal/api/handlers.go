package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
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
	log.Printf("HandleCreateURL: creating entry with short code %s", shortCode)

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

func (api *API) HandleRetrieveURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")
	if err := validateShortCode(shortCode); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("HandleRetrieveURL: retrieving entry with short code %s", shortCode)
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

func (api *API) HandleUpdateURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")
	if err := validateShortCode(shortCode); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	var requestBody postBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	if err := validateURL(requestBody.URL); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("HandleUpdateURL: updating entry with short code %s", shortCode)
	err := api.DB.UpdateOriginalByShortURL(requestBody.URL, shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteErrorAndLog(w, err, http.StatusNotFound)
			return
		}
		response.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	urlResp, err := api.DB.GetOriginalURLByShortURL(shortCode)
	if err != nil {
		response.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	response.WriteJSON(w, urlResp, http.StatusOK)
}

func (api *API) HandleDeleteURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")
	if err := validateShortCode(shortCode); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("HandleDeleteURL: deleting entry with short code %s", shortCode)
	err := api.DB.DeleteShortURL(shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteErrorAndLog(w, err, http.StatusNotFound)
			return
		}
		response.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *API) HandleStatisticsURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")
	if err := validateShortCode(shortCode); err != nil {
		response.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return	
	}

	stats, err := api.DB.GetStatsByShortURL(shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteErrorAndLog(w, err, http.StatusNotFound)
			return
		}
		response.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	response.WriteJSON(w, stats, http.StatusOK)
}
