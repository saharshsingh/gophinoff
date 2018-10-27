package http

import (
	"log"
	"net/http"

	"github.com/saharshsingh/gophinoff/taskmaster/impl"
)

// Server encapsulates the REST API for taskmaster
type Server struct {
	TaskMaster *impl.TaskMaster
}

// Run initializes and starts the server. This function will block till app termination
func (server *Server) Run() {

	// Create controllers
	taskAssignmentsController := &taskAssignmentsController{server.TaskMaster}
	taskController := &taskController{server.TaskMaster}

	// Register handlers
	http.HandleFunc("/task/assignments", taskAssignmentsController.buildHandlers())
	http.HandleFunc("/task", taskController.buildHandlers())

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
