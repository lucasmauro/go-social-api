package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
	err := user.prepare(user.validateForCreation)
	if err != nil {
		return err
	}

	return user.formatPassword()
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

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email format is invalid")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) formatPassword() error {
	hash, err := security.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}
