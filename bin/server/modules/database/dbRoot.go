package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	_ "github.com/go-sql-driver/mysql" //needed
)

func isValidString(some string) bool {
	if strings.Contains(some, "root") {
		return false
	}
	for _, element := range strings.Split(some, "") {
		if element == "%" || element == "*" || element == "?" || element == ";" || element == "\\" {
			fmt.Println("isValidString REACTED")
			return false
		}
	}
	return true
}

//RootCreateUSer --
func (d *Database) RootCreateUSer(req s.CreateDBUser) error {
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

//RootCreateDatabase --
func (d *Database) RootCreateDatabase(req s.CreateDB) error {
	if !isValidString(req.DBname) {
		return errors.New("wrong request")
	}
	db, err := sql.Open("mysql", d.root.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("show databases")
	var swap string
	resultset, err := statement.Query()
	for resultset.Next() {
		resultset.Scan(&swap)
		if swap == req.DBname {
			return errors.New("database exist")
		}
	}
	statement, err = db.Prepare("create database ?")
	_, err = statement.Exec(req.DBname)
	if err != nil {
		return err
	}
	statement, err = db.Prepare("insert into database(username, owner) values(?,?)")
	_, err = statement.Exec(req.DBname, req.Auth.UserID)
	if err != nil {
		return err
	}
	return nil
}

//AsignUser --
func (d *Database) AsignUser(req s.AsignDBUser) error {
	if !isValidString(req.DBname) {
		return errors.New("dbname")
	}
	if req.DBname == "information_schema" || req.DBname == "mysql" || req.DBname == "performance_schema" || req.DBname == "sys" {
		return errors.New("dbname")
	}
	db, err := sql.Open("mysql", d.root.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	var dbid, owner int
	statement, err := db.Prepare("select iddatabase, owner from database where dbname = ?")
	if err != nil {
		return err
	}
	res := statement.QueryRow(req.DBname)
	res.Scan(&dbid, &owner)
	if owner != req.Auth.UserID {
		return errors.New("unauthorized")
	}
	if req.Privileges == "all" {
		statement, err = db.Prepare("GRANT ALL PRIVILEGES ON ?.* TO ?@'%'")
	} else if req.Privileges == "read" {
		statement, err = db.Prepare("GRANT READ PRIVILEGES ON ?.* TO ?@'%'")
	} else if req.Privileges == "write" {
		statement, err = db.Prepare("GRANT WRITE PRIVILEGES ON ?.* TO ?@'%'")
	} else {
		return errors.New("wrong privileges")
	}
	if err != nil {
		return err
	}
	_, err = statement.Exec(req.DBname, req.UserName)
	if err != nil {
		return err
	}
	statement, err = db.Prepare("insert into asigneduser(database, user)")
	_, err = statement.Exec(dbid, owner)
	if err != nil {
		return err
	}
	return nil
}
