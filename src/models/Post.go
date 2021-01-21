package models

import "time"

type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorID,omitempty"`
	AuthorNickname uint64    `json:"authorNickname,omitempty"`
	UpVotes        uint64    `json:"upVotes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}
