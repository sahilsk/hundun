package schema

import (
	"encoding/json"
	"log"
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

func (n *Notes) ToPrettyString() string {
	b, err := json.MarshalIndent(*n, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (n *Notes) ToString() string {
	b, err := json.Marshal(*n)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (n *Note) ToPrettyString() string {
	b, err := json.MarshalIndent(*n, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (n *Note) ToString() string {
	b, err := json.Marshal(*n)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
