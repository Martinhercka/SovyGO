package persistance

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/Martinhercka/SovyGo/bin/server/modules/authentication"
	a "github.com/Martinhercka/SovyGo/bin/server/modules/authentication"
	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
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
func (p *Persistance) NewRecord(sessionID string) string {
	record := a.NewToken()
	p.state[sessionID] = record
	return p.state[sessionID].ServerPem
}

//GetKey returns pem form of public key
func (p *Persistance) GetKey(sessionID string) (string, error) {

	return "", nil
}

//SetSymmetricKey set key (from client) for symmetric encryption
func (p *Persistance) SetSymmetricKey(env s.Envelop) error {
	rng := rand.Reader
	swap := p.state[env.SessionID]
	data, err := rsa.DecryptPKCS1v15(rng, swap.ServerPrivate, env.Body)
	if err != nil {
		return err
	}
	env.Body = data
	fmt.Println(string(env.Body))
	swap.SymmetricKey = data
	p.state[env.SessionID] = swap
	return nil
}
