package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func httpserver(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Fprintf(w, "Hostname: %s", host)
	}

}

func main() {
	http.HandleFunc("/", httpserver)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
