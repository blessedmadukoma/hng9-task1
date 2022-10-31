package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type user struct {
	SlackUsername string `json:"slackUsername"`
	Bio           string `json:"bio"`
	Age           int    `json:"age"`
	Backend       bool   `json:"backend"`
}

func index(w http.ResponseWriter, r *http.Request) {
	user := user{}

	user.Age = 20
	user.Backend = true
	user.Bio = "I am a software engineer with 3+ years experience tasked with demystifying the amazing world of performant systems by designing and building high-quality backend services."
	user.SlackUsername = "blessedmadukoma"

	log.Println(user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return
}

func main() {
	// get PORT number from our environmental variable
	var portNumber = ":" + os.Getenv("PORT")

	// create a new router
	router := http.NewServeMux()
	router.HandleFunc("/", index)

	// create our server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router,
	}

	fmt.Printf("Starting server on port %s\n", portNumber)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
