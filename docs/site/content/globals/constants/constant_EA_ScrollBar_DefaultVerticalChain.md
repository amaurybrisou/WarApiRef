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
| Addons seen in | CDown, Crusher, DAoCBuff, DPSMeter, DeepSleep, DuffTimer, EA_ScenarioGroupWindow, EA_UiModWindow |
| Files seen in | CDownSettingsTabs.xml, Code/CombatLog/CombatLogConfiguration.xml, Code/GroupIcons/GroupIconsConfiguration.xml, Code/Guard/GuardConfiguration.xml, Code/KillSpam/KillSpamConfiguration.xml, Code/TalismanAlerter/TalismanAlerterConfiguration.xml, Code/Timer/TimerConfiguration.xml, Code/UnitFrames/ClickCastingDialog.xml |
| Namespaces detected | EA_ScrollBar_DefaultVerticalChain |
| Source kinds | xml_attributes |
| Example locations | $parentScrollBar, AuctionWindowVertScroll, CDownColorSettingsTab_Scrollbar, CDownGeneralSettingsTab_Scrollbar, CDownNLayoutSettingsTab_Scrollbar, CDownSLayoutSettingsTab_Scrollbar |
| XML usage count | 53 |
| XML attribute usage count | 53 |
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

Engine-supplied XML constant or template class referenced by 28 addons.

## Seen In

- CDown
- Crusher
- DAoCBuff
- DPSMeter
- DeepSleep
- DuffTimer
- EA_ScenarioGroupWindow
- EA_UiModWindow
- Enemy
- EveryBodyGuard
- FozAuction
- Hopper
- Killer
- Miracle Grow Remix
- Motion
- PotionBar
- Pure
- RVMOD_Manager
- RVMOD_Targets
- Shinies
- SocialWindow 2.0
- Tome Titan
- Vectors
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- bigger_MacroWindow
- wbLeadHelper

## Used By

- $parentScrollBar
- AuctionWindowVertScroll
- CDownColorSettingsTab_Scrollbar
- CDownGeneralSettingsTab_Scrollbar
- CDownNLayoutSettingsTab_Scrollbar
- CDownSLayoutSettingsTab_Scrollbar
- CrusherConfigScrollBar
- DAoCBuffFrameSettingsTab_Scrollbar
- DAoCBuffGeneralSettingsTab_Scrollbar
- DAoCBuffListManagerTab_Scrollbar
- DAoCBuffMessageWindowScrollWindowScrollbar
- DAoCBuff_Settings_FilterFrame_Scrollbar
- DPSMeterVerticalScrollbar
- DeepSleep_Settings_ScrollWindow_Scrollbar
- DuffTimerOptions_GeneralTab_Template_Scrollbar
- DuffTimerOptions_WinTab_Scrollbar
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
- EveryBodyGuard_Settings_ScrollWindow_Scrollbar
- HopperConfigScrollBar
- IconsScrollbar
- KillerScoreDetailsWindowListScrollbar
- KillerWindowFeedScrollbar
- MiracleGrow2SettingsConfigScrollScroll
- MiracleGrow2SettingsLayoutScrollScroll
- MiracleGrow2SettingsPresetScrollScroll
- MotionConfigScrollBar
- PureConfigScrollBar
- RVMOD_ManagerModInfoTemplateScrollbar
- RVMOD_ManagerWindowContentListBoxScrollbar
- RVMOD_TargetsSettingsWindowFramesScrollbar
- ShiniesAutoUI_AutoSummaryScrollbar
- ShiniesConfigUI_DisplayBar
- SocialWindowListWindowScrollWindowScrollbar
- TTitanUIMainScrollbar
- TTitanUIVertScroll
- UngroupedPlayerVertScroll
- Vectors_Settings_ScrollWindow_Scrollbar
- WCDBConfigScrollBar
- WCDPConfigScrollBar
- WbLeadHelperMessageContentScrollbar
- WhoHealedMeDetailsContentSpellListScrollbar
- WhoHealedMeWindowContentHealerListScrollbar

## Related APIs

- [VerticalScrollbar](../../xml/element_types/element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type

## Notes

- none
