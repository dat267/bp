Import-Module (Join-Path $PSScriptRoot "../../config/Config.psm1") -Force

function Invoke-Config {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory = $true)]
        $Config
    )

    Write-Host "--- Interactive Configuration Setup ---"

    $appEnv = Read-Host "App Environment [$($Config.AppEnv)]"
    if ($appEnv) { $Config.AppEnv = $appEnv }

    $port = Read-Host "Port [$($Config.Port)]"
    if ($port) { $Config.Port = $port }

    $apiKey = Read-Host "API Key (required) [$($Config.APIKey)]"
    if ($apiKey) { $Config.APIKey = $apiKey }

    $confirm = Read-Host "`nSave changes to config.json? (y/n) [n]"
    if ($confirm -eq "y") {
        try {
            Save-Config -Config $Config
            Write-Host "Configuration saved successfully."
        } catch {
            Write-Error "Failed to save config: $_"
        }
    } else {
        Write-Host "Changes discarded."
    }
}

Export-ModuleMember -Function Invoke-Config
