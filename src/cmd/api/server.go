package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":5000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello root route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Getting teachers")
			return
		}

		switch r.Method {
		case http.MethodGet:
			fmt.Println("Getting teachers")
		case http.MethodPost:
			fmt.Println("Getting teachers")
		case http.MethodPut:
			fmt.Println("Getting teachers")
		case http.MethodDelete:
			fmt.Println("Getting teachers")
		default:
			fmt.Println("Unknown http method")
		}

		w.Write([]byte("Hello teachers Route"))
		fmt.Println("Hello teachers route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello students Route"))
		fmt.Println("Hello students route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello execs Route"))
		fmt.Println("Hello execs route")
	})

	fmt.Println("Server is running on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
