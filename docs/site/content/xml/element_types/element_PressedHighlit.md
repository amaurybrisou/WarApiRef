# PressedHighlit

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AggroMeter, AnywhereTrainer, Aura, BuffHead, CMap, DaemonAssist, EA_UiDebugTools |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:46`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:6`, `/workspace_addons/AggroMeter/AggroMeter.xml:39`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.xml:28`, `/workspace_addons/Aura/Source/Templates.xml:147`, `/workspace_addons/Aura/Source/Templates.xml:176`, `/workspace_addons/Aura/Source/Templates.xml:205`, `/workspace_addons/Aura/Source/Templates.xml:234` |
| Namespaces detected | PressedHighlit |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, AggroMeter: Aggro_Button_Template, AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage, Aura: Aura_Button_DefaultMenuButton, Aura: Aura_Button_DefaultMenuButtonLarge |
| XML usage count | 53 |
| XML attribute usage count | 53 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

Observed XML element type instantiated by 17 addons.

## Common Attributes

- id
- b
- g
- r
- a
- def
- x
- y

## Common Handlers

- none

## Common Inherits

- none

## Seen In

- AdvancedPetAssist
- AggroMeter
- AnywhereTrainer
- Aura
- BuffHead
- CMap
- DaemonAssist
- EA_UiDebugTools
- GuardList
- GuardRange
- LoyalPet
- Miracle Grow Remix
- PotionBar
- RVMOD_Manager
- RoR_SoR
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> PressedHighlit in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> PressedHighlit in Button APA_ComboBoxButtonWide
- AggroMeter: Aggro_Button_Template -> PressedHighlit in Button Aggro_Button_Template
- AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage -> PressedHighlit in Button AnywhereTrainerTabTemplateInactiveImage
- Aura: Aura_Button_DefaultMenuButton -> PressedHighlit in Button Aura_Button_DefaultMenuButton
- Aura: Aura_Button_DefaultMenuButtonLarge -> PressedHighlit in Button Aura_Button_DefaultMenuButtonLarge

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
