package models

import (
	"encoding/json"
	"errors"
	"regexp"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

// Email helpers

type Email string

var (
	ErrEmptyEmail       = errors.New("empty email")
	ErrWrongEmailFormat = errors.New("wrong email format")
)

//Long and strange regexp to validate email format.
var emailRegex = regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@“]+(\.[^<>()\[\]\\.,;:\s@“]+)*)|(“.+“))@((\[[0-9]{1,3}\.
[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

// Validates email
func (e Email) Validate() error {
	if len(e) == 0 {
		return ErrEmptyEmail
	}
	if !emailRegex.MatchString(string(e)) {
		return ErrWrongEmailFormat
	}
	return nil
}

// Password helpers

type Password string

var (
	ErrEmptyPassword       = errors.New("empty password")
	ErrWrongPasswordFormat = errors.New("wrong password format")
)

var passwordRegex = regexp.MustCompile(`^[[:graph:]]{5,14}$`)

// Validates password
func (p Password) Validate() error {
	if len(p) == 0 {
		return ErrEmptyPassword
	}
	if !passwordRegex.MatchString(string(p)) {
		return ErrWrongPasswordFormat
	}
	return nil
}

// Name

type Name string

var nameRegex = regexp.MustCompile(`^[a-zA-Z](.[ a-zA-Z0-9_-]*)$`)

var (
	ErrEmptyName       = errors.New("empty name")
	ErrWrongNameFormat = errors.New("wrong name format")
)

// Validates names
func (n Name) Validate() error {
	if len(n) == 0 {
		return ErrEmptyName
	}
	if !nameRegex.MatchString(string(n)) {
		return ErrWrongNameFormat
	}
	return nil
}

// Text

type Text string

var (
	ErrTextTooLong = errors.New("text too long")

	MaxTextLen = 500
)

// Validate text
func (t Text) Validate() error {
	if utf8.RuneCountInString(string(t)) > MaxTextLen {
		return ErrTextTooLong
	}
	return nil
}

// General Id helpers
type Id struct{ bson.ObjectId }

func (id Id) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Hex())
}

func (id *Id) UnmarshalJSON(data []byte) error {
	var result string
	err := json.Unmarshal(data, &result)
	id.ObjectId = bson.ObjectIdHex(result)
	return err
}

func (id Id) GetBSON() (interface{}, error) {
	return id.ObjectId, nil
}

func (id *Id) SetBSON(raw bson.Raw) error {
	return raw.Unmarshal(&id.ObjectId)
}

var ErrInvalidId = errors.New("invalid id")

// Validates id
func (id Id) Validate() error {
	// NOTE: By  default id.Valid() checks only id len
	// BTW we could pass id like: bson.ObjectId("12_bytes_len")
	if !bson.IsObjectIdHex(id.Hex()) {
		return ErrInvalidId
	}
	return nil
}

// AutoId helpers
type AutoId struct{ Id }

var ErrIdMustBeOmitted = errors.New("id must be omitted")

// Validates generated id
func (id AutoId) Validate() error {
	if id.Hex() != "" {
		return ErrIdMustBeOmitted
	}
	return nil
}

// New auto Id
func NewAutoId() AutoId {
	return AutoId{Id: Id{ObjectId: bson.NewObjectId()}}
}

// RequiredId helpers
type RequiredId struct{ Id }

var ErrIdMustBePresent = errors.New("id must be present")

// Validates required id
func (id RequiredId) Validate() error {
	if id.Hex() == "" {
		return ErrIdMustBePresent
	}

	if err := id.Id.Validate(); err != nil {
		return err
	}
	return nil
}

// New required id
func NewRequiredId(hex string) RequiredId {
	return RequiredId{Id: Id{ObjectId: bson.ObjectIdHex(hex)}}
}

// Optional Id helpers
type OptionalId struct{ Id }

// Validates optional id
func (id OptionalId) Validate() error {
	if id.Hex() == "" {
		return nil
	}

	if err := id.Id.Validate(); err != nil {
		return err
	}
	return nil
}
