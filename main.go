package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("init aadasd")
	http.HandleFunc("/", HelloServer)
	_ = http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	_, _ = fmt.Fprint(w, strings.Join(os.Environ(), "\n \n"))
}
