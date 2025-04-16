# LawPrompt

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/heyoSSam/LawPrompt
   cd LawPrompt
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the project:**
   ```bash
   go run main.go
   ```

## Configuration

The project uses a `config.yaml` file to configure runtime settings. Below is the structure:

```yaml
env:
  PORT: ""           # Port where the service runs
  MODEL: ""        # Model name to use
  OLLAMAURL: ""          # URL to the Ollama inference server
  SEARCHURL: ""          # URL to the search/context retrieval service
```


