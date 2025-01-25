package comment

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Cleaach/cvwo-backend/types"
	"github.com/Cleaach/cvwo-backend/utils"
	"strconv"
	"fmt"
)

type Handler struct {
	store types.CommentStore
}

func NewHandler(store types.CommentStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/threads/{id}/comments", h.handleViewComments).Methods("GET")
	router.HandleFunc("/threads/{id}", h.handleAddComment).Methods("POST")
	router.HandleFunc("/threads/comment/{id}", h.handleDeleteComment).Methods("DELETE")
	router.HandleFunc("/threads/comment/{id}", h.handleEditComment).Methods("PUT")
}

func (h *Handler) handleViewComments(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	threadID, err := strconv.Atoi(vars["id"])
	
	comments, err := h.store.GetCommentsByThreadID(threadID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err) 
		return
	}

	utils.WriteJSON(w, http.StatusOK, comments)
}

func (h *Handler) handleAddComment(w http.ResponseWriter, r *http.Request) {

	var payload types.CreateCommentPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(r)
	threadID, err := strconv.Atoi(vars["id"])

	// Extract user ID from the JWT token
	userID, err := utils.ExtractUserIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	// Add comment
	err = h.store.AddComment(types.Comment{
		AuthorID: userID,
		ThreadID: threadID,
		Content: payload.Content,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with success
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
	return
		
}

func (h *Handler) handleDeleteComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["id"])

	// Find comment
	c, err := h.store.GetCommentByID(commentID)
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

	if c.AuthorID == userID {
		h.store.DeleteCommentByID(c.ID)
		utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "comment deleted successfully"})
		return
	}
	
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("comment deletion unauthorized!"))
	return
		
}

func (h *Handler) handleEditComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["id"])

	// Find comment
	c, err := h.store.GetCommentByID(commentID)
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

	if c.AuthorID == userID {
		
		var payload types.CreateCommentPayload
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		
		err = h.store.EditComment(commentID, types.Comment{
			AuthorID:  c.AuthorID,
			ThreadID:  c.ThreadID,
			Content:   payload.Content,
		})
		
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "comment edited successfully"})
		return
		
	}
	
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("comment edit unauthorized!"))

	return
}