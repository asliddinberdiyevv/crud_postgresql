package v1

import (
	"encoding/json"
	"net/http"
)

type ServerVersion struct {
	Version string `json:"version"`
}


func Version(w http.ResponseWriter, _ *http.Request) {
	versionJSON, err := json.Marshal(ServerVersion{
		Version: "v1",
	})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(versionJSON)
}
