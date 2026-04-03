# WindowSetDrawWhenInterfaceHidden

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Pure, Squared |
| Files seen in | Source/PurePlayer.lua, Source/PurePlayerPet.lua, Source/PurePlayerPetTarget.lua, Source/PureTarget.lua, SquaredUnitFrame.lua |
| Namespaces detected | WindowSetDrawWhenInterfaceHidden |
| Source kinds | lua_calls |
| Example locations | Pure: LoadUnitFrame, Squared: Create, Squared: Update |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
WindowSetDrawWhenInterfaceHidden(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: friendlyFrame:GetName(), hostileFrame:GetName(), petFrame:GetName() |
| arg2 | Observed as a function or method reference. | Observed values: Pure.Get("friendly-frame-showinsiege"), Pure.Get("hostile-frame-showinsiege"), Pure.Get("player-frame-showinsiege") |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Pure
- Squared

## Examples

- Pure: LoadUnitFrame -> WindowSetDrawWhenInterfaceHidden(playerFrame:GetName(), Pure.Get("player-frame-showinsiege"))
- Pure: LoadUnitFrame -> WindowSetDrawWhenInterfaceHidden(targetFrame:GetName(), Pure.Get("playerpettarget-frame-showinsiege"))
- Pure: LoadUnitFrame -> WindowSetDrawWhenInterfaceHidden(hostileFrame:GetName(), Pure.Get("hostile-frame-showinsiege"))
- Pure: LoadUnitFrame -> WindowSetDrawWhenInterfaceHidden(friendlyFrame:GetName(), Pure.Get("friendly-frame-showinsiege"))
- Pure: LoadUnitFrame -> WindowSetDrawWhenInterfaceHidden(petFrame:GetName(), Pure.Get("playerpet-frame-showinsiege"))
- Squared: Create -> WindowSetDrawWhenInterfaceHidden(wName, window.alwaysdraw)

## Affects

- [SystemData.Events.BATTLEGROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_PLAYER_ADDED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_PLAYER_ADDED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_STATUS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_PET_HEALTH_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_PET_HEALTH_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_BEGIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_BEGIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_END](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_JOIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_JOIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field

## Notes

- none
