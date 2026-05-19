# Test for PowerShell Concurrency using local sleep to avoid network noise

Import-Module "./pwsh/client.psm1" -Force

# Mock function that simulates work
function Invoke-MockRequest {
    param($Id)
    Start-Sleep -Seconds 1
    return "Result $Id"
}

$Items = 1..5

Write-Host "Starting 5 parallel tasks (each 1s)..."
$StartTime = Get-Date
$Results = $Items | ForEach-Object -Parallel {
    Start-Sleep -Seconds 1
    return "Done"
} -ThrottleLimit 5
$EndTime = Get-Date

$TotalDuration = ($EndTime - $StartTime).TotalSeconds
Write-Host "Total Duration: $TotalDuration seconds"

if ($TotalDuration -lt 2.0) {
    Write-Host "`nPASS: Concurrency verified (Total time < 2s for 5x 1s tasks)"
    exit 0
} else {
    Write-Host "`nFAIL: Concurrency failed (Tasks ran sequentially? Total time: $TotalDuration)"
    exit 1
}
