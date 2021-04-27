package v1

import (
	"GithubSearch/config"
	"GithubSearch/github"
	some_service "GithubSearch/some-service"
)

// Handler struct containing config
type Handler struct {
	Cfg          *config.Config
	GithubClient github.Github
	SomeService  some_service.SomeService
}
