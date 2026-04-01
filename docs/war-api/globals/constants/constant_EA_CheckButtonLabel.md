# EA_CheckButtonLabel

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoBand, EA_UiDebugTools |
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:192`, `/workspace/Autoband/AutoBandWindowConfig.xml:271`, `/workspace/Autoband/AutoBandWindowConfig.xml:49`, `/workspace/Autoband/AutoBandWindowTemplate.xml:58`, `/workspace/Autoband/AutoBandWindowTools.xml:110`, `/workspace/Autoband/AutoBandWindowTools.xml:62`, `/workspace/Autoband/AutoBandWindowTools.xml:86`, `/workspace/ea_uidebugtools/Source/DebugWindow.xml:1517` |
| Namespaces detected | EA_CheckButtonLabel |
| Source kinds | xml_attributes |
| Example locations | AutoBandWindowConfigAutoKickTimeoutLabel, AutoBandWindowConfigDefaultTemplateLabel, AutoBandWindowConfigMinRankValueLabel, AutoBandWindowTemplateDefaultLabel, AutoBandWindowToolsKickOfflineLabel, AutoBandWindowToolsKickRankLabel |
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

Observed engine XML template or inherited constant referenced by 2 addons.

## Seen In

- AutoBand
- EA_UiDebugTools

## Used By

- AutoBandWindowConfigAutoKickTimeoutLabel
- AutoBandWindowConfigDefaultTemplateLabel
- AutoBandWindowConfigMinRankValueLabel
- AutoBandWindowTemplateDefaultLabel
- AutoBandWindowToolsKickOfflineLabel
- AutoBandWindowToolsKickRankLabel
- AutoBandWindowToolsKickZoneLabel
- EA_LabelCheckButtonSmallCopyLabel

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
