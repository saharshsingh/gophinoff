package main

import (
	"github.com/saharshsingh/gophinoff/taskmaster/http"
	"github.com/saharshsingh/gophinoff/taskmaster/impl"
)

func main() {
	server := &http.Server{TaskMaster: impl.CreateTaskMaster()}
	server.Run()
}
