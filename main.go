package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

type user struct {
	SlackUsername string `json:"slackUsername"`
	Bio           string `json:"bio"`
	Age           int    `json:"age"`
	Backend       bool   `json:"backend"`
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// set content types and header
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

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

// CommonMiddleware sets the CORS needed
func CommonMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter,
		r *http.Request, ps httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next(w, r, ps)
	}
}

func main() {
	// get PORT number from our environmental variable
	var portNumber = ":" + os.Getenv("PORT")

	// create a new router
	// router := http.NewServeMux()
	// router.HandleFunc("/", index)

	router := httprouter.New()

	// router.HandlerFunc("GET", "/", index)
	router.GET("/", CommonMiddleware(index))

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
