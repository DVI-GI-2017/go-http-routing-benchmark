package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

func init() {
	resolvers["Label"] = labelsResolver
}

const (
	LabelAddToTask      = Action("LabelAddToTask")
	LabelsAllOnTask     = Action("LabelsAllOnTask")
	LabelDeleteFromTask = Action("LabelDeleteFromTask")
)

func labelsResolver(action Action) (service ServiceFunc, err error) {
	switch action {

	case LabelAddToTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel, err := models.SafeCastToTaskLabel(data)
			if err != nil {
				return models.LabelsList{}, err
			}
			return services.AddLabelToTask(source, taskLabel.TaskId, taskLabel.Label)
		}
		return

	case LabelsAllOnTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.LabelsList{}, err
			}
			return services.AllLabels(source, id)
		}
		return

	case LabelDeleteFromTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel, err := models.SafeCastToTaskLabel(data)
			if err != nil {
				return models.LabelsList{}, err
			}
			return services.DeleteLabelFromTask(source, taskLabel.TaskId, taskLabel.Label)
		}
		return
	}
	return nil, fmt.Errorf("can not find resolver with action: %v, in labels resolvers", action)
}
