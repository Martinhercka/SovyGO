package linux

import (
	"fmt"

	"github.com/Martinhercka/SovyGo/bin/server/configuration"
)

//Mysql contains methods for mannaging mysql server
type Mysql struct {
	password string
}

func (m *Mysql) initMysql() {
	var err error
	m.password, err = configuration.LoadMysqlRoot()
	if err != nil {
		fmt.Println("Initialization of Linux/Mysql failed")
		return
	}
}

//InsertUser create new user inside mysql server
func (m *Mysql) InsertUser() error {

	return nil
}

func (m *Mysql) CreateDB() error {

	return nil
}
