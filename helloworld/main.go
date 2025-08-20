package main

import "fmt"

const pi = 3.14159

type Username string

type Age int

func greet(name string) string {
	return "Hello " + name
}

func divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	scores := map[string]int{
		"Bob":   90,
		"Jack":  85,
		"Carol": 92,
	}
	scores["Dave"] = 88
	scores["Bob"] = 95
	delete(scores, "Bob")
	fmt.Println("Scores", scores)
}
