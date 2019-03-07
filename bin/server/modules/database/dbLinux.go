package database

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strconv"

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
	//printargs := req.Password + "\n" + req.Password + "\n\n\n\n\n\nY\n"

	cm := exec.Command("sh", "adu.sh", req.Password, req.UserName)
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
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select open from linuxport where port = ?")
	var temopen string
	err = statement.QueryRow(req.Port).Scan(&temopen)
	if temopen != "n" {
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

	statement, err = db.Prepare("update linuxport set open = 'y', changedby = ? where port = ?")
	_, err = statement.Exec(req.Auth.UserID, req.Port)
	if err != nil {
		return err
	}
	return nil
}

//LinuxAvailablePort --
func (d *Database) LinuxAvailablePort(req s.LinuxUSE) (string, error) {
	var err error
	var out string
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return out, errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select port from linuxport where open = 'n'")
	resultset, err := statement.Query()
	out = "{\n\t\"ports\":["
	var swap int
	var first = true
	for resultset.Next() {
		_ = resultset.Scan(&swap)
		if first {
			out += "\n\t\t{\"port\":" + strconv.Itoa(swap) + "}"
			first = false
			continue
		}
		out += ",\n\t\t{\"port\":" + strconv.Itoa(swap) + "}"
	}
	out += "\n\t]\n}"
	return out, nil
}

//LinuxMyPorts --
func (d *Database) LinuxMyPorts(req s.LinuxUSE) (string, error) {
	var err error
	var out string
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return out, errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select port from linuxport where changedby = ? and open = 'y' ")
	resultset, err := statement.Query(req.Auth.UserID)
	out = "{\n\t\"ports\":["
	var swap int
	var first = true
	for resultset.Next() {
		_ = resultset.Scan(&swap)
		if first {
			out += "\n\t\t{\"port\":" + strconv.Itoa(swap) + "}"
			first = false
			continue
		}
		out += ",\n\t\t{\"port\":" + strconv.Itoa(swap) + "}"
	}
	out += "\n\t]\n}"
	return out, nil
}

//LinuxClosePort --
func (d *Database) LinuxClosePort(req s.LinuxUSE) error {
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select open from linuxport where port = ?")
	var temopen string
	err = statement.QueryRow(req.Port).Scan(&temopen)
	if err != nil || temopen != "y" {
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

	statement, err = db.Prepare("update linuxport set open = 'n', changedby = ? where port = ?")
	_, err = statement.Exec(req.Auth.UserID, req.Port)
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
