package main

import "testing"

func TestRecommend(t *testing.T) {
	movies := []Movie{
		{"Inception", "Sci-Fi"},
		{"The Godfather", "Crime"},
		{"Interstellar", "Sci-Fi"},
		{"The Hangover", "Comedy"},
	}
	rc := CategoryRecommender{Movies: movies}
	user := createUser("Jack", []Category{"Comedy", "Sci-Fi"})
	suggestions, err := rc.recommend(user)
	if err != nil {
		t.Error("Expected no error but got ", err)
	}
	if len(suggestions) == 0 {
		t.Error("Expected at least one recommendation, got 0")
	}
}

func TestRecommend_NoMatches(t *testing.T) {
	movies := []Movie{
		{"Inception", "Sci-Fi"},
		{"The Godfather", "Crime"},
		{"Interstellar", "Sci-Fi"},
		{"The Hangover", "Comedy"},
	}
	rc := CategoryRecommender{Movies: movies}
	user := createUser("Jack", []Category{"Horror"})
	suggestions, err := rc.recommend(user)
	if err == nil {
		t.Error("Expected an error when no recommendations are found")
	}
	if len(suggestions) != 0 {
		t.Error("Expected 0 recommendations , but got", len(suggestions))
	}
}
