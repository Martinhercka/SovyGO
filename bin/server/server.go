package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Martinhercka/SovyGo/bin/server/core"
	"github.com/Martinhercka/SovyGo/bin/server/modules/persistance"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Server structure that hold all parts of application
type Server struct {
	r           *mux.Router
	degradation chan int
	state       persistance.Persistance
	core        core.Core
}

//SetupServer prepare new server structure
func (s *Server) SetupServer(degradation chan int) error {
	fmt.Println("Creating server")
	s.degradation = degradation
	s.state = persistance.NewPersistance()
	return nil
}

//StartServer create routes and execute http.listenAndServe
func (s *Server) StartServer() error {
	var err error
	s.core, err = core.NewCore()
	if err != nil {
		return err
	}
	s.r = mux.NewRouter()
	s.r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_files/test.html")
	})
	s.r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "web_files/404.html")
	})
	s.r.HandleFunc("/home", s.core.HomeHandler).Methods("GET")
	s.r.HandleFunc("/register", s.core.RegisterPageHandler).Methods("GET")
	s.r.HandleFunc("/login", s.core.LoginPageHandler).Methods("GET")
	s.r.HandleFunc("/test", s.core.TestHandler).Methods("GET")
	s.r.HandleFunc("/user/listall", s.core.UserListAll).Methods("POST")
	//s.r.HandleFunc().Methods("post")
	s.r.HandleFunc("/project/new", notImplemented).Methods("post")
	s.r.HandleFunc("/project/remove", notImplemented).Methods("post")
	s.r.HandleFunc("/project/adduser", notImplemented).Methods("post")
	s.r.HandleFunc("/project/listall", notImplemented).Methods("post")
	s.r.HandleFunc("/project/list", notImplemented).Methods("post")

	s.r.HandleFunc("/linux/newuser", s.core.LinuxCreateUSer).Methods("post")
	s.r.HandleFunc("/linux/newport", s.core.LinuxOpenPort).Methods("post")
	s.r.HandleFunc("/linux/available", s.core.LinuxAvailablePort).Methods("post")
	s.r.HandleFunc("/linux/closeport", s.core.LinuxClosePort).Methods("post")
	s.r.HandleFunc("/linux/chpasswd", s.core.LinuxChPasswd).Methods("post")

	s.r.HandleFunc("/mysql/newuser", s.core.CreateDBUserHandler).Methods("post")
	s.r.HandleFunc("/mysql/newdatabase", s.core.CreateDBHandler).Methods("post")
	s.r.HandleFunc("/mysql/asignuser", s.core.AsignDBUserHandler).Methods("post")

	s.r.HandleFunc("/auth/register", s.core.RegisterHandler).Methods("POST")
	s.r.HandleFunc("/auth/login", s.core.LoginHandler).Methods("POST")
	s.r.HandleFunc("/auth/logout", s.core.LogoutHnadler).Methods("POST")
	s.r.HandleFunc("/auth/chpasswd", s.core.PasswordChange).Methods("POST")
	s.r.HandleFunc("/key/new/", func(w http.ResponseWriter, r *http.Request) {
		core.NewKey(w, r, &s.state)
	})
	s.r.HandleFunc("/key/aes/", func(w http.ResponseWriter, r *http.Request) {
		core.ImportAESKey(w, r, &s.state)
	})
	s.r.HandleFunc("/off/1234", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		token, ok := r.URL.Query()["token"]
		if !ok || len(token[0]) < 1 {
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintf(w, "no token")
			return
		}
		if token[0] == s.core.Config.Server.ShutdownKey {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("shutdown"))
			s.degradation <- 0
			return
		}
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("wrong key"))
	})
	s.r.HandleFunc("/hello", notImplemented)
	http.Handle("/", s.r)
	log.Fatal(http.ListenAndServe(s.core.Config.Server.Port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(s.r)))
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "web_files/test.html")
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("not implemented yet"))
}
