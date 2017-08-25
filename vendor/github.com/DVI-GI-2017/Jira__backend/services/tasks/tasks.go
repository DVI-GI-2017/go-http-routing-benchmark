package tasks

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
	"gopkg.in/mgo.v2/bson"
)

const cTasks = "tasks"
const cProjects = "projects"

// Checks if task with this 'title == task.Title' exists.
func CheckTaskExists(source db.DataSource, task models.Task) (bool, error) {
	c, err := source.C(cTasks).Find(bson.M{"title": task.Title}).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check if task exists: %v", err)
	}
	return !c, nil
}

// Creates task and returns it.
func AddTaskToProject(source db.DataSource, task models.Task) (models.Task, error) {
	task.Id = models.NewAutoId()

	err := source.C(cTasks).Insert(task)
	if err != nil {
		return models.Task{}, fmt.Errorf("can not create task '%s': %v", task, err)
	}

	if err := pushTask(source, task.Id, task.ProjectId); err != nil {
		return models.Task{}, fmt.Errorf("can add task '%v' to project '%s': %v", task, task.ProjectId.Hex(), err)
	}
	return task, nil
}

// Pushes task to project's tasks array
func pushTask(source db.DataSource, taskId models.AutoId, projectId models.RequiredId) error {
	return source.C(cProjects).Update(
		bson.M{"_id": projectId},
		bson.M{"$push": bson.M{"tasks": taskId}},
	)
}

// Returns all tasks.
func AllTasks(source db.DataSource, projectId models.RequiredId) (tasksList models.TasksList, err error) {
	project, err := projects.FindProjectById(source, projectId)
	if err != nil {
		return models.TasksList{}, fmt.Errorf("can not find project with id %s: %v", projectId.Hex(), err)
	}

	err = source.C(cTasks).Find(bson.M{"": bson.M{"$id": project.Tasks}}).All(&tasksList)
	if err != nil {
		return models.TasksList{}, fmt.Errorf("can not retrieve all tasks: %v", err)
	}
	return tasksList, nil
}

// Returns task with given id
func FindTaskById(source db.DataSource, id models.RequiredId) (task models.Task, err error) {
	err = source.C(cTasks).FindId(id).One(&task)
	if err != nil {
		return models.Task{}, fmt.Errorf("can not find task with id '%s': %s",
			id.Hex(), err)
	}
	return task, nil
}
