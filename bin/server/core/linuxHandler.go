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
		err = c.DB.LinuxCreateUser(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"success\"}")
		return
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
		panic(err)
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.DB.LinuxOpenPort(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err.Error())
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			panic(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"success\"}")
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}

//LinuxAvailablePort -
func (c *Core) LinuxAvailablePort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.LinuxUSE
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		panic(err)
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		out, err := c.DB.LinuxAvailablePort(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			panic(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, out)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}

//LinuxMyPorts -
func (c *Core) LinuxMyPorts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req str.LinuxUSE
	var err error
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
		panic(err)
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		out, err := c.DB.LinuxMyPorts(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			panic(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, out)
		return
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
		panic(err)
		return
	}
	if c.p.AuthenticateSession(req.Auth) {
		err = c.linux.CloseLinuxPort(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			panic(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"success\"}")
		return
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
		err = c.DB.LinuxChPasswd(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"status\" : \"wrong request\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\" : \"success\"}")
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "{\"status\" : \"unauthorized\"}")
}
