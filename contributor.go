package github

// Contributor summarizes one person's contributions to a particular
// Github repository.
type Contributor struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}
