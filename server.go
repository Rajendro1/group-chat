package main

import (
	"encoding/json"
	"fmt"
	"group-chat/db"

	"log"
	"net/http"

	//"database/sql"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

type UsernamePayload struct {
	Username string `json:"username"`
}

func main() {

	http.HandleFunc("/printUsername", printUsernameHandler)

	log.Println("Server will start at http://localhost:8000/")

	route := mux.NewRouter()
	// r.HandleFunc("/insertUser", insertUserHandler).Methods("POST")

	AddApproutes(route)

	log.Fatal(http.ListenAndServe(":8000", route))
}
func printUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var payload UsernamePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Print the received username
	fmt.Println("Received username:", payload.Username)
	createUser, _, createUserErr := db.CreateUsersToDatabase(payload.Username)
	if !createUser{
		log.Println(createUserErr)
	}
	// Send response
	response := map[string]string{"status": "success"}
	json.NewEncoder(w).Encode(response)

}
