# ListData

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
| Namespaces detected | ListData |
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

Structural XML sub-element of ListBox that binds a list to a named Lua backing table. The `table` attribute names the Lua table supplying row data; the optional `populationfunction` attribute names a Lua callback for custom per-row population.

## Common Attributes

- table
- populationfunction

## Common Inherits

- none

## Common Parent Elements

- [ListBox](element_ListBox.md) — 2× (HIGH)

## Common Structural Child Elements

- [ListColumns](element_ListColumns.md) — 2× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `table` | **required** | 100% | TidyChat.LootRoll.itemRollData, TidyRoll.CustomAutoRoll.autoRollList |
| `populationfunction` | optional | 50% | TidyRoll.CustomAutoRoll.PopulateAutoRollList |
## Structural Sub-Elements

### [ListColumns](element_ListColumns.md)

Observed 2 times as an unnamed child.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChatLootRollList -> ListData in ListBox TidyChatLootRollList
- TidyRoll: TRollAutoRollList -> ListData in ListBox TRollAutoRollList

## Related APIs

- [ListBox](element_ListBox.md) (HIGH 100/100) - XML Element Type
- [ListColumns](element_ListColumns.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
