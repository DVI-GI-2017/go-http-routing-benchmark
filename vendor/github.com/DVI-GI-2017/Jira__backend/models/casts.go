package models

import "fmt"

// Helper to generate error message about bad cast.
func ErrInvalidCast(got, expected interface{}) error {
	return fmt.Errorf("can not cast input data with type: %T to %T", got, expected)
}

func SafeCastToTaskLabel(data interface{}) (TaskLabel, error) {
	if val, ok := data.(TaskLabel); ok {
		return val, nil
	}
	return TaskLabel{}, ErrInvalidCast(data, TaskLabel{})
}

func SafeCastToRequiredId(data interface{}) (RequiredId, error) {
	if val, ok := data.(RequiredId); ok {
		return val, nil
	}
	return RequiredId{}, ErrInvalidCast(data, RequiredId{})
}

func SafeCastToProject(data interface{}) (Project, error) {
	if val, ok := data.(Project); ok {
		return val, nil
	}
	return Project{}, ErrInvalidCast(data, Project{})
}

func SafeCastToProjectUser(data interface{}) (ProjectUser, error) {
	if val, ok := data.(ProjectUser); ok {
		return val, nil
	}
	return ProjectUser{}, ErrInvalidCast(data, ProjectUser{})
}

func SafeCastToTask(data interface{}) (Task, error) {
	if val, ok := data.(Task); ok {
		return val, nil
	}
	return Task{}, ErrInvalidCast(data, Task{})
}

func SafeCastToUser(data interface{}) (User, error) {
	if val, ok := data.(User); ok {
		return val, nil
	}
	return User{}, ErrInvalidCast(data, User{})
}

func SafeCastToEmail(data interface{}) (Email, error) {
	if val, ok := data.(Email); ok {
		return val, nil
	}
	return Email(""), ErrInvalidCast(data, Email(""))
}
