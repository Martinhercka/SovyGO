package database

import (
	"database/sql"
	"errors"

	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	_ "github.com/go-sql-driver/mysql" //needed
)

//LinuxCreateUser --
func (d *Database) LinuxCreateUser(req s.CreateDBUser) error {
	return nil
	if !isValidString(req.UserName) {
		return errors.New("wrong request")
	}
	var err error
	db, err := sql.Open("mysql", d.root.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select User from mysql.user")
	var swap string
	resultset, err := statement.Query()
	for resultset.Next() {
		resultset.Scan(&swap)
		if swap == req.UserName {
			return errors.New("user exist")
		}
	}
	statement, err = db.Prepare("create user ?@'%' identified by ?")
	_, err = statement.Exec(req.UserName, req.Password)
	if err != nil {
		return err
	}
	statement, err = db.Prepare("insert into dbuser(username, owner) values(?,?)")
	_, err = statement.Exec(req.UserName, req.Auth.UserID)
	if err != nil {
		return err
	}
	return nil
}
