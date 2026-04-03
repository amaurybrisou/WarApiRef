# CreateWindowFromTemplate

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, PartyCast, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1043`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1146`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1234`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:344`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:399`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:506`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:595` |
| Namespaces detected | CreateWindowFromTemplate |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_Button:New, InfoScroller: LIBGUI_Checkbox:New, InfoScroller: LIBGUI_Closebutton:New, InfoScroller: LIBGUI_Combobox:New, InfoScroller: LIBGUI_Image:New |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 81 |
| Global usage count | 81 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
CreateWindowFromTemplate(windowName, templateName, parentWindow)
```

## Description

Observed instantiating repeated UI elements from an XML template.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a runtime window instance name. | Observed values: WindowName, c_CHANNEL_MENU.."AdviceButton", c_CHANNEL_MENU.."AllianceButton" |
| templateName | Observed as an XML template name. | Observed values: "ChannelMenuButton", "EA_Button_DefaultWindowClose", "EA_ComboBox_DefaultResizable" |
| parentWindow | Observed as a parent window name. | Observed values: "ChatChannelSelectionWindow", "InfoScrollerMainWindow", c_TIDY_ROLL_OPTIONS |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Creates or instantiates UI objects from XML-backed definitions.

## Seen In

- InfoScroller
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: InfoScroller.CreateCard -> CreateWindowFromTemplate(WindowName, "InfoScrollerTemplate", "InfoScrollerMainWindow")
- InfoScroller: LIBGUI_Button:New -> CreateWindowFromTemplate(w.name, base, w.parent)
- InfoScroller: LIBGUI_Checkbox:New -> CreateWindowFromTemplate(w.name, base, w.parent)
- InfoScroller: LIBGUI_Closebutton:New -> CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- InfoScroller: LIBGUI_Combobox:New -> CreateWindowFromTemplate(w.name, base, w.parent)
- InfoScroller: LIBGUI_Image:New -> CreateWindowFromTemplate(w.name, base, w.parent)

## Related APIs

- none

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](../../window_api/functions/window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 98/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
