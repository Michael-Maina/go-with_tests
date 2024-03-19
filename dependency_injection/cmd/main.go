package main

import (
	"dependency_injection"
	"log"
	"net/http"
	"os"
)

func main() {
	dependency_injection.Greet(os.Stdout, "Chris")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency_injection.MyGreeterHandler)))
}