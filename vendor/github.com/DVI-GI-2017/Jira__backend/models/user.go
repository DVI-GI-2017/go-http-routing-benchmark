package models

import "crypto/sha256"

type User struct {
	Id       AutoId       `json:"_id" bson:"_id,omitempty"`
	Email    Email        `json:"email" bson:"email"`
	Password Password     `json:"password,omitempty" bson:"password"`
	Name     Name         `json:"name" bson:"name"`
	Bio      Text         `json:"bio,omitempty" bson:"bio,omitempty"`
	Projects []RequiredId `json:"projects,omitempty" bson:"projects,omitempty"`
}

// Returns validation error or nil if valid
func (u User) Validate() error {
	if err := u.Id.Validate(); err != nil {
		return err
	}
	if err := u.Email.Validate(); err != nil {
		return err
	}
	if err := u.Password.Validate(); err != nil {
		return err
	}
	if err := u.Name.Validate(); err != nil {
		return err
	}
	if err := u.Bio.Validate(); err != nil {
		return err
	}
	return nil
}

func (u *User) Encrypt() {
	hasher := sha256.New()
	hasher.Write([]byte(u.Password))

	u.Password = Password(hasher.Sum(nil))
}

type UsersList []User
