package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func returnAllRoyalMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllRoyalMembers")
	json.NewEncoder(w).Encode(RoyalMembers)
}
func createNewRoyalMember(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var royal RoyalMember
	json.Unmarshal(reqBody, &royal)
	RoyalMembers = append(RoyalMembers, royal)

	json.NewEncoder(w).Encode(royal)
}
func returnSingleRoyalMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our RoyalMembers
	// if the royalMember.Id equals the key we pass in
	// return the royalMember encoded as JSON
	for _, royalMember := range RoyalMembers {
		if royalMember.Id == key {
			json.NewEncoder(w).Encode(royalMember)
		}
	}
}
func deleteRoyalMember(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the royalMember we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our royalMembers
	for index, royalMember := range RoyalMembers {
		// if our id path parameter matches one of our
		// royalMembers
		if royalMember.Id == id {
			// updates our RoyalMembers array to remove the
			// royalMember
			RoyalMembers = append(RoyalMembers[:index], RoyalMembers[index+1:]...)
		}
	}

}
