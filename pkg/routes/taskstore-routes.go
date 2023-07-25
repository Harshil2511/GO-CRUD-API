package routes

import (
	"github.com/gorilla/mux"

	"github.com/Harshil2511/go-api/pkg/controllers"
)

var RegisterTaskStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/task/", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/task/", controllers.GetTask).Methods("GET")
	router.HandleFunc("/task/{taskId}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/task/{taskId}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{taskId}", controllers.DeleteTask).Methods("DELETE")
}
