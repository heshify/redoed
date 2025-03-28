package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/heshify/redoed/internal/auth"
	"github.com/heshify/redoed/internal/models"
	"github.com/heshify/redoed/internal/repository"
	"github.com/heshify/redoed/utils"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Repo *repository.UserRepository
}

func NewAuthHandler(repo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{Repo: repo}
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := utils.ParseJSON(r, &newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateUser(newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	newUser.Password = string(hashedPassword)

	userID, err := h.Repo.CreateUser(&newUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// retrive JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("server misconfiguration"))
		return
	}

	// generate accessToken
	accessToken, err := auth.CreateAccessToken(userID, []byte(jwtSecret))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to generate access token"))
		return
	}

	newUser.Password = ""
	response := map[string]interface{}{
		"message":      "User created successfully",
		"user":         newUser,
		"accessToken":  accessToken,
		"refreshToken": accessToken,
	}

	if err := utils.WriteJSON(w, http.StatusCreated, response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginCredentials models.AuthUser

	if err := utils.ParseJSON(r, &loginCredentials); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate request body
	if err := utils.ValidateLoginPayload(loginCredentials); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if user exists
	user, err := h.Repo.GetUserByEmail(loginCredentials.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid email or password"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to fetch user"))
		}
		return
	}

	// compare hashed passwords
	if !auth.ComparePasswords(user.Password, []byte(loginCredentials.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	// retrive JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("server misconfiguration"))
		return
	}

	// generate accessToken
	accessToken, err := auth.CreateAccessToken(user.ID.String(), []byte(jwtSecret))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to generate access token"))
		return
	}

	// generate refreshToken
	refreshToken, err := auth.CreateRefreshToken(user.ID.String(), []byte(jwtSecret))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to generate refrehs token"))
		return
	}

	// send response
	response := map[string]interface{}{
		"message":      "Login succesful!",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
	if err = utils.WriteJSON(w, http.StatusOK, response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshToken string
	if err := utils.ParseJSON(r, &refreshToken); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// retrive JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("server misconfiguration"))
		return
	}

	userID, err := auth.ValidateRefreshToken(refreshToken, []byte(jwtSecret))
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	accessToken, err := auth.CreateAccessToken(userID, []byte(jwtSecret))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to generate access token"))
		return
	}

	// send response
	response := map[string]interface{}{
		"message":     "Token refresh successfully",
		"accessToken": accessToken,
	}

	if err = utils.WriteJSON(w, http.StatusOK, response); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
}
