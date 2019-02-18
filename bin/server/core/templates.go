package core

import "net/http"

// "encoding/json"
// "fmt"
// "html/template"
// "net/http"

//NotFound default 404 page
func (c *Core) NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web_files/test.html")
}

//Forbiden deffault
func (c *Core) Forbiden() {

}

//Unauthorized deffault
func (c *Core) Unauthorized() {

}
