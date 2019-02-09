package database

import (
	"database/sql"
	"errors"

	scr "github.com/Martinhercka/SovyGo/bin/server/modules/scrypto"
	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	_ "github.com/go-sql-driver/mysql" //needed
)

//Signup provide creation of user in database
func (d *Database) Signup(req s.RegisterRequest) error {
	salted, salt := scr.NewPasswordHash(req.Password)
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("insert into user(username, name, surname, salt, password, auth)values(?,?,?,?,?,?)")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(req.Username, req.Name, req.Surname, salt, salted, "user")
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}

//LoginRead provide read of login data
func (d *Database) LoginRead(req s.LoginRequest) (s.UserIn, error) {
	var u s.UserIn
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return u, errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select iduser, name, surname, salt, password, auth from user where (username = ? or email = ?)")
	defer statement.Close()
	err = statement.QueryRow(req.Username, req.Email).Scan(u.User.UserID, u.User.Name, u.User.Surname, u.User.Salt, u.User.Password, u.User.Authority)
	if err != nil {
		return u, errors.New("failed to read row")
	}
	return u, nil
}

//LoginSucces provide write succes record of login
func (d *Database) LoginSucces(userID int) error {
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("update lastlogin set succes = 'y' where userid = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}

//LoginFail provide write of fail to login into DB
func (d *Database) LoginFail(userID int) error {
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("insert into loginincident(userid)values(?)")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}
