package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//CreateDBHandler -
func (c *Core) CreateDBHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.CreateDB
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.DB.RootCreateDatabase(req)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"succes\"}")
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}

//CreateDBUserHandler -
func (c *Core) CreateDBUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.CreateDBUser
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.DB.RootCreateUSer(req)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"succes\"}")
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}

//AsignDBUserHandler -
func (c *Core) AsignDBUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.AsignDBUser
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.DB.AsignUser(req)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"succes\"}")
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}
