package handlers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"kirkagram/internal/lib/logger/handlers/customResponse"
	"kirkagram/internal/models"
	"kirkagram/internal/storage"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type PhotoUpl interface {
	UploadPhoto(key string, data []byte) error
}

type Post interface {
	CreatePost(post models.CreatePostRequest) error
	GetAllPosts() (*[]models.Posts, error)
	GetPostByID(ID int64) (*models.Posts, error)
	GetAllPostsByUserID(userID int64) (*[]models.Posts, error)
	DeletePost(ID int64) error
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

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post by user ID
// @Tags posts
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} customResponse.CustomStatus
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /post/{userId} [delete]
func (p *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.post.DeletePost"

	log := p.log.With(slog.String("op", op))
	log.Info("starting delete post")

	userID := chi.URLParam(r, "userId")

	if userID == "" {
		log.Error("userID is empty")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("userID is empty"))

		return
	}

	num, err := strconv.Atoi(userID)
	if err != nil {
		log.Error("error converting id to int")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = p.postService.DeletePost(int64(num))
	if err != nil {
		if errors.Is(err, storage.ErrPostNotFound) {
			log.Error("error deleting post", slog.String("userID", userID), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customResponse.NewError(err.Error()))

			return
		}

		log.Error("error deleting post", slog.String("userID", userID), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, customResponse.NewStatus(200))
}

// GetUserPosts godoc
// @Summary Get user's posts
// @Description Get all posts for a specific user
// @Tags posts
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} models.Posts
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /post/user/{userId} [get]
func (p *PostHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.post.GetUserPosts"

	log := p.log.With(slog.String("op", op))
	log.Info("starting get user posts")

	userID := chi.URLParam(r, "userId")

	if userID == "" {
		log.Error("userID is empty")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("userID is empty"))

		return
	}

	num, err := strconv.Atoi(userID)
	if err != nil {
		log.Error("error converting id to int")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	posts, err := p.postService.GetAllPostsByUserID(int64(num))
	if err != nil {
		if errors.Is(err, storage.ErrPostNotFound) {
			log.Error("error getting all posts", slog.String("userId", userID), slog.String("error", err.Error()))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customResponse.NewError(err.Error()))

			return
		}

		log.Error("error getting all posts", slog.String("userId", userID), slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	log.Info("complete get user posts")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, posts)
}

// GetPostByID godoc
// @Summary Get post by ID
// @Description Get details of a specific post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Posts
// @Failure 400 {object} customResponse.Error
// @Failure 404 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /post/{id} [get]
func (p *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.post.GetPostByID"

	log := p.log.With(slog.String("op", op))
	log.Info("starting get post by id post")

	id := chi.URLParam(r, "id")

	if id == "" {
		log.Error("id is empty")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError("id is empty"))

		return
	}

	num, err := strconv.Atoi(id)
	if err != nil {
		log.Error("error converting id to int")

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	post, err := p.postService.GetPostByID(int64(num))
	if err != nil {
		if errors.Is(err, storage.ErrPostNotFound) {
			log.Error("post not found", slog.String("error", err.Error()), slog.String("op", op))

			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, customResponse.NewError(storage.ErrPostNotFound.Error()))

			return
		}
		log.Error("error getting post by id post", slog.String("op", op))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, post)
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with photo upload
// @Tags posts
// @Accept multipart/form-data
// @Produce json
// @Param photo formData file true "Photo file"
// @Param user_id formData int true "User ID"
// @Param caption formData string true "Post caption"
// @Success 201 {object} customResponse.CustomStatus
// @Failure 400 {object} customResponse.Error
// @Failure 500 {object} customResponse.Error
// @Router /post [post]
func (p *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.post.CreatePost"

	log := p.log.With(slog.String("op", op))
	log.Info("starting creating post")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Error("Failed to parse multipart form", slog.String("error", err.Error()))

		render.Status(r, http.StatusLengthRequired)
		render.JSON(w, r, customResponse.NewError(err.Error()))

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
		log.Error("Failed to read file", slog.String("error", err.Error()))

		render.Status(r, http.StatusLengthRequired)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	userID := r.FormValue("user_id")
	caption := r.FormValue("caption")

	filenameURL := "/api/photo/" + filename

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Error("error converting user id to int", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	post := models.CreatePostRequest{
		UserID:   userIDInt,
		Caption:  caption,
		ImageURL: filenameURL,
	}

	err = p.postService.CreatePost(post)
	if err != nil {
		log.Error("Unable to create post", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	err = p.photoService.UploadPhoto(filename, fileRead)
	if err != nil {
		log.Error("Failed to upload file", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	log.Info("finished creation", slog.String("filename", filename))

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, customResponse.NewStatus(201))
}

// GetAllPosts godoc
// @Summary Get all posts
// @Description Get a list of all posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {array} models.Posts
// @Failure 500 {object} customResponse.Error
// @Router /post/all [get]
func (p *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	const op = "rest.handlers.post.GetAllPosts"

	log := p.log.With(slog.String("op", op))
	log.Info("starting getting all posts")

	posts, err := p.postService.GetAllPosts()
	if err != nil {
		log.Error("Failed to get all posts", slog.String("error", err.Error()))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, customResponse.NewError(err.Error()))

		return
	}

	log.Info("finished getting all posts")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, posts)
}
