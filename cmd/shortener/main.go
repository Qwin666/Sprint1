package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var Bd = map[int]string{
	1: "https://google.com",
	2: "https://yandex.ru/",
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Home"))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
		}
		for key, v := range Bd {
			if id == key {
				http.Redirect(w, r, v, http.StatusFound)

			}
		}

		http.NotFound(w, r)
		return
	case http.MethodPost:
		fmt.Println("post")

	default:
		fmt.Fprintf(w, "поддерживает только  GET и POST методы")
		http.Error(w, "поддерживает только  GET и POST методы", http.StatusBadRequest)

	}
}

func handleFunc() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/s", redirect)

	log.Println("Запуск веб-сервера на localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func main() {

	handleFunc()
}
