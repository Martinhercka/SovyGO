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
		return errors.New("windows not supported")
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
	return nil
}

//LinuxOpenPort --
func (d *Database) LinuxOpenPort(req s.LinuxUSE) error {
	if !isValidString(req.UserName) {
		return errors.New("wrong request")
	}
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select available, open from linuxport where port = ?")
	var temavailable, temopen string
	err = statement.QueryRow(req.Port).Scan(&temavailable, &temopen)
	if temavailable != "y" || temopen != "n" {
		return errors.New("invalid port")
	}
	if runtime.GOOS == "windows" {
		return errors.New("windows not supported")
	}
	cm := exec.Command("ufw", "allow", req.Port)
	var b2 bytes.Buffer
	cm.Stdout = &b2
	cm.Start()
	cm.Wait()
	io.Copy(os.Stdout, &b2)

	statement, err = db.Prepare("update linuxport set available = 'n' where port = ?")
	_, err = statement.Exec(req.Port)
	if err != nil {
		return err
	}
	return nil
}

//LinuxClosePort --
func (d *Database) LinuxClosePort(req s.LinuxUSE) error {
	if !isValidString(req.UserName) {
		return errors.New("wrong request")
	}
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select available, open from linuxport where port = ?")
	var temavailable, temopen string
	err = statement.QueryRow(req.Port).Scan(&temavailable, &temopen)
	if err != nil || temavailable != "y" || temopen != "y" {
		return errors.New("invalid port")
	}
	if runtime.GOOS == "windows" {
		return errors.New("windows not supported")
	}
	cm := exec.Command("ufw", "deny", req.Port)
	var b2 bytes.Buffer
	cm.Stdout = &b2
	cm.Start()
	cm.Wait()
	io.Copy(os.Stdout, &b2)

	statement, err = db.Prepare("update linuxport set available = 'y' where port = ?")
	_, err = statement.Exec(req.Port)
	if err != nil {
		return err
	}
	return nil
}

//LinuxChPasswd --
func (d *Database) LinuxChPasswd(req s.LinuxUSE) error {
	if !isValidString(req.UserName) {
		return errors.New("wrong request")
	}
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select createdby from linuxuser where username = ?")
	var temCrearedby int
	err = statement.QueryRow(req.UserName).Scan(&temCrearedby)
	if err != nil || temCrearedby != req.Auth.UserID {
		return errors.New("invalid user")
	}
	if runtime.GOOS == "windows" {
		return errors.New("windows not supported")
	}
	printargs := req.Password + "\n" + req.Password + "\n"

	cm := exec.Command("sh", "pass.sh", printargs, req.UserName)
	var b2 bytes.Buffer
	cm.Stdout = &b2
	cm.Start()
	cm.Wait()
	io.Copy(os.Stdout, &b2)
	return nil
}