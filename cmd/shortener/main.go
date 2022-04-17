package main

import (
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

	w.Write([]byte("Привет"))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {

	}
	for key, v := range Bd {
		if id == key {
			http.Redirect(w, r, v, http.StatusFound)
		}
	}
	http.NotFound(w, r)
	return
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
