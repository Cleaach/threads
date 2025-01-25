package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Cleaach/threads/backend/service/user"
	"github.com/Cleaach/threads/backend/service/thread"
	"github.com/Cleaach/threads/backend/service/category"
	"github.com/Cleaach/threads/backend/service/comment"
	"github.com/gorilla/handlers"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	corsAllowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Initialize UserStore and Handler
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Initialize ThreadStore, CategoryStore, and Handler
	threadStore := thread.NewStore(s.db)
	categoryStore := category.NewStore(s.db)
	threadHandler := thread.NewHandler(threadStore, categoryStore) // Pass both stores
	threadHandler.RegisterRoutes(subrouter)

	// Initialize Category Handler
	categoryHandler := category.NewHandler(categoryStore)
	categoryHandler.RegisterRoutes(subrouter)

	// Initialize Comment Handler
	commentStore := comment.NewStore(s.db)
	commentHandler := comment.NewHandler(commentStore)
	commentHandler.RegisterRoutes(subrouter)

	log.Println("Listening on ", s.addr)

	//return http.ListenAndServe(s.addr, router)

	return http.ListenAndServe(s.addr, handlers.CORS(corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders)(router))
}
