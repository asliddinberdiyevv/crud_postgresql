package config

import "flag"

type AllowCors struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
}

var (
	AppVersion = "v1"

	/* migrate data folder */
	DataDirectory = flag.String("data-directory", "", "Path for loading templates and migration scripts.")

	AllowedOrigins   = []string{"http://localhost:3000"}
	AllowedMethods   = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}
	AllowedHeaders   = []string{"Content-Type"}
	ExposedHeaders   = []string{"Content-Length"}
	AllowCredentials = true
)
