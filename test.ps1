# Unified test script for BP repository (PowerShell)

Write-Host "Starting BP Test Suite..." -ForegroundColor Cyan
$Failed = $false

# 1. Go Tests
Write-Host "`n[1/3] Running Go Tests..."
Push-Location go
go test ./...
if ($LASTEXITCODE -ne 0) {
    Write-Host "Go Tests Failed" -ForegroundColor Red
    $Failed = $true
} else {
    Write-Host "Go Tests Passed" -ForegroundColor Green
}
Pop-Location

# 2. JavaScript Tests
Write-Host "`n[2/3] Running JavaScript Tests..."
Push-Location js
npm test
if ($LASTEXITCODE -ne 0) {
    Write-Host "JS Tests Failed" -ForegroundColor Red
    $Failed = $true
} else {
    Write-Host "JS Tests Passed" -ForegroundColor Green
}
Pop-Location

# 3. PowerShell Tests
Write-Host "`n[3/3] Running PowerShell Tests..."
& pwsh -File ./pwsh/test.ps1
if ($LASTEXITCODE -ne 0) {
    Write-Host "PowerShell Tests Failed" -ForegroundColor Red
    $Failed = $true
} else {
    Write-Host "PowerShell Tests Passed" -ForegroundColor Green
}

Write-Host "`n----------------------------"
if (-not $Failed) {
    Write-Host "ALL TESTS PASSED SUCCESSFULLY" -ForegroundColor Green
    exit 0
} else {
    Write-Host "SOME TESTS FAILED" -ForegroundColor Red
    exit 1
}
