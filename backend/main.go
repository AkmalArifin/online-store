package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is listening on port 8090...")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
