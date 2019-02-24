package database

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"os"
	"os/exec"
	"runtime"

	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	_ "github.com/go-sql-driver/mysql" //needed
)

//LinuxCreateUser --
func (d *Database) LinuxCreateUser(req s.LinuxUSE) error {
	return nil
	if !isValidString(req.UserName) {
		return errors.New("wrong request")
	}
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select count(idlinuxuser) from linuxuser where username = ?")
	var swap int
	err = statement.QueryRow(req.UserName).Scan(&swap)
	if swap != 0 {
		return errors.New("user already exist")
	}
	if runtime.GOOS == "windows" {
		return errors.New("creating users for windows not supported yet")
	}
	/*
		printargs := req.Password + "\n" + req.Password + "\n\n\n\n\n\nY\n"
		printcmd := exec.Command("printf", printargs)
		adduserargs := req.UserName
		adduser := exec.Command("adduser", adduserargs)
		r, w := io.Pipe()
		printcmd.Stdout = w
		adduser.Stdin = r
		printcmd.Start()
		adduser.Start()
		printcmd.Wait()
		w.Close()
		adduser.Wait()
		fmt.Println("done")
	*/
	printargs := req.Password + "\n" + req.Password + "\n\n\n\n\n\nY\n"

	cm := exec.Command("sh", "adu.sh", printargs, req.UserName)
	var b2 bytes.Buffer
	cm.Stdout = &b2
	cm.Start()
	cm.Wait()
	io.Copy(os.Stdout, &b2)

	statement, err = db.Prepare("insert into linuxuser(username, createdby) values(?,?)")
	_, err = statement.Exec(req.UserName, req.Auth.UserID)
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
