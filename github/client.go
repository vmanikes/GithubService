package github

import (
	"GithubSearch/utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flannel-dev-lab/cyclops/v2/logger"
	"net/http"
	"time"
)

const (
	host          = "https://api.github.com"
	versionHeader = "application/vnd.github.v3+json"
	timeout       = 10
)

type Github interface {
	Search(ctx context.Context, user, term string) (*SearchResponse, error)
}

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func New(username, password string) (Github, error) {
	if username == "" || password == "" {
		return nil, errors.New("github username or password should not be empty")
	}
	return &Client{
		apiKey: base64.RawURLEncoding.EncodeToString([]byte(username + ":" + password)),
		httpClient: &http.Client{
			Timeout: timeout * time.Second,
		},
	}, nil
}

// Search takes in a user and term and searches github for those values
func (c *Client) Search(ctx context.Context, user, term string) (*SearchResponse, error) {
	ctx = logger.AddKey(ctx, "scope", utils.GetFuncName())

	request, err := http.NewRequest(http.MethodGet, host+"/search/code", nil)
	if err != nil {
		logger.Error(ctx, "unable to create request", err)
		return nil, err
	}

	request.Header.Set("Accept", versionHeader)
	request.Header.Set("Authorization", "Basic "+c.apiKey)

	q := request.URL.Query()
	q.Add("q", fmt.Sprintf("%s user:%s", term, user))
	request.URL.RawQuery = q.Encode()

	response, err := c.httpClient.Do(request)
	if err != nil {
		logger.Error(ctx, "unable to perform request", err)
		return nil, err
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			logger.Error(ctx, "unable to perform request", err)
		}
	}()

	if response.StatusCode == http.StatusOK {
		var errResponse ErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			logger.Error(ctx, "unable to unmarshal error response", err)
			return nil, err
		}

		var message string

		for _, errorMessage := range errResponse.Errors {
			message += errorMessage.Field + " " + errorMessage.Code
		}

		return nil, errors.New(message)
	}


	var searchResponse SearchResponse

	if err := json.NewDecoder(response.Body).Decode(&searchResponse); err != nil {
		logger.Error(ctx, "unable to unmarshal response", err)
		return nil, err
	}

	return &searchResponse, nil
}
