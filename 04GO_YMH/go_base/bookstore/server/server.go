package server

import (
	"bootstore/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func (bs *BookStoreServer) createBookHandler(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := BookStoreServer{
		s:   s,
		srv: &http.Server{Addr: addr},
	}
	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler)
	return nil
}
