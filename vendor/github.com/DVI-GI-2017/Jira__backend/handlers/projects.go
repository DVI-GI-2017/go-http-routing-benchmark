package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
)

// Creates project
// Post body - project
// Returns created project if OK
func CreateProject(w http.ResponseWriter, req *http.Request) {
	var projectInfo models.Project

	body := mux.Params(req).Body

	err := json.Unmarshal(body, &projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := projectInfo.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	projectInfo.Id = models.NewAutoId()

	exists, err := pool.Dispatch(pool.ProjectExists, projectInfo)

	if err != nil {
		JsonErrorResponse(w, fmt.Errorf("can not check project existence: %v", err),
			http.StatusInternalServerError)
		return
	}

	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("project with title %s already exists", projectInfo.Title),
			http.StatusConflict)
		return
	}

	project, err := pool.Dispatch(pool.ProjectCreate, projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadGateway)
		return
	}

	JsonResponse(w, project)
}

// Returns all projects
func AllProjects(w http.ResponseWriter, _ *http.Request) {
	projects, err := pool.Dispatch(pool.ProjectsAll, nil)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, projects.(models.ProjectsList))
}

// Returns project with given id
// Query param: "id" - project id
func GetProjectById(w http.ResponseWriter, req *http.Request) {
	id := models.NewRequiredId(mux.Params(req).PathParams["id"])

	user, err := pool.Dispatch(pool.ProjectFindById, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user.(models.Project))
	return
}

func GetAllUsersFromProject(w http.ResponseWriter, req *http.Request) {
	id := models.NewRequiredId(mux.Params(req).PathParams["id"])

	user, err := pool.Dispatch(pool.ProjectAllUsers, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user.(models.UsersList))
}

func GetAllTasksFromProject(w http.ResponseWriter, req *http.Request) {
	id := models.NewRequiredId(mux.Params(req).PathParams["id"])

	tasks, err := pool.Dispatch(pool.ProjectAllTasks, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, tasks.(models.TasksList))
}

func AddUserToProject(w http.ResponseWriter, req *http.Request) {
	projectId := models.NewRequiredId(mux.Params(req).PathParams["id"])

	var userId models.RequiredId
	err := json.Unmarshal(mux.Params(req).Body, &userId)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	exists, err := pool.Dispatch(pool.ProjectUserExists,
		models.ProjectUser{
			ProjectId: projectId,
			UserId:    userId,
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("user '%s' already in project '%s'", userId.Hex(), projectId.Hex()),
			http.StatusConflict)
		return
	}

	users, err := pool.Dispatch(pool.ProjectAddUser,
		models.ProjectUser{
			ProjectId: projectId,
			UserId:    userId,
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users.(models.UsersList))
}

func DeleteUserFromProject(w http.ResponseWriter, req *http.Request) {
	projectId := models.NewRequiredId(mux.Params(req).PathParams["id"])

	var userId models.RequiredId
	err := json.Unmarshal(mux.Params(req).Body, &userId)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	exists, err := pool.Dispatch(pool.ProjectUserExists,
		models.ProjectUser{
			ProjectId: projectId,
			UserId:    userId,
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if !exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("user '%s' not in project '%s'", userId.Hex(), projectId.Hex()),
			http.StatusNotFound)
		return
	}

	user, err := pool.Dispatch(pool.ProjectDeleteUser,
		models.ProjectUser{
			ProjectId: projectId,
			UserId:    userId,
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user.(models.UsersList))
}
