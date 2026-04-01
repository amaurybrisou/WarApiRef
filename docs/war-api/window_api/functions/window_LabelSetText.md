# LabelSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 57 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BankArkel, BuffHead, Busted |
| Files seen in | `/workspace/Ace/LibGUI.lua:426`, `/workspace/Ace/LibGUI.lua:439`, `/workspace/AdvancedPetAssist/APAGui.lua:983`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:11`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:181`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:241`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:98`, `/workspace/AggroMeter/AggroMeter.lua:5` |
| Namespaces detected | LabelSetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:Clear, Ace: LIBGUI_Label:SetText, AdvancedPetAssist: APAGui.OnShown, AdvancedPetAssist: APAGui.UpdateFollowTargetHUD, AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD, AdvancedPetAssist: APAGui.UpdateKitingHUD |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1436 |
| Global usage count | 1436 |
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
LabelSetText(windowName, text)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "APAFollowTargetHUDLabel", "APAInstantOnlyHUDLabel", "APAKitingHUDLabel" |
| text | Observed as a text or wstring payload. | Observed values: APA.PetTarget.cachedName, Active and L "zzzZZZzzZZzz" or Check and L "demanding" or L "denied", AddonToShow.RVCredits |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- DeepSleep
- EA_UiDebugTools
- Effigy
- Enemy
- FastInteract
- GCDsaver
- GetStats
- GoldTracker
- JunkDump
- Killer
- LibGroup
- LibWBToggler
- MapMonster
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- PeaceOut
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- TimeInQueue
- TurretRange
- Twister
- WSCT
- WarTriage
- WhoHealedMe
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_Label:Clear -> LabelSetText(self.name, L "")
- Ace: LIBGUI_Label:SetText -> LabelSetText(self.name, towstring(text))
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APAOptionsTitleText", L "AdvancedPetAssist")
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APALabelSectionAutoRecall", L "--- Auto Recall ---")
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APALabelSectionMouseControls", L "--- Mouse Controls ---")
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APALabelSectionLos", L "--- LOS Detection ---")

## Related APIs

- none

## Used With

- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [FullResizeImage](../../xml/element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [ScrollWindowSetOffset](window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [StatusBarSetCurrentValue](window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [SystemData.MousePosition.x](../../systemdata/fields/systemdata_SystemData.MousePosition.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.MousePosition.y](../../systemdata/fields/systemdata_SystemData.MousePosition.y.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [TextEditBoxSetHistory](window_TextEditBoxSetHistory.md) (HIGH 80/100) - Window Function
- [TextEditBoxSetMaxChars](window_TextEditBoxSetMaxChars.md) (HIGH 80/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
