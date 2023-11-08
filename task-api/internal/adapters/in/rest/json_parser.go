package rest

import (
	"encoding/json"
	"net/http"
	"task-api/internal/adapters/in/rest/consts"
)

func respondWithJSON(writer http.ResponseWriter, payload any, status int) {
	writer.Header().Set(consts.ContentTypeHeader, consts.JSONContentType)
	writer.WriteHeader(status)

	json.NewEncoder(writer).Encode(payload)
}
