# SystemData.Events.INTERACT_DONE

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, FozAuction, HealGrid, JunkDump, PartyCast, SNT_CASTBAR, Shinies, SquaredClick |
| Files seen in | Bars/HealGridCastBar.lua, JunkDump.lua, PartyCast.lua, Source/Shinies.lua, Source/SquaredClick.lua, Source/auctionwindow.lua, States/EffigyStateCastbar.lua, snt_castbar.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | Disable, Enable, Init, Initialize, OnDisable, OnEnable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
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

SystemData.SystemData.Events.INTERACT_DONE field accessed by 9 addons; commonly found in Disable and Enable, Init, Initialize, OnDisable, OnEnable, OnShutdown, RegisterStateInfoForCastbar, StartListeners, UpdateDefaultAuctionHouseDisable, entry_point, lua_call contexts.

## Seen In

- Effigy
- FozAuction
- HealGrid
- JunkDump
- PartyCast
- SNT_CASTBAR
- Shinies
- SquaredClick
- xHUD

## Related APIs

- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: Disable, Enable, Init, Initialize, OnDisable, OnEnable
