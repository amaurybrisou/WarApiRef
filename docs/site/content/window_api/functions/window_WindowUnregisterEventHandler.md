# WindowUnregisterEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 14 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionFraction, Atlas, BuddyBind, CaVES, Calling, HealGrid, ManualScenarioRefresh, Map |
| Files seen in | BuddyBind.lua, CallingKeybinding.lua, ManualScenarioRefresh.lua, Map.lua, Map/Main.lua, MiracleGrow.lua, Modules/UI/Shinies-UI-Browse.lua, PhantomLib/PhantomLib.lua |
| Namespaces detected | WindowUnregisterEventHandler |
| Source kinds | lua_calls |
| Example locations | ActionFraction: Shutdown, Atlas: Shutdown, BuddyBind: OnExitBindingMode, BuddyBind: OnRawDeviceInput, CaVES: Shutdown, Calling: OnRawDeviceInput |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 37 |
| Global usage count | 37 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
WindowUnregisterEventHandler(arg1, arg2)
```

## Description

Observed as a window function across 14 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "AtlasFrame", "AuctionWindow", "CaVESWindow" |
| arg2 | Observed as an event identifier. | Observed values: SystemData.Events.AUCTION_BID_RESULT_RECEIVED, SystemData.Events.AUCTION_INIT_RECEIVED, SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- ActionFraction
- Atlas
- BuddyBind
- CaVES
- Calling
- HealGrid
- ManualScenarioRefresh
- Map
- MiracleGrow
- Phantom
- RVMOD_PlayerStatus
- RVMOD_Targets
- ResHelp
- Shinies

## Examples

- ActionFraction: Shutdown -> WindowUnregisterEventHandler(windowName, SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED, "ActionFraction.UpdateActionPoints")
- ActionFraction: Shutdown -> WindowUnregisterEventHandler(windowName, SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED, "ActionFraction.UpdateActionPoints")
- Atlas: Shutdown -> WindowUnregisterEventHandler("AtlasFrame", SystemData.Events.UPDATE_PROCESSED)
- BuddyBind: OnExitBindingMode -> WindowUnregisterEventHandler(BuddyBind.bindingButtonName, SystemData.Events.L_BUTTON_UP_PROCESSED)
- BuddyBind: OnExitBindingMode -> WindowUnregisterEventHandler(BuddyBind.bindingButtonName, SystemData.Events.M_BUTTON_UP_PROCESSED)
- BuddyBind: OnExitBindingMode -> WindowUnregisterEventHandler(BuddyBind.bindingButtonName, SystemData.Events.R_BUTTON_UP_PROCESSED)

## Related APIs

- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HEALTH_FADE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HEALTH_FADE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed
