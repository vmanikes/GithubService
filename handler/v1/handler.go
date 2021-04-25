package v1

import (
	"GithubSearch/config"
	"GithubSearch/github"
)

// Handler struct containing config
type Handler struct {
	Cfg          *config.Config
	GithubClient github.Github
}
