package data

import "github.com/swldmnn/presto/model"

func BuildRequests() []model.HttpRequest {
	requests := []model.HttpRequest{
		model.NewHttpRequest("GET", "https://api.chucknorris.io/jokes/random"),
		model.NewHttpRequest("GET", "https://uselessfacts.jsph.pl/api/v2/facts/random"),
		model.NewHttpRequest("GET", "https://techy-api.vercel.app/api/json"),
		model.NewHttpRequest("GET", "https://corporatebs-generator.sameerkumar.website/"),
		model.NewHttpRequest("GET", "https://whoa.onrender.com/whoas/random"),
		model.NewHttpRequest("GET", "https://v2.jokeapi.dev/joke/Any"),
		model.NewHttpRequest("GET", "https://official-joke-api.appspot.com/random_joke"),
	}

	return requests
}
