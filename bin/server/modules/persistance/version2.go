package persistance

import (
	"errors"
	"fmt"
	"time"

	scr "github.com/Martinhercka/SovyGo/bin/server/modules/scrypto"
)

//Persist --
type Persist struct {
	sesions []Auth
	changed time.Time
}

//Auth standard authentication request
type Auth struct {
	SessionID string `json:"sessionid,omitempty"`
	Username  string `json:"username,omitempty"`
	UserID    int    `json:"iduser,omitempty"`
	Token     string `json:"token,omitempty"`
	Remember  bool   `json:"remember,omitempty"`
}

//CreateSession --
func (s *Persist) CreateSession(a Auth) (Auth, error) {
	if !s.findSession(a.SessionID) {
		a.Token = scr.NewToken()
		s.sesions = append(s.sesions, a)
	}
	return a, errors.New("")
}

func (s *Persist) findSession(sessionid string) bool {
	for _, element := range s.sesions {
		if element.SessionID == sessionid && element.Token != "" {
			return true
		}
	}
	return false
}

//AuthenticateSession --
func (s *Persist) AuthenticateSession(a Auth) bool {
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

func (s *Persist) collectGarbage() {
	now := time.Now()
	now = now.Add(-(5 * time.Minute))
	if s.changed.Before(now) {
		fmt.Println("G")
		var out []Auth
		for _, element := range s.sesions {
			if element.Token != "" || element.Remember {
				out = append(out, element)
			}
		}
		s.sesions = out
		s.changed = time.Now()
	}
}
