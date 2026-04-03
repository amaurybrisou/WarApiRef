# wstring.sub

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CombatTextNames, Enemy, Killer, LibGuard, PartyCast, WSCT |
| Files seen in | `/workspace/data/raw/Enemy/Code/Core/Utils.lua:684`, `/workspace/data/raw/Killer/KillerUtils.lua:169`, `/workspace/data/raw/LibGuard/Source/LibGuard.lua:111`, `/workspace/data/raw/PartyCast/PartyCast.lua:370`, `/workspace/data/raw/PartyCast/PartyCast.lua:377`, `/workspace/data/raw/combattextnames/combattextnames.lua:135`, `/workspace/data/raw/wsct/wsct.lua:596` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | CombatTextNames: CombatTextNames.TruncateAbilityName, Enemy: Enemy.FormatNumber, Killer: Killer.GetCurrentPlayerName, LibGuard: LibGuard.StartCast, PartyCast: PartyCast.GROUP_UPDATED, PartyCast: PartyCast.ON_DEATH |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
wstring.sub(arg1, arg2, arg3)
```

## Description

Observed as a global function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: GameData.Player.name, number_string, raw |
| arg2 | Observed as a numeric value. | Observed values: 0, 1, sub_begin |
| arg3 | Observed as a numeric value. | Observed values: -3, CombatTextNames.Settings.TruncateMinLength, db["TRUNCATESIZE"] |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- CombatTextNames
- Enemy
- Killer
- LibGuard
- PartyCast
- WSCT

## Examples

- CombatTextNames: CombatTextNames.TruncateAbilityName -> wstring.sub(text, 1, CombatTextNames.Settings.TruncateMinLength)
- Enemy: Enemy.FormatNumber -> wstring.sub(number_string, sub_begin, sub_end)
- Enemy: Enemy.FormatNumber -> wstring.sub(number_string, 0, sub_begin-1)
- Killer: Killer.GetCurrentPlayerName -> wstring.sub(raw, 1, -3)
- LibGuard: LibGuard.StartCast -> wstring.sub(towstring(TargetInfo:UnitName("selffriendlytarget")), 1, -3)
- PartyCast: PartyCast.GROUP_UPDATED -> wstring.sub(GameData.Player.name, 1, -3)

## Related APIs

- none

## Used With

- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.len](global_wstring.len.md) (HIGH 100/100) - Global Function
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function

## Triggered By

- [SystemData.Events.ENTER_WORLD](../../events/game_events/game_event_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERFACE_RELOADED](../../events/game_events/game_event_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BEGIN_CAST](../../events/game_events/game_event_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - Game Event

## Affects

- [GameData.Player](../../gamedata/fields/gamedata_GameData.Player.md) (HIGH 100/100) - GameData Field
- [GameData.Player.career.line](../../gamedata/fields/gamedata_GameData.Player.career.line.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
