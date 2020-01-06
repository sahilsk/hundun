package schema

import "encoding/json"

type PriorityResponse struct {
	Priority Priority `json:"priority"`
}

type Priorities struct {
	Priorities []Priority `mapstructure:"priorities"`
	Pagination
}

type Priority struct {
	Entity
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       uint   `json:"order`
	Color       string `json:"color"`
}

func (ir *Priorities) ToPrettyString() ([]byte, error) {
	return json.MarshalIndent(*ir, "", "  ")
}

func (ir *Priorities) ToString() ([]byte, error) {
	return json.Marshal(*ir)
}
