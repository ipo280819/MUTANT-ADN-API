package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	muxDispatcher *mux.Router
}

var ()

func newMuxRouter() Router {
	return &muxRouter{mux.NewRouter()}
}

//	Add handler with slash because
//	for some reason StrictSlash not working but meh! jaja

func (routerMux *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	routerMux.muxDispatcher.HandleFunc(uri, f).Methods("GET")
	routerMux.muxDispatcher.HandleFunc(uri+"/", f).Methods("GET")
}
func (routerMux *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	routerMux.muxDispatcher.HandleFunc(uri, f).Methods("POST")
	routerMux.muxDispatcher.HandleFunc(uri+"/", f).Methods("POST")
}
func (routerMux *muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	routerMux.muxDispatcher.HandleFunc(uri, f).Methods("PUT")
	routerMux.muxDispatcher.HandleFunc(uri+"/", f).Methods("PUT")
}
func (routerMux *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	routerMux.muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
	routerMux.muxDispatcher.HandleFunc(uri+"/", f).Methods("DELETE")
}

func (routerMux *muxRouter) SERVE(port string) {
	fmt.Println("Mux HTTP server listen on port", port)
	http.ListenAndServe(port, routerMux.muxDispatcher)
}
