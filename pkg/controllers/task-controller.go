package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Harshil2511/go-api/pkg/models"
	"github.com/Harshil2511/go-api/pkg/utils"
	"github.com/gorilla/mux"
)

var NewTask models.Task

func GetTask(w http.ResponseWriter, r *http.Request) {
	newTasks := models.GetAllTasks()
	res, _ := json.Marshal(newTasks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	taskDetails, _ := models.GetTaskById(ID)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)
	t := CreateTask.CreateTask()
	res, _ := json.Marshal(t)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task := models.DeleteTask(ID)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{} //create
	utils.ParseBody(r, updateTask)
	vars := mux.Vars(r) //getbyid
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	//if book exist in db find
	taskDetails, db := models.GetTaskById(ID)

	if updateTask.Title != "" {
		taskDetails.Title = updateTask.Title
	}
	if updateTask.Duedate != "" {
		taskDetails.Duedate = updateTask.Duedate
	}
	if updateTask.Description != "" {
		taskDetails.Description = updateTask.Description
	}
	if updateTask.Status != "" {
		taskDetails.Status = updateTask.Status
	}
	db.Save(&taskDetails)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
