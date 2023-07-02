package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/giordano200699/go_rest_api/models"
	"github.com/giordano200699/go_rest_api/db"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte("Tarea no encontrada"))
		json.NewEncoder(w).Encode(map[string]string{
			"Respuesta": "Tarea no encontrada",
		})
	}else{
		json.NewEncoder(w).Encode(&task)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"Respuesta": "Tarea no encontrada",
		})
	}else{
		// db.DB.Unscoped().Delete(&task)
		db.DB.Delete(&task)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&task)
	}
}