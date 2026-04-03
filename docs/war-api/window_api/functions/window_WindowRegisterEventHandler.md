# WindowRegisterEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 35 addons

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
| Addons seen in | ActionFraction, Atlas, BuddyBind, Busted, CMap, CaVES, Calling, CleanUnitFrames |
| Files seen in | BuddyBind.lua, Busted.lua, CMap.lua, CallingKeybinding.lua, CleanTargetWindow.lua, EZCraft.lua, MGRemix.lua, Map.lua |
| Namespaces detected | WindowRegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | ActionFraction: Initialize, Atlas: RegisterEventHandlers, BuddyBind: OnLButtonRawDeviceInput, Busted: Initialize, CMap: Initialize, CaVES: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 154 |
| Global usage count | 154 |
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
| eventId | Observed as a SystemData event identifier. | Observed values: SystemData.Events.ALLIANCE_UPDATED, SystemData.Events.ALL_MODULES_INITIALIZED, SystemData.Events.AUCTION_BID_RESULT_RECEIVED |
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
- CMap
- CaVES
- Calling
- CleanUnitFrames
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraft
- EZCraftX
- FozAuction
- HealGrid
- Map
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MoraleSet
- Moth
- Phantom
- RVMOD_PlayerStatus
- RVMOD_Targets
- ResHelp
- Rolodex
- SNT_CASTBAR
- Shinies
- SocialWindow 2.0
- WARCommander
- WarBoard_WarWhisperer
- XpStatus+G
- bigger_MacroWindow

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

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed
