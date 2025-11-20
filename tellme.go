package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

// Response defines the structure of the answer
type Response struct {
	Answer   string   `json:"answer" jsonschema:"description=A short paragraph answering the question"`
	Examples []string `json:"examples" jsonschema:"description=A couple of examples"`
}

func main() {
	if os.Getenv("GEMINI_API_KEY") == "" {
		fmt.Println("GEMINI_API_KEY environment variable not set. Please set it to your Gemini API key.")
		return
	}

	// If no arguments are provided, print a message and exit
	if len(os.Args) < 2 {
		fmt.Println("You were asking... what?")
		return
	}
	query := os.Args[1]

	ctx := context.Background()

	// Initialize Genkit with the Google AI plugin
	// The API key is expected to be in the GEMINI_API_KEY environment variable
	g := genkit.Init(ctx, genkit.WithPlugins(&googlegenai.GoogleAI{}))

	// Get the current shell
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "unknown shell"
	}

	prompt := fmt.Sprintf("Context: The user is asking a programming question from within the %s shell. Give a short answer and 1 or 2 examples.\nQuestion: %s", shell, query)

	// Generate structured data
	response, err := genkit.Generate(ctx, g,
		ai.WithPrompt(prompt),
		ai.WithModelName("googleai/gemini-2.5-flash"),
	)
	if err != nil {
		log.Fatalf("failed to generate response: %v", err)
	}

	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		// glamour.WithWordWrap(40),
	)

	fmt.Println(r.Render(response.Text()))
}
