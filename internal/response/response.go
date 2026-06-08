package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteErrorAndLog(w http.ResponseWriter, err error, status int) {
	log.Print(err)

	errResp := errorResponse{
		Status:  status,
		Message: err.Error(),
	}

	data, marshalErr := json.Marshal(errResp)
	if marshalErr != nil {
		log.Printf("WriteErrAndLog: %v", marshalErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, writeErr := w.Write(data)
	if writeErr != nil {
		log.Printf("WriteErrAndLog: %v", writeErr)
	}
}

func WriteJSON(w http.ResponseWriter, payload any, status int) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("WriteJSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("WriteJSON: %v", err)
	}
}
