package models

type Task struct {
	Id          AutoId     `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       Name       `json:"title" bson:"title"`
	ProjectId   RequiredId `json:"project_id" bson:"project_id"`
	Description Text       `json:"description,omitempty" bson:"description,omitempty"`
	InitiatorId RequiredId `json:"initiator_id" bson:"initiator_id"`
	AssigneeId  OptionalId `json:"assignee_id,omitempty" bson:"assignee_id,omitempty"`
	Labels      LabelsList `json:"labels,omitempty" bson:"labels,omitempty"`
}

// Validates Task
func (t Task) Validate() error {
	if err := t.Id.Validate(); err != nil {
		return err
	}
	if err := t.Title.Validate(); err != nil {
		return err
	}
	if err := t.ProjectId.Validate(); err != nil {
		return err
	}
	if err := t.Description.Validate(); err != nil {
		return err
	}
	if err := t.InitiatorId.Validate(); err != nil {
		return err
	}
	if err := t.AssigneeId.Validate(); err != nil {
		return err
	}
	return nil
}

type TasksList []Task

type TaskLabel struct {
	TaskId RequiredId
	Label  Label
}
