package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (*muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (*muxRouter) Serve(port string) error {
	log.Printf("Mux HTTP server running on port %v\n", port)
	
	err := http.ListenAndServe(port, muxDispatcher)
	if err != nil {
		return err
	}

	return nil
}
