package services

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const cUsers = "users"

// Checks if user with this credentials.Email exists.
func CheckUserExists(source db.DataSource, credentials models.User) (bool, error) {
	empty, err := source.C(cUsers).Find(bson.M{"email": credentials.Email}).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check if user with credentials '%v' exists: %v", credentials, err)
	}
	return !empty, nil
}

// Checks if user credentials present in users collection.
func AuthorizeUser(source db.DataSource, credentials models.User) (user models.User, err error) {
	err = source.C(cUsers).Find(bson.M{
		"email":    credentials.Email,
		"password": credentials.Password,
	}).One(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not check user credentials '%v': %v", credentials, err)
	}
	return user, nil
}

// Creates user and returns it.
func CreateUser(source db.DataSource, user models.User) (models.User, error) {
	if exists, err := CheckUserExists(source, user); err != nil {
		return models.User{}, err
	} else if exists {
		return models.User{}, fmt.Errorf("user %v already exists", user)
	}

	user.Id = models.NewAutoId()

	err := source.C(cUsers).Insert(user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not create user '%v': %v", user, err)
	}

	user.Password = ""

	return user, nil
}

// Returns all users.
func AllUsers(source db.DataSource) (usersLists models.UsersList, err error) {
	err = source.C(cUsers).Find(nil).Select(bson.M{"password": 0}).All(&usersLists)
	if err != nil {
		return models.UsersList{}, fmt.Errorf("can not retrieve all users: %v", err)
	}
	return usersLists, nil
}

// Returns user with given id.
func FindUserById(source db.DataSource, id models.RequiredId) (user models.User, err error) {
	err = source.C(cUsers).FindId(id).Select(bson.M{"password": 0}).One(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not find user with id '%s': %v", id, err)
	}
	return user, nil
}

// Returns user with given email.
func FindUserByEmail(source db.DataSource, email models.Email) (user models.User, err error) {
	err = source.C(cUsers).Find(bson.M{"email": email}).Select(bson.M{"password": 0}).One(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not find user with email '%s': %v", email, err)
	}
	return user, nil
}

// Returns all users project.
func AllUserProjects(source db.DataSource, id models.RequiredId) (projects models.ProjectsList, err error) {
	var user models.User
	err = source.C(cUsers).FindId(id).One(&user)
	if err != nil {
		return models.ProjectsList{}, fmt.Errorf("can not find user with id '%s': %v", id, err)
	}

	err = source.C(cProjects).Find(bson.M{"_id": bson.M{"$in": user.Projects}}).All(&projects)
	if err != nil {
		return models.ProjectsList{}, fmt.Errorf("can not retrieve all users from project: %s", id.Hex())
	}
	return projects, nil
}
