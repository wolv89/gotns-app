package util

import "net/http"

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
