package user

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Cleaach/threads/backend/types"
	"github.com/Cleaach/threads/backend/utils"
	"github.com/Cleaach/threads/backend/config"
	"fmt"
	"github.com/Cleaach/threads/backend/service/auth"
	"strconv"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/user/{id}", h.handleFindUsername).Methods("GET")
}

func (h *Handler) handleFindUsername(w http.ResponseWriter, r *http.Request) {
	
	// Extract user ID from the URL path
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get user for the given user ID
	user, err := h.store.GetUserByID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with the username
	utils.WriteJSON(w, http.StatusOK, user.Username)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get JSON Payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil { // Pass &payload as a pointer
		utils.WriteError(w, http.StatusBadRequest, err)
		return // Add a return here to prevent further execution on error
	}

	// Find user
	u, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid username"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{"token": token, "userId": u.ID})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON Payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil { // Pass &payload as a pointer
		utils.WriteError(w, http.StatusBadRequest, err)
		return // Add a return here to prevent further execution on error
	}

	// Check if user exists
	_, err := h.store.GetUserByUsername(payload.Username)
	if err == nil { // If no error, user already exists
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s already exists", payload.Username))
		return
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Create user if they don't exist
	err = h.store.CreateUser(types.User{
		Username: payload.Username,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with success
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}
