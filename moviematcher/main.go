package main

import (
	"errors"
	"fmt"
	"os"
)

type Category string

type User struct {
	Name               string
	PrefreedCategories []Category
}

type Movie struct {
	Title    string
	Category Category
}

type Recommender interface {
	recommend(user User) ([]Movie, error)
}

type CategoryRecommender struct {
	Movies []Movie
}

func createUser(name string, categories []Category) User {
	user := User{
		Name:               name,
		PrefreedCategories: categories,
	}
	return user
}

func (cr CategoryRecommender) recommend(user User) ([]Movie, error) {
	suggestions := []Movie{}
	for _, movie := range cr.Movies {
		for _, categrory := range user.PrefreedCategories {
			if movie.Category == categrory {
				suggestions = append(suggestions, movie)
			}
		}
	}
	if len(suggestions) == 0 {
		return nil, errors.New("no matching movies found for user" + user.Name)
	}
	return suggestions, nil
}

func (u User) printUserDetails() {
	fmt.Println(u.Name)
	fmt.Println("User prefereed categories")
	for _, category := range u.PrefreedCategories {
		fmt.Println(category)
	}
}

// &varaible we get the address of that variable
// * on a pointer we access the value at that address

func updateName(u *User, name string) {
	(*u).Name = name
}

func saveToFile(u User) error {
	// Format the user details as string
	// Name : Jack
	// Preferences : []
	data := fmt.Sprintf("Name: %s\nPreferences: %v\n", u.Name, u.PrefreedCategories)

	// Convert the string to a byte slice
	byteData := []byte(data)

	// Write to file (user.Name.txt)
	err := os.WriteFile(u.Name+".txt", byteData, 0644)
	if err != nil {
		return err
	}
	fmt.Println("User data save to ", u.Name+".txt")
	return nil
}

func main() {
	movies := []Movie{
		{"Inception", "Sci-Fi"},
		{"The Godfather", "Crime"},
		{"Interstellar", "Sci-Fi"},
		{"The Hangover", "Comedy"},
	}
	user := createUser("Jack", []Category{"Comedy", "Sci-Fi"})
	recommender := CategoryRecommender{Movies: movies}
	recs, err := recommender.recommend(user)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	for _, m := range recs {
		fmt.Println(m.Title)
	}
	updateName(&user, "Bob")
	saveToFile(user)
	fmt.Println(user.Name)
}
