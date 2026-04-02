# LabelGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | Ace, DAoCBuff, EA_UiDebugTools, Effigy, GCDsaver, LibWBToggler, RoR_SoR, Shinies |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:434`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:215`, `/workspace_addons/Effigy/LibGUI.lua:434`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:434`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:434`, `/workspace_addons/RoR_SoR/RoR_SoR.lua:1055`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:434`, `/workspace_addons/WoH-Reticle/libs/LibGUI.lua:434` |
| Namespaces detected | LabelGetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:GetText, DAoCBuff: DAoCBuffSettings.FilterRowOnLButtonUp, EA_UiDebugTools: DevPad.TestPrint, EA_UiDebugTools: DevPadWindow.ConfirmLoadFile, EA_UiDebugTools: DevPadWindow.CreateContextMenu, EA_UiDebugTools: DevPadWindow.DeleteFile |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 25 |
| Global usage count | 25 |
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
LabelGetText(arg1)
```

## Description

Observed as a window function across 10 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "SoR_"..Window_Name.."KEEP1HEALTH", "SoR_"..Window_Name.."KEEP1SAFETIMER", "SoR_"..Window_Name.."KEEP2HEALTH" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- DAoCBuff
- EA_UiDebugTools
- Effigy
- GCDsaver
- LibWBToggler
- RoR_SoR
- Shinies
- Targets
- WoH-Reticle

## Examples

- Ace: LIBGUI_Label:GetText -> LabelGetText(self.name)
- DAoCBuff: DAoCBuffSettings.FilterRowOnLButtonUp -> LabelGetText(SystemData.ActiveWindow.name.."Name")
- EA_UiDebugTools: DevPad.TestPrint -> LabelGetText(windowName.."ProjectName")
- EA_UiDebugTools: DevPadWindow.ConfirmLoadFile -> LabelGetText(windowName.."ProjectName")
- EA_UiDebugTools: DevPadWindow.CreateContextMenu -> LabelGetText(windowName.."ProjectName")
- EA_UiDebugTools: DevPadWindow.DeleteFile -> LabelGetText(windowName.."ProjectName")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetText](window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.gsub](../../globals/functions/global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
