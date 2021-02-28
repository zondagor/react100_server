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

	res := state.Users[userId - 1];

	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}
