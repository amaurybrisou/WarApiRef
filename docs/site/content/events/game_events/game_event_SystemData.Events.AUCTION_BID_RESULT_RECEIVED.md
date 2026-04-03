# SystemData.Events.AUCTION_BID_RESULT_RECEIVED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Shinies |
| Files seen in | `/workspace/data/raw/Shinies/Source/Shinies.lua:169` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Shinies: LibStub:UpdateDefaultAuctionHouseDisable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 1 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
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

Observed as a shared SystemData runtime event used by 1 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Shinies

## Registrars And Handlers

- AuctionWindow.ReceivedBidResult
- RegisterEventHandler
- ShiniesAPI.Core_OnAuctionBidResultReceived
- WindowRegisterEventHandler
- window

## Examples

- Shinies: LibStub:UpdateDefaultAuctionHouseDisable -> SystemData.Events.AUCTION_BID_RESULT_RECEIVED -> AuctionWindow.ReceivedBidResult
- Shinies: AuctionWindow.ReceivedBidResult -> WindowRegisterEventHandler(SystemData.Events.AUCTION_BID_RESULT_RECEIVED, AuctionWindow.ReceivedBidResult)
- Shinies: ShiniesAPI.Core_OnAuctionBidResultReceived -> RegisterEventHandler(SystemData.Events.AUCTION_BID_RESULT_RECEIVED, ShiniesAPI.Core_OnAuctionBidResultReceived)

## Related APIs

- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current addon-api corpus.
