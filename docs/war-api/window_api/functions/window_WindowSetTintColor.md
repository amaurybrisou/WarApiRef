# WindowSetTintColor

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Effigy |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:249`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:11`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:181`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:241`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:98`, `/workspace_addons/Aura/Source/Aura.lua:417`, `/workspace_addons/Aura/Source/AuraHelpers.lua:33`, `/workspace_addons/Aura/Source/AuraHelpers.lua:55` |
| Namespaces detected | WindowSetTintColor |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Tint, AdvancedPetAssist: APAGui.UpdateFollowTargetHUD, AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD, AdvancedPetAssist: APAGui.UpdateKitingHUD, AdvancedPetAssist: APAGui.UpdatePetTargetHUD, AdvancedRenownTrainer: AdvancedRenownTrainer.local.SelectAdvantage |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 272 |
| Global usage count | 272 |
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
WindowSetTintColor(windowName, arg2, arg3, arg4)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAFollowTargetHUDFill", "APAInstantOnlyHUDFill", "APAKitingHUDFill" |
| arg2 | Observed as a function or method reference. | Observed values: 0, 100, 12 |
| arg3 | Observed as a numeric value. | Observed values: 0, 100, 110 |
| arg4 | Observed as a numeric value. | Observed values: 0, 10, 100 |

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
- CM_ClosetGoblin
- DAoCBuff
- Effigy
- Enemy
- GCDsaver
- GuardLine
- Killer
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MoraleCircle
- Moth
- Pocket Palette
- PotionBar
- QuickWarReport
- RVAPI_ColorDialog
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- Ace: LIBGUI_ELEMENT:Tint -> WindowSetTintColor(self.name, r, g, b)
- AdvancedPetAssist: APAGui.UpdateFollowTargetHUD -> WindowSetTintColor("APAFollowTargetHUDFill", APA.hudColorOnR, APA.hudColorOnG, APA.hudColorOnB)
- AdvancedPetAssist: APAGui.UpdateFollowTargetHUD -> WindowSetTintColor("APAFollowTargetHUDFill", APA.hudColorOffR, APA.hudColorOffG, APA.hudColorOffB)
- AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD -> WindowSetTintColor("APAInstantOnlyHUDFill", 128, 128, 128)
- AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD -> WindowSetTintColor("APAInstantOnlyHUDFill", 0, 200, 0)
- AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD -> WindowSetTintColor("APAInstantOnlyHUDFill", 200, 100, 0)

## Related APIs

- none

## Used With

- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetMovable](window_WindowSetMovable.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStartScaleAnimation](window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
