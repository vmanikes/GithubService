// Package some_service
package some_service

import (
	"GithubSearch/types/v1/api"
	"GithubSearch/utils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	"net/http"
	"time"
)

const hostname = "http://localhost"

type SomeService interface {
	ResultParser(ctx context.Context, searchResponse *api.SearchResponse) error
}

type Client struct {
	httpClient *http.Client
}

func New() SomeService {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// ResultParser parses the search response and returns an error if exists
func (c *Client) ResultParser(ctx context.Context, searchResponse *api.SearchResponse) error {
	ctx = logger.AddKey(ctx, "scope", utils.GetFuncName())

	bodyBytes, err := json.Marshal(searchResponse)
	if err != nil {
		logger.Error(ctx, "unable to marshal request body", err)
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, hostname+"/submit", bytes.NewBuffer(bodyBytes))
	if err != nil {
		logger.Error(ctx, "unable to create request", err)
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		logger.Error(ctx, "unable to send request", err)
		return err
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			logger.Error(ctx, "unable to close body", err)
		}
	}()

	if response.StatusCode != http.StatusCreated {
		// TODO some parsing

		return errors.New("service returned error")
	}

	return nil
}
