# wstring.sub

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 30 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuddyBind, CCTV, CleanUnitFrames, CombatTextNames, DammazKron, Effigy, Emojii, Enemy |
| Files seen in | BuddyBind.lua, CCTV.lua, CleanTargetUnitFrame.lua, Code/Core/Utils.lua, Core/DK_Core.lua, Core/DK_TargetInfo.lua, Elements/EffigyLabel.lua, Emojii.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | BuddyBind: GrabName, BuddyBind: update, CCTV: Initialize, CleanUnitFrames: UpdateUnit, CombatTextNames: TruncateAbilityName, DammazKron: InitializeLocal |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 146 |
| Global usage count | 146 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
wstring.sub(arg1, arg2, arg3)
```

## Description

Observed as a global function across 30 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: Ability_list, Emojii.chooseIconDialogListData[data_index][icon_index], GameData.ChatData.text |
| arg2 | Observed as a numeric value. | Observed values: -1, -2, 0 |
| arg3 | Observed as a numeric value. | Observed values: -1, -2, -3 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- BuddyBind
- CCTV
- CleanUnitFrames
- CombatTextNames
- DammazKron
- Effigy
- Emojii
- Enemy
- GetStats
- HealGrid
- I HATE YOU THIS MUCH
- Info_DeathBlow
- Info_Points
- KeyBar
- Killer
- Kwestor
- LibGuard
- Moth
- NaturalLog
- PartyCast
- Pure
- Queue Queuer
- Refer
- ResHelp
- Sequencer
- Squared
- TomeTracker
- WSCT
- WarBoard_WarWhisperer
- whatsPugSc

## Examples

- BuddyBind: GrabName -> wstring.sub(TargetInfo:UnitName("selffriendlytarget"), 1, -3)
- BuddyBind: update -> wstring.sub(TargetInfo:UnitName("selffriendlytarget"), 1, -3)
- CCTV: Initialize -> wstring.sub(GetStringFromTable("ComponentEffects",13104), 0, -2)
- CCTV: Initialize -> wstring.sub(GetStringFromTable("ComponentEffects",1526), 0, -2)
- CCTV: Initialize -> wstring.sub(GetStringFromTable("ComponentEffects",16201), 0, -2)
- CCTV: Initialize -> wstring.sub(GetStringFromTable("ComponentEffects",1489), 0, -2)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [wstring.len](global_wstring.len.md) (HIGH 100/100) - Global Function
- [wstring.reverse](global_wstring.reverse.md) (HIGH 88/100) - Global Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [wstring.find](global_wstring.find.md) (HIGH 75/100) - Global Function

## Affects

- [GameData.PlayerActions.DO_MACRO](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_MACRO.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
