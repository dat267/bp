# Project: BP (Boilerplate)

This repository serves as a centralized collection of reusable code snippets and architectural patterns for various programming languages (currently Go and JavaScript).

## Core Conventions

- **Organization:** Use language-specific root directories (e.g., `go/`, `js/`, `pwsh/`).
- **Completeness:** Every boilerplate must be accompanied by:
    - A language-specific `README.md` explaining features and usage.
    - Automated tests (Go: `*_test.go`, JS: `*.test.js`, PWSH: `*.test.ps1`).
- **Modular Architecture (Pluggable):** All components must use dependency injection or parameter-based configuration. Avoid hardcoded global imports in command or utility logic to ensure easy extraction.
- **Tech Stack Preferences:**
    - **Go:** Standard library (`slog`, `net/http`, `flag`).
    - **JavaScript:** Node.js 18+ native APIs (`fetch`, `node:test`).
    - **PowerShell:** PowerShell 7+ Core, modular design (`.psm1`).

## Established Patterns

1. **Concurrent API Client:** 
    - Go: Goroutines + `sync.WaitGroup`.
    - JS: `async/await` + `Promise.all`.
    - PowerShell: `ForEach-Object -Parallel` (available in PWSH 7+).
2. **Structured Logging:**
    - Go: `log/slog` (JSON handler).
    - JS: Class-based JSON logger.
3. **Environment Config (Layered):**
    - Priority: Flags > Environment Variables > `config.json` > Defaults.
4. **Extensible CLI Tool (Subcommand Dispatcher):**
    - Go: Interface pattern with `flag.FlagSet`.
    - JS: Dynamic module loader.
    - PowerShell: Module-based with dynamic `Invoke-` naming.
5. **Resilient Retry (Exponential Backoff):**
    - Implementations with jitter and max delay across all languages.
