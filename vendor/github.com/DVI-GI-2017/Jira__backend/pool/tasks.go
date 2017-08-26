package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

func init() {
	resolvers["Task"] = tasksResolver
}

const (
	TaskCreate        = Action("TaskCreate")
	TasksAllOnProject = Action("TasksAllOnProject")
	TaskFindById      = Action("TaskFindById")
)

func tasksResolver(action Action) (service ServiceFunc, err error) {
	switch action {

	case TaskCreate:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			task, err := models.SafeCastToTask(data)
			if err != nil {
				return models.Task{}, err
			}
			return services.AddTaskToProject(source, task)
		}
		return

	case TasksAllOnProject:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.TasksList{}, err
			}
			return services.AllTasks(source, id)
		}
		return

	case TaskFindById:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.Task{}, err
			}
			return services.FindTaskById(source, id)
		}
		return
	}
	return nil, fmt.Errorf("can not find resolver with action: %v, in tasks resolvers", action)

}
