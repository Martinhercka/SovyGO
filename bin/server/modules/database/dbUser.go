package database

import (
	"database/sql"
	"errors"
	"fmt"

	s "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
	_ "github.com/go-sql-driver/mysql" //needed
)

//UserListAll return list of all users
func (d *Database) UserListAll() ([]s.Card, error) {
	out := make([]s.Card, 0)
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return out, errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select iduser, username, email, profilepicture, active from user")
	if err != nil {
		return out, errors.New("failed to prepare statement")
	}
	result, err := statement.Query()
	if err != nil {
		return out, errors.New("error while execution of query")
	}
	for result.Next() {
		var swap s.Card
		var act string
		err = result.Scan(&swap.UserID, &swap.Username, &swap.Email, &swap.Picture, &act)
		//if act == "y" {
		//	out = append(out, swap)
		//} else {
		//	fmt.Println("inactive user")
		//}
		out = append(out, swap)
	}
	fmt.Println(len(out), "users found")
	return out, nil
}

//UserListGroup return list of all users
func (d *Database) UserListGroup(groupID int) ([]s.Card, error) {
	out := make([]s.Card, 0)
	var err error
	db, err := sql.Open("mysql", d.master.acces)
	if err != nil {
		return out, errors.New("failed to open database")
	}
	defer db.Close()
	statement, err := db.Prepare("select iduser, username, email, profilepicture, active from user inner join groupdetail on user.iduser = groupdetail.userid where idgroupdetail = ?")
	if err != nil {
		return out, errors.New("failed to prepare statement")
	}
	result, err := statement.Query(groupID)
	if err != nil {
		return out, errors.New("error while execution of query")
	}
	for result.Next() {
		var swap s.Card
		var act string
		err = result.Scan(&swap.UserID, &swap.Username, &swap.Email, &swap.Picture, &act)
		if act == "y" {
			out = append(out, swap)
		} else {
			fmt.Println("inactive user")
		}

	}
	fmt.Println(len(out), "users found")
	return out, nil
}
