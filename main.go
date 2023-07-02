package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/giordano200699/go_rest_api/routes"
	"github.com/giordano200699/go_rest_api/db"
	"github.com/giordano200699/go_rest_api/models"
	"github.com/go-openapi/runtime/middleware"
)


func main(){
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh) 
 
	//// documentación para compartir
	opts1 := middleware.RedocOpts{SpecURL: "swagger.yaml", Path: "doc"} 
	sh1 := middleware.Redoc(opts1, nil) 
	r.Handle("/doc", sh1)

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}