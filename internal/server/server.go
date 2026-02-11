package server

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct {
	verbose string
	path    string
	f       func(http.ResponseWriter, *http.Request)
}

type Server struct {
	handlers []handler
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) AddHandler(verbose string, path string, f func(http.ResponseWriter, *http.Request)) {
	s.handlers = append(s.handlers, handler{
		verbose: verbose,
		path:    path,
		f:       f,
	})
}

func (s *Server) Run(listener string) error {
	mux := http.NewServeMux()

	for _, h := range s.handlers {
		path := fmt.Sprintf("%s %s", h.verbose, h.path)
		log.Println(" Handler path:", path)
		mux.HandleFunc(path, h.f)
	}

	return http.ListenAndServe(listener, mux)
}
