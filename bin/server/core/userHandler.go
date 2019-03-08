package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//UserListAll --
func (c *Core) UserListAll(w http.ResponseWriter, r *http.Request) {

	var req str.Auth
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"message\":\"wrong request\"}")
		return
	}
	if !c.p.AuthenticateSession(req) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "{\"message\":\"unauthorized\"}")
		return
	}
	data, err := c.DB.UserListAll()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"message\":\"internal error\"}")
		return
	}
	out, err := json.MarshalIndent(data, " ", "  ")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"message\":\"marshaling error\"}")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(out))
}

//UserListGroup -
func (c *Core) UserListGroup(w http.ResponseWriter, r *http.Request) {

}
