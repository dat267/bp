function Get-Config {
    [CmdletBinding()]
    param(
        [string]$ConfigPath
    )

    $config = @{}
    foreach ($key in $script:ConfigSchema.Keys) {
        $config[$key] = $script:ConfigSchema[$key].Default
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
            foreach ($key in $script:ConfigSchema.Keys) {
                # JSON names are app_env, while PowerShell keys are AppEnv
                $jsonKey = $script:ConfigSchema[$key].JsonKey
                if ($fileConfig.$jsonKey) { $config[$key] = $fileConfig.$jsonKey }
            }
        } catch {
            Write-Warning "Failed to parse config.json: $_"
        }
    }

    # 2. Load from environment variables (overrides file/defaults)
    $prefix = "BP_"
    function Get-EnvAuto {
        param($Name)
        $key = $prefix + $Name.ToUpper().Replace("-", "_")
        return Get-Item -Path "Env:$key" -ErrorAction SilentlyContinue | Select-Object -ExpandProperty Value
    }

    foreach ($key in $script:ConfigSchema.Keys) {
        $flagName = $script:ConfigSchema[$key].FlagName
        $envVal = Get-EnvAuto $flagName
        if ($envVal) { 
            if ($script:ConfigSchema[$key].Default -is [bool]) {
                $config[$key] = ($envVal -eq "true")
            } else {
                $config[$key] = $envVal
            }
        }
    }

    return $config
}

$ConfigSchema = [ordered]@{
    AppEnv  = @{ Label = "App Environment"; Default = "development"; JsonKey = "app_env"; FlagName = "app-env" }
    Port    = @{ Label = "Port"; Default = "8080"; Validator = "Test-Port"; JsonKey = "port"; FlagName = "port" }
    APIKey  = @{ Label = "API Key"; Default = ""; Validator = "Test-NotEmpty"; JsonKey = "api_key"; FlagName = "api-key" }
    Verbose = @{ Label = "Verbose Mode"; Default = $false; JsonKey = "verbose"; FlagName = "verbose" }
}

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

Export-ModuleMember -Function Get-Config, Save-Config, Test-Port, Test-NotEmpty
