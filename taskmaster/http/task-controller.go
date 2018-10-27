package http

import (
	"net/http"

	"github.com/saharshsingh/gophinoff/taskmaster/impl"
)

type taskController struct {
	taskMaster *impl.TaskMaster
}

func (controller *taskController) buildHandlers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controller.handleGet(w, r)
		case "POST":
			controller.handlePost(w, r)
		default:
			setErrorResponse(w, &errorMessage{http.StatusMethodNotAllowed,
				"Method Not Allowed", r.Method + " is not supported"})
		}
	}
}

func (controller *taskController) handleGet(w http.ResponseWriter, r *http.Request) {
	setJSONResponse(w, http.StatusOK, transformTaskForTransport(controller.taskMaster.PeekNext()))
}

func (controller *taskController) handlePost(w http.ResponseWriter, r *http.Request) {

	task := &task{}
	unmarshallError := parseJSONRequest(r, task, w)
	if unmarshallError != nil {
		return
	}

	controller.taskMaster.Create(task.Name, task.Description, task.Priority)
	w.WriteHeader(http.StatusCreated)
}
