function Invoke-Hello {
    [CmdletBinding()]
    param(
        [Parameter(ValueFromRemainingArguments = $true)]
        $Args,
        
        [Parameter(Mandatory = $true)]
        $Config
    )

    $defaultName = if ($Config.AppEnv -eq "production") { "Production User" } else { "World" }

    # Simple manual flag parsing to mirror other languages
    $name = $defaultName
    foreach ($arg in $Args) {
        if ($arg -like "--name=*") {
            $name = $arg.Split("=")[1]
        }
    }

    return "Hello, $name!"
}

Export-ModuleMember -Function Invoke-Hello
