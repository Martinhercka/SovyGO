package core

import (
	"encoding/json"
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
		sendSimpleMsg(w, 300, "Wrong request")
		panic(err)
	}
	user, err = c.DB.UserLoginRead(req)
	if err != nil {
		if err.Error() == "wrong password" {
			sendSimpleMsg(w, http.StatusUnauthorized, "wrong password")
		} else {
			sendSimpleMsg(w, 500, "internal error")
			panic(err)
		}
	}

	sendSimpleMsg(w, http.StatusAccepted, "status OK")
}

//RegisterHandler serve main htm page
func (c *Core) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reg str.RegisterRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		sendSimpleMsg(w, 300, "Wrong request")
		panic(err)
	}
	var email string
	email = reg.Email
	err = c.DB.UserSignup(reg)
	if err != nil {
		sendSimpleMsg(w, 500, "internal error")
		panic(err)
	}
	m.Activationmail(email)
	sendSimpleMsg(w, http.StatusCreated, "created")
	//Test print
	//c.Templates["register"].Execute(w, nil)

}
