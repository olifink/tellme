# TellMe

**TellMe** is a smart CLI companion that answers your programming questions directly in your terminal. Powered by Google's Gemini AI (via Firebase Genkit), it provides concise answers and examples tailored to your current shell environment.

## Features

- 🧠 **AI-Powered**: Uses the `googleai/gemini-2.5-flash` model for fast and accurate responses.
- 🐚 **Context Aware**: Automatically detects your shell (bash, zsh, fish, etc.) to provide relevant command examples.
- 💅 **Beautiful Output**: Renders responses in the terminal with syntax highlighting and markdown formatting using [Glamour](https://github.com/charmbracelet/glamour).

## Prerequisites

- [Go](https://go.dev/dl/) 1.23 or later.
- A Google AI Studio API Key. You can get one [here](https://aistudio.google.com/).

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd tellme
   ```

2. Install dependencies and build the binary:
   ```bash
   go build .
   ```

3. (Optional) Move the binary to your PATH:
   ```bash
   go install .
   ```

## Configuration

You must set the `GEMINI_API_KEY` environment variable for the tool to work.

```bash
export GEMINI_API_KEY="your-api-key-here"
```

You can add this line to your shell's configuration file (e.g., `.bashrc`, `.zshrc`) to make it permanent.

## Usage

Simply run `tellme` followed by your question in quotes:

```bash
tellme "how do I find all files larger than 100MB?"
```

```bash
tellme "how to undo the last git commit"
```

The tool will detect your shell (e.g., `/bin/zsh`) and tailor the answer to it.

## Built With

- [Firebase Genkit for Go](https://github.com/firebase/genkit/tree/main/go)
- [Google AI SDK](https://github.com/google/generative-ai-go)
- [Glamour](https://github.com/charmbracelet/glamour)
