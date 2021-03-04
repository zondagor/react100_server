package profile

import (
	"encoding/json"
	"net/http"
	"react100_server/state"
	"strconv"
	"strings"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	startIndexOfId := strings.LastIndex(urlPath, "/") + 1
	userId, _ :=  strconv.Atoi(urlPath[startIndexOfId:])

	res := state.Users[userId - 1]

	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func GetUserStatus(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	startIndexOfId := strings.LastIndex(urlPath, "/") + 1
	userId, _ :=  strconv.Atoi(urlPath[startIndexOfId:])

	status := state.Users[userId - 1].Status

	resJSON, err := json.Marshal(struct {ResultCode byte; UserStatus string}{0,status})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	startIndexOfId := strings.LastIndex(urlPath, "/") + 1
	userId, _ :=  strconv.Atoi(urlPath[startIndexOfId:])

	var userStatusFromFrontEnd struct{Status string}
	err := json.NewDecoder(r.Body).Decode(&userStatusFromFrontEnd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newStatus := userStatusFromFrontEnd.Status;
	state.Users[userId - 1].Status = newStatus;

	resJSON, err := json.Marshal(struct {ResultCode byte; UserStatus string}{0,newStatus})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

