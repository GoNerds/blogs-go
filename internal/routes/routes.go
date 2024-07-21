package routes

import (
	"blogs-go/handlers"
	"blogs-go/repository"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(userRepo *repository.UserRepository) *mux.Router {
	userHandler := handlers.NewUserHandler(userRepo)
	
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return r
}

func InitializeRoutes(db *mongo.Database) *mux.Router {
	userRepo := repository.NewUserRepository(db)
	return NewRouter(userRepo)
}