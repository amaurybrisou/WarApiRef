# LabelSetTextColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 35 addons

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
| Addons seen in | Ace, AdvancedPetAssist, Aura, BuffHead, Busted, CM_ClosetGoblin, CombatTextNames, DAoCBuff |
| Files seen in | `/workspace/Ace/LibGUI.lua:463`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:11`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:181`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:241`, `/workspace/AdvancedPetAssist/APAGuiHUD.lua:98`, `/workspace/Aura/Source/Aura.lua:505`, `/workspace/Aura/Source/AuraTooltip.lua:32`, `/workspace/BuffHead/EffectFrame.lua:52` |
| Namespaces detected | LabelSetTextColor |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:Color, AdvancedPetAssist: APAGui.UpdateFollowTargetHUD, AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD, AdvancedPetAssist: APAGui.UpdateKitingHUD, AdvancedPetAssist: APAGui.UpdatePetTargetHUD, Aura: Aura:UpdateTimerWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 203 |
| Global usage count | 203 |
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
LabelSetTextColor(arg1, arg2, arg3, arg4)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "APAFollowTargetHUDLabel", "APAInstantOnlyHUDLabel", "APAKitingHUDLabel" |
| arg2 | Observed as a numeric value. | Observed values: 0, 100, 128 |
| arg3 | Observed as a numeric value. | Observed values: 0, 100, 110 |
| arg4 | Observed as a numeric value. | Observed values: 0, 10, 100 |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- Aura
- BuffHead
- Busted
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- DeepSleep
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- Killer
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- PeaceOut
- PotionBar
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TurretRange
- WSCT
- WarTriage
- WhoHealedMe
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_Label:Color -> LabelSetTextColor(self.name, red, green, blue)
- AdvancedPetAssist: APAGui.UpdateFollowTargetHUD -> LabelSetTextColor("APAFollowTargetHUDLabel", 255, 255, 255)
- AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD -> LabelSetTextColor("APAInstantOnlyHUDLabel", 192, 192, 192)
- AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD -> LabelSetTextColor("APAInstantOnlyHUDLabel", 255, 255, 255)
- AdvancedPetAssist: APAGui.UpdateKitingHUD -> LabelSetTextColor("APAKitingHUDLabel", 192, 192, 192)
- AdvancedPetAssist: APAGui.UpdateKitingHUD -> LabelSetTextColor("APAKitingHUDLabel", 255, 255, 255)

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
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetTextColor](window_LabelGetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [StatusBarSetCurrentValue](window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
