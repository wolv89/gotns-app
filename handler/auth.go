package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/wolv89/gotnsapp/util"
)


func Ready(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "GO")
}



func Login(w http.ResponseWriter, req *http.Request) {

	loginUser, loginPass, ok := req.BasicAuth()

	if !ok {
		util.HttpBadRequest(w, "No credentials provided")
		return
	}

	const userCount = 3
	var username, password string
	credentialsFound := false

	for i := 0; i < userCount; i++ {
		username = os.Getenv(fmt.Sprintf("USER_%d_LOGIN", i))
		if username != loginUser {
			continue
		}
		password = os.Getenv(fmt.Sprintf("USER_%d_PASS", i))
		if password == loginPass {
			credentialsFound = true
		}
	}

	if !credentialsFound {
		util.HttpBadRequest(w, "Invalid username or password")
		return
	}

	sessToken, err := util.RandomHex(64)
	if err != nil {
		util.HttpBadRequest(w, "Unable to generate session token")
		fmt.Println(err.Error())
		return
	}

	session := util.Session {
		Token: sessToken,
		Expiry: time.Now().Unix() + 43200,
	}

	util.Sessions = append(util.Sessions, session)

	util.HttpSuccess(w, sessToken)

}


func Logout(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Env: %s", os.Getenv("USER_1_LOGIN"))
}
