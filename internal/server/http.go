package server

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
)


func NewHTTPServer(addr string) *http.Server{

	httpServer := newHTTPServer()
	r := mux.NewRouter()

	r.HandleFunc("/", httpServer.handleProduce).Methods("POST")
	r.HandleFunc("/", httpServer.handleConsume).Methods("GET")
	return &http.Server{
		Addr: addr,
		Handler: r,
	}

}

type httpServer struct {
	Log *Log
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

