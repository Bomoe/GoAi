# GoAi

GoAi is a Go CLI tool to query different Large Language Models (LLMs).

## Prerequisites

- Go installed on your system
- OpenAI API key

## Installation Instructions

1. Clone the repository:

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Build your Go program:
   ```
   go build -o goai
   ```
   This will create an executable named `goai` in your current directory.

4. Move the executable to a directory in your system's PATH:
   ```
   sudo mv goai /usr/local/bin/
   ```
   This moves the `goai` executable to `/usr/local/bin/`, which is typically in the system's PATH.

5. Make sure the executable has the correct permissions:
   ```
   sudo chmod +x /usr/local/bin/goai
   ```

6. Create a `.goai.env` file in your home directory with your OpenAI API key:
   ```
   echo "OPENAI_API_KEY=your_api_key_here" > ~/.goai.env
   ```

## Usage

Once installed, you can use the tool by running:
```
goai "Your query here"
```

For example:
```
goai "What is the capital of France?"
```

## Troubleshooting

- If you encounter permission issues, ensure your OpenAI API key is correct in `~/.goai.env`
- Make sure `/usr/local/bin` is in your PATH
- If the command is not found, try restarting your terminal or running `source ~/.bashrc` (or equivalent for your shell)

## Updating

To update the tool to the latest version:

1. Navigate to the GoAi directory
2. Pull the latest changes:
   ```
   git pull origin main
   ```
3. Repeat steps 2-5 from the Installation Instructions
