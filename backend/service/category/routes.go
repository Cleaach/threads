package category

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Cleaach/cvwo-backend/types"
	"github.com/Cleaach/cvwo-backend/utils"
)

type Handler struct {
	store types.CategoryStore
}

func NewHandler(store types.CategoryStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/categories/{id}/threads", h.handleViewThreadsByCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", h.handleGetCategoryName).Methods("GET")
}

func (h *Handler) handleGetCategoryName(w http.ResponseWriter, r *http.Request) {
	// Extract category ID from the URL path
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	cat, err := h.store.GetCategoryNameByID(categoryID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, cat.Name)
}

// handleViewThreadsByCategory retrieves threads belonging to a specific category
func (h *Handler) handleViewThreadsByCategory(w http.ResponseWriter, r *http.Request) {
	// Extract category ID from the URL path
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get threads for the given category ID
	threads, err := h.store.GetThreadsByCategoryID(categoryID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with the list of threads
	utils.WriteJSON(w, http.StatusOK, threads)
}
