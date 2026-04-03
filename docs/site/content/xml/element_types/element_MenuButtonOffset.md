# MenuButtonOffset

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
| Addons seen in | AdvancedPetAssist, Aura, Shinies, WSCT |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/Shinies/Source/ShiniesUITemplates.xml:0`, `/workspace/data/raw/wsct/wsct_options/wsct_options.xml:0` |
| Namespaces detected | MenuButtonOffset |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBox, AdvancedPetAssist: APA_ComboBoxWide, Aura: Aura_ComboBox_DefaultResizable, Aura: Aura_ComboBox_DefaultResizableLarge, Aura: Aura_ComboBox_DefaultResizableTiny, Shinies: Shinies_ComboBox_DefaultResizableLarge |
| XML usage count | 7 |
| XML attribute usage count | 7 |
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

MenuButtonOffset is a XML UI element. It commonly appears under ComboBox.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [ComboBox](element_ComboBox.md) — 7× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 5 |
| `y` | **required** | 100% | 5 |
## Seen In

- AdvancedPetAssist
- Aura
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APA_ComboBox -> MenuButtonOffset in ComboBox APA_ComboBox
- AdvancedPetAssist: APA_ComboBoxWide -> MenuButtonOffset in ComboBox APA_ComboBoxWide
- Aura: Aura_ComboBox_DefaultResizable -> MenuButtonOffset in ComboBox Aura_ComboBox_DefaultResizable
- Aura: Aura_ComboBox_DefaultResizableLarge -> MenuButtonOffset in ComboBox Aura_ComboBox_DefaultResizableLarge
- Aura: Aura_ComboBox_DefaultResizableTiny -> MenuButtonOffset in ComboBox Aura_ComboBox_DefaultResizableTiny
- Shinies: Shinies_ComboBox_DefaultResizableLarge -> MenuButtonOffset in ComboBox Shinies_ComboBox_DefaultResizableLarge

## Related APIs

- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
