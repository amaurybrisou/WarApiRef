param(
    [string]$WorkspaceRoot = (Resolve-Path (Join-Path $PSScriptRoot "..\..")).Path,
    [string]$SourceRoot = "API_Ref",
    [string]$OutputRoot = "WAR_API_Ref"
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

Push-Location $WorkspaceRoot
try {
    docker compose build api-doc-gen
    docker compose run --rm api-doc-gen --mode platform "/workspace/$SourceRoot" "/workspace/$OutputRoot"
}
finally {
    Pop-Location
}