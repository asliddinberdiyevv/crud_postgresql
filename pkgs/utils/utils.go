package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Generec error response struct
type GenerecError struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln("Error write response")
	}
}

func WriteError(w http.ResponseWriter, code int, message string, data interface{}) {
	response := GenerecError{
		Error: message,
		Code:  code,
		Data:  data,
	}

	WriteJSON(w, code, response)
}

func ResponseErr(err error, w http.ResponseWriter, msg string, status int) {
	logrus.WithError(err).Warn(msg)
	WriteError(w, status, msg, nil)
}

func ResponseErrWithMap(err error, w http.ResponseWriter, msg string, status int) {
	logrus.WithError(err).Warn(msg)
	WriteError(w, status, msg, map[string]string{
		"error": err.Error(),
	})
}
