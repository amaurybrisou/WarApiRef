# OverlayOffset

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
| Addons seen in | AdvancedPetAssist, Aura, AutoBand, Crusher, DaemonAssist, EA_UiDebugTools, GDes, Ges |
| Files seen in | Configuration/Config.xml, Configuration/HopperConfig.xml, Configuration/WCDBConfig.xml, Configuration/WCDPConfig.xml, Source/DebugWindowVerticalScroll.xml, Source/PureUIElementTemplates.xml, Source/ShiniesUITemplates.xml, Source/SocialWindow.xml |
| Namespaces detected | OverlayOffset |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Aura: Aura_Button_DefaultResizableTinyComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge, AutoBand: AutoBandWindowToolsNoMicCheckBox |
| XML usage count | 35 |
| XML attribute usage count | 35 |
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

OverlayOffset is a XML UI element. It commonly appears under Button and ComboBox.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 31× (HIGH)
- [ComboBox](element_ComboBox.md) — 4× (MEDIUM)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 93, 193, 148, 223, ... |
| `y` | **required** | 100% | 0, 1, 7, 9, ... |
## Seen In

- AdvancedPetAssist
- Aura
- AutoBand
- Crusher
- DaemonAssist
- EA_UiDebugTools
- GDes
- Ges
- Hopper
- LoyalPet
- Motion
- PartyAd
- Pure
- RaidMeter
- Shinies
- SocialWindow 2.0
- SpamBayes
- TalismanGenie
- WSCT
- Warbuilder
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> OverlayOffset in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> OverlayOffset in Button APA_ComboBoxButtonWide
- Aura: Aura_Button_DefaultResizableTinyComboBoxSelected -> OverlayOffset in Button Aura_Button_DefaultResizableTinyComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelected -> OverlayOffset in Button Aura_ComboBox_DefaultResizableComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge -> OverlayOffset in Button Aura_ComboBox_DefaultResizableComboBoxSelectedLarge
- AutoBand: AutoBandWindowToolsNoMicCheckBox -> OverlayOffset in Button AutoBandWindowToolsNoMicCheckBox

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type
