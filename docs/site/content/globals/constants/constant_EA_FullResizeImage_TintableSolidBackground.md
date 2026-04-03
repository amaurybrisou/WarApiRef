# EA_FullResizeImage_TintableSolidBackground

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
| Addons seen in | AggroMeter, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, Killer, PartyCast |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTooltip.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0` |
| Namespaces detected | EA_FullResizeImage_TintableSolidBackground |
| Source kinds | xml_attributes |
| Example locations | $parentScrollRightPaneBackground, AggroMeterWindow_AggroWindow1BorderCheck, AggroMeterWindow_AggroWindow1Seperator1, AuraShareTooltipFooterSeperator, AuraShareTooltipHeaderSeperator, AuraSharesRowBackground |
| XML usage count | 89 |
| XML attribute usage count | 89 |
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

Observed engine XML template or inherited constant referenced by 17 addons.

## Seen In

- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- PartyCast
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WarBoard
- WhoHealedMe

## Used By

- $parentScrollRightPaneBackground
- AggroMeterWindow_AggroWindow1BorderCheck
- AggroMeterWindow_AggroWindow1Seperator1
- AuraShareTooltipFooterSeperator
- AuraShareTooltipHeaderSeperator
- AuraSharesRowBackground
- AuraTooltipFooterSeperator
- AuraTooltipHeaderSeperator
- AuraWindowRowBackground
- AutoRollRowTemplateBackground
- BuffHeadLayoutFrameWindowBackground
- BuffHeadSetupAdvancedCompressionItemRowTemplateBackground
- BuffHeadSetupAdvancedCompressionRowTemplateBackground
- BuffHeadSetupAdvancedContainersItemRowTemplateBackground
- BuffHeadSetupAdvancedContainersRowTemplateBackground
- BuffHeadSetupEffectCacheRowTemplateBackground
- BuffHeadSetupFilterRowTemplateBackground
- BuffHeadSetupLayoutWindowLayerWindowBackground
- BuffHeadSetupMenuRowTemplateBackground
- BuffHeadSetupPriorityEffectsItemRowTemplateBackground
- BuffHeadSetupPriorityEffectsRowTemplateBackground
- BuffHeadSetupSelectColorWindowBackground
- ClosetGoblinSetRowBackgroundName
- ClosetGoblinZoneRowBackgroundZone
- DAoCBuffFrameSettings_FilterRowBackground
- DAoCBuff_FakeSettingsRowBackground
- DAoCBuff_FrameSettingsRowBackground
- EnemyCombatLogIDSAnchorBackground
- EnemyCombatLogIDSRowBackground
- EnemyCombatLogStatsWindowSnapshot_ListRowTemplateBackground
- EnemyCombatLogStatsWindow_ListRowTemplateBackground
- EnemyCombatLogTargetDefeseTotalWindowBackground
- EnemyCombatLogTargetDefeseWindowBackground
- EnemyConfigurationWindow_PropertyColorTemplateExample
- EnemyDebugBackground
- EnemyEffectsIndicatorDialogContentScrollChildEffectFiltersListBackground
- EnemyKillSpamDialogBackground
- EnemyKillSpam_AreaStatsRowTemplateFlash
- EnemyKillSpam_RowTemplateFlash
- EnemyMarksWindowBackground
- EnemyScenarioInfoDialogBackground
- EnemySimpleListRowTemplateBackground
- EnemyUnitFramesAnchor1Background
- EnemyUnitFramesAnchor2Background
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsListBackground
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsListBackground
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsListBackground
- Frame_BG_TemlateBackground3
- KillerPersonalCounterBg
- KillerTooltipSep
- KillerWindowFeedDivider
- LayoutTemplateBackground
- ModBackgroundTemplate
- PartyCastWindow_Template_LargeBG
- PartyCastWindow_Template_PlainBG
- PotionBarButtonTemplateCooldown
- RoR_SoR_PopperBG
- RoR_SoR_PopperBG2
- RoR_SoR_RealmTemplateKEEP1HPBARBARCOUNTER
- RoR_SoR_RealmTemplateKEEP1HPBARBG
- RoR_SoR_RealmTemplateKEEP1PROGRESSBARCOUNTER
- RoR_SoR_RealmTemplateKEEP1PROGRESSBG
- RoR_SoR_RealmTemplateKEEP1PROGRESSBLINK
- RoR_SoR_RealmTemplateKEEP2HPBARBARCOUNTER
- RoR_SoR_RealmTemplateKEEP2HPBARBG
- RoR_SoR_RealmTemplateKEEP2PROGRESSBARCOUNTER
- RoR_SoR_RealmTemplateKEEP2PROGRESSBG
- RoR_SoR_RealmTemplateKEEP2PROGRESSBLINK
- RoR_SoR_RealmTemplateLockTint
- ShiniesAuctionsUI_ResultsRowBackground
- ShiniesAutoUI_ItemRowBackground
- ShiniesBrowseUI_ResultsRowBackground
- ShiniesBrowseUI_SearchesRowBackground
- ShiniesConfigPrice_PriorityRowBackground
- ShiniesConfigUI_ModulesRowBackground
- ShiniesPostUI_ResultsRowBackground
- TexturedButtonsSetupMenuRowTemplateBackground
- TexturedButtonsSetupSelectColorWindowBackground
- TidyChatLootRollRowTemplateBackground
- TurretRangeSetupDisplayWindowSelectColorBackground
- TurretRangeSetupDistancesRowTemplateBackground
- WarBoardBackground
- WhoHealedMeDetailRowTemplateFill
- WhoHealedMeDetailsContentFill
- WhoHealedMeHealerRowTemplateFill
- WhoHealedMeOptionsContentBackgroundFillPreviewFill
- WhoHealedMeOptionsContentFill
- WhoHealedMeWindowContentFill
- WhoHealedMeWindowContentSummaryFill

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
