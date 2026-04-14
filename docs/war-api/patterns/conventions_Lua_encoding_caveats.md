# Lua encoding caveats

- Category: conventions
- Confidence: HIGH

## Description

Implementation-validated encoding and tooling findings for WAR Lua addon files.

## Involved APIs

- [Text](../xml/element_types/element_Text.md) (HIGH 100/100) - XML Element Type
- [Start](../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event
- [Stop](../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Flow Diagram

```text
Start -> RegisterEventHandler
```

## Example Code

```lua
`HIGH`: The WAR engine Lua parser rejects .lua files that begin with a UTF-8 BOM (byte sequence 0xEF 0xBB 0xBF). The error message is: Error Reading Lua buffer <file>: [string ...]:1: '=' expected near '»'. The '»' character is the visual artefact of the BOM rendered in the error log.
```

## Evidence

- `HIGH`: The WAR engine Lua parser rejects .lua files that begin with a UTF-8 BOM (byte sequence 0xEF 0xBB 0xBF). The error message is: Error Reading Lua buffer <file>: [string ...]:1: '=' expected near '»'. The '»' character is the visual artefact of the BOM rendered in the error log.
- Guidance: Always save Lua addon files as UTF-8 without BOM. In PowerShell 5.1, do NOT use Set-Content -Encoding UTF8 (it adds a BOM). Use [System.IO.File]::WriteAllText(path, content, [System.Text.UTF8Encoding]::new($false)) instead. To verify: read the first 3 bytes — they must NOT be 239 187 191.
- `HIGH`: WAR's SavedVariables engine persists partial or malformed config tables across sessions. If a previous version of an addon wrote an incomplete table (missing expected sub-keys), the next version receives the partial table from LoadCharacterSettings and all missing keys are nil — even though the top-level character entry exists. Nil checks on the entry itself are insufficient.
- Guidance: After calling the character-profile load function (e.g. AuraProfile.LoadCharacterSettings), validate every expected sub-key and replace nil values with their defaults before using them. Guard at least: RuntimeAuras (init to {}), RuntimeSettings (copy DefaultSettings), Enabled (copy DefaultConfiguration.Enabled), Debug (copy DefaultConfiguration.Debug). This is the correct boundary — not inside the library.
- `HIGH`: Any Lua function that iterates a WAR addon config table with pairs() or ipairs() must guard the table for nil before the call, even if the table was initialized earlier. SavedVariables corruption or partial-load scenarios mean the table can be nil at any call site that runs before a full validation pass.
- Guidance: Add an explicit `if table == nil then return end` (or equivalent early-return) before every pairs()/ipairs() call site that touches config data loaded from SavedVariables. Affected patterns in Aura: PrepareSettingsForRuntime, CleanInternalMembers (ipairs over RuntimeAuras), CreateListDisplayData (pairs over RuntimeAuras).
- `HIGH`: AuraShares iterates AuraAddon.Configuration for all characters and accesses data.Auras. Not all character entries are guaranteed to have an Auras sub-table — entries written by a broken or partial implementation will have a nil Auras key. The guard must be: `if( data ~= nil and data.Auras ~= nil )` — not just `if( data ~= nil )`.
- Guidance: In any AuraShares loop over the per-character config table, check both the entry and the Auras sub-table before iterating. Failure to guard data.Auras produces a `attempt to call a nil value (pairs)` crash at runtime.
- `HIGH`: AuraAddon.CleanInternalMembers() iterates RuntimeAuras and calls Aura:CleanInternalMembers() on each, which destroys all runtime overlay windows (internal-runtimewindowid, internal-runtimetimerwindowid, internal-runtimeflashwindowid) and clears their IDs from aura.Data. If the engine is still running when this is called, the windows are gone but engineRunning remains true, so AuraEngine.Start() becomes a no-op and the windows are never rebuilt. Aura icons stop displaying.
- Guidance: Never call AuraAddon.CleanInternalMembers() while the engine is running unless you immediately follow it with AuraEngine.Stop() + AuraEngine.Start() (as SwitchProfile correctly does). For snapshot/save operations that do not switch profiles, copy RuntimeAuras with AuraHelpers.CopyTable() and strip internal-* keys from the copy only — do not touch the live runtime table.
- `HIGH`: To safely snapshot RuntimeAuras for persistence (without affecting the live engine), copy with AuraHelpers.CopyTable(), then iterate the copy and remove keys matching '^internal.*' from each aura's Data sub-table. Use a two-pass approach (collect keys in one table, then remove) to avoid modifying the table while iterating it.
- Guidance: Pattern: `local copy = AuraHelpers.CopyTable(AuraAddon.RuntimeAuras)` then for each entry `for _, auraCopy in ipairs(copy) do -- collect internal keys, then nil them in auraCopy.Data end`. This is the correct substitute for calling CleanInternalMembers() before a CopyTable when the engine must remain running.
