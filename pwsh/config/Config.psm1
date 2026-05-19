function Get-Config {
    [CmdletBinding()]
    param(
        [string]$ConfigPath
    )

    $config = @{
        AppEnv = "development"
        Port   = "8080"
        APIKey = ""
    }

    # 1. Load from file (lowest priority)
    $finalPath = if ($ConfigPath) { $ConfigPath } else { Join-Path $PSScriptRoot "config.json" }
    if (Test-Path $finalPath) {
        try {
            $fileConfig = Get-Content $finalPath -Raw | ConvertFrom-Json
            if ($fileConfig.app_env) { $config.AppEnv = $fileConfig.app_env }
            if ($fileConfig.port) { $config.Port = $fileConfig.port }
            if ($fileConfig.api_key) { $config.APIKey = $fileConfig.api_key }
        } catch {
            Write-Warning "Failed to parse config.json: $_"
        }
    }

    # 2. Load from environment variables (overrides file/defaults)
    if ($env:APP_ENV) { $config.AppEnv = $env:APP_ENV }
    if ($env:PORT) { $config.Port = $env:PORT }
    if ($env:API_KEY) { $config.APIKey = $env:API_KEY }

    return $config
}

Export-ModuleMember -Function Get-Config
