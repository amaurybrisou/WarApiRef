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
| Files seen in | `/workspace/AdvancedPetAssist/APAGui.xml:1389`, `/workspace/AdvancedPetAssist/APAGui.xml:1430`, `/workspace/AdvancedPetAssist/APAGui.xml:1471`, `/workspace/AdvancedPetAssist/APAGui.xml:1509`, `/workspace/AggroMeter/AggroMeter.xml:126`, `/workspace/Aura/Source/AuraConfig.xml:390`, `/workspace/Aura/Source/AuraConfig.xml:400`, `/workspace/Aura/Source/AuraConfig.xml:410` |
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

## Common Handlers

- none

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
