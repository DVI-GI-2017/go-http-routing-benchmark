package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
)

// Adds task to project
// Path param - "project_id"
// Post body - task
// Returns created task if OK
func AddTaskToProject(w http.ResponseWriter, req *http.Request) {
	params := mux.Params(req)

	projectId := models.NewRequiredId(params.PathParams["project_id"])

	var task models.Task
	if err := json.Unmarshal(params.Body, &task); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	task.ProjectId = projectId

	if err := task.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	task.Id = models.NewAutoId()

	newTask, err := pool.Dispatch(pool.TaskCreate, task)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadGateway)
		return
	}

	JsonResponse(w, newTask)
}

// Returns all tasks
// Path param - "project_id"
func AllTasksInProject(w http.ResponseWriter, req *http.Request) {
	projectId := models.NewRequiredId(mux.Params(req).PathParams["project_id"])

	tasks, err := pool.Dispatch(pool.TasksAllOnProject, projectId)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, tasks.(models.TasksList))
}

// Returns task with given id
// Path params: "id" - task id.
func GetTaskById(w http.ResponseWriter, req *http.Request) {

	id := models.NewRequiredId(mux.Params(req).PathParams["id"])

	task, err := pool.Dispatch(pool.TaskFindById, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, task.(models.Task))
	return
}
