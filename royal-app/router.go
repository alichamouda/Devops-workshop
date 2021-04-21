package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/royals", returnAllRoyalMembers)
	myRouter.HandleFunc("/royal", createNewRoyalMember).Methods("POST")
	myRouter.HandleFunc("/royal/{id}", deleteRoyalMember).Methods("DELETE")
	myRouter.HandleFunc("/royal/{id}", returnSingleRoyalMember)
	log.Printf("Server started ...")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
