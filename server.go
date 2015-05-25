package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tahasabih/BNDBackEnd/GETS"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Name Name
	Age  Age
}

type Name map[string]string
type Age map[string]int

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", GETS.HomeHandler)

	// Posts collection
	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)
	posts.Methods("POST").HandlerFunc(PostsCreateHandler)

	// Posts singular
	post := r.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostShowHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var a = make(map[string]string)
	var b = make(map[string]int)
	a["taha"] = "Sabih"
	a["jameel"] = "Rehman"
	b["Twenty"] = 20
	b["Twenty-one"] = 21

	var c = Data{a, b}
	var d = Payload{c}
	fmt.Println(d)
	u, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(u))
	fmt.Fprintln(rw, string(u))
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts create")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post edit")
}
