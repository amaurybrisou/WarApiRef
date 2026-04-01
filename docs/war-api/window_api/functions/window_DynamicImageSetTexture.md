# DynamicImageSetTexture

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
| Addons seen in | Ace, AdvancedRenownTrainer, Aura, AutoMark, BankArkel, BuffHead, CM_ClosetGoblin, CombatTextNames |
| Files seen in | `/workspace/Ace/LibGUI.lua:1270`, `/workspace/Aura/Source/AuraHelpers.lua:33`, `/workspace/AutoMark/Source/AutoMark.lua:7`, `/workspace/BankArkel/BankArkel.lua:242`, `/workspace/BankArkel/BankArkel.lua:350`, `/workspace/BuffHead/EffectFrame.lua:359`, `/workspace/BuffHead/EffectFrame.lua:52`, `/workspace/BuffHead/Setup/SetupLayoutProperties.lua:117` |
| Namespaces detected | DynamicImageSetTexture |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:Texture, AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow, AdvancedRenownTrainer: CreateAbilityWindow, Aura: AuraHelpers.SetDynamicImageTexture, AutoMark: AutoMark.local.CreateMarker, AutoMark: CreateMarker |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 127 |
| Global usage count | 127 |
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
DynamicImageSetTexture(windowName, texture, textureX, textureY)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target image control name. | Observed values: "ClosetGoblinCharacterWindowContentsEquipmentSlot"..slot.."IconBase", "EnemyClickCastingDialogContentScrollChildActionConfig1Ability", "EnemyIcon" |
| texture | Observed as a texture resource name. | Observed values: "", "BuffHead_Effect_Border_Frame"..frame, "BuffHead_Effect_Border_Frame1" |
| textureX | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 115, 185 |
| textureY | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 105, 128 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedRenownTrainer
- Aura
- AutoMark
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Effigy
- Enemy
- GCDsaver
- GetStats
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- Moth
- PotionBar
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- Twister
- WSCT
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:Texture -> DynamicImageSetTexture(self.name, texture, x, y)
- AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow -> DynamicImageSetTexture(t.windowName.."Square", GetIconData(t.icon), 0, 0)
- AdvancedRenownTrainer: CreateAbilityWindow -> DynamicImageSetTexture(t.windowName.."Square", GetIconData(t.icon), 0, 0)
- Aura: AuraHelpers.SetDynamicImageTexture -> DynamicImageSetTexture(window, texture, dx, dy)
- AutoMark: AutoMark.local.CreateMarker -> DynamicImageSetTexture(window_name.."_Icon", texture_name, texture_x, texture_y)
- AutoMark: CreateMarker -> DynamicImageSetTexture(window_name.."_Icon", texture_name, texture_x, texture_y)

## Related APIs

- none

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [StatusBarSetCurrentValue](window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetId](window_WindowSetId.md) (HIGH 100/100) - Window Function
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

- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
