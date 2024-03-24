package middleware

import (
	"net/http"
	"github.com/wolv89/gotnsapp/util"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3060")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		if r.Method == "OPTIONS" {
			util.HttpSuccess(w,"")
			return
		}
		next.ServeHTTP(w,r)
	})
}
