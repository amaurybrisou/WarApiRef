param(
  [string]$RepoRoot = ".",
  [string]$AddonRoot = "addons",
  [string]$AddonApiRoot = "docs/addon-api",
  [string]$AnomaliesCsv = "",
  [string]$InformationalCsv = ""
)

$ErrorActionPreference = "Stop"

function Canon([string]$s) {
  if ([string]::IsNullOrWhiteSpace($s)) { return "" }
  return (($s.ToLower()) -replace "[^a-z0-9]", "")
}

function CollectFamily([string]$dir) {
  $byCanon = @{}
  if (-not (Test-Path -LiteralPath $dir)) { return $byCanon }

  Get-ChildItem -Path $dir -Recurse -Filter "*.json" | ForEach-Object {
    try {
      $jsonObj = Get-Content -LiteralPath $_.FullName -Raw | ConvertFrom-Json
      $addon = [string]$jsonObj.addon
      if ([string]::IsNullOrWhiteSpace($addon)) { return }

      $canon = Canon $addon
      if (-not $byCanon.ContainsKey($canon)) {
        $byCanon[$canon] = [ordered]@{
          Names       = @()
          Files       = @()
          HandlerLinks = 0
        }
      }

      $rec = $byCanon[$canon]
      $rec.Names += $addon
      $rec.Files += $_.FullName
      if ($null -ne $jsonObj.handler_links) {
        $rec.HandlerLinks += @($jsonObj.handler_links).Count
      }
    } catch {
      # ignore malformed artifact payloads in this audit
    }
  }

  foreach ($k in @($byCanon.Keys)) {
    $byCanon[$k].Names = @($byCanon[$k].Names | Sort-Object -Unique)
    $byCanon[$k].Files = @($byCanon[$k].Files | Sort-Object -Unique)
  }
  return $byCanon
}

$repo = (Resolve-Path $RepoRoot).Path
Set-Location $repo

$sourceRoot = Join-Path $repo $AddonRoot
$addonApi = Join-Path $repo $AddonApiRoot

$sourceDirs = Get-ChildItem -LiteralPath $sourceRoot -Directory | Where-Object { -not $_.Name.StartsWith('.') }
$sourceByCanon = @{}

foreach ($dir in $sourceDirs) {
  $mod = Get-ChildItem -LiteralPath $dir.FullName -Filter "*.mod" -File | Select-Object -First 1
  $toc = Get-ChildItem -LiteralPath $dir.FullName -Filter "*.toc" -File | Select-Object -First 1

  $manifestType = ""
  $manifestName = ""
  $filesListed = @()

  if ($mod) {
    $manifestType = "mod"
    try {
      [xml]$xmlDoc = Get-Content -LiteralPath $mod.FullName -Raw
      $uiMod = $xmlDoc.SelectSingleNode("//UiMod")
      if ($uiMod -and $uiMod.name) { $manifestName = [string]$uiMod.name }
      $fileNodes = $xmlDoc.SelectNodes("//Files/File")
      foreach ($n in $fileNodes) {
        if ($n.name) { $filesListed += [string]$n.name }
      }
    } catch {
      # keep partial data for audit visibility
    }
  } elseif ($toc) {
    $manifestType = "toc"
    $manifestName = [IO.Path]::GetFileNameWithoutExtension($toc.Name)
    $lines = Get-Content -LiteralPath $toc.FullName
    foreach ($ln in $lines) {
      $t = $ln.Trim()
      if ($t -eq "" -or $t.StartsWith("#") -or $t.StartsWith("##")) { continue }
      $filesListed += $t
    }
  }

  $listedXml = @($filesListed | Where-Object { $_.ToLower().EndsWith(".xml") })
  $listedLua = @($filesListed | Where-Object { $_.ToLower().EndsWith(".lua") })
  $listedXmlExisting = @($listedXml | Where-Object { Test-Path -LiteralPath (Join-Path $dir.FullName $_) })
  $listedLuaExisting = @($listedLua | Where-Object { Test-Path -LiteralPath (Join-Path $dir.FullName $_) })

  $canon = Canon $dir.Name
  if (-not $sourceByCanon.ContainsKey($canon)) {
    $sourceByCanon[$canon] = [ordered]@{
      Canon                    = $canon
      SourceDir                = $dir.Name
      ManifestName             = $manifestName
      HasMod                   = ($manifestType -eq "mod")
      ManifestType             = $manifestType
      HasManifestListedXML     = ($listedXml.Count -gt 0)
      HasManifestListedLua     = ($listedLua.Count -gt 0)
      ListedXMLCount           = $listedXml.Count
      ListedLuaCount           = $listedLua.Count
      ExistingListedXMLCount   = $listedXmlExisting.Count
      ExistingListedLuaCount   = $listedLuaExisting.Count
      SourceNames              = @($dir.Name)
    }
  }
}

$xmlFam = CollectFamily (Join-Path $addonApi "xml-tree")
$luaFam = CollectFamily (Join-Path $addonApi "lua-analysis")
$linkFam = CollectFamily (Join-Path $addonApi "xml-lua-links")

$allKeys = @($sourceByCanon.Keys + $xmlFam.Keys + $luaFam.Keys + $linkFam.Keys | Sort-Object -Unique)

$rows = @()
foreach ($k in $allKeys) {
  $srcRec = $sourceByCanon[$k]
  $xmlRec = $xmlFam[$k]
  $luaRec = $luaFam[$k]
  $linkRec = $linkFam[$k]

  $present = $null -ne $srcRec
  $hasMod = if ($present) { [bool]$srcRec.HasMod } else { $false }
  $hasListedXml = if ($present) { [bool]$srcRec.HasManifestListedXML } else { $false }
  $hasListedLua = if ($present) { [bool]$srcRec.HasManifestListedLua } else { $false }
  $existingListedXml = if ($present) { [int]$srcRec.ExistingListedXMLCount } else { 0 }
  $existingListedLua = if ($present) { [int]$srcRec.ExistingListedLuaCount } else { 0 }

  $hasXmlTree = $null -ne $xmlRec
  $hasLuaAnalysis = $null -ne $luaRec
  $hasLinks = $null -ne $linkRec

  $zeroLinks = $hasLinks -and ($linkRec.HandlerLinks -eq 0)
  $realLinks = $hasLinks -and ($linkRec.HandlerLinks -gt 0)

  $displayNames = @()
  if ($present) {
    $displayNames += $srcRec.SourceDir
    if ($srcRec.ManifestName) { $displayNames += $srcRec.ManifestName }
  }
  if ($hasXmlTree) { $displayNames += $xmlRec.Names }
  if ($hasLuaAnalysis) { $displayNames += $luaRec.Names }
  if ($hasLinks) { $displayNames += $linkRec.Names }
  $displayNames = @($displayNames | Where-Object { $_ -and $_.Trim() -ne "" } | Sort-Object -Unique)

  $classification = ""
  $reasons = @()

  if ($classification -eq "" -and (-not $present) -and ($hasXmlTree -or $hasLuaAnalysis -or $hasLinks)) {
    $classification = "stale artifact"
    $reasons += "artifact exists without source addon directory"
  }

  if ($realLinks -and (-not $hasXmlTree -or -not $hasLuaAnalysis)) {
    if ($classification -eq "") { $classification = "true contract inconsistency" }
    $reasons += "real xml-lua links require both xml-tree and lua-analysis artifacts"
  }

  if ($present -and $hasMod -and $hasListedXml -and $existingListedXml -gt 0 -and (-not $hasXmlTree)) {
    if ($classification -eq "") { $classification = "manifest/discovery issue" }
    $reasons += "mod+existing listed XML but xml-tree artifact missing"
  }

  if ($present -and $hasMod -and $hasListedLua -and $existingListedLua -gt 0 -and (-not $hasLuaAnalysis)) {
    if ($classification -eq "") { $classification = "manifest/discovery issue" }
    $reasons += "mod+existing listed Lua but lua-analysis artifact missing"
  }

  if ($hasLinks -and $zeroLinks -and (-not $hasLuaAnalysis) -and $present -and (($hasListedLua -eq $false) -or ($existingListedLua -eq 0))) {
    if ($classification -eq "") { $classification = "legitimate empty artifact" }
    $reasons += "xml-lua-links empty and no effective Lua source contract"
  }

  if ($hasLinks -and $zeroLinks -and $present -and $hasListedLua -and $existingListedLua -gt 0 -and (-not $hasLuaAnalysis)) {
    if ($classification -eq "") { $classification = "generator bug" }
    $reasons += "effective Lua source exists but lua-analysis missing while links artifact exists"
  }

  $addonLabel = ""
  if ($present) {
    $addonLabel = $srcRec.SourceDir
  } elseif ($hasLinks) {
    $addonLabel = $linkRec.Names[0]
  } elseif ($hasXmlTree) {
    $addonLabel = $xmlRec.Names[0]
  } else {
    $addonLabel = $luaRec.Names[0]
  }

  $rows += [pscustomobject]@{
    Addon = $addonLabel
    CanonicalKey = $k
    PresentInSourceInventory = $present
    HasMod = $hasMod
    HasManifestListedXML = $hasListedXml
    HasManifestListedLua = $hasListedLua
    ExistingListedXMLCount = $existingListedXml
    ExistingListedLuaCount = $existingListedLua
    HasXMLTreeArtifact = $hasXmlTree
    HasLuaAnalysisArtifact = $hasLuaAnalysis
    HasXMLLuaLinksArtifact = $hasLinks
    XMLLuaLinksZeroLinks = $zeroLinks
    XMLLuaLinksRealLinks = $realLinks
    DisplayNameVariation = $false
    Informational = ""
    Classification = $classification
    NameVariants = ($displayNames -join "; ")
    Notes = ($reasons -join " | ")
  }
}

$anomalies = @($rows | Where-Object { $_.Classification -ne "" })
$informationalRows = @()

$summary = [pscustomobject]@{
  TotalAddonsUniverse = $rows.Count
  SourceInventory = @($rows | Where-Object { $_.PresentInSourceInventory }).Count
  WithMod = @($rows | Where-Object { $_.HasMod }).Count
  WithManifestListedXML = @($rows | Where-Object { $_.HasManifestListedXML }).Count
  WithManifestListedLua = @($rows | Where-Object { $_.HasManifestListedLua }).Count
  XMLTreeArtifacts = @($rows | Where-Object { $_.HasXMLTreeArtifact }).Count
  LuaAnalysisArtifacts = @($rows | Where-Object { $_.HasLuaAnalysisArtifact }).Count
  XMLLuaLinksArtifacts = @($rows | Where-Object { $_.HasXMLLuaLinksArtifact }).Count
  LinksZero = @($rows | Where-Object { $_.XMLLuaLinksZeroLinks }).Count
  LinksReal = @($rows | Where-Object { $_.XMLLuaLinksRealLinks }).Count
  InformationalDisplayNameVariationOnly = 0
  IdentitySplit = 0
  TotalAnomalies = $anomalies.Count
}

if ($AnomaliesCsv -eq "") {
  $AnomaliesCsv = Join-Path $env:TEMP "contract_audit_anomalies.csv"
}
if ($InformationalCsv -eq "") {
  $InformationalCsv = Join-Path $env:TEMP "contract_audit_informational.csv"
}

$anomalies | Sort-Object Addon | Export-Csv -LiteralPath $AnomaliesCsv -NoTypeInformation -Encoding UTF8
$informationalRows | Sort-Object Addon | Export-Csv -LiteralPath $InformationalCsv -NoTypeInformation -Encoding UTF8

"===SUMMARY==="
$summary | Format-List
"===ANOMALY_COUNTS==="
$anomalies | Group-Object Classification | Sort-Object Count -Descending | Select-Object Name,Count | Format-Table -AutoSize
"===INFORMATIONAL_COUNTS==="
if ($informationalRows.Count -eq 0) {
  "(none)"
} else {
  $informationalRows | Group-Object Informational | Select-Object Name,Count | Format-Table -AutoSize
}
"AnomaliesCsv=$AnomaliesCsv"
"InformationalCsv=$InformationalCsv"
