package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	return err
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	if err := Greet(w, "world"); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
