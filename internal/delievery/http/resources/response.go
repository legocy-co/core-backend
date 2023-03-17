package v1

import (
	"encoding/json"
	"net/http"
)

type DataMetaResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func Respond(w http.ResponseWriter, data DataMetaResponse) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ErrorRespond(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		DataMetaResponse{
			Data: nil,
			Meta: map[string]interface{}{
				"status": 400,
				"msg":    msg,
			},
		},
	)
}
