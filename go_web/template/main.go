package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"username"`
	Age  string `json:"age"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, name)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("./hello.html")
	if err != nil {
		log.Fatal(err)
	}
	r.ParseForm()
	username := r.PostForm["username"]
	age := r.PostForm["age"]
	log.Println(username)
	log.Println(age)

	user := User{
		Name: username[0],
		Age:  age[0],
	}
	json.NewEncoder(w).Encode(user)
}

func Start(listenAddr string) error {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", SayHello).Methods("GET")
	r.HandleFunc("/user/login", UserLogin).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         listenAddr,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	return srv.ListenAndServe()
}

func main() {
	listenAddr := flag.String("listenAddr", ":9001", "set server addr")
	flag.Parse()

	fmt.Println("the server is running on port", *listenAddr)

	log.Fatal(Start(*listenAddr))
}
