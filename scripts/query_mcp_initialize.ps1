param(
    [ValidateSet("stdio", "http")]
    [string]$Mode = "stdio",
    [string]$WorkspacePath = "c:/Return of Reckoning/Interface/AddOns/WAR_API_Ref",
    [string]$Image = "ror-mcp-server:local",
    [string]$DocsRoot = "/workspace/docs/war-api",
    [string]$HttpUrl = "http://127.0.0.1:8091/mcp"
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

function Read-StdioFrame {
    param(
        [Parameter(Mandatory = $true)]
        [System.IO.Stream]$Stream
    )

    $headerBytes = New-Object System.Collections.Generic.List[byte]
    $delimiter = [byte[]](13, 10, 13, 10)

    while ($true) {
        $b = $Stream.ReadByte()
        if ($b -lt 0) {
            throw "Reached EOF while reading MCP frame headers."
        }

        $headerBytes.Add([byte]$b)

        if ($headerBytes.Count -ge 4) {
            $n = $headerBytes.Count
            if (
                $headerBytes[$n - 4] -eq $delimiter[0] -and
                $headerBytes[$n - 3] -eq $delimiter[1] -and
                $headerBytes[$n - 2] -eq $delimiter[2] -and
                $headerBytes[$n - 1] -eq $delimiter[3]
            ) {
                break
            }
        }
    }

    $headerText = [System.Text.Encoding]::ASCII.GetString($headerBytes.ToArray())
    $contentLength = 0
    foreach ($line in ($headerText -split "`r?`n")) {
        if ($line -match "^Content-Length:\s*(\d+)\s*$") {
            $contentLength = [int]$matches[1]
            break
        }
    }

    if ($contentLength -le 0) {
        throw "Missing or invalid Content-Length in MCP response frame. Header: $headerText"
    }

    $payload = New-Object byte[] $contentLength
    $offset = 0
    while ($offset -lt $contentLength) {
        $read = $Stream.Read($payload, $offset, $contentLength - $offset)
        if ($read -le 0) {
            throw "Reached EOF while reading MCP response payload."
        }
        $offset += $read
    }

    return [System.Text.Encoding]::UTF8.GetString($payload)
}

$requestObject = @{
    jsonrpc = "2.0"
    id = 1
    method = "initialize"
    params = @{}
}
$requestJson = $requestObject | ConvertTo-Json -Compress

if ($Mode -eq "http") {
    $response = $null
    $lastError = $null
    for ($i = 0; $i -lt 10; $i++) {
        try {
            $response = Invoke-RestMethod -Method Post -Uri $HttpUrl -ContentType "application/json" -Body $requestJson
            break
        }
        catch {
            $lastError = $_
            Start-Sleep -Milliseconds 500
        }
    }

    if ($null -eq $response) {
        throw "HTTP initialize request failed after retries. Last error: $lastError"
    }

    $response | ConvertTo-Json -Depth 20
    exit 0
}

$psi = New-Object System.Diagnostics.ProcessStartInfo
$psi.FileName = "docker"
$psi.Arguments = "run --rm -i -v `"${WorkspacePath}:/workspace:ro`" $Image --transport stdio --docs-root $DocsRoot"
$psi.RedirectStandardInput = $true
$psi.RedirectStandardOutput = $true
$psi.RedirectStandardError = $true
$psi.UseShellExecute = $false
$psi.CreateNoWindow = $true

$process = New-Object System.Diagnostics.Process
$process.StartInfo = $psi
$null = $process.Start()

try {
    $requestBytesLen = [System.Text.Encoding]::UTF8.GetByteCount($requestJson)
    $frame = "Content-Length: $requestBytesLen`r`n`r`n$requestJson"

    $process.StandardInput.Write($frame)
    $process.StandardInput.Flush()

    $responsePayload = Read-StdioFrame -Stream $process.StandardOutput.BaseStream

    try {
        $responsePayload | ConvertFrom-Json | ConvertTo-Json -Depth 20
    }
    catch {
        Write-Output $responsePayload
    }
}
finally {
    if (-not $process.HasExited) {
        $process.Kill()
        $process.WaitForExit()
    }

    if ($process.StandardError.Peek() -ge 0) {
        $stderrText = $process.StandardError.ReadToEnd()
        if (-not [string]::IsNullOrWhiteSpace($stderrText)) {
            Write-Verbose $stderrText
        }
    }
}
