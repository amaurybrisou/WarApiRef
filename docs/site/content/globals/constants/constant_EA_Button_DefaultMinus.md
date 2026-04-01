# EA_Button_DefaultMinus

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
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:220`, `/workspace/Autoband/AutoBandWindowConfig.xml:296`, `/workspace/Busted/Busted.xml:50`, `/workspace/MGRemix/layout.xml:444`, `/workspace/MGRemix/layout.xml:770`, `/workspace/ObjectInspector/ObjectInspector.xml:123`, `/workspace/WarBoard/WarBoardOptions.xml:482`, `/workspace/WarBoard/WarBoardOptions.xml:547` |
| Namespaces detected | EA_Button_DefaultMinus |
| Source kinds | xml_attributes |
| Example locations | AutoBandWindowConfigButtonMinusRank, AutoBandWindowConfigButtonMinusTime, BustedGUIPrevError, MiracleGrow2LayoutProgTemplateBarMinus, MiracleGrow2LayoutSettingsTemplateVisMinus, ObjectInspectorDepthMinusButton |
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

- AutoBandWindowConfigButtonMinusRank
- AutoBandWindowConfigButtonMinusTime
- BustedGUIPrevError
- MiracleGrow2LayoutProgTemplateBarMinus
- MiracleGrow2LayoutSettingsTemplateVisMinus
- ObjectInspectorDepthMinusButton
- PlusMinusMinButton
- PlusMinusVertMinButton

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
