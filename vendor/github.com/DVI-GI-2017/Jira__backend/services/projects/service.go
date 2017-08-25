package projects

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const cProjects = "projects"
const cUsers = "users"
const cTasks = "tasks"

// Check if project with title == project.Title exists
func CheckProjectExists(source db.DataSource, project models.Project) (bool, error) {
	empty, err := source.C(cProjects).Find(bson.M{"title": project.Title}).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check if project '%v' exists: %v", project, err)
	}
	return !empty, err
}

// Creates project and returns it.

func CreateProject(source db.DataSource, project models.Project) (models.Project, error) {
	project.Id = models.NewAutoId()

	err := source.C(cProjects).Insert(project)
	if err != nil {
		return models.Project{}, fmt.Errorf("can not create project '%v': %v", project, err)
	}
	return project, nil
}

// Returns all projects.
func AllProjects(source db.DataSource) (result models.ProjectsList, err error) {
	err = source.C(cProjects).Find(bson.M{}).All(&result)
	if err != nil {
		return models.ProjectsList{}, fmt.Errorf("can not retrieve all projects: %v", err)
	}
	return result, nil
}

// Returns task with given id.
func FindProjectById(mongo db.DataSource, id models.RequiredId) (result models.Project, err error) {
	err = mongo.C(cProjects).FindId(id).One(&result)
	if err != nil {
		return models.Project{}, fmt.Errorf("can not find project with id '%s': %s", id.Hex(), err)
	}
	return result, nil
}

// Returns all users in project
func AllUsersInProject(mongo db.DataSource, id models.RequiredId) (result models.UsersList, err error) {
	var project models.Project
	err = mongo.C(cProjects).FindId(id).One(&project)
	if err != nil {
		return models.UsersList{}, fmt.Errorf("can not find project with id '%s': %v", id, err)
	}

	err = mongo.C(cUsers).Find(bson.M{"_id": bson.M{"$in": project.Users}}).Select(bson.M{"password": 0}).All(&result)
	if err != nil {
		return models.UsersList{}, fmt.Errorf("can not retrieve all users from project: %s", id.Hex())
	}
	return result, nil
}

// Returns all tasks in project
func AllTasksInProject(mongo db.DataSource, id models.RequiredId) (result models.TasksList, err error) {
	err = mongo.C(cTasks).Find(bson.M{"project_id": id}).All(&result)
	if err != nil {
		fmt.Println(err)
		return models.TasksList{}, fmt.Errorf("can not retrieve all tasks from project: %s", id.Hex())
	}
	return result, nil
}

func AddUserToProject(source db.DataSource, projectId, userId models.RequiredId) (result models.UsersList, err error) {
	if err := pushUser(source, projectId, userId); err != nil {
		return models.UsersList{},
			fmt.Errorf("can not add user '%v' to project '%s': %v", userId, projectId, err)
	}
	if err := pushProject(source, userId, projectId); err != nil {
		return models.UsersList{},
			fmt.Errorf("can not add project '%v' to user '%s': %v", projectId.Hex(), userId.Hex(), err)
	}

	return AllUsersInProject(source, projectId)
}

// Pushes user to project's users array.
func pushUser(source db.DataSource, projectId, userId models.RequiredId) error {
	return source.C(cProjects).Update(
		bson.M{"_id": projectId},
		bson.M{"$push": bson.M{"users": userId}},
	)
}

// Pushes project to user's projects array
func pushProject(source db.DataSource, userId, projectId models.RequiredId) error {
	return source.C(cUsers).Update(
		bson.M{"_id": userId},
		bson.M{"$push": bson.M{"projects": projectId}},
	)
}

func DeleteUserFromProject(source db.DataSource, projectId, userId models.RequiredId) (result models.UsersList, err error) {
	if err := pullUser(source, projectId, userId); err != nil {
		return models.UsersList{},
			fmt.Errorf("can not delete user '%v' from project '%s': %v", userId.Hex(), projectId.Hex(), err)
	}
	if err := pullProject(source, userId, projectId); err != nil {
		return models.UsersList{},
			fmt.Errorf("can not delete project '%v' from user '%s': %v", projectId.Hex(), userId.Hex(), err)
	}

	return AllUsersInProject(source, projectId)
}

// Pulls user from project's users array.
func pullUser(source db.DataSource, projectId, userId models.RequiredId) error {
	return source.C(cProjects).Update(
		bson.M{"_id": projectId},
		bson.M{"$pull": bson.M{"users": userId}},
	)
}

// Pulls project from user's projects array.
func pullProject(source db.DataSource, userId, projectId models.RequiredId) error {
	return source.C(cUsers).Update(
		bson.M{"_id": userId},
		bson.M{"$pull": bson.M{"projects": projectId}},
	)
}

// Checks if user already in current project
func CheckUserInProject(source db.DataSource, userId, projectId models.RequiredId) (bool, error) {
	var project models.Project
	err := source.C(cProjects).FindId(projectId).One(&project)
	if err != nil {
		return false, fmt.Errorf("can not find project with id '%s': %v", projectId.Hex(), err)
	}
	for _, id := range project.Users {
		if id == userId {
			return true, nil
		}
	}

	return false, nil
}
