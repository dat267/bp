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

# 1. Parse global flags and identify subcommand
$ConfigPath = $null
$Verbose = $false
$filteredArgs = @()
$subcommand = $null

foreach ($arg in $args) {
    if ($arg -like "--config=*") {
        $ConfigPath = $arg.Split("=")[1]
    } elseif ($arg -eq "--verbose" -or $arg -eq "-v") {
        $Verbose = $true
    } elseif (-not $subcommand -and $arg -notlike "-*") {
        $subcommand = $arg.ToLower()
    } else {
        $filteredArgs += $arg
    }
}

if (-not $subcommand) {
    Show-Help
    exit 1
}

Import-Module (Join-Path $PSScriptRoot "../config/Config.psm1") -Force
$Config = Get-Config -ConfigPath $ConfigPath
if ($Verbose) { $Config.Verbose = $true }

if ($commands.ContainsKey($subcommand)) {
    Import-Module $commands[$subcommand] -Force
    # Construct the function name (e.g., info -> Invoke-Info)
    $functionName = "Invoke-$($subcommand.Substring(0,1).ToUpper())$($subcommand.Substring(1))"
    & $functionName -Args $filteredArgs -Config $Config
}
else {
    Write-Error "Unknown command: $subcommand"
    Show-Help
    exit 1
}
