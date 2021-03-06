package database

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"

	mail "github.com/Martinhercka/SovyGo/bin/server/modules/mailer"
	scr "github.com/Martinhercka/SovyGo/bin/server/modules/scrypto"
	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	_ "github.com/go-sql-driver/mysql" //needed
)

//UserSignup provide creation of user in database
func (d *Database) UserSignup(req s.RegisterRequest) error {
	salted, salt := scr.NewPasswordHash(req.Password)
	var iduser int
	fmt.Println(d.master.acces)
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select count(iduser) as iduser from user where username = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	err = statement.QueryRow(req.Username).Scan(&iduser)
	if err != nil {
		return errors.New("error while execution of query")
	}
	if iduser != 0 {
		return errors.New("user exist")
	}
	statement, err = db.Prepare("insert into user(username, salt, password, auth, email) values(?,?,?,?,?)")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(req.Username, salt, salted, "user", req.Email)
	if err != nil {
		return errors.New("error while execution of query")
	}
	statement, err = db.Prepare("select iduser from user where username = ?")
	err = statement.QueryRow(req.Username).Scan(&iduser)
	if err != nil {
		return errors.New("error while execution of query")
	}
	statement, err = db.Prepare("insert into userdetail(userid, name, surname, email, class) values (?,?,?,?,?)")
	_, err = statement.Exec(iduser, req.Name, req.Surname, req.Email, req.Class)
	statement, err = db.Prepare("insert into lastlogin(userid, succes) values (?, 'n')")
	_, err = statement.Exec(iduser)
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}

//UserLoginRead provide read of login data
func (d *Database) UserLoginRead(req s.LoginRequest) (s.UserIn, error) {
	var u s.UserIn
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return u, errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select iduser, salt, password, auth, profilepicture, active from user where (username = ? or email = ?)")
	defer statement.Close()
	row := statement.QueryRow(req.Username, req.Email)
	row.Scan(&u.User.UserID, &u.User.Salt, &u.User.Password, &u.User.Authority, &u.User.ProfilePicture, &u.User.Active)
	if err != nil {
		return u, errors.New("failed to read row")
	}
	if u.User.Active == "n" {
		return u, errors.New("not active")
	}
	accepted := scr.MatchPasswordHash(req.Password, u.User.Salt, u.User.Password)
	if accepted {
		d.userLoginSucces(u.User.UserID)
		return u, nil
	}
	d.userLoginFail(u.User.UserID)
	return u, errors.New("wrong password")
}

//UserChangePassword provide write succes record of login
func (d *Database) UserChangePassword(userID int, newPass string, oldpass string) error {
	var salt, salted, passw string
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select salt, password from user where iduser = ?")
	defer statement.Close()
	row := statement.QueryRow(userID)
	row.Scan(&salt, &passw)
	accept := scr.MatchPasswordHash(oldpass, salt, passw)
	if !accept {
		return errors.New("wrong password")
	}
	salted, salt = scr.NewPasswordHash(newPass)
	statement, err = db.Prepare("update user set password = ?, salt = ? where iduser = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(salted, salt, userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}

//UserResetPassword provide write succes record of login
func (d *Database) UserResetPassword(userID int) (string, error) {
	newPass := scr.NewRandomPassword()
	salted, salt := scr.NewPasswordHash(newPass)
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return "", errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("update user set password = ?, salt = ?, where iduser = ?")
	if err != nil {
		return "", errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(salted, salt, userID)
	if err != nil {
		return "", errors.New("error while execution of query")
	}
	return newPass, nil
}

//UserArchive provide write succes record of login
func (d *Database) UserArchive(userID int) error {
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("update user set password = '', salt = '', auth = '' where iduser = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}

//UserUnArchive provide write succes record of login
func (d *Database) UserUnArchive(userID int) (string, error) {
	newPass := scr.NewRandomPassword()
	salted, salt := scr.NewPasswordHash(newPass)
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return "", errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("update user set password = ?, salt = ?, auth = 'user' where iduser = ?")
	if err != nil {
		return "", errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(salted, salt, userID)
	if err != nil {
		return "", errors.New("error while execution of query")
	}
	return newPass, nil
}

//UserLoginSucces provide write succes record of login
func (d *Database) userLoginSucces(userID int) error {
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("update lastlogin set succes = ' ' where userid = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	statement, err = db.Prepare("update lastlogin set succes = 'y' where userid = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}

//UserLoginFail provide write of fail to login into DB
func (d *Database) userLoginFail(userID int) error {
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
	statement, err = db.Prepare("update lastlogin set succes = ' ' where userid = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	statement, err = db.Prepare("update lastlogin set succes = 'n' where userid = ?")
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	_, err = statement.Exec(userID)
	if err != nil {
		return errors.New("error while execution of query")
	}
	return nil
}

//UserActivation --
func (d *Database) UserActivation(req s.RegisterRequest, mailer s.Mail) (s.RegisterRequest, error) {
	db, err := sql.Open("mysql", d.master.acces)
	var usrid int
	if err != nil {
		return req, errors.New("failed to open database")
	}

	defer db.Close()
	statement, err := db.Prepare("select iduser from user where username = ?")
	if err != nil {
		return req, errors.New("failed to prepare statement")
	}
	err = statement.QueryRow(req.Username).Scan(&usrid)

	statement, err = db.Prepare("insert into activationtoken(userid,activationtoken)values(?,?)")
	if err != nil {
		return req, errors.New("failed to prepare statement")
	}

	b := make([]byte, 8)

	rand.Read(b)
	req.ActivationToken = fmt.Sprintf("%x", b)
	_, err = statement.Exec(usrid, req.ActivationToken)
	mail.Activationmail(req.Email, req.ActivationToken, mailer)
	if err != nil {
		return req, errors.New("error while execution of query")
	}
	return req, nil
}

//SetUserActive --
func (d *Database) SetUserActive(tkn string) error {
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select userid from activationtoken where activationtoken = ?")
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to prepare statementasdasd")
	}

	var usrid int
	err = statement.QueryRow(tkn).Scan(&usrid)
	if err != nil {
		fmt.Println(err)
		return errors.New("invalid token")
	}
	statement, err = db.Prepare("delete from activationtoken where activationtoken = ?")
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to prepare statementasdasd")
	}
	_, err = statement.Exec(tkn)
	statement, err = db.Prepare("update user set active = 'y' where iduser = ?")
	fmt.Println(tkn)
	_, err = statement.Exec(usrid)
	return err
}
