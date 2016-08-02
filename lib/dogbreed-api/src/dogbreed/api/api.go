package api

import (
	"fmt"
	"github.com/cespare/go-apachelog"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	// Status Strings and a colon which will seperate them from specific reason

	NotAuthenticatedStatus          = "Authentication not succesful: "
	IncorrectContentTypeStatus      = "Request failed, the content-type must be application/json"
	CouldNotReadRequestDataStatus   = "Failed to read request data: "
	CouldNotCompleteOperationStatus = "The requested operation could not be completed: "
	ResourceDoesNotExistStatus      = "The requested resources does not exist"
	RequestSuccesfulStatus          = "The request was succesfull"
)

const (
	// Status codes supported for the API

	OkStatusCode            = 200
	BadRequestStatusCode    = 400
	NotAuthorizedStatusCode = 401
	NotFoundStatusCode      = 404
	ServerErrorCode         = 500
)

var (
	notFoundHTML = "404 Not Found"
	lAddr        = ""
	server       *http.Server
	HandlerFunc  = func(h http.Handler) http.Handler {
		return h
	}
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, notFoundHTML)
}

func Init(listenAddr string) error {

	gmux := mux.NewRouter()

	gmux.NotFoundHandler = http.HandlerFunc(NotFound)
	gmux.HandleFunc("/healthz", Healthz).Methods("GET")

	//dogbreed api endpoints
	gmux.HandleFunc("/dogs", ListDogs).Methods("GET")

	gmux.HandleFunc("/dogs/breed/{Breed}", GetDogsByBreed).Methods("GET")

	gmux.HandleFunc("/dogs/{URL}", GetDog).Methods("GET")
	gmux.HandleFunc("/dogs/{URL}", SetDog).Methods("PUT")
	gmux.HandleFunc("/dogs/{URL}", AddDog).Methods("POST")
	gmux.HandleFunc("/dogs/{URL}", DeleteDog).Methods("DELETE")

	//maybe instead:
	//dogs/{URL}/?upvote=true
	//dogs/{URL}/?downvote=true
	//dogs/{URL}/?favorite=true
	gmux.HandleFunc("/dogs/favorite/{URL}", FavoriteDog).Methods("PUT")
	gmux.HandleFunc("/dogs/upvote/{URL}", UpvoteDog).Methods("PUT")
	gmux.HandleFunc("/dogs/downvote/{URL}", DownvoteDog).Methods("PUT")

	handler := apachelog.NewHandler(HandlerFunc(gmux), os.Stderr)
	server = &http.Server{Addr: "0.0.0.0:" + listenAddr, Handler: handler}
	lAddr = listenAddr
	return nil

}

func Listen() {

	if server == nil {
		log.Fatal("NOT INITIALIZED")
		panic("Not Initialized.")
	}
	log.Println("[API] Listening on", lAddr)
	log.Println(server.ListenAndServe())
}
