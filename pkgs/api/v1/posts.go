package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"posts/pkgs/database"
	"posts/pkgs/models"
	"posts/pkgs/utils"

	"github.com/gorilla/mux"
)

type PostAPI struct {
	DB database.Database
}

func SetPostAPI(db database.Database, router *mux.Router) {
	api := PostAPI{
		DB: db,
	}

	apis := []API{
		NewAPI("/posts", "POST", api.Create),
	}

	for _, api := range apis {
		router.HandleFunc(api.Path, api.Func).Methods(api.Method)
	}
}

// POST - /users/{userID}/categories
func (api *PostAPI) Create(w http.ResponseWriter, r *http.Request) {
	// Decode parameters
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.ResponseErrWithMap(err, w, "Could not decode parametrs.", http.StatusBadRequest)
		return
	}

	if err := post.Verify(); err != nil {
		utils.ResponseErrWithMap(err, w, "Not all fields found.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := api.DB.CreatePost(ctx, &post); err != nil {
		log.Println("Error creating post.")
		utils.WriteError(w, http.StatusInternalServerError, "Error creating post.", nil)
		return
	}

	log.Printf("postID: %v, create.\n", post.ID)
	utils.WriteJSON(w, http.StatusCreated, post)
}
