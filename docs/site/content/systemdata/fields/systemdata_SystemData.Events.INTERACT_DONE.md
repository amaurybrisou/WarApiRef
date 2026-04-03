# SystemData.Events.INTERACT_DONE

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 186

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | PartyCast, PartyCast, Shinies, Shinies |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.lua:153`, `/workspace/data/raw/PartyCast/PartyCast.lua:51`, `/workspace/data/raw/Shinies/Source/Shinies.lua:169`, `/workspace/data/raw/Shinies/Source/Shinies.lua:206`, `/workspace/data/raw/Shinies/Source/Shinies.lua:219` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | AuctionWindow.Hide, LibStub:OnDisable, LibStub:OnEnable, LibStub:UpdateDefaultAuctionHouseDisable, PartyCast.EndCast, PartyCast.Init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 2 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | yes |
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

Observed SystemData field used by 3 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- PartyCast
- PartyCast, Shinies
- Shinies

## Related APIs

- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 90/100) - Window Function

## Used With

- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](systemdata_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_INIT_RECEIVED](systemdata_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](systemdata_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AuctionWindow.Hide, LibStub:OnDisable, LibStub:OnEnable, LibStub:UpdateDefaultAuctionHouseDisable, PartyCast.EndCast, PartyCast.Init
