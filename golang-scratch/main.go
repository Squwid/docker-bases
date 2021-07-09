package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var port = "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello, World!")) })

	fmt.Printf("Listening on port :%v\n", port)
	http.ListenAndServe(":"+port, r)
}
