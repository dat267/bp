# Project: BP (Boilerplate)

This repository serves as a centralized collection of reusable code snippets and architectural patterns for Go, JavaScript, and PowerShell.

## Core Conventions

- **Organization:** Use language-specific root directories (e.g., `go/`, `js/`, `pwsh/`).
- **Completeness:** Every boilerplate must include:
    - A language-specific `README.md` explaining features and usage.
    - Automated tests (Go: `*_test.go`, JS: `*.test.js`, PWSH: `*.test.ps1`).
    - Unified validation via root-level `test.sh` (Linux) and `test.ps1` (Windows).
- **Modular Architecture (Pluggable):** All components must use dependency injection or parameter-based configuration. Avoid global imports to ensure easy "copy-paste" extraction.
- **Zero-Config Startup:** Applications should auto-generate a default `config.json` if one is missing, providing an immediate template for users.
- **Production Hardening:** Code must be resilient against common edge cases:
    - **Config Resilience**: Fallback to defaults if `config.json` is malformed/corrupt.
    - **Resource Safety**: Retries and async tasks must respect context/timeout cancellation.
    - **Input Validation**: Interactive prompts must use schema-defined validators.
- **Tech Stack Preferences:**
    - **Go:** Standard library (`slog`, `net/http`, `flag`, `reflect`).
    - **JavaScript:** Node.js 18+ native APIs (`fetch`, `node:test`, `readline`).
    - **PowerShell:** PowerShell 7+ Core, modular design (`.psm1`), `[ordered]` hashtables.

## Established Patterns

1. **Concurrent API Client:** 
    - Pattern: High-performance parallel execution with throttling.
    - Implementation: Goroutines (Go), `Promise.all` (JS), `ForEach-Object -Parallel` (PWSH).
2. **Resilient Retry (Exponential Backoff):**
    - Features: Jitter, max delay, and interruptible sleep.
3. **DRY Layered Configuration:**
    - **Source of Truth**: Schema-driven (Struct tags in Go, static objects in JS, ordered hashtables in PWSH).
    - **Precedence**: Command Flags > `BP_` Env Variables > `config.json` > Hardcoded Defaults.
    - **Position Independence**: Global flags (e.g., `--config`, `--verbose`) must be recognized anywhere in the command string.
4. **Interactive Setup Dispatcher:**
    - Pattern: Menu-driven CLI tool that reflects the configuration schema to automatically build its UI and validation logic.
5. **Structured Logging:**
    - Standard: Class-based or `slog` JSON handlers with configurable levels.
