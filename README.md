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

### How to build an extensible CLI tool?
- **Go:** See [go/cmd/cli/main.go](./go/cmd/cli/main.go) - Command interface pattern.
- **JavaScript:** See [js/cli/cli.js](./js/cli/cli.js) - Modular command dispatcher.
- **PowerShell:** See [pwsh/cli/cli.ps1](./pwsh/cli/cli.ps1) - Module-based command dispatcher.

### How to handle transient failures (Retries)?
- **Go:** See [go/utils/retry.go](./go/utils/retry.go) - Exponential backoff with jitter.
- **JavaScript:** See [js/utils/retry.js](./js/utils/retry.js) - Async/await retry utility.
- **PowerShell:** See [pwsh/utils/Retry.psm1](./pwsh/utils/Retry.psm1) - Script block wrapper.

## Getting Started

Refer to the README in each subdirectory for specific usage and testing instructions.
