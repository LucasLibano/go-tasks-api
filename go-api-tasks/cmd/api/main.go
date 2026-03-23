package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/lucaslibano/go-api-tasks/config"
	"github.com/lucaslibano/go-api-tasks/handlers"
)

func main() {

	//Conecta ao banco

	db := config.ConnectDB()
	taskHandler := handlers.NewTaskHandler(db)

	//O router é o "porteiro da API", olha o endereço e o encaminha para a sala correta.

	r := mux.NewRouter()

	//Middleware é o Segurança do corredor, ele vê todas as requisições que entram e saem.

	r.Use(loggingMiddleware)

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	//Rotas das tasks

	r.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")

	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/task/{id}", taskHandler.GetTaskByID).Methods("GET")

	//Porta da variavel de ambiente.
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
