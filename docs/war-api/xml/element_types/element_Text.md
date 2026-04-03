# Text

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 98/100

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Killer, WhoHealedMe |
| Files seen in | `/workspace/data/raw/Killer/Killer.xml:0`, `/workspace/data/raw/WhoHealedMe/WHMGui.xml:0` |
| Namespaces detected | Text |
| Source kinds | xml_frames |
| Example locations | Killer: KillerScoreDetailsWindowTitle, Killer: KillerSettingsWindowContentDisplaySection, Killer: KillerSettingsWindowContentFeedFontLabel, Killer: KillerSettingsWindowContentFeedHistoryLimitLabel, Killer: KillerSettingsWindowContentHistoryMaxZonesLabel, Killer: KillerSettingsWindowContentIntro |
| XML usage count | 60 |
| XML attribute usage count | 60 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- alignment
- text

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 60× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `text` | **required** | 88% | Killer Settings, Changes apply immediately. Zone K/D history uses 0 for unlimited saved zone leaderboards., Display, Personal, ... |
| `alignment` | optional | 35% | left |
## Seen In

- Killer
- WhoHealedMe

## Examples

- Killer: KillerScoreDetailsWindowTitle -> Text in Label KillerScoreDetailsWindowTitle
- Killer: KillerSettingsWindowContentDisplaySection -> Text in Label KillerSettingsWindowContentDisplaySection
- Killer: KillerSettingsWindowContentFeedFontLabel -> Text in Label KillerSettingsWindowContentFeedFontLabel
- Killer: KillerSettingsWindowContentFeedHistoryLimitLabel -> Text in Label KillerSettingsWindowContentFeedHistoryLimitLabel
- Killer: KillerSettingsWindowContentHistoryMaxZonesLabel -> Text in Label KillerSettingsWindowContentHistoryMaxZonesLabel
- Killer: KillerSettingsWindowContentIntro -> Text in Label KillerSettingsWindowContentIntro

## Related APIs

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
