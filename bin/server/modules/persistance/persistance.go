package persistance

import (
	"github.com/michalnov/SovyGo/bin/server/modules/authentication"
)

//Persistance struct
type Persistance struct {
	state map[string]authentication.Token
}

//NewPersistance create new persistance structure
func NewPersistance() Persistance {
	out := Persistance{}
	out.state = make(map[string]authentication.Token, 0)
	return out
}
