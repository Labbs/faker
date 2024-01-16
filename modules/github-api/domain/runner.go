package domain

type GithubActionsRunner struct {
	Id     int           `json:"id"`
	Name   string        `json:"name"`
	Os     string        `json:"os"`
	Status string        `json:"status"`
	Busy   bool          `json:"busy"`
	Labels []GithubLabel `json:"labels"`
}
