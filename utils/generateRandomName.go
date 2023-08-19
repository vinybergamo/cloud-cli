package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomName() string {
	adjectives := []string{
		"Small", "Tiny", "Quick", "Fast", "Brave", "Clever", "Happy",
		"Funny", "Smart", "Kind", "Lucky", "Strong", "Gentle", "Crazy",
	}

	nouns := []string{
		"Dog", "Cat", "Bird", "Fish", "Fox", "Bear", "Lion", "Tiger",
		"Star", "Moon", "Sun", "Tree", "Mountain", "River", "Ocean",
	}

	rand.Seed(time.Now().UnixNano())
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	hash := rand.Intn(1000000)

	return fmt.Sprintf("%s-%s-%06d", adjective, noun, hash)
}
