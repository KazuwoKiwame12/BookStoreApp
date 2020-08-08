package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	book "github.com/KazuwoKiwame12/book_store_app_api/DB/Model"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/get", booksList).Methods("GET")
	// TODO idしかdbに入らない
	r.HandleFunc("/api/add", addBook).Methods("POST")
	r.HandleFunc("/api/delete/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/test", test)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// 起動
	log.Fatal(srv.ListenAndServe())
}

func test(w http.ResponseWriter, r *http.Request) {
	log.Println("test関数")
	w.Write([]byte("Hello"))
}

func booksList(w http.ResponseWriter, r *http.Request) {
	books := book.Get()
	//json形式に変換します
	bytes, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(string(bytes)))
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("add関数")
	e := r.ParseForm()
	log.Print(e)
	title := r.Form.Get("title")
	desc := r.Form.Get("desciption")
	log.Printf("title:%s, desc: %s", title, desc)
	result := book.Add(title, desc)
	answer := "false!"
	if result {
		answer = "success!"
	}
	w.Write([]byte(answer))
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("delete")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Printf(vars["id"])

	result := book.Delete(id)
	answer := "false!"
	if result {
		answer = "success!"
	}
	w.Write([]byte(answer))
}
