package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saharshsingh/gophinoff/taskmaster/impl"
)

//-----------
// task model

type task struct {
	Name        string
	Description string
	Priority    int
}

func transformTaskForTransport(rawTask *impl.Task) *task {
	if rawTask == nil {
		return nil
	}
	return &task{rawTask.Name, rawTask.Description, rawTask.Priority()}
}

//--------------------
// error message model

type errorMessage struct {
	StatusCode int
	Message    string
	Detail     string
}

func setErrorResponse(w http.ResponseWriter, errorMsg *errorMessage) {
	fmt.Printf("ErrorMessage: '%v'\n", errorMsg)
	setJSONResponse(w, errorMsg.StatusCode, errorMsg)
}

//----------------------------------------
// helpful request/response util functions

func setJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) {

	bodyJSON, marshalError := json.Marshal(body)
	if marshalError != nil {
		setErrorResponse(w,
			&errorMessage{http.StatusInternalServerError, "Internal Server Error", marshalError.Error()})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bodyJSON)
}

func parseJSONRequest(r *http.Request, value interface{}, w http.ResponseWriter) error {

	decoder := json.NewDecoder(r.Body)

	unmarshalError := decoder.Decode(value)
	if unmarshalError != nil {
		setErrorResponse(w, &errorMessage{http.StatusBadRequest, "Bad Request", "Malformed JSON body"})
		return unmarshalError
	}

	return nil
}
