package http

import (
	"net/http"

	"github.com/saharshsingh/gophinoff/taskmaster/impl"
)

// TODO: Don't feel like implementing query params, so mocking it for now
const user string = "saharsh"

type taskAssignmentsController struct {
	taskMaster *impl.TaskMaster
}

func (controller *taskAssignmentsController) buildHandlers() http.HandlerFunc {
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

func (controller *taskAssignmentsController) handleGet(w http.ResponseWriter, r *http.Request) {
	controller.taskMaster.Assign(user)
	setJSONResponse(w, http.StatusOK, transformTaskForTransport(controller.taskMaster.PeekAssigned(user)))
}

func (controller *taskAssignmentsController) handlePost(w http.ResponseWriter, r *http.Request) {
	error := controller.taskMaster.MarkDone(user)
	if error != nil {
		setErrorResponse(w, &errorMessage{http.StatusBadRequest, "Unable to mark done", error.Error()})
	}
	w.WriteHeader(http.StatusOK)
}
