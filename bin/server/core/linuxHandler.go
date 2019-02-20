package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//LinuxCreateUSer -
func (c *Core) LinuxCreateUSer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.LinuxUSE
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.linux.CreateLinuxUser(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"succes\"}")
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}

//LinuxOpenPort -
func (c *Core) LinuxOpenPort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.LinuxUSE
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.linux.OpenLinuxPort(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"succes\"}")
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}

//LinuxClosePort -
func (c *Core) LinuxClosePort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.LinuxUSE
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.linux.CloseLinuxPort(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"succes\"}")
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}

//LinuxChPasswd -
func (c *Core) LinuxChPasswd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.LinuxUSE
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.linux.ChangeLinuxPassword(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"succes\"}")
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}
