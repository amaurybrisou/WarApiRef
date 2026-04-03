# ListColumns

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 55/100

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

- Rationale: unknown

## Evidence Signals

- none

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | none |
| Source kinds | none |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

ListColumns is a structural XML container that declares list column mappings. It commonly appears under ListBox and contains ListColumn entries.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [ListData](element_ListData.md) — 74× (HIGH)

## Common Structural Child Elements

- [ListColumn](element_ListColumn.md) — 192× (HIGH)

## Structural Sub-Elements

### [ListColumn](element_ListColumn.md)

Observed 192 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `format` | **required** | wstring, number, icon, comma |
| `variable` | **required** | Name, RankN, Character, Type |
| `windowname` | **required** | Name, Rank, Character, Type |
| `style` | optional | comma |
## Recursive Hierarchy

- Root: [ListColumns](element_ListColumns.md)
- [ListColumn](element_ListColumn.md) (structural, 192×, HIGH)

## Seen In

- AggroMeter
- Aura
- BlackBook
- BuffHead
- CM_ClosetGoblin
- CastSequence
- Crusher
- DAoCBuff
- Dascore
- Deathblow
- Deathblow2
- DetauntHelper
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiModWindow
- Enemy
- HealGrid
- Hopper
- Kwestor
- LibAddonButton
- LootAlert
- Motion
- NerfedButtons
- Obsidian
- PieTracker
- Pocket Palette
- Pure
- QuickTacticSwitch
- RealmStatus
- SessionRPs
- Shinies
- SocialWindow 2.0
- TastyButtons
- TaxPayer
- TexturedButtons
- TidyChat
- TidyRoll
- Tome Titan
- TomeTracker
- Tortall_DPS
- TurretRange
- WTes
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- nLootLink
- zMailMod

## Examples

- none

## Related APIs

- [ListData](element_ListData.md) (HIGH 100/100) - XML Element Type
- [ListColumn](element_ListColumn.md) (MEDIUM 45/100) - XML Element Type
