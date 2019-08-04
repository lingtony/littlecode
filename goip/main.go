package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lingtony/littlecode1/goip/ipaddr"
)

func httpserver(w http.ResponseWriter, r *http.Request) {

	ipaddrs := ipaddr.GetIpv4()
	for index, ip := range ipaddrs {
		fmt.Fprintf(w, "ipadd%d: %s\n", index+1, ip)

	}

}

func main() {
	http.HandleFunc("/", httpserver)
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
