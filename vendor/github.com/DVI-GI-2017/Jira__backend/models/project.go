package models

type Project struct {
	Id          AutoId       `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       Name         `json:"title" bson:"title"`
	Description Text         `json:"description,omitempty" bson:"description,omitempty"`
	Users       []RequiredId `json:"users,omitempty" bson:"users,omitempty"`
	Tasks       []RequiredId `json:"tasks,omitempty" bson:"tasks,omitempty"`
}

// Helper for both project and user ids.
type ProjectUser struct {
	ProjectId RequiredId
	UserId    RequiredId
}

type ProjectsList []Project

// Validates project.
func (p Project) Validate() error {
	if err := p.Id.Validate(); err != nil {
		return err
	}
	if err := p.Title.Validate(); err != nil {
		return err
	}
	if err := p.Description.Validate(); err != nil {
		return err
	}
	return nil
}
