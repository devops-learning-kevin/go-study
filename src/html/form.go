package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key is: ", k)
		fmt.Println("val is: ", v)
	}
	fmt.Fprintf(w, "e17d27b29fd488799bacd5e73e20a7f68f27eb7b4057c387bdf3b4da1eb667d1")
}
func main() {
	http.HandleFunc("/", GetForm)
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
