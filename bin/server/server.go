package server

import (
	"log"

	"github.com/Martinhercka/SovyGo/bin/server/core"
	"github.com/Martinhercka/SovyGo/bin/server/modules/persistance"

	"fmt"
	"net/http"

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
	s.r.HandleFunc("/home", s.core.HomeHandler).Methods("GET")
	s.r.HandleFunc("/auth/activate", s.core.ActivationHandler).Methods("GET")
	s.r.HandleFunc("/register", s.core.RegisterPageHandler).Methods("GET")
	s.r.HandleFunc("/login", s.core.LoginPageHandler).Methods("GET")
	s.r.HandleFunc("/test", s.core.TestHandler).Methods("GET")
	s.r.HandleFunc("/auth/register", s.core.RegisterHandler).Methods("POST")
	s.r.HandleFunc("/auth/login", s.core.LoginHandler).Methods("POST")
	s.r.HandleFunc("/key/new/", func(w http.ResponseWriter, r *http.Request) {
		core.NewKey(w, r, &s.state)
	})
	s.r.HandleFunc("/key/aes/", func(w http.ResponseWriter, r *http.Request) {
		core.ImportAESKey(w, r, &s.state)
	})
	s.r.HandleFunc("/off/1234", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("shutdown"))
		s.degradation <- 0
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
