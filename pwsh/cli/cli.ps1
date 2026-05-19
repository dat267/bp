#!/usr/bin/env pwsh

$commandsDir = Join-Path $PSScriptRoot "commands"
$commandFiles = Get-ChildItem $commandsDir -Filter "*.psm1"

$commands = @{}
foreach ($file in $commandFiles) {
    $name = $file.BaseName.ToLower()
    $commands[$name] = $file.FullName
}

function Show-Help {
    Write-Host "Usage: ./cli.ps1 <command> [arguments]"
    Write-Host "`nAvailable commands:"
    foreach ($cmd in $commands.Keys) {
        Write-Host "  $($cmd.PadRight(10))"
    }
}

if ($args.Count -lt 1) {
    Show-Help
    exit 1
}

$subcommand = $args[0].ToLower()
$remainingArgs = $args[1..($args.Count - 1)]

Import-Module (Join-Path $PSScriptRoot "../config/Config.psm1") -Force
$Config = Get-Config

if ($commands.ContainsKey($subcommand)) {
    Import-Module $commands[$subcommand] -Force
    # Construct the function name (e.g., info -> Invoke-Info)
    $functionName = "Invoke-$($subcommand.Substring(0,1).ToUpper())$($subcommand.Substring(1))"
    & $functionName $remainingArgs -Config $Config
}
else {
    Write-Error "Unknown command: $subcommand"
    Show-Help
    exit 1
}
