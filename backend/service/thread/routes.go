package thread

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Cleaach/threads/backend/types"
	"github.com/Cleaach/threads/backend/utils"
	"strconv"
	"fmt"
)

type Handler struct {
	threadStore types.ThreadStore
	categoryStore types.CategoryStore
}

func NewHandler(threadStore types.ThreadStore, categoryStore types.CategoryStore) *Handler {
	return &Handler{
		threadStore: threadStore,
		categoryStore: categoryStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/threads", h.handleViewThreads).Methods("GET")
	router.HandleFunc("/threads", h.handleCreateThread).Methods("POST")
	router.HandleFunc("/threads/{id}", h.handleViewThread).Methods("GET")
	router.HandleFunc("/threads/{id}", h.handleDeleteThread).Methods("DELETE")
	router.HandleFunc("/threads/{id}", h.handleEditThread).Methods("PUT")
}

func (h *Handler) handleViewThreads(w http.ResponseWriter, r *http.Request) {
	threads, err := h.threadStore.GetThreads()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err) 
		return
	}

	utils.WriteJSON(w, http.StatusOK, threads)
}

func (h *Handler) handleViewThread(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	threadID, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Find thread
	t, err := h.threadStore.GetThreadByID(threadID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, t)
}

func (h *Handler) handleCreateThread(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateThreadPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Extract user ID from the JWT token
	userID, err := utils.ExtractUserIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	// Fetch or create category
	category, err := h.categoryStore.GetCategoryIDByName(payload.Category)
	if err != nil {
		err = h.categoryStore.CreateCategory(types.Category{Name: payload.Category})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		category, _ = h.categoryStore.GetCategoryIDByName(payload.Category)
	}

	// Create thread
	err = h.threadStore.CreateThread(types.Thread{
		AuthorID:  userID,
		CategoryID: category.ID,
		Title:      payload.Title,
		Content:    payload.Content,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Thread created successfully"})
}

func (h *Handler) handleDeleteThread(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	threadID, err := strconv.Atoi(vars["id"])

	// Find thread
	t, err := h.threadStore.GetThreadByID(threadID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	// Extract user ID from the JWT token
	userID, err := utils.ExtractUserIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	if t.AuthorID == userID {
		err = h.threadStore.DeleteThreadByID(threadID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "thread deleted successfully"})
		return
	}
	
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("thread deletion unauthorized!"))
	return
		
}

func (h *Handler) handleEditThread(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	threadID, err := strconv.Atoi(vars["id"])

	// Find thread
	t, err := h.threadStore.GetThreadByID(threadID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	// Extract user ID from the JWT token
	userID, err := utils.ExtractUserIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	if t.AuthorID == userID {
		
		var payload types.CreateThreadPayload
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		// Fetch or create category
		category, err := h.categoryStore.GetCategoryIDByName(payload.Category)
		if err != nil {
			err = h.categoryStore.CreateCategory(types.Category{Name: payload.Category})
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}
			category, _ = h.categoryStore.GetCategoryIDByName(payload.Category)
		}
		
		err = h.threadStore.EditThread(t.ID, types.Thread{
			AuthorID:  userID,
			CategoryID: category.ID,
			Title:     payload.Title,
			Content:   payload.Content,
		})
		
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		

		utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "thread edited successfully"})
		return
	}
	
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("thread edit unauthorized!"))

	return
}
