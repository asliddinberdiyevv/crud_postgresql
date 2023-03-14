package v1

import (
	"net/http"
	"posts/pkgs/config"
	"posts/pkgs/utils"
)

type ServerVersion struct {
	Version string `json:"version"`
}

func Version(w http.ResponseWriter, _ *http.Request) {
	versionJSON := ServerVersion{
		Version: config.AppVersion,
	}

	utils.WriteJSON(w, http.StatusOK, versionJSON)
}
