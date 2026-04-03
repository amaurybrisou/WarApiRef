param(
    [string]$RepoRoot = ".",
    [string]$SourceRoot = "addons",
    [string]$AddonApiOutput = "docs/addon-api-debug",
    [string]$WarApiOutput = "docs/war-api-debug",
    [string]$SiteOutput = "docs/site-debug/content",
    [string]$SubsetSourceRoot = ".debug/source-subset",
    [string[]]$Addons = @(),
    [string]$AddonListFile = "",
    [switch]$IncludeSite,
    [switch]$RunAudit,
    [switch]$NoClean
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

function Get-ResolvedPathOrThrow {
    param(
        [Parameter(Mandatory = $true)]
        [string]$PathValue,
        [string]$Label
    )

    if (-not (Test-Path -LiteralPath $PathValue)) {
        throw "$Label does not exist: $PathValue"
    }

    return (Resolve-Path -LiteralPath $PathValue).Path
}

function Invoke-Step {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Name,
        [Parameter(Mandatory = $true)]
        [scriptblock]$Action
    )

    $stepWatch = [System.Diagnostics.Stopwatch]::StartNew()
    Write-Host ""
    Write-Host "== $Name =="
    & $Action
    $stepWatch.Stop()
    Write-Host "Completed $Name in $($stepWatch.Elapsed.ToString())"
}

function Collect-Addons {
    param(
        [string[]]$InlineAddons,
        [string]$ListFile
    )

    $items = New-Object System.Collections.Generic.List[string]
    foreach ($name in $InlineAddons) {
        foreach ($part in ($name -split "[,;]")) {
            $trimmed = $part.Trim()
            if ($trimmed -ne "") {
                [void]$items.Add($trimmed)
            }
        }
    }

    if ($ListFile -ne "") {
        $resolvedList = Get-ResolvedPathOrThrow -PathValue $ListFile -Label "Addon list file"
        foreach ($line in Get-Content -LiteralPath $resolvedList) {
            $trimmed = $line.Trim()
            if ($trimmed -eq "" -or $trimmed.StartsWith("#")) {
                continue
            }
            [void]$items.Add($trimmed)
        }
    }

    return @($items | Sort-Object -Unique)
}

$repo = Get-ResolvedPathOrThrow -PathValue $RepoRoot -Label "Repo root"
Set-Location $repo

$selectedAddons = @(Collect-Addons -InlineAddons $Addons -ListFile $AddonListFile)
if ($selectedAddons.Length -eq 0) {
    throw "No addons selected. Use -Addons <name1,name2> or -AddonListFile <path>."
}

$sourceRootResolved = Get-ResolvedPathOrThrow -PathValue $SourceRoot -Label "Source root"
$addonApiResolved = Join-Path $repo $AddonApiOutput
$warApiResolved = Join-Path $repo $WarApiOutput
$siteResolved = Join-Path $repo $SiteOutput
$subsetSourceResolved = Join-Path $repo $SubsetSourceRoot

if (-not $NoClean) {
    Invoke-Step -Name "clean debug outputs" -Action {
        Remove-Item -LiteralPath $addonApiResolved -Recurse -Force -ErrorAction SilentlyContinue
        Remove-Item -LiteralPath $warApiResolved -Recurse -Force -ErrorAction SilentlyContinue
        Remove-Item -LiteralPath $subsetSourceResolved -Recurse -Force -ErrorAction SilentlyContinue
        if ($IncludeSite) {
            Remove-Item -LiteralPath $siteResolved -Recurse -Force -ErrorAction SilentlyContinue
        }
    }
}

Invoke-Step -Name "prepare source subset" -Action {
    New-Item -ItemType Directory -Path $subsetSourceResolved -Force | Out-Null

    foreach ($addon in $selectedAddons) {
        $sourceAddonPath = Join-Path $sourceRootResolved $addon
        if (-not (Test-Path -LiteralPath $sourceAddonPath)) {
            throw "Selected addon not found under source root: $addon"
        }

        $targetAddonPath = Join-Path $subsetSourceResolved $addon
        Copy-Item -LiteralPath $sourceAddonPath -Destination $targetAddonPath -Recurse -Force
    }
}

Write-Host ""
Write-Host "Selected addons ($($selectedAddons.Count)): $($selectedAddons -join ', ')"
Write-Host "Source root: $sourceRootResolved"
Write-Host "Subset source root: $subsetSourceResolved"
Write-Host "Addon API output: $addonApiResolved"
Write-Host "WAR API output: $warApiResolved"
if ($IncludeSite) {
    Write-Host "Site output: $siteResolved"
}

$allWatch = [System.Diagnostics.Stopwatch]::StartNew()

Invoke-Step -Name "generate addon artifacts (subset)" -Action {
    $cmd = @("run", ".\tools\api_doc_gen\main.go", "generate", "addon", $subsetSourceResolved, $addonApiResolved) + $selectedAddons
    & go $cmd
    if ($LASTEXITCODE -ne 0) {
        throw "addon generation failed"
    }
}

Invoke-Step -Name "generate platform docs (subset)" -Action {
    $cmd = @("run", ".\tools\api_doc_gen\main.go", "generate", "platform", $addonApiResolved, $warApiResolved, "--source-root", $subsetSourceResolved)
    & go $cmd
    if ($LASTEXITCODE -ne 0) {
        throw "platform generation failed"
    }
}

if ($IncludeSite) {
    Invoke-Step -Name "generate site content (subset)" -Action {
        $cmd = @("run", ".\tools\api_doc_gen\main.go", "generate", "site", $warApiResolved, $siteResolved)
        & go $cmd
        if ($LASTEXITCODE -ne 0) {
            throw "site generation failed"
        }
    }
}

if ($RunAudit) {
    Invoke-Step -Name "run contract audit on subset outputs" -Action {
        powershell -ExecutionPolicy Bypass -File .\scripts\contract_family_audit.ps1 -RepoRoot $repo -AddonRoot $SubsetSourceRoot -AddonApiRoot $AddonApiOutput
        if ($LASTEXITCODE -ne 0) {
            throw "contract audit failed"
        }
    }
}

$allWatch.Stop()
Write-Host ""
Write-Host "Subset debug pipeline completed in $($allWatch.Elapsed.ToString())"
Write-Host "Debug WAR docs root: $warApiResolved"
