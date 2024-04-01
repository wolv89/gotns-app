package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"
	"slices"

	"github.com/wolv89/gotnsapp/util"
)



func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if ! IsAdmin(r) {
			util.HttpUnauthorized(w)
			return
		}

		next.ServeHTTP(w, r)

	})
}



func IsAdmin(r *http.Request) bool {

	authorization := r.Header.Get("Authorization")

	if !strings.HasPrefix(authorization, "Bearer ") {
		return false
	}

	encodedToken := strings.TrimPrefix(authorization, "Bearer ")

	token, err := base64.StdEncoding.DecodeString(encodedToken)
	if err != nil {
		return false
	}

	// Dev hack -- remove for prod!
	if string(token) == "localhost" {
		return true
	}

	isValid := false
	now := time.Now().Unix()
	oldSessions := make([]int, 0, len(util.Sessions))

	for s, sess := range util.Sessions {
		if sess.Expiry < now {
			oldSessions = append(oldSessions, s)
			continue
		}
		if sess.Token == string(token) {
			isValid = true
			break
		}
	}

	for oldSession := range oldSessions {
		util.Sessions = slices.Delete(util.Sessions, oldSession, 1)
	}

	if !isValid {
		return false
	}

	return true

}

