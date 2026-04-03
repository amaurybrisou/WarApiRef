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
| Addons seen in | DAoCBuff, Enemy, Killer, PotionBar, Shinies, WhoHealedMe, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/DAoCBuff/Source/DAoCBuffMsgWindow.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2134`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2464`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:5`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:988`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogConfiguration.xml:0`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIconsConfiguration.xml:0`, `/workspace/data/raw/Enemy/Code/Guard/GuardConfiguration.xml:0` |
| Namespaces detected | EA_ScrollBar_DefaultVerticalChain |
| Source kinds | xml_attributes |
| Example locations | $parentScrollBar, DAoCBuffFrameSettingsTab_Scrollbar, DAoCBuffGeneralSettingsTab_Scrollbar, DAoCBuffListManagerTab_Scrollbar, DAoCBuffMessageWindowScrollWindowScrollbar, DAoCBuff_Settings_FilterFrame_Scrollbar |
| XML usage count | 24 |
| XML attribute usage count | 24 |
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

Observed engine XML template or inherited constant referenced by 7 addons.

## Seen In

- DAoCBuff
- Enemy
- Killer
- PotionBar
- Shinies
- WhoHealedMe
- bigger_MacroWindow

## Used By

- $parentScrollBar
- DAoCBuffFrameSettingsTab_Scrollbar
- DAoCBuffGeneralSettingsTab_Scrollbar
- DAoCBuffListManagerTab_Scrollbar
- DAoCBuffMessageWindowScrollWindowScrollbar
- DAoCBuff_Settings_FilterFrame_Scrollbar
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
- ShiniesAutoUI_AutoSummaryScrollbar
- ShiniesConfigUI_DisplayBar
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
