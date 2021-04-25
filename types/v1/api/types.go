// Package api contains the types for request, responses from the API handlers
package api

type Result struct {
	FileUrl string `json:"file_url"`
	Repo    string `json:"repo"`
}

type SearchResponse struct {
	Results []Result `json:"results"`
}
