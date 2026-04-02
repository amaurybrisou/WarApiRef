# WindowRegisterEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | Busted, EA_UiDebugTools, Miracle Grow Remix, MiracleGrow, Moth, Shinies |
| Files seen in | `/workspace_addons/Busted/Busted.lua:324`, `/workspace_addons/MGRemix/MGRemix.lua:195`, `/workspace_addons/MiracleGrow/MiracleGrow.lua:155`, `/workspace_addons/Moth/Moth.lua:721`, `/workspace_addons/Shinies/Modules/UI/Shinies-UI-Browse.lua:463`, `/workspace_addons/Shinies/Source/Shinies.lua:169`, `/workspace_addons/ea_uidebugtools/Source/bust/Busted.lua:331` |
| Namespaces detected | WindowRegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | Busted: Busted.Initialize, EA_UiDebugTools: Busted.Initialize, Miracle Grow Remix: MiracleGrow2.Initialize, MiracleGrow: MiracleGrow.enableAddon, Moth: Moth.Initialize, Shinies: LibStub:UpdateDefaultAuctionHouseDisable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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
| windowName | Observed as a target window name. | Observed values: "AuctionWindow", "BustedProxy", "Moth" |
| eventId | Observed as a SystemData event identifier. | Observed values: SystemData.Events.AUCTION_BID_RESULT_RECEIVED, SystemData.Events.AUCTION_INIT_RECEIVED, SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED |
| handlerName | Observed as a Lua handler reference. | Observed values: "AuctionWindow.Hide", "AuctionWindow.OnSearchResultsReceived", "AuctionWindow.ReceivedBidResult" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Busted
- EA_UiDebugTools
- Miracle Grow Remix
- MiracleGrow
- Moth
- Shinies

## Examples

- Busted: Busted.Initialize -> WindowRegisterEventHandler("BustedProxy", uiLogEventId, "Busted.ProcessUiLogUpdateProxy")
- EA_UiDebugTools: Busted.Initialize -> WindowRegisterEventHandler("BustedProxy", uiLogEventId, "Busted.ProcessUiLogUpdateProxy")
- Miracle Grow Remix: MiracleGrow2.Initialize -> WindowRegisterEventHandler(mg.sWindowName, SystemData.Events.PLAYER_CULTIVATION_UPDATED, "MiracleGrow2.onCultUpdate")
- MiracleGrow: MiracleGrow.enableAddon -> WindowRegisterEventHandler(windowName, SystemData.Events.PLAYER_CULTIVATION_UPDATED, "MiracleGrow.fired")
- Moth: Moth.Initialize -> WindowRegisterEventHandler("Moth", SystemData.Events.PLAYER_TARGET_UPDATED, "Moth.UpdateTarget")
- Shinies: LibStub:UpdateDefaultAuctionHouseDisable -> WindowRegisterEventHandler("AuctionWindow", SystemData.Events.AUCTION_INIT_RECEIVED, "AuctionWindow.Show")

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_INIT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CULTIVATION_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CULTIVATION_UPDATED.md) (HIGH 100/100) - SystemData Field
- [TextLogGetUpdateEventId](../../events/game_events/game_event_TextLogGetUpdateEventId.md) (HIGH 100/100) - Game Event
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event

## Affects

- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_INIT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CRAFTING_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CRAFTING_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CULTIVATION_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CULTIVATION_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
