package main

import (
	"cloud.google.com/go/compute/metadata"
	"fmt"
	"log"
	"net/http"
	"os"
	"smtest/example"
)

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	name := fmt.Sprintf("projects/%s/secrets/MY_SECRET/versions/latest", projectNumber())
	err := example.AccessSecretVersion(os.Stdout, name)
	if err != nil {
		log.Fatalf("error accessing secret %s", err)
		return
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = fmt.Fprint(w, "Hello, World!")
}

func projectNumber() string {
	projectNumber, err := metadata.NumericProjectID()
	if err != nil {
		log.Fatalf("unable to get project number: %s", err)
	}
	return projectNumber
}
