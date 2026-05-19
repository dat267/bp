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
if ($cfg.Port -eq "5555") { $Results += "PASS: File Fallback" } else { $Results += "FAIL: File Fallback ($($cfg.Port))" }
Remove-Item $configPath

# Test 6: Save-Config
$savePath = "./pwsh/config/test_save.json"
$testCfg = @{ AppEnv = "prod"; Port = "4444"; APIKey = "key" }
Save-Config -Config $testCfg -ConfigPath $savePath
$saved = Get-Content $savePath -Raw | ConvertFrom-Json
if ($saved.port -eq "4444") { $Results += "PASS: Save-Config" } else { $Results += "FAIL: Save-Config" }
Remove-Item $savePath

# Test 7: Validators
# Note: Validators are inside the module scope, but we can test them via the module members if exported
$vPort1 = Test-Port "8080"
$vPort2 = Test-Port "abc"
$vEmpty = Test-NotEmpty ""
if (-not $vPort1 -and $vPort2 -and $vEmpty) { $Results += "PASS: Validators" } else { $Results += "FAIL: Validators" }

# Test 8: Config Corruption
$corruptPath = "./pwsh/config/corrupt.json"
'{ "bad": "json" ' | Out-File $corruptPath # Missing closing brace
$cfg = Get-Config -ConfigPath $corruptPath
if ($cfg.Port -eq "8080") { $Results += "PASS: Config Corruption" } else { $Results += "FAIL: Config Corruption ($($cfg.Port))" }
Remove-Item $corruptPath

# Test 4: CLI Hello
$out = & ./pwsh/cli/cli.ps1 hello --name=Tester | Out-String
if ($out -match "Hello, Tester!") { $Results += "PASS: CLI Hello Flag" } else { $Results += "FAIL: CLI Hello Flag (Output: $out)" }

# Test 9: Flag Position Independence
$outBefore = & ./pwsh/cli/cli.ps1 --verbose info | Out-String
$outAfter = & ./pwsh/cli/cli.ps1 info --verbose | Out-String
if ($outBefore -match "Verbose:     True" -and $outAfter -match "Verbose:     True") {
    $Results += "PASS: Flag Position Independence"
} else {
    $Results += "FAIL: Flag Position Independence"
}

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
