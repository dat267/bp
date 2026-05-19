# Simple test script for PowerShell boilerplate

$Results = @()

# Test 1: Config Defaults
Import-Module "./pwsh/config/Config.psm1" -Force
$cfg = Get-Config
if ($cfg.AppEnv -eq "development") { $Results += "PASS: Config Defaults" } else { $Results += "FAIL: Config Defaults" }

# Test 2: Env Precedence
$env:BP_PORT = "1234"
$cfg = Get-Config
if ($cfg.Port -eq "1234") { $Results += "PASS: Env Precedence" } else { $Results += "FAIL: Env Precedence ($($cfg.Port))" }
Remove-Item Env:BP_PORT

# Test 3: File Fallback
$configPath = "./pwsh/config/config.json"
'{"port": "5555"}' | Out-File $configPath
$cfg = Get-Config
if ($cfg.Port -eq "5555") { $Results += "PASS: File Fallback" } else { $Results += "FAIL: File Fallback" }
Remove-Item $configPath

# Test 4: CLI Hello
$out = & ./pwsh/cli/cli.ps1 hello --name=Tester | Out-String
if ($out -match "Hello, Tester!") { $Results += "PASS: CLI Hello Flag" } else { $Results += "FAIL: CLI Hello Flag (Output: $out)" }

# Test 5: Retry Utility
Import-Module "./pwsh/utils/Retry.psm1" -Force
$attempts = 0
Invoke-WithRetry -ScriptBlock {
    $global:attempts++
    if ($global:attempts -lt 2) { throw "fail" }
    return "success"
} -InitialDelayMs 1 -MaxAttempts 3 > $null
if ($global:attempts -eq 2) { $Results += "PASS: Retry Utility" } else { $Results += "FAIL: Retry Utility ($global:attempts)" }

Write-Host "`nTest Results:"
$Results | ForEach-Object { Write-Host $_ }

if ($Results -match "FAIL") { exit 1 } else { exit 0 }
