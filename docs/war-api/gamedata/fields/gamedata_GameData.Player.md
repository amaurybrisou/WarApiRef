# GameData.Player

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Killer, Swift Assist |
| Files seen in | `/workspace/data/raw/Killer/KillerRenown.lua:108`, `/workspace/data/raw/Killer/KillerRenown.lua:127`, `/workspace/data/raw/Killer/KillerRenown.lua:61`, `/workspace/data/raw/Killer/KillerZoneHistory.lua:329`, `/workspace/data/raw/Killer/KillerZoneHistory.lua:370`, `/workspace/data/raw/Killer/KillerZoneHistory.lua:446`, `/workspace/data/raw/Killer/KillerZoneHistory.lua:5`, `/workspace/data/raw/swift-assist/SwiftAssist.lua:271` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | Killer.CaptureInitialWarCrests, Killer.GetCurrentZoneHistoryEntry, Killer.LoadCurrentZoneFromHistory, Killer.OnCurrencyUpdated, Killer.OnRenownUpdated, Killer.UpsertZoneHistory |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed GameData field used by 2 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Killer
- Swift Assist

## Related APIs

- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 100/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: Killer.CaptureInitialWarCrests, Killer.GetCurrentZoneHistoryEntry, Killer.LoadCurrentZoneFromHistory, Killer.OnCurrencyUpdated, Killer.OnRenownUpdated, Killer.UpsertZoneHistory
