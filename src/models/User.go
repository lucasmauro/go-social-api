package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) PrepareForCreation() error {
	return user.prepare(user.validateForCreation)
}

func (user *User) PrepareForUpdate() error {
	return user.prepare(user.validateForUpdate)
}

func (user *User) prepare(validate func() error) error {
	if err := validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validateForCreation() error {
	if err := user.validateBasicFields(); err != nil {
		return err
	}

	if user.Password == "" {
		return errors.New("Password is mandatory")
	}

	return nil
}

func (user *User) validateForUpdate() error {
	return user.validateBasicFields()
}

func (user *User) validateBasicFields() error {
	if user.Name == "" {
		return errors.New("Name is mandatory")
	}

	if user.Nickname == "" {
		return errors.New("Nickame is mandatory")
	}

	if user.Email == "" {
		return errors.New("Email is mandatory")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
}
