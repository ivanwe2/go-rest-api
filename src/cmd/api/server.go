package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"restapi/src/internal/api/middlewares"
	"strings"
	"time"
)

func main() {
	port := ":5000"
	cert := "cert.pem"
	key := "key.pem"
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers", teachersHandler)

	mux.HandleFunc("/students", studentsHandler)

	mux.HandleFunc("/execs", execsHandler)

	rl := middlewares.NewRateLimiter(5, time.Minute)
	hppOptions := middlewares.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	server := &http.Server{
		Addr:      port,
		Handler:   middlewares.Hpp(hppOptions)(rl.Middleware(middlewares.Compression(middlewares.ResponseTimeMiddleware(middlewares.Cors(middlewares.SecurityHeaders(mux)))))),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello root route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Path params
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userId := strings.TrimSuffix(path, "/")

		fmt.Println("The id is: ", userId)

		// QUery params
		fmt.Println(r.URL.Query())
		query := r.URL.Query()
		sortby := query.Get("sortby")

		fmt.Println("The query params is: ", sortby)

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
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Path params
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/students/")
		userId := strings.TrimSuffix(path, "/")

		fmt.Println("The id is: ", userId)

		// QUery params
		fmt.Println(r.URL.Query())
		query := r.URL.Query()
		sortby := query.Get("sortby")

		fmt.Println("The query params is: ", sortby)

		fmt.Println("Getting students")
	case http.MethodPost:
		fmt.Println("Getting students")
	case http.MethodPut:
		fmt.Println("Getting students")
	case http.MethodDelete:
		fmt.Println("Getting students")
	default:
		fmt.Println("Unknown http method")
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Path params
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/execs/")
		userId := strings.TrimSuffix(path, "/")

		fmt.Println("The id is: ", userId)

		// QUery params
		fmt.Println(r.URL.Query())
		query := r.URL.Query()
		sortby := query.Get("sortby")

		fmt.Println("The query params is: ", sortby)

		fmt.Println("Getting execs")
	case http.MethodPost:
		fmt.Println("Getting execs")
	case http.MethodPut:
		fmt.Println("Getting teachexecsers")
	case http.MethodDelete:
		fmt.Println("Getting teacexecshers")
	default:
		fmt.Println("Unknown http method")
	}
}
