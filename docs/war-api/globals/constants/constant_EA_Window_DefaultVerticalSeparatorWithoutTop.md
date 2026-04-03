# EA_Window_DefaultVerticalSeparatorWithoutTop

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
| Addons seen in | EA_OpenPartyWindow, GDes, Ges, Res, TalismanGenie |
| Files seen in | GDesOptions.xml, GesOptions.xml, Res.xml, Source/OpenPartyWindowTabNearby.xml, Source/OpenPartyWindowTabWorld.xml, TalismanGenie.xml |
| Namespaces detected | EA_Window_DefaultVerticalSeparatorWithoutTop |
| Source kinds | xml_attributes |
| Example locations | EA_Window_OpenPartyNearbySocketDefSeparator, EA_Window_OpenPartyWorldSocketDefSeparator, GDesOptionsColourVertDivide, GDesOptionsExtrasVertDivide, GDesOptionsGeneralVertDivide, GesOptionsColourVertDivide |
| XML usage count | 10 |
| XML attribute usage count | 10 |
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

Engine-supplied XML constant or template class referenced by 5 addons.

## Seen In

- EA_OpenPartyWindow
- GDes
- Ges
- Res
- TalismanGenie

## Used By

- EA_Window_OpenPartyNearbySocketDefSeparator
- EA_Window_OpenPartyWorldSocketDefSeparator
- GDesOptionsColourVertDivide
- GDesOptionsExtrasVertDivide
- GDesOptionsGeneralVertDivide
- GesOptionsColourVertDivide
- GesOptionsExtrasVertDivide
- GesOptionsGeneralVertDivide
- ResOptionsVertDivide1
- TalismanGenieVertDivide

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
