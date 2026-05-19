function Invoke-WithRetry {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory = $true)]
        [scriptblock]$ScriptBlock,

        [int]$MaxAttempts = 3,
        [int]$InitialDelayMs = 1000,
        [int]$MaxDelayMs = 30000,
        [double]$BackoffFactor = 2.0,
        [switch]$UseJitter = $true
    )

    $lastError = $null
    $delay = $InitialDelayMs

    for ($attempt = 1; $attempt -le $MaxAttempts; $attempt++) {
        try {
            return &$ScriptBlock
        }
        catch {
            $lastError = $_
            if ($attempt -eq $MaxAttempts) { break }

            $sleepTime = $delay
            if ($UseJitter -and $delay -gt 1) {
                # Add random jitter: [0, delay / 2]
                $maxJitter = [int]($delay / 2)
                $jitter = (Get-Random -Minimum 0 -Maximum $maxJitter)
                $sleepTime += $jitter
            }

            Start-Sleep -Milliseconds $sleepTime
            $delay = [Math]::Min($delay * $BackoffFactor, $MaxDelayMs)
        }
    }

    throw $lastError
}

Export-ModuleMember -Function Invoke-WithRetry
