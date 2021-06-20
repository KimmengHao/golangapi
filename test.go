package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/utahta/go-linenotify/auth"
)

// EDIT THIS
var (
	BaseURL      = "0.0.0.0"
	ClientID     = "BhWus13WIhI4HI7loycM42"
	ClientSecret = "HJHgKSrqUuhIepCzNkCx7E82RSTN1m47dqPoS1Lf6VA"
	port         = ":9090"
)

func Authorize(w http.ResponseWriter, req *http.Request) {
	c, err := auth.New(ClientID, BaseURL+port+"/callback")
	if err != nil {
		fmt.Fprintf(w, "error:%v", err)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "state", Value: c.State, Expires: time.Now().Add(60 * time.Second)})

	c.Redirect(w, req)
}

func Callback(w http.ResponseWriter, req *http.Request) {
	c, err := auth.New(ClientID, BaseURL+"/callback")
	if err != nil {
		fmt.Fprintf(w, "error:%v", err)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "state", Value: c.State, Expires: time.Now().Add(60 * time.Second)})

	fmt.Fprintf(w, "callback:")
}

func main() {
	fmt.Println("Hello, world.")
	godotenv.Load()
	port = os.Getenv("PORT")

	router := mux.NewRouter()
	headers := handlers.AllowedOrigins([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedOrigins([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/auth", Authorize)
	router.HandleFunc("/callback", Callback)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router)))
}
