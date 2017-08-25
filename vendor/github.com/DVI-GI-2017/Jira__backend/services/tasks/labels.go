package tasks

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

// Returns all labels from given task.
func AllLabels(source db.DataSource, taskId models.RequiredId) (models.LabelsList, error) {
	var container struct {
		models.LabelsList `bson:"labels"`
	}

	err := queryLabels(source, taskId).One(&container)
	if err != nil {
		return models.LabelsList{},
			fmt.Errorf("can not retrieve all labels on task %s: %v", taskId.Hex(), err)
	}

	return container.LabelsList, nil
}

// Selects labels from tasks query
func queryLabels(source db.DataSource, taskId models.RequiredId) db.Query {
	return source.C(cTasks).Find(
		bson.M{
			"_id": taskId,
		}).Select(bson.M{"labels": 1})
}

// Checks if label already set on this task.
func CheckLabelAlreadySet(source db.DataSource, id models.RequiredId, label models.Label) (bool, error) {
	notset, err := queryLabel(source, id, label).IsEmpty()
	if err != nil {
		return false, err
	}

	return !notset, nil
}

// Selects label from collection.
func queryLabel(source db.DataSource, taskId models.RequiredId, label models.Label) db.Query {
	return source.C(cTasks).Find(
		bson.M{
			"_id":    taskId,
			"labels": label,
		}).Select(bson.M{"labels": 1})
}

// Adds label to task and returns new list of labels on this task.
func AddLabelToTask(source db.DataSource, taskId models.RequiredId, label models.Label) (models.LabelsList, error) {
	err := pushLabel(source, taskId, label)
	if err != nil {
		return models.LabelsList{},
			fmt.Errorf("can not add label '%v' to task '%s': %v", label, taskId.Hex(), err)
	}

	return AllLabels(source, taskId)
}

// Pushes label in task's labels array.
func pushLabel(source db.DataSource, taskId models.RequiredId, label models.Label) error {
	return source.C(cTasks).Update(
		bson.M{"_id": taskId},
		bson.M{"$push": bson.M{"labels": label}},
	)
}

// Deletes label from task and returns new list of labels on this task
func DeleteLabelFromTask(source db.DataSource, taskId models.RequiredId, label models.Label) (models.LabelsList, error) {
	err := pullLabel(source, taskId, label)
	if err != nil {
		return models.LabelsList{},
			fmt.Errorf("can not delete label '%v' from task '%s': %v", label, taskId.Hex(), err)
	}

	return AllLabels(source, taskId)
}

// Pulls label from task's labels array.
func pullLabel(source db.DataSource, taskId models.RequiredId, label models.Label) error {
	return source.C(cTasks).Update(
		bson.M{"_id": taskId},
		bson.M{"$pull": bson.M{"labels": label}},
	)
}
