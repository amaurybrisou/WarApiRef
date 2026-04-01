param(
    [string]$WorkspaceRoot = (Resolve-Path (Join-Path $PSScriptRoot "..\..")).Path,
    [int]$Port = 8081,
    [switch]$OpenBrowser
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

Push-Location $WorkspaceRoot
try {
    $env:WAR_API_REF_PORT = [string]$Port
    docker compose up -d war-api-ref-preview

    if ($OpenBrowser) {
        Start-Process "http://localhost:$Port/"
    }
}
finally {
    Pop-Location
}