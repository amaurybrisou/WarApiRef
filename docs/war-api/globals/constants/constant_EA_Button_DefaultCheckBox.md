# EA_Button_DefaultCheckBox

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
| Addons seen in | AdvancedRenownTrainer, ArmorGraphicToggle, Atlas, AuctionStats, Aura, AutoBand, AutoSalvage, CCTV |
| Files seen in | AGTSettingsWindow.xml, AdvancedRenownTrainingPresets.xml, AutoBandWindowConfig.xml, AutoBandWindowTemplate.xml, AutoBandWindowTools.xml, CCTV.xml, Checkbox.xml, Cheeseboard.xml |
| Namespaces detected | EA_Button_DefaultCheckBox |
| Source kinds | xml_attributes |
| Example locations | AGTSettingsWindowCloakInCombat, AGTSettingsWindowCloakNoCombat, AGTSettingsWindowHelmInCombat, AGTSettingsWindowHelmNoCombat, AGTSettingsWindowHeraldryInCombat, AGTSettingsWindowHeraldryNoCombat |
| XML usage count | 264 |
| XML attribute usage count | 264 |
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

Engine-supplied XML constant or template class referenced by 64 addons.

## Seen In

- AdvancedRenownTrainer
- ArmorGraphicToggle
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoSalvage
- CCTV
- CM_ClosetGoblin
- CMap
- Cheeseboard
- Dascore
- Deathblow
- Deathblow2
- DuffTimer
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiModWindow
- EZCraftX
- Enemy
- FastFriends
- GDes
- Ges
- Group Icons SG
- HealGrid
- Killer
- MapMonster
- MapPin
- MarkBuff
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- PartyAd
- Phantom
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RememberIt
- Res
- SNT_CASTBAR
- SNT_INFO
- ScenarioStats
- Sequencer
- Shinies
- SocialWindow 2.0
- Statdoll Remix
- TalismanGenie
- TaxPayer
- TidyChat
- TidyQueue
- TidyRoll
- Tome Titan
- WARCommander
- WTes
- WarBoard_FPS
- emotes
- nLootLink
- wbLeadHelper
- zMailMod

## Used By

- AGTSettingsWindowCloakInCombat
- AGTSettingsWindowCloakNoCombat
- AGTSettingsWindowHelmInCombat
- AGTSettingsWindowHelmNoCombat
- AGTSettingsWindowHeraldryInCombat
- AGTSettingsWindowHeraldryNoCombat
- AdvancedRenownTrainingPresetsWindowSaveSelectedCheckBox
- AuStatsOptionsBodyExitOnComplete
- AuStatsOptionsBodyItemIndexChk
- AuStatsOptionsBodyItemShowToolTip
- AuctionStatsCheckboxTemplateButton
- Aura_LabelCheckButtonButton
- Aura_LargeLabelCheckButtonButton
- AutoBandWindowConfigAltCheckCheckBox
- AutoBandWindowConfigBWSorcAsMDPSCheckBox
- AutoBandWindowConfigBackfillCheckBox
- AutoBandWindowConfigCommonRaceNamesCheckBox
- AutoBandWindowConfigDpsWeightingCheckBox
- AutoBandWindowConfigExcludeRealmHealerAltSpecCheckBox
- AutoBandWindowConfigGuildPriorityCheckBox
- AutoBandWindowConfigKickCheckBox
- AutoBandWindowConfigKickIgnoreCheckBox
- AutoBandWindowConfigKickLowRankCheckBox
- AutoBandWindowConfigKickRvRZonesCheckBox
- AutoBandWindowConfigKickStealthersCheckBox
- AutoBandWindowConfigKickToofarCheckBox
- AutoBandWindowConfigRestrictRaceCheckBox
- AutoBandWindowTemplateDefaultCheckBox
- AutoBandWindowTemplateRightClickTemplateMenuCheckBox
- AutoBandWindowToolsAutoFormSearchCheckBox
- AutoBandWindowToolsAutoPartyNoteCheckBox
- AutoBandWindowToolsAutoSearchCheckBox
- AutoBandWindowToolsDiscordReqCheckBox
- AutoBandWindowToolsKickOfflineCheckBox
- AutoBandWindowToolsKickRankCheckBox
- AutoBandWindowToolsKickZoneCheckBox
- AutoBandWindowToolsNoMicCheckBox
- AutoBandWindowToolsNotifyBuffsCheckBox
- AutoBandWindowToolsPrintRoleCheckBox
- AutoBandWindowToolsRightClickOrganizeCheckBox
- AutoBandWindowToolsSearchRoleChan1CheckBox
- AutoBandWindowToolsSearchRoleChan5CheckBox
- AutoBandWindowToolsSearchRoleChanT4CheckBox
- AutoSalvage_Stat_TemplateCheckBox
- CCTVSettingsCheckBox1
- CCTVSettingsCheckBox_DISARM
- CCTVSettingsCheckBox_KD
- CCTVSettingsCheckBox_ROOT
- CCTVSettingsCheckBox_SILENCE
- CCTVSettingsCheckBox_SNARE
- CCTVSettingsCheckBox_STAGGER
- CCTVSettingsCheckBox_THROW
- CheeseboardWindowDefaultCheck
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB2
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB3
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB4
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB5
- ClosetGoblinCharacterWindowContentsCheckboxUseItemRack
- ClosetGoblinCharacterWindowContentsEquipmentShowCloak
- ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry
- ClosetGoblinCharacterWindowContentsEquipmentShowHelm
- Cmap_MapPinFilterEntryButton1
- Cmap_MapPinFilterEntryButton2
- Cmap_MapPinGutterEntryButton1
- Cmap_MapPinGutterEntryButton2
- DascoreWin1WindowTemplatecheckboxCheckbox
- DuffTimerOptions_CheckBox_chk
- EA_Button_UiModuleListRowTemplateEnabled
- EA_SocialFilterEntryButton
- EA_Window_OpenPartyManageSocketDefWarband1Show
- EA_Window_OpenPartyManageSocketDefWarband2Show
- EA_Window_OpenPartyManageSocketDefWarband3Show
- EA_Window_OpenPartyManageSocketDefWarband4Show
- EZCraftXWindow.useCraftingResult.Selector
- EZCraftXWindow.usePowerPreview.Selector
- EnemyAssistConfigurationMarkNewTarget
- EnemyAssistConfigurationNewTargetSound
- EnemyAssistConfigurationTargetOnNotifyClick
- EnemyClickCastingDialogContentScrollChildArchetype1
- EnemyClickCastingDialogContentScrollChildArchetype2
- EnemyClickCastingDialogContentScrollChildArchetype3
- EnemyClickCastingDialogContentScrollChildExceptMe
- EnemyClickCastingDialogContentScrollChildKeyModifier1
- EnemyClickCastingDialogContentScrollChildKeyModifier2
- EnemyClickCastingDialogContentScrollChildKeyModifier3
- EnemyConfigurationWindow_PropertyBoolTemplateValue
- EnemyEffectsIndicatorDialogContentScrollChildArchetype1
- EnemyEffectsIndicatorDialogContentScrollChildArchetype2
- EnemyEffectsIndicatorDialogContentScrollChildArchetype3
- EnemyEffectsIndicatorDialogContentScrollChildCircleIcon
- EnemyEffectsIndicatorDialogContentScrollChildExceptMe
- EnemyEffectsIndicatorDialogContentScrollChildLScaleCheckBox
- EnemyScenarioAlerterConfigurationEnabled
- EnemyUnitFramePartDialogContentScrollChildArchetype1
- EnemyUnitFramePartDialogContentScrollChildArchetype2
- EnemyUnitFramePartDialogContentScrollChildArchetype3
- EnemyUnitFramePartDialogContentScrollChildExceptMe
- EnemyUnitFramesConfigurationContentScrollChildEnabled
- EnemyUnitFramesConfigurationContentScrollChildSortingEnabled
- FastFriendsConfigMainCharOverrideDefaultRadio
- FastFriendsConfigMainCharOverrideOffRadio
- FastFriendsConfigMainCharOverrideOnRadio
- FastFriendsConfigMainFriendsCheckbox
- FastFriendsConfigMainIgnoreCheckbox
- FastFriendsConfigMainModeAllRadio
- FastFriendsConfigMainModeLevelRadio
- FastFriendsConfigMainModeOptInRadio
- FastFriendsConfigMainSyncMasterCheckbox
- GDesOptionsColourIconHealthCheckBox
- GDesOptionsExtrasFeatureEnableCheckBox
- GDesOptionsGeneralModCheckBox
- GDesOptionsGeneralSoundEnableCheckBox
- GDesOptionsGeneralToggleCheckBox
- GesOptionsExtrasFeatureEnableCheckBox
- GesOptionsGeneralModCheckBox
- GesOptionsGeneralMonitorClickCheckBox
- GesOptionsGeneralMonitorHealthCheckBox
- GesOptionsGeneralMonitorIconCheckBox
- GesOptionsGeneralSoundEnableCheckBox
- GesOptionsGeneralToggleCheckBox
- GesOptionsGeneralTrackerHealthCheckBox
- GesOptionsGeneralTrackerIconCheckBox
- GroupIconsSGOptionsChk1
- GroupIconsSGOptionsChk2
- GroupIconsSGOptionsChk3
- GroupIconsSGOptionsChk4
- GroupIconsSGOptionsChk5
- HGG_LabelCheckButtonButton
- KillerSettingsWindowContentPersonalEnabled
- MapFilterContextMenuChoiceCheckBox
- MapMonster_EditorWindowIsLockedCheckBox
- MapMonster_EditorWindowIsPrivateCheckBox
- MapMonster_PinTypeEditorWindowIsLockedCheckBox
- MapMonster_PinTypeEditorWindowIsPrivateCheckBox
- MapPin_SetupBGBox
- MarkBuffContextMenuItemBuffCheckBoxCheckBox
- MegaphoneMainHighlightLeaderCheckbox
- MegaphoneMainShowLeaderCheckbox
- MinmapMapPinFilterEntryButton
- MinmapMapScenarioEntryButton
- MiracleGrow2LayoutSettingsTemplateStage0
- MiracleGrow2LayoutSettingsTemplateStage1
- MiracleGrow2LayoutSettingsTemplateStage2
- MiracleGrow2LayoutSettingsTemplateStage3
- MiracleGrow2LayoutSettingsTemplateStage4
- PPMainSaveSettingsButton
- PartyAdWindowUseDetectedNeedsButton
- PartyAdWindowUseSpecCheckButton
- PhantomSettingsHideBarLock
- PhantomSettingsHideGroupBuffs
- PhantomSettingsHideMainAssist
- PhantomSettingsHideMapArea
- PhantomSettingsHideMapCity
- PhantomSettingsHideMapFrame
- PhantomSettingsHideMapMail
- PhantomSettingsHideMapPins
- PhantomSettingsHideMapRally
- PhantomSettingsHideMapScen
- PhantomSettingsHideMapWorld
- PhantomSettingsHideMapZoom
- PhantomSettingsHidePet
- PhantomSettingsHidePlayerBuffs
- PhantomSettingsHideSocial
- PhantomSettingsHideTargetBuffs
- PieTrackerOptionsGuildCheckBox
- PieTrackerOptionsModCheckBox
- PotionBarTypeTemplateAutohideCheck
- PotionBarTypeTemplateCheck
- PotionBarTypeTemplateDontSplitNameCheck
- PotionBarTypeTemplateShow1Check
- PotionBarTypeTemplateShowLastCheck
- PureLabelCheckButtonButton
- RVAPI_RangeSettingsWindowCheckBoxMapDistancesEnabled
- RVMOD_3DPortraitSettingsWindowCheckBoxEnabled
- RVMOD_PlayerStatusSettingsCheckBoxEnable
- RVMOD_SquaredDistancesSettingsWindowCheckBoxUseGlobalScale
- RememberItSettingsOnDeath
- RememberItSettingsOnKill
- RememberItSettingsOnRankUp
- RememberItSettingsOnRenownRankUp
- RememberItSettingsOnScenarioEnd
- ResOptionsEnableCastCheckBox
- ResOptionsEnableModCheckBox
- ResOptionsEnableToggleCheckBox
- SNT_CASTBAR_SETUP_WINDOW1_chk
- SNT_CASTBAR_SETUP_WINDOW2_chk
- SNT_CASTBAR_SETUP_WINDOW3_chk
- SNT_CASTBAR_SETUP_WINDOW4_chk
- SNT_CASTBAR_SETUP_WINDOW5_chk
- SNT_CASTBAR_SETUP_WINDOW6_chk
- SNT_CASTBAR_SETUP_WINDOW7_chk
- SNT_CASTBAR_SETUP_WINDOW8_chk
- SNT_INFO_SETUP_WINDOW_B0
- SNT_INFO_SETUP_WINDOW_B1
- SNT_INFO_SETUP_WINDOW_B2
- SNT_INFO_SETUP_WINDOW_B3
- SNT_INFO_SETUP_WINDOW_F0
- SNT_INFO_SETUP_WINDOW_F1
- SNT_INFO_SETUP_WINDOW_F2
- SNT_INFO_SETUP_WINDOW_F3
- SNT_INFO_SETUP_WINDOW_F4
- SNT_INFO_SETUP_WINDOW_F5
- SNT_INFO_SETUP_WINDOW_F6
- SNT_INFO_SETUP_WINDOW_F7
- SNT_INFO_SETUP_WINDOW_I0
- SNT_INFO_SETUP_WINDOW_I1
- SNT_INFO_SETUP_WINDOW_I2
- SNT_INFO_SETUP_WINDOW_I3
- ScenarioConfigWindowCheckBox1
- ScenarioConfigWindowCheckBox2
- ScenarioConfigWindowCheckBox3
- ScenarioConfigWindowCheckBox4
- ScenarioConfigWindowCheckBox5
- ScenarioConfigWindowCheckBox6
- ScenarioConfigWindowCheckBox7
- ScenarioConfigWindowCheckBox8
- ScenarioGroupMiddleBarGroupToggleButton
- Sequencer_WindowCheckBoxCombat
- Sequencer_WindowCheckBoxCooldown
- Sequencer_WindowCheckBoxEnemy
- Sequencer_WindowCheckBoxFriend
- Sequencer_WindowCheckBoxTimer
- SettingsWindowCheckBox1
- SettingsWindowCheckBox2
- SettingsWindowCheckBox3
- SettingsWindowCheckBox4
- ShiniesConfigPrice_CheckBox
- ShiniesConfigPrice_PriorityRowEnable
- Shinies_CheckButtonWithLabelButton
- SiegeChatFiltersRowTemplateCheckBox
- SingleSGroupTemplateVisibleButton
- SocialWindowTabOptionsSocketAdvisorPreference
- SocialWindowTabOptionsSocketAnonymousPreference
- SocialWindowTabOptionsSocketDisableBuddyListPreference
- SocialWindowTabOptionsSocketHiddenPreference
- SocialWindowTabOptionsSocketPrivatePartyPreference
- StatdollOptionsCheckboxTemplateButton
- TAtlasCheckboxButton
- TChatCheckboxTemplateButton
- TQueueCheckboxTemplateButton
- TRollCheckboxTemplateButton
- TTitanUICheckBoxButton
- TalismanGenieContainerCheckBox
- TalismanGenieCurioCheckBox
- TalismanGenieEssenceCheckBox
- TalismanGenieGoldDustCheckBox
- TaxPayerOptionsGuildCheckBox
- TaxPayerOptionsModCheckBox
- WARCommanderConfig_ListenCheckBoxTemplate
- WARCommander_ConfigWindowUpdatesCheckBox
- WARCommander_ConfigWindowUpdatesTitleCheckBox
- WARCommander_ConfigWindowWaypointCheckBox
- WTesBackpackCheckBox
- WTesGlyphPosCheckBox
- WTesMoraleClickCheckBox
- WTesScenStartPosCheckBox
- WTesSiegeChatCheckBox
- WarBoard_FPSOptions_chkShowAvg
- emotes_MapPinFilterEntryCheckbox
- nLootLinkOptionsCheckRarityConfirmButton
- wbLeadHelperConfigTabLfgIconsCheckBox
- zMailModCheckboxTemplate
- zMailModOptionsCheckbox

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- none
