package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
)

func init() {
	resolvers["Label"] = labelsResolver
}

const (
	LabelAddToTask      = Action("LabelAddToTask")
	LabelsAllOnTask     = Action("LabelsAllOnTask")
	LabelAlreadySet     = Action("LabelAlreadySet")
	LabelDeleteFromTask = Action("LabelDeleteFromTask")
)

func labelsResolver(action Action) (service ServiceFunc) {
	switch action {

	case LabelAddToTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.AddLabelToTask(source, taskLabel.TaskId, taskLabel.Label)
		}
		return

	case LabelsAllOnTask:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return tasks.AllLabels(source, id.(models.RequiredId))
		}
		return

	case LabelAlreadySet:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.CheckLabelAlreadySet(source, taskLabel.TaskId, taskLabel.Label)
		}
		return
	case LabelDeleteFromTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.DeleteLabelFromTask(source, taskLabel.TaskId, taskLabel.Label)
		}
		return

	default:
		log.Panicf("can not find resolver with action: %v, in labels resolvers", action)
		return
	}
}
