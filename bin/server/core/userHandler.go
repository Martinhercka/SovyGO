package core

import (
	"encoding/json"
	"net/http"
)

//UserListAll --
func (c *Core) UserListAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.DB.UserListAll()
	if err != nil {
		sendSimpleMsg(w, 400, "user list all error")
		return
	}
	out, err := json.Marshal(data)
	if err != nil {
		sendSimpleMsg(w, 400, "marshal error")
	}
	sendSimpleMsg(w, 200, string(out))
}

//UserListGroup -
func (c *Core) UserListGroup(w http.ResponseWriter, r *http.Request) {

}
