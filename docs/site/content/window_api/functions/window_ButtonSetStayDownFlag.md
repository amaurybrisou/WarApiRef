# ButtonSetStayDownFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 12 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Bloody Mess, BuddyBind, Calling, DetauntHelper, HealGrid, KillTracker, Kwestor, RvRStats |
| Files seen in | Bloody Mess.lua, BuddyBind.lua, CallingKeybinding.lua, CallingSetup.lua, Gui/HealGridGuiUtility.lua, Gui/KwestorGui.lua, Source/Config/Config.lua, Source/KillTracker.lua |
| Namespaces detected | ButtonSetStayDownFlag |
| Source kinds | lua_calls |
| Example locations | Bloody Mess: OnInitialize, BuddyBind: OnExitBindingMode, BuddyBind: OnLButtonRawDeviceInput, BuddyBind: OnRawDeviceInput, Calling: EnableSubMenu, Calling: OnRawDeviceInput |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 29 |
| Global usage count | 29 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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
ButtonSetStayDownFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BloodyMessOptionsEnableCheckBoxButton", "CallingSetupCategories"..subName, "CharacterWindowTabRvRStats" |
| arg2 | Observed as a boolean toggle. | Observed values: false, flag, setit |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Bloody Mess
- BuddyBind
- Calling
- DetauntHelper
- HealGrid
- KillTracker
- Kwestor
- RvRStats
- RvRStatsTab
- Shinies
- TastyButtons
- nLootLink

## Examples

- Bloody Mess: OnInitialize -> ButtonSetStayDownFlag("BloodyMessOptionsEnableCheckBoxButton", true)
- BuddyBind: OnExitBindingMode -> ButtonSetStayDownFlag(BuddyBind.bindingButtonName, false)
- BuddyBind: OnLButtonRawDeviceInput -> ButtonSetStayDownFlag(BuddyBind.bindingButtonName, true)
- BuddyBind: OnRawDeviceInput -> ButtonSetStayDownFlag(BuddyBind.bindingButtonName, false)
- Calling: EnableSubMenu -> ButtonSetStayDownFlag("CallingSetupCategories"..subName, val)
- Calling: OnRawDeviceInput -> ButtonSetStayDownFlag(KeyMappingWindow.bindingButtonName, false)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [ButtonSetCheckButtonFlag](window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.USER_SETTINGS_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.USER_SETTINGS_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field

## Notes

- none
