package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh/spinner"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

type Interaction struct {
	prompt   string
	response string
}

// Global interaction variable to more easily enable spinner ui
var interaction Interaction

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
	query := strings.Join(os.Args[1:], " ")

	// Get the current shell
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "unknown shell"
	}

	interaction.prompt = fmt.Sprintf("Context: The user is asking a programming question from within the %s shell. Give a short answer and 1 or 2 examples.\nQuestion: %s", shell, query)

	_ = spinner.New().
		Title("thinking about it...").
		Action(getThisAnswer).
		Run()

	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
	)

	fmt.Println(r.Render(interaction.response))
}

func getThisAnswer() {
	getAnswer(&interaction)
}

func getAnswer(interaction *Interaction) {
	ctx := context.Background()

	// Initialize Genkit with the Google AI plugin
	// The API key is expected to be in the GEMINI_API_KEY environment variable
	g := genkit.Init(ctx, genkit.WithPlugins(&googlegenai.GoogleAI{}))

	// Generate structured data
	response, err := genkit.Generate(ctx, g,
		ai.WithPrompt(interaction.prompt),
		ai.WithModelName("googleai/gemini-2.5-flash"),
	)
	if err != nil {
		log.Fatalf("failed to generate response: %v", err)
	}

	interaction.response = response.Text()
}
