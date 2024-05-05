package main
import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mukesh")
	})
	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CSE")
	})
	http.HandleFunc("/reg-number", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211CS218")
	})
	http.HandleFunc("/favourite-colour", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Black")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
