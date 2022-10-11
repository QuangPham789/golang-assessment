package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/quangpham789/golang-assessment/service"
	"github.com/quangpham789/golang-assessment/utils"
	"net/http"
	"net/mail"
	"strings"
)

type UserHandler struct {
	userService service.UserServ
}

func NewUserHandler(db *sql.DB) UserHandler {
	return UserHandler{
		userService: service.NewUserService(db),
	}
}

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	IsActive  bool   `json:"is_active"`
}

func (user UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	userReq := UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		JsonResponseError(w, err)
		return
	}

	// Validate body user request
	input, err := validateUserInput(userReq)
	if err != nil {
		JsonResponseError(w, err)
		return
	}

	// Call service
	result, err := user.userService.CreateUser(r.Context(), input)
	if err != nil {
		JsonResponseError(w, err)
		return
	}
	utils.JsonResponse(w, http.StatusCreated, result)
}

func validateUserInput(user UserRequest) (service.CreateUserInput, error) {
	firstName := strings.TrimSpace(user.FirstName)
	if firstName == "" {
		return service.CreateUserInput{}, errNameCannotBeBlank
	}

	lastName := strings.TrimSpace(user.LastName)
	if lastName == "" {
		return service.CreateUserInput{}, errNameCannotBeBlank
	}

	email := strings.TrimSpace(user.Email)
	if email == "" {
		return service.CreateUserInput{}, errEmailCannotBeBlank
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return service.CreateUserInput{}, errInvalidEmail
	}

	return service.CreateUserInput{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
	}, nil
}
