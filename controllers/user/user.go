package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type user struct {
	Id       int
	Photos   struct{ Small string }
	Followed bool
	Name     string
	Status   string
}

var users = []user{
{
Id: 1,
Followed: false,
Name: "bro",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 2,
Followed: false,
Name: "igor",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 3,
Followed: false,
Name: "igor",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 4,
Followed: false,
Name: "igor",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 5,
Followed: false,
Name: "foxie",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 6,
Followed: false,
Name: "bobbie",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 7,
Followed: false,
Name: "igor",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 8,
Followed: false,
Name: "igor",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 9,
Followed: false,
Name: "igor",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 10,
Followed: false,
Name: "igor",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
{
Id: 11,
Followed: false,
Name: "jasmin",
Status: "single",
Photos: struct{Small string} {Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	totalNumOfUsers := len(users)

	query := r.URL.Query()
	page := query.Get("page")
	pageSequentialNumber, _ := strconv.Atoi(page)
	countStr := query.Get("count")
	userCountOnOnePage, _ := strconv.Atoi(countStr)

	lastUserIndex := pageSequentialNumber * userCountOnOnePage
	firstUserIndex := lastUserIndex - userCountOnOnePage

	var usersSliced []user
	if lastUserIndex > totalNumOfUsers {
		usersSliced = users[firstUserIndex:]
	} else {
		usersSliced = users[firstUserIndex:lastUserIndex]
	}

	res := struct {TotalCount int; Users []user} {
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

	users[userId].Followed = true

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

	users[userId].Followed = false

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