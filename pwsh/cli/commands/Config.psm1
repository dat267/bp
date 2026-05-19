Import-Module (Join-Path $PSScriptRoot "../../config/Config.psm1") -Force

function Invoke-Config {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory = $true)]
        $Config
    )

    # Access the schema from the module scope
    $schema = $script:ConfigSchema

    function Show-Menu {
        Write-Host "`n--- rclone-style Configuration Menu ---"
        Write-Host "v) View current configuration"
        for ($i = 0; $i -lt $schema.Count; $i++) {
            Write-Host "$($i + 1)) Edit $($schema[$i].Label)"
        }
        Write-Host "s) Save and Exit"
        Write-Host "q) Quit without saving"
    }

    function Get-ValidatedInput {
        param($Label, $Current, $ValidatorName)
        while ($true) {
            $input = Read-Host "$Label [$Current]"
            if (-not $input) { return $Current }
            
            if ($ValidatorName) {
                # Look up the validator function in the module
                $error = & $ValidatorName $input
                if ($error) {
                    Write-Host "Error: $error" -ForegroundColor Red
                    continue
                }
            }
            return $input
        }
    }

    while ($true) {
        Show-Menu
        $choice = Read-Host "Choose option"

        $keys = $schema.Keys | ForEach-Object { $_ }

        if ($choice -eq "v") {
            Write-Host "`nCurrent Configuration:"
            foreach ($key in $keys) {
                $field = $schema[$key]
                $label = ($field.Label + ":").PadRight(16)
                Write-Host "  $label $($Config[$key])"
            }
            continue
        }
        if ($choice -eq "s") {
            Save-Config -Config $Config
            Write-Host "Configuration saved."
            return
        }
        if ($choice -eq "q") {
            Write-Host "Exiting without saving."
            return
        }

        if ($choice -as [int] -and $choice -gt 0 -and $choice -le $keys.Count) {
            $key = $keys[[int]$choice - 1]
            $field = $schema[$key]
            $Config[$key] = Get-ValidatedInput $field.Label $Config[$key] $field.Validator
        } else {
            Write-Host "Invalid option."
        }
    }
}

Export-ModuleMember -Function Invoke-Config
