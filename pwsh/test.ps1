$Results = @()
Import-Module "$PSScriptRoot/config/Config.psm1" -Force
$cfg = Get-Config
if ($cfg.AppEnv -eq "development") { $Results += "PASS: Config Defaults" } else { $Results += "FAIL: Config Defaults" }
$env:BP_PORT = "1234"
$cfg = Get-Config
if ($cfg.Port -eq "1234") { $Results += "PASS: Env Precedence" } else { $Results += "FAIL: Env Precedence ($($cfg.Port))" }
Remove-Item Env:BP_PORT
$configPath = "$PSScriptRoot/config/config.json"
'{"port": "5555"}' | Out-File $configPath
$cfg = Get-Config
if ($cfg.Port -eq "5555") { $Results += "PASS: File Fallback" } else { $Results += "FAIL: File Fallback ($($cfg.Port))" }
Remove-Item $configPath
$savePath = "$PSScriptRoot/config/test_save.json"
$testCfg = @{ AppEnv = "prod"; Port = "4444"; APIKey = "key" }
Save-Config -Config $testCfg -ConfigPath $savePath
$saved = Get-Content $savePath -Raw | ConvertFrom-Json
if ($saved.port -eq "4444") { $Results += "PASS: Save-Config" } else { $Results += "FAIL: Save-Config" }
Remove-Item $savePath
$vPort1 = Test-Port "8080"
$vPort2 = Test-Port "abc"
$vEmpty = Test-NotEmpty ""
if (-not $vPort1 -and $vPort2 -and $vEmpty) { $Results += "PASS: Validators" } else { $Results += "FAIL: Validators" }
$corruptPath = "$PSScriptRoot/config/corrupt.json"
'{ "bad": "json" ' | Out-File $corruptPath
$cfg = Get-Config -ConfigPath $corruptPath
if ($cfg.Port -eq "8080") { $Results += "PASS: Config Corruption" } else { $Results += "FAIL: Config Corruption ($($cfg.Port))" }
Remove-Item $corruptPath
$out = & "$PSScriptRoot/cli/cli.ps1" hello --name=Tester | Out-String
if ($out -match "Hello, Tester!") { $Results += "PASS: CLI Hello Flag" } else { $Results += "FAIL: CLI Hello Flag (Output: $out)" }
$outBefore = & "$PSScriptRoot/cli/cli.ps1" --verbose info | Out-String
$outAfter = & "$PSScriptRoot/cli/cli.ps1" info --verbose | Out-String
if ($outBefore -match "Verbose:     True" -and $outAfter -match "Verbose:     True") {
    $Results += "PASS: Flag Position Independence"
} else {
    $Results += "FAIL: Flag Position Independence"
}
Import-Module "$PSScriptRoot/utils/Retry.psm1" -Force
$attempts = 0
Invoke-WithRetry -ScriptBlock {
    $global:attempts++
    if ($global:attempts -lt 2) { throw "fail" }
    return "success"
} -InitialDelayMs 1 -MaxAttempts 3 > $null
if ($global:attempts -eq 2) { $Results += "PASS: Retry Utility" } else { $Results += "FAIL: Retry Utility ($global:attempts)" }
Import-Module "$PSScriptRoot/client.psm1" -Force
$job = Start-Job -ScriptBlock {
    $listener = [System.Net.HttpListener]::new()
    $listener.Prefixes.Add("http://localhost:8099/")
    $listener.Start()
    $ctx = $listener.GetContext()
    $buf = [System.Text.Encoding]::UTF8.GetBytes("ok")
    $ctx.Response.ContentLength64 = $buf.Length
    $ctx.Response.OutputStream.Write($buf, 0, $buf.Length)
    $ctx.Response.Close()
    $listener.Stop()
}
Start-Sleep -Milliseconds 500
$reqs = @( @{ Url = "http://localhost:8099/test"; Method = "GET" } )
$res = Invoke-ConcurrentRequest -Requests $reqs -ThrottleLimit 1
Stop-Job $job -ErrorAction SilentlyContinue
Remove-Job $job -ErrorAction SilentlyContinue
if ($res.StatusCode -eq 200 -and $res.Data -eq "ok") { $Results += "PASS: Concurrent Requests" } else { $Results += "FAIL: Concurrent Requests" }
Write-Host "`nTest Results:"
$Results | ForEach-Object { Write-Host $_ }
if ($Results -match "FAIL") { exit 1 } else { exit 0 }
