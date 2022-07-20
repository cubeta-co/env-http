package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("init app")
	http.HandleFunc("/", HelloServer)
	_ = http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	_, _ = fmt.Fprint(w, strings.Join(os.Environ(), "\n \n"))
}
