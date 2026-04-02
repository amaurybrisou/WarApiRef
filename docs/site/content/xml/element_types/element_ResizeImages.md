# ResizeImages

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
| Addons seen in | AdvancedPetAssist, Aura, BuffHead, DaemonAssist, EA_UiDebugTools, LoyalPet, Miracle Grow Remix, Shinies |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:46`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:6`, `/workspace_addons/Aura/Source/Templates.xml:147`, `/workspace_addons/Aura/Source/Templates.xml:176`, `/workspace_addons/Aura/Source/Templates.xml:205`, `/workspace_addons/Aura/Source/Templates.xml:234`, `/workspace_addons/Aura/Source/Templates.xml:262`, `/workspace_addons/Aura/Source/Templates.xml:291` |
| Namespaces detected | ResizeImages |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Aura: Aura_Button_DefaultMenuButton, Aura: Aura_Button_DefaultMenuButtonLarge, Aura: Aura_Button_DefaultMenuButtonTiny, Aura: Aura_Button_DefaultResizableTinyComboBoxSelected |
| XML usage count | 21 |
| XML attribute usage count | 21 |
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

Observed XML element type instantiated by 9 addons.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 21× (HIGH)

## Seen In

- AdvancedPetAssist
- Aura
- BuffHead
- DaemonAssist
- EA_UiDebugTools
- LoyalPet
- Miracle Grow Remix
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> ResizeImages in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> ResizeImages in Button APA_ComboBoxButtonWide
- Aura: Aura_Button_DefaultMenuButton -> ResizeImages in Button Aura_Button_DefaultMenuButton
- Aura: Aura_Button_DefaultMenuButtonLarge -> ResizeImages in Button Aura_Button_DefaultMenuButtonLarge
- Aura: Aura_Button_DefaultMenuButtonTiny -> ResizeImages in Button Aura_Button_DefaultMenuButtonTiny
- Aura: Aura_Button_DefaultResizableTinyComboBoxSelected -> ResizeImages in Button Aura_Button_DefaultResizableTinyComboBoxSelected

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
