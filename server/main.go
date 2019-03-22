package main

import (
	"fmt"
	"log"
	"net/http"
)

var isDev bool        // Is in development mode
var mySecrets Secrets // mySecrets

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	getSecrets(isDev, &mySecrets)

	fmt.Println("isDev", isDev)

	fmt.Printf("%+v\n", mySecrets)
	fmt.Println(mySecrets.Port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
