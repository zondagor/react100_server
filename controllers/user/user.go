package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"react100_server/state"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	totalNumOfUsers := len(state.Users)

	query := r.URL.Query()
	page := query.Get("page")
	pageSequentialNumber, _ := strconv.Atoi(page)
	countStr := query.Get("count")
	userCountOnOnePage, _ := strconv.Atoi(countStr)

	lastUserIndex := pageSequentialNumber * userCountOnOnePage
	firstUserIndex := lastUserIndex - userCountOnOnePage

	var usersSliced []state.User
	if lastUserIndex > totalNumOfUsers {
		usersSliced = state.Users[firstUserIndex:]
	} else {
		usersSliced = state.Users[firstUserIndex:lastUserIndex]
	}

	res := struct {TotalCount int; Users []state.User} {
		TotalCount: totalNumOfUsers,
		Users:      usersSliced,
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	startIndexOfId := strings.LastIndex(urlPath, "/") + 1
	userId, _ :=  strconv.Atoi(urlPath[startIndexOfId:])

	state.Users[userId - 1].Followed = true

	res := struct {ResultCode int} {
		ResultCode: 0,
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	startIndexOfUserId := strings.LastIndex(urlPath, "/") + 1
	userId, _ :=  strconv.Atoi(urlPath[startIndexOfUserId:])

	state.Users[userId - 1].Followed = false

	res := struct {ResultCode int} {
		ResultCode: 0,
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}
