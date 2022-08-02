package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/fatih/camelcase"
	"github.com/gorilla/mux"
)

type Commits struct {
	Commit string `json:"http-server"`
}

//this block creates a function for the handler that responds to the /helloworld endpoint
// the helloworld function listens for the request and handles the response for this endpoint
//then goes on to print the 
func Helloworld(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Stranger\n")
}

//This block creates the function for the second endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")        
	var output string
	splittedWord := camelcase.Split(name)
	for _, word := range splittedWord {
		output += word + " "
	}
	fmt.Fprintf(w,"%s\n", output)

}

//This block handles the versions endpoint
func Versionz(w http.ResponseWriter, r *http.Request) {
	GitCommit := os.Getenv("GIT_COMMIT")

    gitcommit :=Commits{
		Commit:GitCommit,
	}
	json.NewEncoder(w).Encode(gitcommit)
}


func main() {
     
	addr := os.Getenv("ADDR")   //this gets the env variable that 
    //if the env returns empty, the server resolves to port 8080
	if len(addr) == 0{
		addr = ":8080"
	}

	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/helloworld", Helloworld).Methods("GET")
	router.HandleFunc("/helloworld?name=AlfredNeumann", Index).Methods("GET")
    router.HandleFunc("/versionz",Versionz).Methods("GET")


   log.Printf("Server listening on %v", addr)
   log.Fatal(http.ListenAndServe(addr, router))
}