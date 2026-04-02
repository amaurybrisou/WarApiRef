# HorizontalResizeImage

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AggroMeter, Aura, BuffHead, EA_ThreePartBar, EA_UiDebugTools, MapMonster, PotionBar |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:1389`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1430`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1471`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1509`, `/workspace_addons/AggroMeter/AggroMeter.xml:126`, `/workspace_addons/Aura/Source/AuraConfig.xml:390`, `/workspace_addons/Aura/Source/AuraConfig.xml:400`, `/workspace_addons/Aura/Source/AuraConfig.xml:410` |
| Namespaces detected | HorizontalResizeImage |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, AggroMeter: AggroMeterGrayWindowSeparatorMiddle, Aura: AuraConfigTabsSeparatorActivation |
| XML usage count | 31 |
| XML attribute usage count | 31 |
| Lua usage count | 0 |
| Global usage count | 0 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed XML element type instantiated by 12 addons.

## Common Attributes

- name
- inherits
- texture
- handleinput
- layer
- savesettings
- alpha

## Common Inherits

- EA_HorizontalResizeImage_TabSeparatorMiddle
- BuffHeadLayoutHorizontalResizeImage
- EA_BrownHorizontalRule
- EA_HorizontalResizeImage_DefaultTopFrame
- RVAPI_ColorDialogHorizontalLineTintable
- RewardPoolFilledBarLoser
- RewardPoolFilledBarVictor
- VictoryPointsFilledBarDestruction
- VictoryPointsFilledBarOrder

## Common Parent Elements

- [Window](element_Window.md)


## Structural Sub-Elements

### [Middle](element_Middle.md)

- Observed in 15 parent frames
- Attributes: `id`, `x`, `y`
  - `id`: `Loser-Bar-horiz`, `Victor-Bar-horiz`, `Horiz-Background`, `Dest-VP-Bar-horiz`, `Order-VP-Bar-horiz`
  - `x`: `0`, `424`, `7`, `346`
  - `y`: `0`, `762`, `655`

### [Left](element_Left.md)

- Observed in 10 parent frames
- Attributes: `id`, `x`, `y`
  - `id`: `Loser-Bar-horiz-End-Cap`, `Dest-VP-Bar-horiz-End-Cap`
  - `x`: `0`, `346`
  - `y`: `0`, `655`

### [Right](element_Right.md)

- Observed in 10 parent frames
- Attributes: `id`, `x`, `y`
  - `id`: `Victor-Bar-horiz-End-Cap`, `Order-VP-Bar-horiz-End-Cap`
  - `x`: `0`, `59`, `396`
  - `y`: `30`, `24`, `0`, `50`, `655`

### [TintColor](element_TintColor.md)

- Observed in 10 parent frames
- Attributes: `b`, `g`, `r`
  - `b`: `0`, `200`, `100`, `255`
  - `g`: `0`, `200`, `255`, `100`
  - `r`: `255`, `0`, `180`, `200`, `100`

### [TexSlices](element_TexSlices.md)

- Observed in 5 parent frames

## Typical XML Structure

```xml
<HorizontalResizeImage name="..." texture="EA_VictoryPoints01_32b" handleinput="false">
  <TexSlices>
    <Left id="Loser-Bar-horiz-End-Cap"/>
    <Middle id="Loser-Bar-horiz"/>
  </TexSlices>
</HorizontalResizeImage>
```


## Attribute Reference

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `APAFollowTargetHUDFill`, `APAInstantOnlyHUDFill`, `APAKitingHUDFill`, `APAPetTargetHUDBg`, … | 31 |
| `inherits` | frame-ref | `EA_HorizontalResizeImage_TabSeparatorMiddle`, `BuffHeadLayoutHorizontalResizeImage`, `VictoryPointsFilledBarDestruction`, `RewardPoolFilledBarLoser`, … | 16 |
| `texture` | string | `EA_TintableImage`, `EA_VictoryPoints01_32b`, `shared_01`, `EA_HUD_01` | 15 |
| `handleinput` | boolean | `false` | 12 |
| `layer` | string | `background` | 5 |
| `savesettings` | boolean | `false` | 5 |
| `alpha` | number | `1` | 1 |

## Seen In

- AdvancedPetAssist
- AggroMeter
- Aura
- BuffHead
- EA_ThreePartBar
- EA_UiDebugTools
- MapMonster
- PotionBar
- RVAPI_ColorDialog
- RoR_SoR
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> HorizontalResizeImage APAPetTargetHUDBg
- AggroMeter: AggroMeterGrayWindowSeparatorMiddle -> HorizontalResizeImage AggroMeterGrayWindowSeparatorMiddle
- Aura: AuraConfigTabsSeparatorActivation -> HorizontalResizeImage AuraConfigTabsSeparatorActivation

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
