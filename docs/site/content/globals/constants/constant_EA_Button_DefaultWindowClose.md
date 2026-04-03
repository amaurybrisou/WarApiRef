# EA_Button_DefaultWindowClose

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
| Addons seen in | Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/Soloq/ui/Overview.xml:65`, `/workspace/data/raw/TidyChat/TidyChat.xml:63`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:32`, `/workspace/data/raw/TidyChat/TidyChatLootRoll.xml:58`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:138`, `/workspace/data/raw/minesweep/minesweep.xml:89` |
| Namespaces detected | EA_Button_DefaultWindowClose |
| Source kinds | xml_attributes |
| Example locations | MineSweepWindowClose, SoloqOverviewWindowCloseButton, TRollAutoRollCloseButton, TidyChatCopyClose, TidyChatLootRollClose, TidyChatOptionsCloseButton |
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

Observed engine XML template or inherited constant referenced by 4 addons.

## Seen In

- Soloq
- TidyChat
- TidyRoll
- minesweep

## Used By

- MineSweepWindowClose
- SoloqOverviewWindowCloseButton
- TRollAutoRollCloseButton
- TidyChatCopyClose
- TidyChatLootRollClose
- TidyChatOptionsCloseButton

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
