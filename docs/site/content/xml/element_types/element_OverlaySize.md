# OverlaySize

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
| Addons seen in | AdvancedPetAssist, Aura, DaemonAssist, EA_UiDebugTools, LoyalPet, Shinies, WSCT |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:46`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:6`, `/workspace_addons/Aura/Source/Templates.xml:147`, `/workspace_addons/Aura/Source/Templates.xml:205`, `/workspace_addons/Aura/Source/Templates.xml:262`, `/workspace_addons/DaemonAssist/DaemonAssist.xml:6`, `/workspace_addons/LoyalPet/gui/lpet_gui.xml:45`, `/workspace_addons/LoyalPet/gui/lpet_gui.xml:5` |
| Namespaces detected | OverlaySize |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Aura: Aura_Button_DefaultResizableTinyComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge, DaemonAssist: DaemonAssist_ComboBoxButtonWide |
| XML usage count | 14 |
| XML attribute usage count | 14 |
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

Observed XML element type instantiated by 7 addons.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% |  |
| `y` | **required** | 100% |  |
## Seen In

- AdvancedPetAssist
- Aura
- DaemonAssist
- EA_UiDebugTools
- LoyalPet
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> OverlaySize in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> OverlaySize in Button APA_ComboBoxButtonWide
- Aura: Aura_Button_DefaultResizableTinyComboBoxSelected -> OverlaySize in Button Aura_Button_DefaultResizableTinyComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelected -> OverlaySize in Button Aura_ComboBox_DefaultResizableComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge -> OverlaySize in Button Aura_ComboBox_DefaultResizableComboBoxSelectedLarge
- DaemonAssist: DaemonAssist_ComboBoxButtonWide -> OverlaySize in Button DaemonAssist_ComboBoxButtonWide

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
