# SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

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
| Addons seen in | CM_ClosetGoblin, CraftingWillard, Crusher, Enemy, KeyBar, Miracle Grow Remix, Motion, NerfedButtons |
| Files seen in | ClosetGoblin.lua, Code/TalismanAlerter/TalismanAlerter.lua, CraftingWillard.lua, KeyBar.lua, MGRemix.lua, Modules/API/Shinies-API.lua, Modules/Data/Shinies-Data-Inventory.lua, Modules/UI/Shinies-UI-Post.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | Initialize, OnDisable, OnEnable, OnInitialize, OnLoad, OnShutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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

SystemData.SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED field accessed by 17 addons; commonly found in Initialize and OnDisable, OnEnable, OnInitialize, OnLoad, OnShutdown, RegisterEventsHandlers, Shutdown, _TalismanAlerterEnabledChanged, initEvents, lua_call contexts.

## Seen In

- CM_ClosetGoblin
- CraftingWillard
- Crusher
- Enemy
- KeyBar
- Miracle Grow Remix
- Motion
- NerfedButtons
- Pocket Palette
- PotionBar
- RoR_debolster
- Shinies
- WarBoard_TaliMon
- War_RU
- nLootLink
- talisman-monitor
- zMailMod

## Related APIs

- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: Initialize, OnDisable, OnEnable, OnInitialize, OnLoad, OnShutdown
