# EA_Label_DefaultSubHeading

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
| Addons seen in | CMap, EA_UiModWindow, PotionBar |
| Files seen in | `/workspace_addons/PotionBar/settings/Settings.xml:526`, `/workspace_addons/cmap/CMap.xml:118`, `/workspace_addons/cmap/CMap.xml:127`, `/workspace_addons/ea_uimodwindow/Source/VersionMismatchWindow.xml:92` |
| Namespaces detected | EA_Label_DefaultSubHeading |
| Source kinds | xml_attributes |
| Example locations | $parentRightPaneTitle, CMapWindowAreaNameText, CMapWindowDigitinf, UiModVersionMismatchWindowCurrentVersionText |
| XML usage count | 4 |
| XML attribute usage count | 4 |
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

Observed engine XML template or inherited constant referenced by 3 addons.

## Seen In

- CMap
- EA_UiModWindow
- PotionBar

## Used By

- $parentRightPaneTitle
- CMapWindowAreaNameText
- CMapWindowDigitinf
- UiModVersionMismatchWindowCurrentVersionText

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
