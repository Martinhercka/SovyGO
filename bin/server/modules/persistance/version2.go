package persistance

import (
	"errors"
	"fmt"
	"time"

	scr "github.com/Martinhercka/SovyGo/bin/server/modules/scrypto"
	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//Persist --
type Persist struct {
	sesions []str.Auth
	changed time.Time
}

//NewPv2 create new object of persistance
func NewPv2() Persist {
	out := Persist{}
	return out
}

//CreateSession --
func (s *Persist) CreateSession(a str.Auth) (str.Auth, error) {
	if !s.findSession(a.SessionID, a.UserID) {
		a.Token = scr.NewToken()
		s.sesions = append(s.sesions, a)
		return a, nil
	}
	return a, errors.New("session exist")
}

func (s *Persist) findSession(sessionid string, userid int) bool {
	for _, element := range s.sesions {
		if element.SessionID == sessionid && (element.Token != "" || element.UserID != userid) {
			return true
		}
	}
	return false
}

//AuthenticateSession --
func (s *Persist) AuthenticateSession(a str.Auth) bool {
	for _, element := range s.sesions {
		if element.SessionID == a.SessionID {
			if element.Token == a.Token && element.UserID == a.UserID && element.Username == a.Username {
				return true
			}
			break
		}
	}
	return false
}

//LogoutSession --
func (s *Persist) LogoutSession(a str.Auth) {
	for index, element := range s.sesions {
		if element.SessionID == a.SessionID {
			if element.Token == a.Token && element.UserID == a.UserID && element.Username == a.Username {
				s.sesions[index].Token = ""
				s.sesions[index].SessionID = ""
				s.sesions[index].Remember = false
			}
			break
		}
	}
}

func (s *Persist) collectGarbage() {
	now := time.Now()
	now = now.Add(-(5 * time.Minute))
	if s.changed.Before(now) {
		fmt.Println("G")
		var out []str.Auth
		for _, element := range s.sesions {
			if element.Token != "" || element.Remember {
				out = append(out, element)
			}
		}
		s.sesions = out
		s.changed = time.Now()
	}
}
