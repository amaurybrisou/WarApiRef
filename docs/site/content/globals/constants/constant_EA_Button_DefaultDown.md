# EA_Button_DefaultDown

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
| Addons seen in | CM_ClosetGoblin, GDes, Ges, KeyBar |
| Files seen in | ClosetGoblin.xml, GDesOptions.xml, GesOptions.xml, KeyBar.xml |
| Namespaces detected | EA_Button_DefaultDown |
| Source kinds | xml_attributes |
| Example locations | ClosetGoblinActionBarPageSelectorDown, GDesOptionsExtrasFeatureOffseDownButton, GDesOptionsExtrasIconsOffseDownButton, GesOptionsExtrasFeatureOffseDownButton, GesOptionsExtrasTrackerOffseDownButton, KeyBarPageSelectorDown |
| XML usage count | 6 |
| XML attribute usage count | 6 |
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

Engine-supplied XML constant or template class referenced by 4 addons.

## Seen In

- CM_ClosetGoblin
- GDes
- Ges
- KeyBar

## Used By

- ClosetGoblinActionBarPageSelectorDown
- GDesOptionsExtrasFeatureOffseDownButton
- GDesOptionsExtrasIconsOffseDownButton
- GesOptionsExtrasFeatureOffseDownButton
- GesOptionsExtrasTrackerOffseDownButton
- KeyBarPageSelectorDown

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- none
