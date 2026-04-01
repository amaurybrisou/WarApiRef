# GameData.Player.worldObjNum

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

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
| Addons seen in | BuffHead, Effigy, LibGuard |
| Files seen in | `/workspace/BuffHead/Core.lua:1088`, `/workspace/BuffHead/Core.lua:207`, `/workspace/Effigy/States/EffigyStatePlayer.lua:244`, `/workspace/Effigy/States/EffigyStatePlayer.lua:271`, `/workspace/Effigy/States/EffigyStatePlayer.lua:283`, `/workspace/Effigy/States/EffigyStatePlayer.lua:294`, `/workspace/Effigy/States/EffigyStatePlayer.lua:373`, `/workspace/Effigy/States/EffigyStatePlayer.lua:387` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | BuffHead.OnGroupEffectsUpdated, BuffHead.SetSelfTracking, CB_Update, Effigy.UpdateName, Effigy.UpdatePlayerAP, Effigy.UpdatePlayerCareer |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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

Observed GameData field used by 3 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- BuffHead
- Effigy
- LibGuard

## Related APIs

- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function
- [wstring.find](../../globals/functions/global_wstring.find.md) (MEDIUM 63/100) - Global Function

## Used With

- [GameData.Player.activeTitle](gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function
- [wstring.find](../../globals/functions/global_wstring.find.md) (MEDIUM 63/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: BuffHead.OnGroupEffectsUpdated, BuffHead.SetSelfTracking, CB_Update, Effigy.UpdateName, Effigy.UpdatePlayerAP, Effigy.UpdatePlayerCareer
