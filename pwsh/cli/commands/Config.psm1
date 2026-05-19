Import-Module (Join-Path $PSScriptRoot "../../config/Config.psm1") -Force

function Invoke-Config {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory = $true)]
        $Config
    )

    function Show-Menu {
        Write-Host "`n--- rclone-style Configuration Menu ---"
        Write-Host "1) View current configuration"
        Write-Host "2) Edit App Environment"
        Write-Host "3) Edit Port"
        Write-Host "4) Edit API Key"
        Write-Host "s) Save and Exit"
        Write-Host "q) Quit without saving"
    }

    function Get-ValidatedInput {
        param($Label, $Current, $Validator)
        while ($true) {
            $input = Read-Host "$Label [$Current]"
            if (-not $input) { return $Current }
            
            if ($Validator) {
                $error = &$Validator $input
                if ($error) {
                    Write-Host "Error: $error" -ForegroundColor Red
                    continue
                }
            }
            return $input
        }
    }

    $validatePort = {
        param($val)
        if ($val -notmatch "^\d+$") { return "Port must be a number" }
        $p = [int]$val
        if ($p -lt 1 -or $p -gt 65535) { return "Port must be between 1 and 65535" }
        return $null
    }

    $validateNotEmpty = {
        param($val)
        if (-not $val -or $val.Trim() -eq "") { return "Value cannot be empty" }
        return $null
    }

    while ($true) {
        Show-Menu
        $choice = Read-Host "Choose option"

        switch ($choice) {
            "1" {
                Write-Host "`nCurrent Configuration:"
                Write-Host "  AppEnv:  $($Config.AppEnv)"
                Write-Host "  Port:    $($Config.Port)"
                Write-Host "  APIKey:  $($Config.APIKey)"
            }
            "2" {
                $Config.AppEnv = Get-ValidatedInput "App Environment" $Config.AppEnv
            }
            "3" {
                $Config.Port = Get-ValidatedInput "Port" $Config.Port $validatePort
            }
            "4" {
                $Config.APIKey = Get-ValidatedInput "API Key" $Config.APIKey $validateNotEmpty
            }
            "s" {
                Save-Config -Config $Config
                Write-Host "Configuration saved."
                return
            }
            "q" {
                Write-Host "Exiting without saving."
                return
            }
            Default {
                Write-Host "Invalid option."
            }
        }
    }
}

Export-ModuleMember -Function Invoke-Config
