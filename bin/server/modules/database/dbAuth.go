package database

import (
	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	_ "github.com/go-sql-driver/mysql" //needed
)

//Signup provide creation of user in database
func (d *Database) Signup(req s.RegisterRequest) {

}

//LoginRead provide read of login data
func (d *Database) LoginRead(req s.LoginRequest) (s.LoginRequest, error) {

	return req, nil
}

//LoginSucces provide write succes record of login
func (d *Database) LoginSucces(req s.LoginRequest) {

}

//LoginFail provide write of fail to login into DB
func (d *Database) LoginFail(req s.LoginRequest) {

}
