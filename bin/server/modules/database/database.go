package database

import (
	"database/sql"
	"fmt"

	"github.com/Martinhercka/SovyGo/bin/server/modules/configuration"

	_ "github.com/go-sql-driver/mysql" //needed
)

//Database is structure of database manipulator
type Database struct {
	master masterDb
	root   db
	log    db
}

type masterDb struct {
	acces string
	slave db
}

type db struct {
	active bool
	acces  string
}

//NewDatabase return new structure of database manipulator
func NewDatabase() (Database, error) {
	out := Database{}
	master, slave, root, err := configuration.InitializeDb()
	if err != nil {
		return out, err
	}
	out.master.acces = master
	if slave != "" {
		out.master.slave.acces = slave
		out.master.slave.active = true
	}
	if root != "" {
		out.root.acces = root
		out.root.active = true
	}
	fmt.Println(out.master.acces)
	fmt.Println(out.root.acces)
	fmt.Println(out.log.acces)
	return out, nil
}

//TestConnection test of connection while creating DB
func (d *Database) TestConnection() bool {
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return false
	}
	db.Close()
	return true
}

//CreateConnection open connection to db
func (d *Database) createConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return nil, err
	}
	db.Close()
	return db, nil
}
