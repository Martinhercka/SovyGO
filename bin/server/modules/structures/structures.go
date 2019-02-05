package structures

import (
	"encoding/json"
)

//LoginRequest req
type LoginRequest struct {
	SessionID string `json:"sessionid,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Salt      string
}

//LoginResponse res
type LoginResponse struct {
	Message string `json:"message,omitempty"`
}

//RegisterRequest req
type RegisterRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	//Class    string `json:"class,omitempty"`
	//Salt     string
}

//SessionRequest req
type SessionRequest struct {
	SessionID string `json:"sessionid,omitempty"`
}

//DecodeLogin method for decoding structure from request
func (l *LoginRequest) DecodeLogin(data []byte) error {
	err := json.Unmarshal(data, l)
	return err
}

//DecodeSession method for decoding session structure from request
func (l *SessionRequest) DecodeSession(data []byte) error {
	err := json.Unmarshal(data, l)
	return err
}
