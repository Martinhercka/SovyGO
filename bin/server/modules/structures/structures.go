package structures

import (
	"encoding/json"
)

//LoginRequest req
type LoginRequest struct {
	SessionID string `json:"sessionid,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Salt      string `json:"salt,omitempty"`
}

//LoginResponse res
type LoginResponse struct {
	Message string `json:"message,omitempty"`
}

//UserIn internal structure for user representation
type UserIn struct {
	User     User
	Groups   []Group
	Projects []Project
	Follow   []Card
}

//User basic data of user
type User struct {
	UserID    int    `json:"iduser,omitempty"`
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Class     string `json:"class,omitempty"`
	Salt      string `json:"salt,omitempty"`
	Authority string
}

//Group --
type Group struct {
	GroupName string
	GroupID   int
}

//Project basic information about project
type Project struct {
	ProjectName string
	ProjectID   int
	Role        string
	Active      string
	LinkDeploy  string
	LinkGit     string
}

//Card "ID card" contains basic inforamtion about user
type Card struct {
	Username    string
	UserID      int
	USerPicture string
}

//RegisterRequest req
type RegisterRequest struct {
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Class    string `json:"class,omitempty"`
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
