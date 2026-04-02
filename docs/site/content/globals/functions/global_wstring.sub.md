# wstring.sub

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 8 addons

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
| Addons seen in | CombatTextNames, Effigy, Enemy, GetStats, Killer, LibGuard, Moth, WSCT |
| Files seen in | `/workspace_addons/Effigy/Elements/EffigyLabel.lua:16`, `/workspace_addons/Effigy/Elements/EffigyLabel.lua:40`, `/workspace_addons/Effigy/States/EffigyStatePlayer.lua:294`, `/workspace_addons/Enemy/Code/Core/Utils.lua:684`, `/workspace_addons/GetStats/GetStats.lua:119`, `/workspace_addons/Killer/KillerUtils.lua:169`, `/workspace_addons/LibGuard/Source/LibGuard.lua:111`, `/workspace_addons/Moth/Moth.lua:270` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | CombatTextNames: CombatTextNames.TruncateAbilityName, Effigy: Effigy.UpdateTitle, Effigy: Effigy.local.WStringSplitInclDelimiter, Effigy: Effigy.local.wstringStartsWith, Effigy: WStringSplitInclDelimiter, Effigy: wstringStartsWith |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
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

Observed as a global function across 8 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: WString, inString, number_string |
| arg2 | Observed as a numeric value. | Observed values: 0, 1, 11 |
| arg3 | Observed as a numeric value. | Observed values: -3, 1, CombatTextNames.Settings.TruncateMinLength |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- CombatTextNames
- Effigy
- Enemy
- GetStats
- Killer
- LibGuard
- Moth
- WSCT

## Examples

- CombatTextNames: CombatTextNames.TruncateAbilityName -> wstring.sub(text, 1, CombatTextNames.Settings.TruncateMinLength)
- Effigy: Effigy.UpdateTitle -> wstring.sub(titleData.name, offset+1)
- Effigy: Effigy.local.WStringSplitInclDelimiter -> wstring.sub(inString, pos-1, first-1)
- Effigy: Effigy.local.WStringSplitInclDelimiter -> wstring.sub(inString, pos-1)
- Effigy: Effigy.local.wstringStartsWith -> wstring.sub(WString, 1, wstring.len(Start))
- Effigy: WStringSplitInclDelimiter -> wstring.sub(inString, pos-1, first-1)

## Related APIs

- none

## Used With

- [GameData.Player.activeTitle](../../gamedata/fields/gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [GameData.Player.worldObjNum](../../gamedata/fields/gamedata_GameData.Player.worldObjNum.md) (HIGH 100/100) - GameData Field
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.len](global_wstring.len.md) (HIGH 100/100) - Global Function
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.find](global_wstring.find.md) (MEDIUM 63/100) - Global Function
- [wstring.lower](global_wstring.lower.md) (MEDIUM 45/100) - Global Function
- [wstring.upper](global_wstring.upper.md) (MEDIUM 45/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.Player.activeTitle](../../gamedata/fields/gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [GameData.Player.worldObjNum](../../gamedata/fields/gamedata_GameData.Player.worldObjNum.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
