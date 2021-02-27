package auth

import (
	// "fmt"
	"encoding/json"
	"net/http"
)

type authData struct {
	Id    int
	Email string
	Login string
}

type response struct {
	ResultCode int
	Data       authData
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
