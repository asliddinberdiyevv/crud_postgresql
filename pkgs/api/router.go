package api

import (
	"net/http"
	v1 "posts/pkgs/api/v1"
	"posts/pkgs/config"
	"posts/pkgs/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter(db database.Database) (http.Handler, error) {
	/* ---------- ROUTER ---------- */
	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api/" + config.AppVersion).Subrouter()

	/* ---------- ROUTES ---------- */
	router.HandleFunc("/", v1.Version)
	v1.SetPostAPI(db, apiRouter)

	/* ---------- CORS ---------- */
	Cors := cors.New(cors.Options{
		AllowedOrigins:   config.AllowedOrigins,
		AllowedMethods:   config.AllowedMethods,
		AllowedHeaders:   config.AllowedHeaders,
		ExposedHeaders:   config.ExposedHeaders,
		AllowCredentials: config.AllowCredentials,
	})

	return Cors.Handler(router), nil
}
