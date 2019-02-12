package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/Martinhercka/SovyGo/bin/server/modules/mailer"
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
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprintf(w, "not activated user")
			//return
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
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, "internal error")
			panic(err)
		}
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
			w.WriteHeader(400)
			fmt.Fprintf(w, "user already exist")
			return
		}
		w.WriteHeader(400)
		fmt.Fprintf(w, "Wrong request")
		panic(err)
	}
	var email string
	email = reg.Email
	err = c.DB.UserSignup(reg)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal error")
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Created")
	m.Activationmail(email)
	//Test print
	//c.Templates["register"].Execute(w, nil)

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
