package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type User struct {
	Uname string
	Pass  string
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup 関数実行")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login 関数実行")
}

//I've nevever done anything yet bout this 10/23
func main() {
	//これのことをハンドラーとしてるみたい
	router2 := mux.NewRouter()

	router2.HandleFunc("/singup", signup).Methods("POST")
	router2.HandleFunc("/login", login).Methods("POST")

	// 何らかの service

	log.Println("サーバ起動")
	log.Fatal(http.ListenAndServe(":8000", router2))
}
