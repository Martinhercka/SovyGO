package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//LoginHandler serve main htm page
func (c *Core) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req str.LoginRequest
	var err error
	var user str.UserIn
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "wrong request")
		return
	}
	user, err = c.DB.UserLoginRead(req)
	if err != nil {
		if err.Error() == "wrong password" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "wrong password")
			return
		} else if err.Error() == "not active" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "not activated user")
			return
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, "internal error")
			panic(err)
		}
	}
	auth := str.Auth{Username: user.User.Username, UserID: user.User.UserID, SessionID: req.SessionID}
	auth, err = c.p.CreateSession(auth)
	if err != nil {
		if err.Error() == "session exist" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "session exist")
			return
		}
		w.WriteHeader(500)
		fmt.Fprintf(w, "internal error")
		panic(err)

	}
	out, err := json.Marshal(auth)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(out))
	fmt.Println(user.User.Email)
}

//RegisterHandler serve main htm page
func (c *Core) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reg str.RegisterRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&reg)

	if err != nil {
		if err.Error() == "user exist" {
			sendSimpleMsg(w, 400, "user already exist")
			return
		}
		sendSimpleMsg(w, 300, "Wrong request")
		panic(err)
	}

	err = c.DB.UserSignup(reg)
	if err != nil {
		sendSimpleMsg(w, 500, "internal error")
		panic(err)
	}

	c.DB.UserActivation(reg, c.mail)

	sendSimpleMsg(w, http.StatusCreated, "created")
	//Test print
	//c.Templates["register"].Execute(w, nil)

}

//ActivationHandler --
func (c *Core) ActivationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token, ok := r.URL.Query()["token"]
	if !ok || len(token[0]) < 1 {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "no token")
		return
	}
	var err error
	var tkn string
	tkn = token[0]
	err = c.DB.SetUserActive(tkn)
	if err != nil {
		panic(err)
	}
}

//PasswordResetRequire -
func (c *Core) PasswordResetRequire(w http.ResponseWriter, r *http.Request) {

}

//PasswordReset -
func (c *Core) PasswordReset(w http.ResponseWriter, r *http.Request) {

}

//PasswordChange -
func (c *Core) PasswordChange(w http.ResponseWriter, r *http.Request) {

}
