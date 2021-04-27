package v1

import (
	"GithubSearch/types/v1/api"
	"github.com/flannel-dev-lab/cyclops/v2/response"
	"net/http"
)

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("search_term")
	user := r.URL.Query().Get("user")

	ctx := r.Context()

	if searchTerm == "" || user == "" {
		response.ErrorResponse(http.StatusBadRequest, "search term or user is empty", w)
		return
	}

	githubSearchResponse, err := h.GithubClient.Search(ctx, user, searchTerm)
	if err != nil {
		response.ErrorResponse(http.StatusInternalServerError, "something is wrong", w)
		return
	}

	var searchResponse api.SearchResponse

	results := make([]api.Result, len(githubSearchResponse.Items))

	for idx, searchResult := range githubSearchResponse.Items {
		var result api.Result

		result.Repo = searchResult.Repository.HTMLURL
		result.FileUrl = searchResult.HTMLURL

		results[idx] = result
	}

	searchResponse.Results = results

	err = h.SomeService.ResultParser(ctx, &searchResponse)
	if err != nil {
		response.ErrorResponse(http.StatusInternalServerError, "something is wrong", w)
		return
	}

	response.SuccessResponse(http.StatusOK, w, results)
	return
}
