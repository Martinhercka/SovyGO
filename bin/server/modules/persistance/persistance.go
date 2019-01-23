package persistance

import (
	"github.com/michalnov/SovyGo/bin/server/modules/authentication"
	a "github.com/michalnov/SovyGo/bin/server/modules/authentication"
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

//NewRecord create new Token in Persistance map
func (p *Persistance) NewRecord(sessionID string) {

	record := a.NewToken()
	p.state[sessionID] = record
}

//GetKey returns pem form of public key
func (p *Persistance) GetKey(sessionID string) (string, error) {

	return "", nil
}

//SetSymmetricKey set key (from client) for symmetric encryption
func (p *Persistance) SetSymmetricKey(sessionID string, key string) {

}
