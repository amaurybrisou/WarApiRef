# Text

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DaemonAssist, Killer, QuickWarReport, WhoHealedMe |
| Files seen in | `/workspace_addons/DaemonAssist/DaemonAssist.xml:100`, `/workspace_addons/DaemonAssist/DaemonAssist.xml:113`, `/workspace_addons/DaemonAssist/DaemonAssist.xml:135`, `/workspace_addons/DaemonAssist/DaemonAssist.xml:67`, `/workspace_addons/DaemonAssist/DaemonAssist.xml:78`, `/workspace_addons/DaemonAssist/DaemonAssist.xml:89`, `/workspace_addons/Killer/Killer.xml:105`, `/workspace_addons/Killer/Killer.xml:116` |
| Namespaces detected | Text |
| Source kinds | xml_frames |
| Example locations | DaemonAssist: DaemonAssistWindowAttackLabel, DaemonAssist: DaemonAssistWindowHeader, DaemonAssist: DaemonAssistWindowHeelLabel, DaemonAssist: DaemonAssistWindowStatusLabel, DaemonAssist: DaemonAssistWindowStatusValue, DaemonAssist: DaemonAssistWindowToggleButton |
| XML usage count | 72 |
| XML attribute usage count | 72 |
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

Observed XML element type instantiated by 4 addons.

## Common Attributes

- text
- alignment

## Common Handlers

- none

## Common Inherits

- none

## Seen In

- DaemonAssist
- Killer
- QuickWarReport
- WhoHealedMe

## Examples

- DaemonAssist: DaemonAssistWindowAttackLabel -> Text in Label DaemonAssistWindowAttackLabel
- DaemonAssist: DaemonAssistWindowHeader -> Text in Label DaemonAssistWindowHeader
- DaemonAssist: DaemonAssistWindowHeelLabel -> Text in Label DaemonAssistWindowHeelLabel
- DaemonAssist: DaemonAssistWindowStatusLabel -> Text in Label DaemonAssistWindowStatusLabel
- DaemonAssist: DaemonAssistWindowStatusValue -> Text in Label DaemonAssistWindowStatusValue
- DaemonAssist: DaemonAssistWindowToggleButton -> Text in Button DaemonAssistWindowToggleButton

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
