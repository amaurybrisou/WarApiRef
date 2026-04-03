# WindowRegisterEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 108

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Shinies |
| Files seen in | `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Browse.lua:463`, `/workspace/data/raw/Shinies/Source/Shinies.lua:169` |
| Namespaces detected | WindowRegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | Shinies: LibStub:UpdateDefaultAuctionHouseDisable, Shinies: _G.Shinies:NewModule.OnSelChanged_Criteria_MultiSelCombo |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | no |
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
| windowName | Observed as a target window name. | Observed values: "AuctionWindow", lastComboSelected |
| eventId | Observed as a SystemData event identifier. | Observed values: SystemData.Events.AUCTION_BID_RESULT_RECEIVED, SystemData.Events.AUCTION_INIT_RECEIVED, SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED |
| handlerName | Observed as a Lua handler reference. | Observed values: "AuctionWindow.Hide", "AuctionWindow.OnSearchResultsReceived", "AuctionWindow.ReceivedBidResult" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Shinies

## Examples

- Shinies: LibStub:UpdateDefaultAuctionHouseDisable -> WindowRegisterEventHandler("AuctionWindow", SystemData.Events.AUCTION_INIT_RECEIVED, "AuctionWindow.Show")
- Shinies: LibStub:UpdateDefaultAuctionHouseDisable -> WindowRegisterEventHandler("AuctionWindow", SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED, "AuctionWindow.OnSearchResultsReceived")
- Shinies: LibStub:UpdateDefaultAuctionHouseDisable -> WindowRegisterEventHandler("AuctionWindow", SystemData.Events.AUCTION_BID_RESULT_RECEIVED, "AuctionWindow.ReceivedBidResult")
- Shinies: LibStub:UpdateDefaultAuctionHouseDisable -> WindowRegisterEventHandler("AuctionWindow", SystemData.Events.INTERACT_DONE, "AuctionWindow.Hide")
- Shinies: _G.Shinies:NewModule.OnSelChanged_Criteria_MultiSelCombo -> WindowRegisterEventHandler(lastComboSelected, SystemData.Events.L_BUTTON_UP_PROCESSED, "ShiniesBrowseUI.Criteria_ReopenComboBox")

## Related APIs

- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](../../events/game_events/game_event_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.AUCTION_INIT_RECEIVED](../../events/game_events/game_event_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../events/game_events/game_event_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_DONE](../../events/game_events/game_event_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../events/game_events/game_event_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - Game Event

## Used With

- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_INIT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_INIT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
