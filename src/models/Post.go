package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorID,omitempty"`
	AuthorNickname string    `json:"authorNickname,omitempty"`
	UpVotes        uint64    `json:"upVotes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

func (post *Post) PrepareForCreation() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("Title is mandatory")
	}

	if post.Content == "" {
		return errors.New("Content is mandatory")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
