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
		NewAPI("/posts", "GET", api.List),
		NewAPI("/posts/{postID}", "GET", api.Get),
		NewAPI("/posts/{postID}", "PATCH", api.Update),
		NewAPI("/posts/{postID}", "DELETE", api.Delete),
	}

	for _, api := range apis {
		router.HandleFunc(api.Path, api.Func).Methods(api.Method)
	}
}

func (api *PostAPI) Create(w http.ResponseWriter, r *http.Request) {
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

	createdPost, err := api.DB.GetPostByID(ctx, post.ID)
	if err != nil {
		utils.ResponseErr(err, w, "Error creating user.", http.StatusConflict)
		return
	}

	log.Printf("postID: %v, create.\n", post.ID)
	utils.WriteJSON(w, http.StatusCreated, createdPost)
}

func (api *PostAPI) List(w http.ResponseWriter, r *http.Request) {
	var posts []*models.Post
	var err error

	ctx := r.Context()
	paramName := r.URL.Query().Get("name")
	
	if paramName != "" {
		posts, err = api.DB.SearchByName(ctx, paramName)
		if err != nil {
			utils.ResponseErr(err, w, "Error getting posts.", http.StatusConflict)
			return
		}
	} else {
		posts, err = api.DB.GetListPost(ctx)
		if err != nil {
			utils.ResponseErr(err, w, "Error getting posts.", http.StatusConflict)
			return
		}
	}

	if posts == nil {
		posts = make([]*models.Post, 0)
	}
	log.Println("returned posts")
	utils.WriteJSON(w, http.StatusOK, posts)
}

func (api *PostAPI) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := models.PostID(vars["postID"])

	ctx := r.Context()
	post, err := api.DB.GetPostByID(ctx, postID)
	if err != nil {
		utils.ResponseErr(err, w, "Not found post.", http.StatusNotFound)
		return
	}

	log.Printf("postID: %v, returned post\n", post.ID)
	utils.WriteJSON(w, http.StatusOK, post)
}

func (api *PostAPI) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := models.PostID(vars["postID"])

	var postRequest models.Post
	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		utils.ResponseErrWithMap(err, w, "Could not decode parametrs.", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	post, err := api.DB.GetPostByID(ctx, postID)
	if err != nil {
		utils.ResponseErr(err, w, "Error get post.", http.StatusConflict)
		return
	}

	if postRequest.Name != nil || len(*postRequest.Name) != 0 {
		post.Name = postRequest.Name
	}

	if postRequest.Like != nil {
		post.Like = postRequest.Like
	}

	if postRequest.Star != nil {
		post.Star = postRequest.Star
	}

	if err := api.DB.UpdatePost(ctx, post); err != nil {
		log.Println("Error updating post.")
		utils.WriteError(w, http.StatusInternalServerError, "Error updating post.", nil)
		return
	}

	log.Println("Post update")
	utils.WriteJSON(w, http.StatusOK, post)
}

func (api *PostAPI) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := models.PostID(vars["postID"])

	ctx := r.Context()
	post, err := api.DB.DeletePost(ctx, postID)
	if err != nil {
		utils.ResponseErr(err, w, "Error delete post.", http.StatusConflict)
		return
	}

	log.Printf("postID: %v, deleted post\n", postID)
	utils.WriteJSON(w, http.StatusOK, ActDeleted{
		Deleted: post,
	})
}
