package auth

import (
	"encoding/json"
	"net/http"
)

type authData struct {
	Id    int    `json:"userId"`
	Email string `json:"email"`
	Login string `json:"login"`
}

type response struct {
	ResultCode int      `json:"resultCode"`
	Data       authData `json:"data"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	res := response{
		ResultCode: 0,
		Data:       authData{Id: 0, Email: "bbb@gfffl.com", Login: "golang=))"},
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	resJSON, err := json.Marshal(struct {
		Res byte `json:"resultCode"`
	}{0})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	resJSON, err := json.Marshal(struct {
		Res byte `json:"resultCode"`
	}{0})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}
