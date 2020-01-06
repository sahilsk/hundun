package schema

type Entity struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HtmlUrl string `json:"html_url"`
}

type Pagination struct {
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
	Total  uint `json:"total"`
	More   bool `json:"more"`
}
