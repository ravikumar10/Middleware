package main

import (
  "fmt"
  "net/http"
)

type middleware struct {
  mux      http.Handler // for the router/mux. (Gorilla, httprouter, etc)
  handlers []http.HandlerFunc // middleware funcs to run.
}

func main() {
  mux := http.NewServeMux() // create a new http mux (router)
  mux.HandleFunc("/1", barHandler) // add a route.
  
  m := New() // create some middleware
  m.Add(fooHandler, foo2Handler) // add our middleware functions.
  m.AddMux(mux) // add our router.
  
  http.ListenAndServe(":8080", m) // let's run this thing.
}

// Inits a new middleware chain.
func New() *middleware {
  return &middleware{handlers: make([]http.HandlerFunc, 0, 0)}
}

// Add adds a variable number of handlers using variadic arguments.
func (m *middleware) Add(h ...http.HandlerFunc) {
  m.handlers = append(m.handlers, h...)
}

// AddMux adds our mux to run.
func (m *middleware) AddMux(mux http.Handler) {
  m.mux = mux
}

// So we can satisfy the http.Handler interface.
func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  for _, h := range m.handlers {
    h.ServeHTTP(w, r)
  }
  m.mux.ServeHTTP(w, r)
}

func fooHandler(rw http.ResponseWriter, r *http.Request) {
  fmt.Println("running foo.")
	
}

func foo2Handler(rw http.ResponseWriter, r *http.Request) {
  fmt.Println("running foo 2.")
}

func barHandler(rw http.ResponseWriter, r *http.Request) {
  fmt.Println("running bar.")
  fmt.Fprint(rw, "bar!\n")
}

