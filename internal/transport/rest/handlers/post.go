package handlers

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-chi/render"
	"io"
	"kirkagram/internal/lib/logger/handlers/customErrors"
	"kirkagram/internal/models"
	"log/slog"
	"net/http"
	"time"
)

type PhotoUpl interface {
	UploadPhoto(key string, data []byte) error
}

type Post interface {
	CreatePost(post models.CreatePostRequest) error
	GetAllPosts() (*[]models.Posts, error)
}

type PostHandler struct {
	postService  Post
	photoService PhotoUpl
	log          *slog.Logger
}

func NewPostHandler(postService Post, photoService PhotoUpl, log *slog.Logger) *PostHandler {
	return &PostHandler{
		postService:  postService,
		photoService: photoService,
		log:          log,
	}
}

func (p *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.post.CreatePost"

	p.log.With(slog.String("op", op))
	p.log.Info("starting creating post")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		p.log.Error("Failed to parse multipart form", slog.String("error", err.Error()))

		render.Status(r, http.StatusLengthRequired)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	file, header, err := r.FormFile("photo")
	currentTime := time.Now()
	timestampString := currentTime.Format("2006-01-02_15-04-05.000000")
	filename := header.Filename
	filename = filename + timestampString
	hash := sha256.Sum256([]byte(filename))
	filename = fmt.Sprintf("%x", hash[:8])

	fileRead, err := io.ReadAll(file)
	if err != nil {
		p.log.Error("Failed to read file", slog.String("error", err.Error()))

		render.Status(r, http.StatusLengthRequired)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	userID := r.FormValue("user_id")
	caption := r.FormValue("caption")

	filenameURL := "/api/photo/" + filename

	post := models.CreatePostRequest{
		UserID:   userID,
		Caption:  caption,
		ImageURL: filenameURL,
	}

	err = p.postService.CreatePost(post)
	if err != nil {
		p.log.Error("Unable to create post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	err = p.photoService.UploadPhoto(filename, fileRead)
	if err != nil {
		p.log.Error("Failed to upload file", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	p.log.Info("finished creation", slog.String("filename", filename))

	render.Status(r, http.StatusCreated)
}

func (p *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.post.GetAllPosts"

	p.log.With(slog.String("op", op))
	p.log.Info("starting getting all posts")

	posts, err := p.postService.GetAllPosts()
	if err != nil {
		p.log.Error("Failed to get all posts", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customErrors.NewError(err.Error()))

		return
	}

	p.log.Info("finished getting all posts")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, posts)
}
