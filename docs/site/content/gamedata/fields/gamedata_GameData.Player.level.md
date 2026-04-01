# GameData.Player.level

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 105

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy |
| Files seen in | `/workspace/Effigy/States/EffigyStatePlayer.lua:244`, `/workspace/Effigy/States/EffigyStatePlayer.lua:320`, `/workspace/Effigy/States/EffigyStatePlayer.lua:332`, `/workspace/Effigy/States/EffigyStatePlayer.lua:364`, `/workspace/Effigy/States/EffigyStateTargets.lua:325` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | Effigy.UpdateExp, Effigy.UpdatePlayerHP, Effigy.UpdatePlayerLevel, Effigy.UpdateRestedExp, Effigy.UpdateTarget, lua_call |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

Observed GameData field used by 1 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Effigy

## Related APIs

- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionTrigger](../../window_api/functions/window_WindowSetGameActionTrigger.md) (HIGH 100/100) - Window Function

## Used With

- [GameData.PlayerActions.SET_TARGET](gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionTrigger](../../window_api/functions/window_WindowSetGameActionTrigger.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: Effigy.UpdateExp, Effigy.UpdatePlayerHP, Effigy.UpdatePlayerLevel, Effigy.UpdateRestedExp, Effigy.UpdateTarget, lua_call
