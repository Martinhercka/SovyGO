package structures

import (
	"encoding/json"
)

//LinuxUSE --
type LinuxUSE struct {
	Auth     Auth   `json:"auth,omitempty"`
	UserName string `json:"username,omitempty"`
	Port     string `json:"port,omitempty"`
	Password string `json:"password,omitempty"`
}

//PasswordChange --
type PasswordChange struct {
	Auth    Auth   `json:"auth,omitempty"`
	OldPass string `json:"oldpass,omitempty"`
	NewPass string `json:"newpass,omitempty"`
}

//CreateDB --
type CreateDB struct {
	Auth   Auth   `json:"auth,omitempty"`
	DBname string `json:"dbname,omitempty"`
}

//CreateDBUser --
type CreateDBUser struct {
	Auth     Auth   `json:"auth,omitempty"`
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

//AsignDBUser --
type AsignDBUser struct {
	Auth       Auth   `json:"auth,omitempty"`
	UserName   string `json:"username,omitempty"`
	DBname     string `json:"dbname,omitempty"`
	Privileges string `json:"privileges,omitempty"`
}

//Mail --
type Mail struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
}

//Auth standard authentication request
type Auth struct {
	SessionID string `json:"sessionid,omitempty"`
	Username  string `json:"username,omitempty"`
	UserID    int    `json:"iduser,omitempty"`
	Token     string `json:"token,omitempty"`
	Remember  bool   `json:"remember,omitempty"`
}

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
	User     User      `json:"user,omitempty"`
	Groups   []Group   `json:"groups,omitempty"`
	Projects []Project `json:"projects,omitempty"`
	Follow   []Card    `json:"follow,omitempty"`
}

//User basic data of user
type User struct {
	UserID         int    `json:"iduser,omitempty"`
	Username       string `json:"username,omitempty"`
	Name           string `json:"name,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	Class          string `json:"class,omitempty"`
	Salt           string `json:"salt,omitempty"`
	Authority      string `json:"authority,omitempty"`
	Active         string `json:"activated,omitempty"`
	ProfilePicture string `json:"profilepicture,omitempty"`
}

//Group --
type Group struct {
	GroupName string `json:"groupname,omitempty"`
	GroupID   int    `json:"groupid,omitempty"`
}

//Project basic information about project
type Project struct {
	ProjectName string `json:"projectname,omitempty"`
	ProjectID   int    `json:"projectid,omitempty"`
	Role        string `json:"role,omitempty"`
	Active      string `json:"active,omitempty"`
	LinkDeploy  string `json:"linkdeploy,omitempty"`
	LinkGit     string `json:"linkgit,omitempty"`
}

//Card "ID card" contains basic inforamtion about user
type Card struct {
	Username string `json:"username,omitempty"`
	UserID   int    `json:"userid,omitempty"`
	Picture  string `json:"picture,omitempty"`
	Email    string `json:"email,omitempty"`
}

//RegisterRequest req
type RegisterRequest struct {
	Username        string `json:"username,omitempty"`
	Name            string `json:"name,omitempty"`
	Surname         string `json:"surname,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	Class           string `json:"class,omitempty"`
	ActivationToken string `json:"activationtoken,omitempty"`
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
