package main

import (
	"GithubSearch/config"
	"GithubSearch/github"
	v1 "GithubSearch/handler/v1"
	"GithubSearch/routes"
	"GithubSearch/utils"
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/flannel-dev-lab/cyclops/v2"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	"os"
)

func main() {
	ctx := context.Background()
	ctx = logger.AddKey(ctx, "scope", utils.GetFuncName())

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		logger.Error(ctx, "error parsing configuration", err)
		os.Exit(1)
	}

	githubClient := github.New(cfg.Username, cfg.Password)

	handler := v1.Handler{
		Cfg:          &cfg,
		GithubClient: githubClient,
	}

	routerObj := routes.GetRoutes(&handler)

	cyclops.StartServer(":8081", routerObj)
}
