# SystemData.Events.L_BUTTON_UP_PROCESSED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 14 addons

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
| Addons seen in | BuddyBind, BuffHead, Calling, Dascore, EA_OpenPartyWindow, FozAuction, Miracle Grow Remix, RVAPI_ColorDialog |
| Files seen in | BuddyBind.lua, CallingKeybinding.lua, Dascore.lua, Modules/UI/Shinies-UI-Browse.lua, RVAPI_ColorDialog.lua, RVMOD_Manager.lua, RVMOD_SquaredDistances.lua, RVMOD_Targets.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | Criteria_ReopenComboBox, InitLayout, Initialize, OnCheckItemSlot, OnCheckItemType, OnComboBoxConditionClick |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
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

SystemData.SystemData.Events.L_BUTTON_UP_PROCESSED field accessed by 14 addons; commonly found in Criteria_ReopenComboBox and InitLayout, Initialize, OnCheckItemSlot, OnCheckItemType, OnComboBoxConditionClick, OnExitBindingMode, OnHidden, OnInitialize, OnLButtonRawDeviceInput, OnRawDeviceInput, OnSelChanged_Criteria_MultiSelCombo, ReopenComboBox, ReopenItemSlots, ReopenItemTypes, SetupHooks, Show, Shutdown, StartBinding, UpdateAdditonalFiltersButton, lua_call contexts.

## Seen In

- BuddyBind
- BuffHead
- Calling
- Dascore
- EA_OpenPartyWindow
- FozAuction
- Miracle Grow Remix
- RVAPI_ColorDialog
- RVMOD_Manager
- RVMOD_SquaredDistances
- RVMOD_Targets
- Shinies
- TastyButtons
- TidyChat

## Related APIs

- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxExternalOpenMenu](../../window_api/functions/window_ComboBoxExternalOpenMenu.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [DefaultColor.LabelSetTextColor](../../globals/functions/global_DefaultColor.LabelSetTextColor.md) (MEDIUM 55/100) - Global Function

## Used With

- [ButtonSetCheckButtonFlag](../../window_api/functions/window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxExternalOpenMenu](../../window_api/functions/window_ComboBoxExternalOpenMenu.md) (HIGH 100/100) - Window Function
- [SystemData.Events.BATTLEGROUP_MEMBER_UPDATED](systemdata_SystemData.Events.BATTLEGROUP_MEMBER_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.BATTLEGROUP_UPDATED](systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_SETTINGS_UPDATED](systemdata_SystemData.Events.GROUP_SETTINGS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function

## Notes

- Observed in contexts: Criteria_ReopenComboBox, InitLayout, Initialize, OnCheckItemSlot, OnCheckItemType, OnComboBoxConditionClick
