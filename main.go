package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"paxos-store/consensus"
	"paxos-store/handler"
)

var port int

func init() {
	port = 8080
}

func main() {
	r := mux.NewRouter()
	consensus.NewPaxosEnvironment()

	r.HandleFunc("/insert", handler.PostHandler).Methods(http.MethodPost)
	r.HandleFunc("/get", handler.GetHandler).Methods(http.MethodGet)
	r.HandleFunc("/getData", handler.WeakConsistent).Methods(http.MethodGet)

	logrus.Infof("Starting server at port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		logrus.WithError(err).Fatal("Cannot spawn server")
	}
}


