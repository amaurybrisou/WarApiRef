# UnregisterEventHandler

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 81/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 81/100

- Rationale: Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, used in event registration or dispatch.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | PartyCast, TidyRoll |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.lua:153`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:249` |
| Namespaces detected | UnregisterEventHandler |
| Source kinds | lua_calls |
| Example locations | PartyCast: PartyCast.OnShutdown, TidyRoll: TidyRoll.Shutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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
UnregisterEventHandler(eventId, handlerName)
```

## Description

Observed removing previously registered global runtime handlers.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: SystemData.Events.ENTER_WORLD, SystemData.Events.INTERACT_DONE, SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA |
| handlerName | Observed as a Lua handler function reference. | Observed values: "PartyCast.EndCast", "PartyCast.GROUP_UPDATED", "PartyCast.ON_DEATH" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- PartyCast
- TidyRoll

## Examples

- PartyCast: PartyCast.OnShutdown -> UnregisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
- PartyCast: PartyCast.OnShutdown -> UnregisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")
- PartyCast: PartyCast.OnShutdown -> UnregisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, "PartyCast.StartCast")
- PartyCast: PartyCast.OnShutdown -> UnregisterEventHandler(SystemData.Events.PLAYER_END_CAST, "PartyCast.EndCast")
- PartyCast: PartyCast.OnShutdown -> UnregisterEventHandler(SystemData.Events.PLAYER_CAST_TIMER_SETBACK, "PartyCast.SetbackCast")
- PartyCast: PartyCast.OnShutdown -> UnregisterEventHandler(SystemData.Events.PLAYER_DEATH, "PartyCast.ON_DEATH")

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
