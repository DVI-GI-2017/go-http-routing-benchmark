package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
)

// Returns all labels from task
// Path parameter: "task_id" - task id.
func AllLabelsOnTask(w http.ResponseWriter, req *http.Request) {
	id := models.NewRequiredId(mux.Params(req).PathParams["task_id"])

	labels, err := pool.Dispatch(pool.LabelsAllOnTask, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, labels)
}

// Adds label to task.
// Query parameter: "task_id" - task id.
// Post body - label.
func AddLabelToTask(w http.ResponseWriter, req *http.Request) {
	params := mux.Params(req)

	var label models.Label
	err := json.Unmarshal(params.Body, &label)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = label.Validate()
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	taskId := models.NewRequiredId(params.PathParams["task_id"])
	taskLabel := models.TaskLabel{TaskId: taskId, Label: label}

	exists, err := pool.Dispatch(pool.LabelAlreadySet, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("label '%v' already set on project '%s'", label, taskId.Hex()),
			http.StatusConflict)
		return
	}

	labels, err := pool.Dispatch(pool.LabelAddToTask, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, labels)
}

// Deletes label from task and returns new labels
// Path parameter: "task_id" - task id.
// Post body - label
func DeleteLabelFromTask(w http.ResponseWriter, req *http.Request) {
	params := mux.Params(req)

	var label models.Label
	err := json.Unmarshal(params.Body, &label)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	taskId := models.NewRequiredId(params.PathParams["task_id"])

	taskLabel := models.TaskLabel{TaskId: taskId, Label: label}

	labels, err := pool.Dispatch(pool.LabelDeleteFromTask, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, labels)
}
