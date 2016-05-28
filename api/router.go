package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/maleck13/testapp/api/middleware"
	
)

type HttpHandler func(wr http.ResponseWriter, req *http.Request) HttpError

func NewRouter() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	//dedicated sub route for /api which will use the PathPrefix
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	//recovery middleware for any panics in the handlers
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	//add middleware for all routes
	n := negroni.New(recovery)
	//add some top level routes

	r.HandleFunc("/sys/info/health",RouteErrorHandler(HealthHandler))
	r.HandleFunc("/sys/info/ping",RouteErrorHandler(Ping))


	r.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(middleware.ExampleMiddleware),
		negroni.Wrap(apiRouter),
	))

	

        apiRouter.HandleFunc("/", RouteErrorHandler(IndexHandler)).Methods("GET")
	
	

        
	//wire up middleware and router
	n.UseHandler(r)

	return n  //negroni implements the http.Handler interface
}


