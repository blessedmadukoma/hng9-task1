package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type user struct {
	slackUsername string
	bio string
	age int
	backend bool
}

// home handler is unexported due to no other package requiring the handler
func home(w http.ResponseWriter, r *http.Request) {
	user := user{}

	user.age = 20
	user.backend = true
	user.bio = "I am a software engineer with 3+ years experience tasked with demystifying the amazing world of performant systems by designing and building high-quality backend services."
	user.slackUsername = "skillz"

	log.Println(user)

	// json.NewEncoder(w).Encode("{slackUsername: 'blessed', bio:'I am a software engineer with 3+ years experience tasked with demystifying the amazing world of performant systems by designing and building high-quality backend services.', 'backend': true, 'age': 20}")
	json.NewEncoder(w).Encode(user)
	return
}

func lol(w http.ResponseWriter, r *http.Request) {
	user := user{}

	user.age = 20
	user.backend = true
	user.bio = "I am a software engineer with 3+ years experience tasked with demystifying the amazing world of performant systems by designing and building high-quality backend services."
	user.slackUsername = "skillz"

	log.Println(user)

	// json.NewEncoder(w).Encode("{slackUsername: 'blessed', bio:'I am a software engineer with 3+ years experience tasked with demystifying the amazing world of performant systems by designing and building high-quality backend services.', 'backend': true, 'age': 20}")
	json.NewEncoder(w).Encode(user)
	return
}

// loadEnv loads our .env file: we will use this to test locally
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file")
	}
}

func main() {

	// load our env: to be commented out when we push live
	// loadEnv()

	// get PORT number from our environmental variable
	var portNumber = ":" +os.Getenv("PORT")
	// portNumber = ":" + portNumber

	// create a new router
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/lol", lol)

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
