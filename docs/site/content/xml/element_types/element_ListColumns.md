# ListColumns

- Category: XML Element Type
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
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChatLootRoll.xml:108`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:228` |
| Namespaces detected | ListColumns |
| Source kinds | xml_frames |
| Example locations | TidyChat: TidyChatLootRollList, TidyRoll: TRollAutoRollList |
| XML usage count | 2 |
| XML attribute usage count | 2 |
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

Structural XML container inside ListBox that groups one or more ListColumn declarations, mapping backing-table fields to row-template child windows.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [ListData](element_ListData.md) — 2× (HIGH)

## Common Structural Child Elements

- [ListColumn](element_ListColumn.md) — 5× (HIGH)

## Structural Sub-Elements

### [ListColumn](element_ListColumn.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `format` | **required** | wstring, icon, number |
| `variable` | **required** | 1, 2, iconNum, name |
| `windowname` | **required** | PlayerName, RollNumber, Icon, Name |
## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChatLootRollList -> ListColumns in ListBox TidyChatLootRollList
- TidyRoll: TRollAutoRollList -> ListColumns in ListBox TRollAutoRollList

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
