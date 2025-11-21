package utils

import (
	"math/rand"
)

// funnyWaitingMessages is a list of funny short strings to display while waiting.
var funnyWaitingMessages = []string{
	"Hitchhiking through the cosmos...",
	"Polishing my thinking cap...",
	"Consulting the digital oracle...",
	"Bribing the electrons with coffee...",
	"Untangling the quantum spaghetti...",
	"Summoning the data spirits...",
	"Warming up the neural networks...",
	"Just a moment, the hamsters are running...",
	"Percolating thoughts into brilliance...",
	"Searching for meaning in the matrix...",
}

// GetRandomWaitingMessage returns a random funny waiting message.
func GetRandomWaitingMessage() string {
	return funnyWaitingMessages[rand.Intn(len(funnyWaitingMessages))]
}
