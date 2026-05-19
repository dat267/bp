# BP - Boilerplate Repository

A collection of reusable code snippets and architectural patterns for various programming languages.

## Project Structure

- `go/`: Go-based boilerplates (e.g., Concurrent API Client).
- `js/`: JavaScript/Node.js-based boilerplates.
- `pwsh/`: PowerShell-based boilerplates.

## Common Patterns (How-Tos)

### How to call multiple APIs concurrently?
- **Go:** See [go/README.md](./go/README.md) - Uses **goroutines** and `sync.WaitGroup`.
- **JavaScript:** See [js/README.md](./js/README.md) - Uses **`Promise.all`** and `async/await`.
- **PowerShell:** See [pwsh/client.psm1](./pwsh/client.psm1) - Uses **`ForEach-Object -Parallel`**.

### How to build an extensible CLI tool?
- **Go:** See [go/cmd/cli/main.go](./go/cmd/cli/main.go) - Command interface pattern.
- **JavaScript:** See [js/cli/cli.js](./js/cli/cli.js) - Modular command dispatcher.
- **PowerShell:** See [pwsh/cli/cli.ps1](./pwsh/cli/cli.ps1) - Module-based command dispatcher.

### How to handle transient failures (Retries)?
- **Go:** See [go/utils/retry.go](./go/utils/retry.go) - Exponential backoff with jitter.
- **JavaScript:** See [js/utils/retry.js](./js/utils/retry.js) - Async/await retry utility.
- **PowerShell:** See [pwsh/utils/Retry.psm1](./pwsh/utils/Retry.psm1) - Script block wrapper.

### How to interactively setup configuration?
- **All Languages:** Run the `config` subcommand (e.g., `go run ./cmd/cli config`, `node cli/cli.js config`, or `pwsh ./cli/cli.ps1 config`). It will prompt for required variables and save them to `config.json`.

## Getting Started

Refer to the README in each subdirectory for specific usage and testing instructions.

### Running all tests
You can run the entire test suite (Go, JS, and PowerShell) using the provided unified scripts:

**Linux/macOS:**
```bash
./test.sh
```

**Windows/PowerShell:**
```powershell
./test.ps1
```
