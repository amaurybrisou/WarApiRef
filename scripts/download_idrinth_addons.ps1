param(
    [Parameter(Mandatory = $true)]
    [string]$OutputPath,

    [string]$ListingUrl = "https://tools.idrinth.de/addons/",

    [switch]$Force
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$listingUri = [System.Uri]$ListingUrl
$siteRoot = [System.Uri]($listingUri.GetLeftPart([System.UriPartial]::Authority) + "/")

if (-not (Test-Path -LiteralPath $OutputPath)) {
    New-Item -ItemType Directory -Path $OutputPath | Out-Null
}

Write-Host "Fetching addon listing from $ListingUrl ..."
$response = Invoke-WebRequest -Uri $ListingUrl

$searchResults = $null
if ($null -ne $response.ParsedHtml) {
    $searchResults = $response.ParsedHtml.getElementById("searchResults")
}

if ($null -eq $searchResults) {
    throw "Could not find HTML element with id 'searchResults'."
}

$actionButtons = @(
    $searchResults.getElementsByClassName("actionButton") |
        Where-Object { $null -ne $_ -and -not [string]::IsNullOrWhiteSpace([string]$_.getAttribute("href")) }
)

if ($actionButtons.Count -eq 0) {
    throw "No addon download links found in #searchResults."
}

Write-Host "Found $($actionButtons.Count) addons to download."

$downloaded = 0
$failed = 0

function Expand-AddonArchive {
    param(
        [Parameter(Mandatory = $true)]
        [string]$ZipPath,

        [Parameter(Mandatory = $true)]
        [string]$DestinationPath
    )

    $tempExtractPath = Join-Path ([System.IO.Path]::GetTempPath()) ("ror-addon-extract-" + [System.Guid]::NewGuid().ToString("N"))
    New-Item -ItemType Directory -Path $tempExtractPath | Out-Null

    try {
        Expand-Archive -LiteralPath $ZipPath -DestinationPath $tempExtractPath -Force

        $topLevel = @(Get-ChildItem -LiteralPath $tempExtractPath -Force)
        $topLevelDirs = @($topLevel | Where-Object { $_.PSIsContainer })
        $topLevelFiles = @($topLevel | Where-Object { -not $_.PSIsContainer })

        if ($topLevelDirs.Count -eq 1 -and $topLevelFiles.Count -eq 0) {
            Move-Item -LiteralPath $topLevelDirs[0].FullName -Destination $DestinationPath
        }
        else {
            New-Item -ItemType Directory -Path $DestinationPath | Out-Null
            foreach ($item in $topLevel) {
                Move-Item -LiteralPath $item.FullName -Destination $DestinationPath
            }
        }
    }
    finally {
        if (Test-Path -LiteralPath $tempExtractPath) {
            Remove-Item -LiteralPath $tempExtractPath -Recurse -Force
        }
    }
}

foreach ($button in $actionButtons) {
    $rawHref = [string]$button.getAttribute("href")

    try {
        $downloadUri = [System.Uri]::new($siteRoot, $rawHref)

        $path = $downloadUri.AbsolutePath
        $slug = "addon"
        if ($path -match "/addons/([^/]+)/download/") {
            $slug = $matches[1]
        }

        $version = "latest"
        if ($path -match "/download/([^/]+)/?$") {
            $version = $matches[1]
        }

        $zipName = "$slug-$version.zip"
        $zipPath = Join-Path $OutputPath $zipName
        $extractPath = Join-Path $OutputPath $slug

        if (Test-Path -LiteralPath $extractPath) {
            if ($Force) {
                Remove-Item -LiteralPath $extractPath -Recurse -Force
            }
            else {
                Write-Warning "Skipping $slug ($version): destination already exists at $extractPath. Use -Force to overwrite."
                continue
            }
        }

        Write-Host "Downloading $slug ($version) from $downloadUri"
        Invoke-WebRequest -Uri $downloadUri -OutFile $zipPath

        Expand-AddonArchive -ZipPath $zipPath -DestinationPath $extractPath

        Remove-Item -LiteralPath $zipPath -Force
        $downloaded++
    }
    catch {
        $failed++
        Write-Warning "Failed to process '$rawHref': $($_.Exception.Message)"
    }
}

Write-Host "Completed. Downloaded and extracted: $downloaded. Failed: $failed."
