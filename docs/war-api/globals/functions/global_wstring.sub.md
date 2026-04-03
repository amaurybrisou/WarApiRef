# wstring.sub

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 63/100
- Seen in: 2 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 63/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, PartyCast |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:267`, `/workspace/data/raw/Moth/MothHelpers.lua:39`, `/workspace/data/raw/PartyCast/PartyCast.lua:370`, `/workspace/data/raw/PartyCast/PartyCast.lua:377` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.UpdateLevel, Moth: MothHelpers.CapitalizeWString, PartyCast: PartyCast.GROUP_UPDATED, PartyCast: PartyCast.ON_DEATH |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: GameData.Player.name, unitTierDesc, wstr |
| arg2 | Observed as a numeric value. | Observed values: 1, 2 |
| arg3 | Observed as a numeric value. | Observed values: -3, 1 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Moth
- PartyCast

## Examples

- Moth: Moth.UpdateLevel -> wstring.sub(unitTierDesc, 1, 1)
- Moth: MothHelpers.CapitalizeWString -> wstring.sub(wstr, 2)
- Moth: MothHelpers.CapitalizeWString -> wstring.sub(wstr, 1, 1)
- PartyCast: PartyCast.GROUP_UPDATED -> wstring.sub(GameData.Player.name, 1, -3)
- PartyCast: PartyCast.ON_DEATH -> wstring.sub(GameData.Player.name, 1, -3)

## Related APIs

- none

## Used With

- [wstring.lower](global_wstring.lower.md) (MEDIUM 45/100) - Global Function
- [wstring.upper](global_wstring.upper.md) (MEDIUM 45/100) - Global Function

## Triggered By

- [SystemData.Events.ENTER_WORLD](../../events/game_events/game_event_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERFACE_RELOADED](../../events/game_events/game_event_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_DEATH](../../events/game_events/game_event_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - Game Event

## Affects

- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
