# PowerShell Boilerplate

This directory contains reusable PowerShell patterns mirroring the Go and JavaScript implementations.

## Features

1.  **Modular CLI:** Command dispatcher in `cli.ps1` that auto-loads commands from the `commands/` folder.
2.  **Layered Config:** Fallback logic (Flags > Env > File > Defaults) implemented in `config/Config.psm1`.

## Usage

### Run the CLI
```powershell
# Show help
./cli/cli.ps1

# Run a command
./cli/cli.ps1 hello --name=Gemini
```

### Configuration
Create a `config.json` in `pwsh/config/`:
```json
{
    "app_env": "production",
    "port": "9000"
}
```

## Testing
```powershell
# Run the test script
./pwsh.test.ps1
```
