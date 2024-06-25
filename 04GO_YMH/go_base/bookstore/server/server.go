package server

import (
	"bootstore/server/middleware"
	"bootstore/store"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := &BookStoreServer{
		s:   s,
		srv: &http.Server{Addr: addr},
	}
	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.updateBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.getBookHandler).Methods("GET")
	router.HandleFunc("/book", srv.getAllBookHandler).Methods("GET")
	router.HandleFunc("/book/{id}", srv.delBookHandler).Methods("DELETE")
	srv.srv.Handler = middleware.Logging(middleware.Validating(router))
	return srv
}

func (bs *BookStoreServer) createBookHandler(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := bs.s.Create(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func (bs *BookStoreServer) updateBookHandler(w http.ResponseWriter, req *http.Request) {

}

func (bs *BookStoreServer) getBookHandler(w http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}
	book, err := bs.s.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response(w, book)
}

func (bs *BookStoreServer) getAllBookHandler(w http.ResponseWriter, req *http.Request) {

}

func (bs *BookStoreServer) delBookHandler(w http.ResponseWriter, req *http.Request) {

}

func (bs *BookStoreServer) ListenAndServer() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = bs.srv.ListenAndServe()
		errChan <- err
	}()
	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}

}

func (bs *BookStoreServer) Shutdown(contex context.Context) {

}

/*
*

	这些处理函数都是大同小异，都是先获取http请求包数据，然后通过标准库json将这些数据解码（decode）
	为我们需要的store.Book结构实例，再通过Store接口对数据进行存储操作。
*/
func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
