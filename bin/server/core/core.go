package core

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	auth "github.com/Martinhercka/SovyGo/bin/server/modules/authentication"
	conf "github.com/Martinhercka/SovyGo/bin/server/modules/configuration"
	dtb "github.com/Martinhercka/SovyGo/bin/server/modules/database"
	str "github.com/Martinhercka/SovyGo/bin/server/modules/structures"
)

//Core --
type Core struct {
	Config    conf.Config
	clients   []session
	Templates map[string]*template.Template
	DB        dtb.Database
}

type session struct {
	sessionID string
	login     bool
	token     auth.Token
	user      str.UserIn
}

//NewCore ---
func NewCore() (Core, error) {
	var core Core
	var err error
	core.Config, err = conf.ReadConfig()
	if err != nil {
		fmt.Println("error read config")
		panic(err)
	}
	core.DB, err = dtb.NewDatabase()
	if err != nil {
		fmt.Println("error in creating of database structure")
	}
	core.loadTemplates()
	if err != nil {
		fmt.Printf("error while loading templates")
		panic(err)
		//return core, err
	}
	core.DB, err = dtb.NewDatabase()
	if err != nil {
		fmt.Printf("error while loading Database")
		panic(err)
	}
	fmt.Println("Result of test database: ", core.DB.TestConnection())

	return core, nil

}

func (c *Core) loadTemplates() error {
	var err error
	c.Templates = make(map[string]*template.Template, 0)
	var swap = make(map[string]*template.Template, 0)

	swap["index"] = laodTemplate("index.html")
	if err != nil {
		return err
	}
	swap["login"] = laodTemplate("login.html")
	if err != nil {
		return err
	}
	swap["register"] = laodTemplate("register.html")
	if err != nil {
		return err
	}
	swap["test"] = laodTemplate("test.html")
	if err != nil {
		return err
	}
	c.Templates = swap

	return nil
}

func laodTemplate(path string) *template.Template {
	absPath, err := filepath.Abs("build/web_files/" + path)
	//fmt.Println(absPath)
	tmpl := template.Must(template.ParseFiles(absPath))
	if err != nil {
		fmt.Println("EEE 22")
		return nil
	}
	//fmt.Println("iii")
	return tmpl

}

//HomeHandler serve main htm page
func (c *Core) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EEE 110")
	//http.ServeFile(w, r, "web_files/test.html")
	c.Templates["index"].Execute(w, "")

}

//LoginPageHandler serve main htm page
func (c *Core) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	c.Templates["login"].Execute(w, nil)
}

//RegisterPageHandler serve main htm page
func (c *Core) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	c.Templates["register"].Execute(w, nil)

}

//TestHandler serve main htm page
func (c *Core) TestHandler(w http.ResponseWriter, r *http.Request) {
	c.Templates["test"].Execute(w, nil)
}

func sendSimpleMsg(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Write()
}
