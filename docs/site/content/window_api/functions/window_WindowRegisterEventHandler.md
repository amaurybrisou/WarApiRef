# WindowRegisterEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 24 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionFraction, Atlas, BuddyBind, Busted, CaVES, Calling, CleanUnitFrames, EA_ScenarioGroupWindow |
| Files seen in | BuddyBind.lua, Busted.lua, CallingKeybinding.lua, CleanTargetWindow.lua, MGRemix.lua, Map.lua, Map/Main.lua, MiracleGrow.lua |
| Namespaces detected | WindowRegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | ActionFraction: Initialize, Atlas: RegisterEventHandlers, BuddyBind: OnLButtonRawDeviceInput, Busted: Initialize, CaVES: Initialize, Calling: StartBinding |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 80 |
| Global usage count | 80 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
WindowRegisterEventHandler(windowName, eventId, handlerName)
```

## Description

Observed binding SystemData events directly to a named window.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AtlasFrame", "AuctionWindow", "BustedProxy" |
| eventId | Observed as a SystemData event identifier. | Observed values: SystemData.Events.AUCTION_BID_RESULT_RECEIVED, SystemData.Events.AUCTION_INIT_RECEIVED, SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED |
| handlerName | Observed as a Lua handler reference. | Observed values: "ActionFractionWindow.OnAgroModeUpdated", "ActionFractionWindow.UpdateActionPoints", "ActionFractionWindow.UpdateBasedOnUserSettings" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- ActionFraction
- Atlas
- BuddyBind
- Busted
- CaVES
- Calling
- CleanUnitFrames
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- HealGrid
- Map
- Minmap
- Miracle Grow Remix
- MiracleGrow
- Moth
- Phantom
- RVMOD_PlayerStatus
- RVMOD_Targets
- ResHelp
- SNT_CASTBAR
- Shinies
- WARCommander
- XpStatus+G

## Examples

- ActionFraction: Initialize -> WindowRegisterEventHandler(windowName, SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED, "ActionFractionWindow.UpdateActionPoints")
- ActionFraction: Initialize -> WindowRegisterEventHandler(windowName, SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED, "ActionFractionWindow.UpdateActionPoints")
- ActionFraction: Initialize -> WindowRegisterEventHandler(windowName, SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED, "ActionFractionWindow.UpdateCurrentHitPoints")
- ActionFraction: Initialize -> WindowRegisterEventHandler(windowName, SystemData.Events.PLAYER_AGRO_MODE_UPDATED, "ActionFractionWindow.OnAgroModeUpdated")
- ActionFraction: Initialize -> WindowRegisterEventHandler(windowName, SystemData.Events.PLAYER_HEALTH_FADE_UPDATED, "ActionFractionWindow.UpdateBasedOnUserSettings")
- Atlas: RegisterEventHandlers -> WindowRegisterEventHandler("AtlasFrame", SystemData.Events.UPDATE_PROCESSED, "AtlasMap.OnUpdate")

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HEALTH_FADE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HEALTH_FADE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_POSITION_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_POSITION_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed
