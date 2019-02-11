package persistance

import (
	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//Pv2 --
type Pv2 struct {
	sesions     map[string]session
	activeUsers []string
}

type session struct {
	user  s.Card
	token string
}

func (s *session) createToken() {

}
