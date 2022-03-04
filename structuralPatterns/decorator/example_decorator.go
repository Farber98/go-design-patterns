package decorator

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
	Real life example of a server middleware using decorator
*/

type Server struct {
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (m *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprintf(s.LogWriter, "--------------------\n")
	s.Handler.ServeHTTP(w, r)

}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintf(w, "Error trying to retrieve data from Basic auth")
	}
}

func MyServer(selection int) {
	var mySuperServer http.Handler
	switch selection {
	case 1:
		mySuperServer = new(Server)
	case 2:
		mySuperServer = &LoggerServer{
			LogWriter: os.Stdout,
			Handler:   new(Server),
		}
	case 3:
		var user, password string
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)
		mySuperServer = &LoggerServer{
			LogWriter: os.Stdout,
			Handler: &BasicAuthMiddleware{
				User:     user,
				Password: password,
				Handler:  new(Server),
			},
		}
	default:
		mySuperServer = new(Server)
	}

	http.Handle("/", mySuperServer)

	log.Fatal(http.ListenAndServe(":8089", nil))
}
