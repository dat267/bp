function Invoke-ConcurrentRequest {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory = $true)]
        [array]$Requests,

        [int]$ThrottleLimit = 5
    )

    # ForEach-Object -Parallel runs on threads, similar to Go's goroutines
    $Results = $Requests | ForEach-Object -Parallel {
        $req = $_
        $params = @{
            Uri         = $req.Url
            Method      = $req.Method
            ErrorAction = "Stop"
        }
        if ($req.Headers) { $params.Headers = $req.Headers }
        if ($req.Body) { $params.Body = ($req.Body | ConvertTo-Json) }

        try {
            $startTime = Get-Date
            $response = Invoke-RestMethod @params
            $endTime = Get-Date
            
            return [PSCustomObject]@{
                Method     = $req.Method
                Url        = $req.Url
                StatusCode = 200
                Data       = $response
                DurationMs = ($endTime - $startTime).TotalMilliseconds
                Error      = $null
            }
        }
        catch {
            return [PSCustomObject]@{
                Method     = $req.Method
                Url        = $req.Url
                StatusCode = if ($_.Exception.Response) { [int]$_.Exception.Response.StatusCode } else { 0 }
                Data       = $null
                DurationMs = 0
                Error      = $_.Exception.Message
            }
        }
    } -ThrottleLimit $ThrottleLimit

    return $Results
}

Export-ModuleMember -Function Invoke-ConcurrentRequest
