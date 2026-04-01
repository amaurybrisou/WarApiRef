# EA_Button_DefaultPlus

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
| Addons seen in | AutoBand, Busted, EA_UiDebugTools, Miracle Grow Remix, ObjectInspector, WarBoard |
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:208`, `/workspace/Autoband/AutoBandWindowConfig.xml:284`, `/workspace/Busted/Busted.xml:73`, `/workspace/MGRemix/layout.xml:434`, `/workspace/MGRemix/layout.xml:760`, `/workspace/ObjectInspector/ObjectInspector.xml:113`, `/workspace/WarBoard/WarBoardOptions.xml:472`, `/workspace/WarBoard/WarBoardOptions.xml:537` |
| Namespaces detected | EA_Button_DefaultPlus |
| Source kinds | xml_attributes |
| Example locations | AutoBandWindowConfigButtonPlusRank, AutoBandWindowConfigButtonPlusTime, BustedGUINextError, MiracleGrow2LayoutProgTemplateBarPlus, MiracleGrow2LayoutSettingsTemplateVisPlus, ObjectInspectorDepthPlusButton |
| XML usage count | 8 |
| XML attribute usage count | 8 |
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

Observed engine XML template or inherited constant referenced by 6 addons.

## Seen In

- AutoBand
- Busted
- EA_UiDebugTools
- Miracle Grow Remix
- ObjectInspector
- WarBoard

## Used By

- AutoBandWindowConfigButtonPlusRank
- AutoBandWindowConfigButtonPlusTime
- BustedGUINextError
- MiracleGrow2LayoutProgTemplateBarPlus
- MiracleGrow2LayoutSettingsTemplateVisPlus
- ObjectInspectorDepthPlusButton
- PlusMinusMaxButton
- PlusMinusVertMaxButton

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
