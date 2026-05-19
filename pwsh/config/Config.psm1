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
    $isDefault = -not $ConfigPath
    $finalPath = if ($ConfigPath) { $ConfigPath } else { Join-Path $PSScriptRoot "config.json" }
    
    if (-not (Test-Path $finalPath) -and $isDefault) {
        $example = @{
            app_env = $config.AppEnv
            port    = $config.Port
            api_key = $config.APIKey
        }
        $example | ConvertTo-Json | Out-File $finalPath
    }

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

function Save-Config {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory = $true)]
        $Config,

        [string]$ConfigPath
    )

    $finalPath = if ($ConfigPath) { $ConfigPath } else { Join-Path $PSScriptRoot "config.json" }
    
    $data = @{
        app_env = $Config.AppEnv
        port    = $Config.Port
        api_key = $Config.APIKey
    }

    $data | ConvertTo-Json | Out-File $finalPath
}

Export-ModuleMember -Function Get-Config, Save-Config
