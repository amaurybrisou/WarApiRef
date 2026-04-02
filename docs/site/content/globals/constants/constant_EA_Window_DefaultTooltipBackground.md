# EA_Window_DefaultTooltipBackground

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
| Addons seen in | Aura, DAoCBuff, Enemy, GetStats, Killer, MapMonster, MapPin, Miracle Grow Remix |
| Files seen in | `/workspace_addons/Aura/Source/AuraTooltip.xml:161`, `/workspace_addons/Aura/Source/AuraTooltip.xml:22`, `/workspace_addons/DAoCBuff/Source/DAoCTooltips.xml:38`, `/workspace_addons/Enemy/Code/CombatLog/CombatLogStatsWindow.xml:436`, `/workspace_addons/GetStats/GetStats.xml:115`, `/workspace_addons/GetStats/GetStats.xml:41`, `/workspace_addons/Killer/Killer.xml:216`, `/workspace_addons/MGRemix/MGRemix.xml:337` |
| Namespaces detected | EA_Window_DefaultTooltipBackground |
| Source kinds | xml_attributes |
| Example locations | AuraShareTooltipBackground, AuraTooltipBackground, DAoCBuffCondenseTooltipBackground, EnemyCombatLogStatsWindowTooltipBackground, GetStatsCompareWindowBackground, GetStatsWindowBackground |
| XML usage count | 12 |
| XML attribute usage count | 12 |
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

Observed engine XML template or inherited constant referenced by 9 addons.

## Seen In

- Aura
- DAoCBuff
- Enemy
- GetStats
- Killer
- MapMonster
- MapPin
- Miracle Grow Remix
- Shinies

## Used By

- AuraShareTooltipBackground
- AuraTooltipBackground
- DAoCBuffCondenseTooltipBackground
- EnemyCombatLogStatsWindowTooltipBackground
- GetStatsCompareWindowBackground
- GetStatsWindowBackground
- KillerTooltipBackground
- MapMonsterTooltipBackground
- MapPinCallTemplateWindowBackground
- MapPinTooltipsBackground
- MiracleGrow2RepeatTipBackground
- ShiniesTooltipWindowBackground

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
