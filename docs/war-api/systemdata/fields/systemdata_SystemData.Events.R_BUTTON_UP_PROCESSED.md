# SystemData.Events.R_BUTTON_UP_PROCESSED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 176

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuddyBind, Calling, Dascore |
| Files seen in | BuddyBind.lua, CallingEvents.lua, CallingKeybinding.lua, CallingSetup.lua, Dascore.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | CallingSetup.OnRButtonUp, Initialize, OnExitBindingMode, OnInitialize, OnLButtonRawDeviceInput, OnRawDeviceInput |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 3 |
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

SystemData.SystemData.Events.R_BUTTON_UP_PROCESSED field accessed by 3 addons; commonly found in CallingSetup.OnRButtonUp and Initialize, OnExitBindingMode, OnInitialize, OnLButtonRawDeviceInput, OnRawDeviceInput, StartBinding, SystemData.Events.R_BUTTON_UP_PROCESSED, event_page, event_registration, lua_call contexts.

## Seen In

- BuddyBind
- Calling
- Dascore

## Related APIs

- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [ButtonSetCheckButtonFlag](../../window_api/functions/window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [SystemData.Events.L_BUTTON_UP_PROCESSED](systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Observed in contexts: CallingSetup.OnRButtonUp, Initialize, OnExitBindingMode, OnInitialize, OnLButtonRawDeviceInput, OnRawDeviceInput
