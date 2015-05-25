package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tahasabih/BNDBackEnd/GETS"
	"net/http"
)

type payload struct {
	data Data
}

type Data struct {
	name map[string]string
	age  map[int]string
}

var a = make(map[string]string)
var b = make(map[int]string)

func main() {

	a["taha"] = "Sabih"
	a["jameel"] = "Rehman"
	b[20] = "Twenty"
	b[21] = "Twenty-One"

	var c = Data(a, b)
	var d = payload(c)

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
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts index")
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
