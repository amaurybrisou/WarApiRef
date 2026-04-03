# EA_Window_Backpack.GetItemsFromBackpack

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 156

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BankArkel, Killer, Shinies |
| Files seen in | `/workspace/data/raw/BankArkel/BankArkel.lua:371`, `/workspace/data/raw/Killer/KillerRenown.lua:108`, `/workspace/data/raw/Killer/KillerRenown.lua:127`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Browse.lua:298` |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | globals, lua_calls |
| Example locations | BankArkel: BankArkel.UpdateData, Killer: Killer.CaptureInitialWarCrests, Killer: Killer.OnCurrencyUpdated, Shinies: _G.Shinies:NewModule.EA_Window_Backpack_EquipmentLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
EA_Window_Backpack.GetItemsFromBackpack(arg1)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: EA_Window_Backpack.TYPE_CURRENCY, EA_Window_Backpack.currentMode, Read[i] |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- BankArkel
- Killer
- Shinies

## Examples

- BankArkel: BankArkel.UpdateData -> EA_Window_Backpack.GetItemsFromBackpack(Read[i])
- Killer: Killer.CaptureInitialWarCrests -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.TYPE_CURRENCY)
- Killer: Killer.OnCurrencyUpdated -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.TYPE_CURRENCY)
- Shinies: _G.Shinies:NewModule.EA_Window_Backpack_EquipmentLButtonDown -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.currentMode)

## Related APIs

- none

## Used With

- none

## Triggered By

- [SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED.md) (HIGH 100/100) - Game Event

## Affects

- [GameData.Player](../../gamedata/fields/gamedata_GameData.Player.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
