package database

import (
	"github.com/Martinhercka/SovyGo/bin/server/configuration"
)

//Database is structure of database manipulator
type Database struct {
	master masterDb
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
	master, slave, err := configuration.InitializeDb()
	if err != nil {
		return out, err
	}
	out.master.acces = master
	if slave != "" {
		out.master.slave.acces = slave
		out.master.slave.active = true
	}
	return out, nil
}
