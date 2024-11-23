package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"kimu_backend/cmd/app/resolvers"
	"kimu_backend/config"
	"kimu_backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	resolver := &resolvers.Resolver{DB: db}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/", logRequestsMiddleware(srv))

	log.Printf("server running on http://localhost:%s/ 🚀🚀", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func logRequestsMiddleware(next http.Handler) http.Handler {
	// Open or create the log file
	logFile, err := os.OpenFile("logs/requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	// Create a new logger that writes to the file
	logger := log.New(logFile, "", log.LstdFlags)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and URL
		logger.Printf("Request: %s %s", r.Method, r.URL.Path)

		// Read and log the request body
		if r.Body != nil {
			body, err := io.ReadAll(r.Body)
			if err == nil {
				logger.Printf("Request Body: %s", string(body))
				// Replace the body so the handler can still read it
				r.Body = io.NopCloser(bytes.NewReader(body))
			} else {
				logger.Printf("Error reading body: %v", err)
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
