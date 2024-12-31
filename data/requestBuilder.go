package data

import "github.com/swldmnn/presto/model"

func BuildRequests() []model.HttpRequest {
	requests := []model.HttpRequest{
		model.NewHttpRequest("GET", "https://api.chucknorris.io/jokes/random"),
		model.NewHttpRequest("GET", "https://uselessfacts.jsph.pl/api/v2/facts/random"),
	}

	return requests
}
