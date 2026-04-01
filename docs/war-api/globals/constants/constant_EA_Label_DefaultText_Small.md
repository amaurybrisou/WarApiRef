# EA_Label_DefaultText_Small

- Category: Constant
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
| Addons seen in | AutoBand, LoyalPet, MegaphonePlusPlus, Pocket Palette, WSCT |
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:367`, `/workspace/LoyalPet/gui/lpet_gui.xml:102`, `/workspace/PocketPalette/PocketPalette.xml:100`, `/workspace/PocketPalette/PocketPalette.xml:160`, `/workspace/PocketPalette/PocketPalette.xml:171`, `/workspace/PocketPalette/PocketPalette.xml:301`, `/workspace/PocketPalette/PocketPalette.xml:477`, `/workspace/PocketPalette/PocketPalette.xml:489` |
| Namespaces detected | EA_Label_DefaultText_Small |
| Source kinds | xml_attributes |
| Example locations | AutoBandWindowConfigVersionLabel, DyeListBoxRowTemplateCount, DyeListBoxRowTemplateName, DyeWindowSelectedDyeCount, DyeWindowSelectedDyeName, ItemWindowGuide |
| XML usage count | 19 |
| XML attribute usage count | 19 |
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

Observed engine XML template or inherited constant referenced by 5 addons.

## Seen In

- AutoBand
- LoyalPet
- MegaphonePlusPlus
- Pocket Palette
- WSCT

## Used By

- AutoBandWindowConfigVersionLabel
- DyeListBoxRowTemplateCount
- DyeListBoxRowTemplateName
- DyeWindowSelectedDyeCount
- DyeWindowSelectedDyeName
- ItemWindowGuide
- LPETComboBoxTemplateLabel
- MegaphoneMainFontTitleLabel
- MegaphoneMainHighlightLeaderLabel
- MegaphoneMainSFXTitleLabel
- MegaphoneMainShowLeaderLabel
- PPMainIntroText
- PPMainSaveSettingsLabel
- WSCTComboBoxTemplateLabel
- WSCTOptionsColorPickerWindowCustomColorBlueText
- WSCTOptionsColorPickerWindowCustomColorGreenText
- WSCTOptionsColorPickerWindowCustomColorRedText
- WSCTSliderTemplateLabel
- WSCTSliderTemplateValue

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
