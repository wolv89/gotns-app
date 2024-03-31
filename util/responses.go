package util

import (
	"encoding/json"
	"net/http"
)


type SimpleRes struct {
	Result 		bool 	`json:"result"`
	Response 	string 	`json:"response"`
	Return 		string 	`json:"return"`
}


func HttpJ(w http.ResponseWriter, res bool, resp string, ret string) {

	sr := SimpleRes{res, resp, ret}

	output, err := json.Marshal(sr)

	if err != nil {
		HttpBadRequest(w, "")
		return
	}

	HttpSuccess(w, string(output))

}


func HttpSuccess(w http.ResponseWriter, s string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func HttpUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func HttpBadRequest(w http.ResponseWriter, s string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(s))
}

func HttpServerError(w http.ResponseWriter, s string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(s))
}

func HttpNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(http.StatusText(http.StatusNotFound)))
}
