package rest

import (
	"kirkagram/internal/models"
	"kirkagram/internal/transport/rest/handlers"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type Photo interface {
	GetPhoto(key string) ([]byte, error)
	UploadPhoto(key string, data []byte) error
}

type Post interface {
	CreatePost(post models.CreatePostRequest) error
}

type Handler struct {
	userHandler  *handlers.UserHandler
	photoHandler *handlers.PhotoHandler
	postHandler  *handlers.PostHandler
	log          *slog.Logger
}

func NewHandler(log *slog.Logger, userHandler *handlers.UserHandler, photoHandler *handlers.PhotoHandler, postHandler *handlers.PostHandler) *Handler {
	return &Handler{userHandler: userHandler, photoHandler: photoHandler, postHandler: postHandler, log: log}
}

func (h *Handler) InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		h.log.Info("Hello")
		w.Write([]byte("Hello"))
	})

	router.Route("/api", func(r chi.Router) {
		h.log.Info("Init api routes")
		r.Get("/user/{id}", h.userHandler.GetUser)
		r.Put("/user", h.userHandler.UpdateUser)
		r.Get("/user/{userID}/followers", h.userHandler.GetAllFollowers)
		r.Get("/user/{userID}/following", h.userHandler.GetAllFollowing)
		r.Get("/photo/{key}", h.photoHandler.GetPhotoURL)
		r.Post("/photo", h.photoHandler.UploadPhoto)
		r.Post("/post", h.postHandler.CreatePost)
		r.Get("/post/all", h.postHandler.GetAllPosts)
	})

	return router
}
