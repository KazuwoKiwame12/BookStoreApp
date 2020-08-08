package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	book "github.com/KazuwoKiwame12/book_store_app_api/DB/Model"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Form struct {
	Title       string
	Description string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/get", booksList).Methods("GET")
	r.HandleFunc("/api/add", addBook).Methods("POST")
	r.HandleFunc("/api/delete/{id}", deleteBook).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// 起動
	log.Fatal(srv.ListenAndServe())
}

//==================API処理======================
func booksList(w http.ResponseWriter, r *http.Request) {
	books := book.Get()

	// Json化
	bytes, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(string(bytes)))
}

func addBook(w http.ResponseWriter, r *http.Request) {
	/* 参照:
	https://github.com/gorilla/schema#example
	https://golang.org/pkg/net/http/#Request.ParseMultipartForm
	*/
	err := r.ParseMultipartForm(1000)
	if err != nil {
		log.Fatal(err)
	}

	decoder := schema.NewDecoder()
	var form Form
	err = decoder.Decode(&form, r.PostForm)
	if err != nil {
		log.Fatal(err)
	}
	result := book.Add(form.Title, form.Description)
	answer := "false!"
	if result {
		answer = "success!"
	}
	w.Write([]byte(answer))
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	result := book.Delete(id)
	answer := "false!"
	if result {
		answer = "success!"
	}
	w.Write([]byte(answer))
}
