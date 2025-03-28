package handlers

import (
	"errors"
	"net/http"

	"github.com/heshify/redoed/internal/repository"
	"github.com/heshify/redoed/utils"
	"gorm.io/gorm"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing user ID"))
		return
	}

	user, err := h.Repo.GetUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusNotFound, errors.New("user not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to fetch user"))
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, users); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
