package schema

import (
	"encoding/json"
	"time"
)

type Notes struct {
	Notes []Note `json:"notes"`
}

type Note struct {
	Id         string    `json:"id"`
	User       User      `json:"user"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

type User struct {
	Entity
}

func (n *Notes) ToPrettyString() ([]byte, error) {
	return json.MarshalIndent(*n, "", "  ")
}

func (n *Notes) ToString() ([]byte, error) {
	return json.Marshal(*n)
}
