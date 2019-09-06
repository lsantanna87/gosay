package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetHostNameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := GetHostname()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error when trying to get Hostname: %s \n", err.Error())
	} else {
		fmt.Fprintf(w, "Hello World! %s \n", hostname)
	}
}

func GetHostname() (string, error) {
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		return "", errors.New("Hostname is empty!")
	}

	return hostname, nil
}

func main() {
	log.Println("Initializing Application!")
	http.HandleFunc("/", GetHostNameHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
