# LabelSetTextColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 32 addons

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
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:463`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:11`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:181`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:241`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:98`, `/workspace_addons/Aura/Source/Aura.lua:505`, `/workspace_addons/Aura/Source/AuraTooltip.lua:32`, `/workspace_addons/BuffHead/EffectFrame.lua:52` |
| Namespaces detected | LabelSetTextColor |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:Color, AdvancedPetAssist: APAGui.UpdateFollowTargetHUD, AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD, AdvancedPetAssist: APAGui.UpdateKitingHUD, AdvancedPetAssist: APAGui.UpdatePetTargetHUD, Aura: Aura:UpdateTimerWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 199 |
| Global usage count | 199 |
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

- Not confidently inferable from addon-api docs alone.

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

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetTextColor](window_LabelGetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [StatusBarSetCurrentValue](window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
