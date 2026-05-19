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
        Save-Config -Config $config -ConfigPath $finalPath
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
    # rclone style: auto-mapping BP_PORT, BP_APP_ENV, etc.
    $prefix = "BP_"
    function Get-EnvAuto {
        param($Name)
        $key = $prefix + $Name.ToUpper().Replace("-", "_")
        return Get-Item -Path "Env:$key" -ErrorAction SilentlyContinue | Select-Object -ExpandProperty Value
    }

    $envAppEnv = Get-EnvAuto "app-env"
    if ($envAppEnv) { $config.AppEnv = $envAppEnv }

    $envPort = Get-EnvAuto "port"
    if ($envPort) { $config.Port = $envPort }

    $envApiKey = Get-EnvAuto "api-key"
    if ($envApiKey) { $config.APIKey = $envApiKey }

    return $config
}

$ConfigSchema = @(
    @{ Key = "AppEnv"; Label = "App Environment" }
    @{ Key = "Port"; Label = "Port"; Validator = "Test-Port" }
    @{ Key = "APIKey"; Label = "API Key"; Validator = "Test-NotEmpty" }
)

function Test-Port {
    param($val)
    if ($val -notmatch "^\d+$") { return "Port must be a number" }
    $p = [int]$val
    if ($p -lt 1 -or $p -gt 65535) { return "Port must be between 1 and 65535" }
    return $null
}

function Test-NotEmpty {
    param($val)
    if (-not $val -or $val.Trim() -eq "") { return "Value cannot be empty" }
    return $null
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
