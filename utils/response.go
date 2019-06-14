package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJson(w http.ResponseWriter, code int, message string, data interface{}) {
	result, err := json.Marshal(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
	if err != nil {
		log.Printf("json encode failed. err:%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
