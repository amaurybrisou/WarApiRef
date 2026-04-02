# TextEditBoxSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 29 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, Aura, BuffHead, Busted, DAoCBuff, EA_UiDebugTools |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:656`, `/workspace_addons/Ace/LibGUI.lua:669`, `/workspace_addons/Ace/LibGUI.lua:711`, `/workspace_addons/Ace/LibGUI.lua:724`, `/workspace_addons/AdvancedPetAssist/APAGui.lua:983`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:147`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:163`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:220` |
| Namespaces detected | TextEditBoxSetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_MultiTextbox:Clear, Ace: LIBGUI_MultiTextbox:SetText, Ace: LIBGUI_Textbox:Clear, Ace: LIBGUI_Textbox:SetText, AdvancedPetAssist: APAGui.ApplyHUDColorOff, AdvancedPetAssist: APAGui.ApplyHUDColorOn |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 327 |
| Global usage count | 327 |
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
TextEditBoxSetText(windowName, text)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "APAEditHUDOffB", "APAEditHUDOffG", "APAEditHUDOffR" |
| text | Observed as a text or wstring payload. | Observed values: AddonToShow.RVProjectURL, DAoCBuffSettings.TmpFilter[activefilter.index].name, DAoCBuffVar.Frames[activewindow.index].name |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BuffHead
- Busted
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- JunkDump
- Killer
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WhoHealedMe
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_MultiTextbox:Clear -> TextEditBoxSetText(self.name, L "")
- Ace: LIBGUI_MultiTextbox:SetText -> TextEditBoxSetText(self.name, towstring(text))
- Ace: LIBGUI_Textbox:Clear -> TextEditBoxSetText(self.name, L "")
- Ace: LIBGUI_Textbox:SetText -> TextEditBoxSetText(self.name, towstring(text))
- AdvancedPetAssist: APAGui.ApplyHUDColorOff -> TextEditBoxSetText("APAEditHUDOffR", towstring(APA.hudColorOffR))
- AdvancedPetAssist: APAGui.ApplyHUDColorOff -> TextEditBoxSetText("APAEditHUDOffG", towstring(APA.hudColorOffG))

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [ScrollWindowSetOffset](window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [TextEditBoxSetHistory](window_TextEditBoxSetHistory.md) (HIGH 80/100) - Window Function
- [TextEditBoxSetMaxChars](window_TextEditBoxSetMaxChars.md) (HIGH 80/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
