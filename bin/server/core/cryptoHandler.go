package core

import (
	"fmt"
	"net/http"

	"github.com/michalnov/SovyGo/bin/server/modules/persistance"
	s "github.com/michalnov/SovyGo/bin/server/modules/structures"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//NewKey register new session and create new key-pair for it
func NewKey(w http.ResponseWriter, r *http.Request, p *persistance.Persistance) {
	var env s.Envelop
	err := env.FromEnvelop(r)
	fmt.Println("1")
	checkErr(err)
	var session s.SessionRequest
	err = session.DecodeSession(env.Body)
	fmt.Println("2")
	checkErr(err)
	env.Key = p.NewRecord(session.SessionID)
	env.Encryption = false
	w.WriteHeader(http.StatusOK)
	resp, err := env.ToEnvelop(env)
	fmt.Println("3")
	checkErr(err)
	fmt.Println(string(resp))
	fmt.Fprintf(w, string(resp))
}
