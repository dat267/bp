function Invoke-Info {
    [CmdletBinding()]
    param(
        [Parameter(ValueFromRemainingArguments = $true)]
        $Args,

        [Parameter(Mandatory = $true)]
        $Config
    )

    $apiKeyStatus = if ($Config.APIKey) { "********" } else { "not set" }
    
    return @"
Environment: $($Config.AppEnv)
Port:        $($Config.Port)
API Key:     $apiKeyStatus
"@
}

Export-ModuleMember -Function Invoke-Info
