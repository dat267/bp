# BP - Boilerplate Repository

A collection of reusable code snippets and architectural patterns for various programming languages.

## Quick Start: Create a new CLI

Each language directory is designed to be **pluggable**. To start a new project, copy the following core files:

### 🚀 Go Starter Kit
Copy these files from the `go/` directory:
1.  `cmd/cli/main.go` - The entry point and global flag parser.
2.  `config/config.go` - The reflective "Source of Truth" configuration module.
3.  `cmd/cli/hello.go` - A sample command (copy as a template for your own commands).

### 🚀 JavaScript Starter Kit
Copy these files from the `js/` directory:
1.  `cli/cli.js` - The dynamic command loader and dispatcher.
2.  `cli/args.js` - The lightweight flag parsing utility.
3.  `config/config.js` - The schema-driven configuration class.
4.  `cli/commands/hello.js` - A sample command template.

### 🚀 PowerShell Starter Kit
Copy these files from the `pwsh/` directory:
1.  `cli/cli.ps1` - The module-based command dispatcher.
2.  `config/Config.psm1` - The ordered-schema configuration module.
3.  `cli/commands/Hello.psm1` - A sample command template.

---

## Project Structure
```text
.
├── go/                 # Go Boilerplates
│   ├── cmd/cli/        # CLI Entry point & Command implementations
│   ├── config/         # Configuration logic (Reflection-based)
│   └── utils/          # Shared utilities (Retries, etc.)
├── js/                 # JavaScript/Node.js Boilerplates
│   ├── cli/            # Dispatcher and Flag utility
│   │   └── commands/   # Command implementations (Auto-loaded)
│   ├── config/         # Schema-driven configuration class
│   └── utils/          # Async utilities
├── pwsh/               # PowerShell Boilerplates
│   ├── cli/            # Script dispatcher
│   │   └── commands/   # Command modules (.psm1)
│   ├── config/         # Hashtable-based config module
│   └── utils/          # Script-block wrappers
└── test.sh / test.ps1  # Unified test automation scripts
```

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
