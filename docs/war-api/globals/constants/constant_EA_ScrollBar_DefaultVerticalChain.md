# EA_ScrollBar_DefaultVerticalChain

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
| Addons seen in | DAoCBuff, EA_UiModWindow, Enemy, Killer, Miracle Grow Remix, PotionBar, RVMOD_Manager, Shinies |
| Files seen in | `/workspace_addons/DAoCBuff/Source/DAoCBuffMsgWindow.xml:63`, `/workspace_addons/Enemy/Code/CombatLog/CombatLogConfiguration.xml:13`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIconsConfiguration.xml:13`, `/workspace_addons/Enemy/Code/Guard/GuardConfiguration.xml:13`, `/workspace_addons/Enemy/Code/KillSpam/KillSpamConfiguration.xml:13`, `/workspace_addons/Enemy/Code/TalismanAlerter/TalismanAlerterConfiguration.xml:13`, `/workspace_addons/Enemy/Code/Timer/TimerConfiguration.xml:13`, `/workspace_addons/Enemy/Code/UnitFrames/ClickCastingDialog.xml:46` |
| Namespaces detected | EA_ScrollBar_DefaultVerticalChain |
| Source kinds | xml_attributes |
| Example locations | $parentScrollBar, DAoCBuffMessageWindowScrollWindowScrollbar, EA_ScrollWindow_ModInfoTemplateScrollbar, EA_Window_MacroMacrosScrollbar, EnemyClickCastingDialogContentScrollbar, EnemyCombatLogConfigurationContentScrollbar |
| XML usage count | 27 |
| XML attribute usage count | 27 |
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

Observed engine XML template or inherited constant referenced by 11 addons.

## Seen In

- DAoCBuff
- EA_UiModWindow
- Enemy
- Killer
- Miracle Grow Remix
- PotionBar
- RVMOD_Manager
- Shinies
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Used By

- $parentScrollBar
- DAoCBuffMessageWindowScrollWindowScrollbar
- EA_ScrollWindow_ModInfoTemplateScrollbar
- EA_Window_MacroMacrosScrollbar
- EnemyClickCastingDialogContentScrollbar
- EnemyCombatLogConfigurationContentScrollbar
- EnemyEffectsIndicatorDialogContentScrollbar
- EnemyGroupIconsConfigurationContentScrollbar
- EnemyGuardConfigurationContentScrollbar
- EnemyKillSpamConfigurationContentScrollbar
- EnemyTalismanAlerterConfigurationContentScrollbar
- EnemyTimerConfigurationContentScrollbar
- EnemyUnitFramePartDialogContentScrollbar
- EnemyUnitFramesConfigurationContentScrollbar
- IconsScrollbar
- KillerScoreDetailsWindowListScrollbar
- KillerWindowFeedScrollbar
- MiracleGrow2SettingsConfigScrollScroll
- MiracleGrow2SettingsLayoutScrollScroll
- MiracleGrow2SettingsPresetScrollScroll
- RVMOD_ManagerModInfoTemplateScrollbar
- RVMOD_ManagerWindowContentListBoxScrollbar
- ShiniesAutoUI_AutoSummaryScrollbar
- ShiniesConfigUI_DisplayBar
- WbLeadHelperMessageContentScrollbar
- WhoHealedMeDetailsContentSpellListScrollbar
- WhoHealedMeWindowContentHealerListScrollbar

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
