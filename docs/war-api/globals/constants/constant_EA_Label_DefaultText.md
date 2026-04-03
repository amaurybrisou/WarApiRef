# EA_Label_DefaultText

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
| Addons seen in | AdvancedPetAssist, AuctionStats, AutoBand, AutoSalvage, DetauntHelper, EA_OpenPartyWindow, EA_ScenarioGroupWindow, FastFriends |
| Files seen in | APAGui.xml, AutoBandWindowConfig.xml, AutoBandWindowTemplate.xml, AutoBandWindowTools.xml, Checkbox.xml, CustomAutoRoll.xml, FastFriends_ConfigUI.xml, JunkDumpOptions.xml |
| Namespaces detected | EA_Label_DefaultText |
| Source kinds | xml_attributes |
| Example locations | APAFollowTargetHUDLabel, APAInstantOnlyHUDLabel, APAKitingHUDLabel, APALabelAttackBind, APALabelAutoReattack, APALabelAutoReattackDelay |
| XML usage count | 338 |
| XML attribute usage count | 338 |
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

Engine-supplied XML constant or template class referenced by 40 addons.

## Seen In

- AdvancedPetAssist
- AuctionStats
- AutoBand
- AutoSalvage
- DetauntHelper
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- FastFriends
- FozAuction
- JunkDump
- LoyalPet
- MapMonster
- MapPin
- PartyAd
- PeaceOut
- PieTracker
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- Rolodex
- RvRStats
- RvRStatsTab
- Statdoll Remix
- TaxPayer
- ThinkOutLoud
- TidyChat
- TidyRoll
- TokenMachine
- Tome Titan
- TomeTracker
- WARCommander
- WSCT
- Wargames
- nLootLink
- wbLeadHelper
- zMailMod

## Used By

- APAFollowTargetHUDLabel
- APAInstantOnlyHUDLabel
- APAKitingHUDLabel
- APALabelAttackBind
- APALabelAutoReattack
- APALabelAutoReattackDelay
- APALabelCastDelay
- APALabelCastOnAcquire
- APALabelCombatExitDelay
- APALabelDebug
- APALabelEnabled
- APALabelFTAutoArmMount
- APALabelFTFollowOnMount
- APALabelFTPendingDelay
- APALabelFTRedirectCooldown
- APALabelFTTargetFilter
- APALabelFollowTarget
- APALabelGlobalCooldown
- APALabelHUDColorOff
- APALabelHUDColorOn
- APALabelHUDSection
- APALabelHUDVisible
- APALabelHeelBind
- APALabelHeelKeepsFT
- APALabelInstantOnlyHUDSection
- APALabelInstantOnlyHUDVisible
- APALabelInstantPriorityWindow
- APALabelInstantReserveWindow
- APALabelKiting
- APALabelKitingDuration
- APALabelKitingGrace
- APALabelKitingGrowth
- APALabelKitingHUDColorOff
- APALabelKitingHUDColorOn
- APALabelKitingHUDSection
- APALabelKitingHUDVisible
- APALabelKitingInstantOnly
- APALabelKitingInterval
- APALabelKitingMaxMult
- APALabelKitingPostCastFollow
- APALabelLeash
- APALabelLeashHardCeiling
- APALabelLeashInterval
- APALabelLeashRetries
- APALabelLos
- APALabelLosFollow
- APALabelLosFollowGrowth
- APALabelLosFollowMaxMult
- APALabelLosInterval
- APALabelLosNudges
- APALabelNoTargetGrace
- APALabelPTColor
- APALabelPTEnabled
- APALabelPTSection
- APALabelPassiveAutoAttack
- APALabelRmbArmsFT
- APALabelSectionAutoRecall
- APALabelSectionKiting
- APALabelSectionLos
- APALabelSectionMouseControls
- APALabelStance
- APAPetTargetHUDHP
- APAPetTargetHUDName
- AuStatsOptionsBodyExitOnCompleteLabel
- AuStatsOptionsBodyItemIndexLabel
- AuStatsOptionsBodyItemShowToolTipLabel
- AuStatsOptionsBodyProgressLabel
- AuctionSellControlsBuyOutPriceHeader
- AuctionSellControlsDepositHeader
- AuctionSellControlsItemName
- AuctionSellControlsVendorPriceHeader
- AuctionStatsCheckboxTemplateLabel
- AuctionWindowErrorText
- AuctionWindowListLabel
- AutoBandWindowConfigAltCheckLabel
- AutoBandWindowConfigAutoKickLabel
- AutoBandWindowConfigAutoKickLowRankLabel
- AutoBandWindowConfigAutoKickRvRZonesLabel
- AutoBandWindowConfigAutoKickStealthersLabel
- AutoBandWindowConfigAutoKickToofarLabel
- AutoBandWindowConfigBWSorcAsMDPSLabel
- AutoBandWindowConfigBackfillLabel
- AutoBandWindowConfigComboBoxLabel
- AutoBandWindowConfigDefaultLabel
- AutoBandWindowConfigDpsWeightingLabel
- AutoBandWindowConfigExcludeRealmHealerAltSpecLabel
- AutoBandWindowConfigGuildPriorityLabel
- AutoBandWindowConfigHealerRankReqLabel
- AutoBandWindowConfigKickIgnoreLabel
- AutoBandWindowConfigKickTimeLabel
- AutoBandWindowConfigMaxMdpsLabel
- AutoBandWindowConfigMaxRdpsLabel
- AutoBandWindowConfigRankReqLabel
- AutoBandWindowConfigRestrictRaceLabel
- AutoBandWindowTemplateComboBoxLabel
- AutoBandWindowTemplateDeleteConfirmMessageLabel
- AutoBandWindowTemplateDpsLabel
- AutoBandWindowTemplateHealersLabel
- AutoBandWindowTemplateRolesLabel
- AutoBandWindowTemplateSaveNameLabel
- AutoBandWindowTemplateTanksLabel
- AutoBandWindowToolsKickLabel
- AutoBandWindowToolsNoMicLabel
- AutoBandWindowToolsSearchRoleLabel
- AutoSalvage_Stat_TemplateLabel
- AutoSalvage_Stat_TemplatePosition
- CharacterWindowRvRStatsClassA
- CharacterWindowRvRStatsClassA_KillCount
- CharacterWindowRvRStatsClassA_Name
- CharacterWindowRvRStatsClassB
- CharacterWindowRvRStatsClassB_KillCount
- CharacterWindowRvRStatsClassB_Name
- CharacterWindowRvRStatsClassC
- CharacterWindowRvRStatsClassC_KillCount
- CharacterWindowRvRStatsClassC_Name
- CharacterWindowRvRStatsClassD
- CharacterWindowRvRStatsClassD_KillCount
- CharacterWindowRvRStatsClassD_Name
- CharacterWindowRvRStatsClassE
- CharacterWindowRvRStatsClassE_KillCount
- CharacterWindowRvRStatsClassE_Name
- CharacterWindowRvRStatsClassF
- CharacterWindowRvRStatsClassF_KillCount
- CharacterWindowRvRStatsClassF_Name
- CharacterWindowRvRStatsClassG
- CharacterWindowRvRStatsClassG_KillCount
- CharacterWindowRvRStatsClassG_Name
- CharacterWindowRvRStatsClassH
- CharacterWindowRvRStatsClassH_KillCount
- CharacterWindowRvRStatsClassH_Name
- CharacterWindowRvRStatsClassI
- CharacterWindowRvRStatsClassI_KillCount
- CharacterWindowRvRStatsClassI_Name
- CharacterWindowRvRStatsClassJ
- CharacterWindowRvRStatsClassJ_KillCount
- CharacterWindowRvRStatsClassJ_Name
- CharacterWindowRvRStatsClassK
- CharacterWindowRvRStatsClassK_KillCount
- CharacterWindowRvRStatsClassK_Name
- CharacterWindowRvRStatsClassKills
- CharacterWindowRvRStatsClassL
- CharacterWindowRvRStatsClassL_KillCount
- CharacterWindowRvRStatsClassL_Name
- CharacterWindowRvRStatsLifetimeD
- CharacterWindowRvRStatsLifetimeDB
- CharacterWindowRvRStatsLifetimeK
- CharacterWindowRvRStatsLifetimeRatioDBD
- CharacterWindowRvRStatsLifetimeRatioKD
- CharacterWindowRvRStatsRenownBonus
- CharacterWindowRvRStatsSessionK
- CharacterWindowRvRStatsXPBonus
- DTC_TARGETS_RowTemplatePlayerClass
- DTC_TARGETS_RowTemplatePlayerDPS
- DTC_TARGETS_RowTemplatePlayerLast
- DTC_TARGETS_RowTemplatePlayerName
- DTC_TARGETS_RowTemplatePlayerRank
- EA_Template_AutoRollRarityHeader
- EA_Template_AutoRollTitle
- EA_Template_OpenParty_Label
- EA_Window_OpenPartyManageSocketDefLegendAssistantText
- EA_Window_OpenPartyManageSocketDefLegendLeaderMarkText
- EA_Window_OpenPartyManageSocketDefLegendLeaderText
- EA_Window_OpenPartyManageSocketDefLegendMainAssistText
- EA_Window_OpenPartyManageSocketDefLegendMasterLooterText
- EA_Window_OpenPartyManageSocketDefLootModeTitle
- EA_Window_OpenPartyManageSocketDefLootThresholdTitle
- EA_Window_OpenPartyManageSocketDefMasterLooterTitle
- EA_Window_OpenPartyManageSocketDefWarband1Label
- EA_Window_OpenPartyManageSocketDefWarband2Label
- EA_Window_OpenPartyManageSocketDefWarband3Label
- EA_Window_OpenPartyManageSocketDefWarband4Label
- EA_Window_OpenPartyNearbySocketDefLegendAllianceText
- EA_Window_OpenPartyNearbySocketDefLegendFriendText
- EA_Window_OpenPartyNearbySocketDefLegendGuildText
- EA_Window_OpenPartyNearbySocketDefLegendIgnoredText
- FastFriendsConfigMainCharOverrideLabel
- FastFriendsConfigMainSyncMasterLabel
- FastFriendsConfigMainSyncModeLabel
- JunkDumpOptionsWinFriendsName
- LPETOptionsAbilityHealthLabel
- LPETOptionsAbilityModeLabel
- LPETOptionsAbilityNameLabel
- LPETOptionsAbilityPriorityLabel
- LPETOptionsAttackRangeCheckLabel
- LPETOptionsAutoAttackLabel
- LPETOptionsAutoDefendLabel
- LPETOptionsAutoSwitchLabel
- LPETOptionsBiteLabel
- LPETOptionsClawSweepLabel
- LPETOptionsCoruscatingEnergyLabel
- LPETOptionsDaemonicConsumptionLabel
- LPETOptionsDaemonicFireLabel
- LPETOptionsDeathFromAboveLabel
- LPETOptionsDefendRangeCheckLabel
- LPETOptionsFangAndClawLabel
- LPETOptionsFlameOfTzeentchLabel
- LPETOptionsFlamesOfChangeLabel
- LPETOptionsFlamethrowerLabel
- LPETOptionsGoopShootinLabel
- LPETOptionsGoreLabel
- LPETOptionsGutRipperLabel
- LPETOptionsHeadButtLabel
- LPETOptionsHighExplosiveGrenadeLabel
- LPETOptionsLegTearLabel
- LPETOptionsLionsRoarLabel
- LPETOptionsMachineGunLabel
- LPETOptionsMaulLabel
- LPETOptionsPenetratingRoundLabel
- LPETOptionsPetAttackLabel
- LPETOptionsPetFollowLabel
- LPETOptionsPoisonedSpineLabel
- LPETOptionsSelfFollowLabel
- LPETOptionsShockGrenadeLabel
- LPETOptionsShredLabel
- LPETOptionsSpineFlingLabel
- LPETOptionsSporeCloudLabel
- LPETOptionsSquigSquealLabel
- LPETOptionsSteamVentLabel
- LPETOptionsSwitchRangeCheckLabel
- LPETOptionsTerrifyingRoarLabel
- LPETOptionsWarpingEnergyLabel
- MapMonsterEditorWindowHeaderDefault
- MapMonsterEditorWindowLabelDefault
- MapMonsterPinTypeEditorWindowHeaderDefault
- MapMonsterPinTypeEditorWindowLabelDefault
- MapMonster_EditorWindowDatestampLabel
- MapMonster_EditorWindowMessageBox
- MapMonster_PinTypeEditorWindowMessageBox
- MapPin_SetupRotateFactorLabel
- MapPin_SetupRotateFactorValue
- MapPin_SetupScaleFactorLabel
- MapPin_SetupScaleFactorValue
- PartyAdWindowAdvertisePreviewLabel
- PartyAdWindowAdvertisePreviewTextLabel
- PartyAdWindowCurrentColumnLabel
- PartyAdWindowDpsCurrentValue
- PartyAdWindowDpsLabel
- PartyAdWindowHealerCurrentValue
- PartyAdWindowHealerLabel
- PartyAdWindowOptionsLabel
- PartyAdWindowPreviewLabel
- PartyAdWindowPreviewNoteStatusLabel
- PartyAdWindowPreviewNoteTextLabel
- PartyAdWindowPurposeLabel
- PartyAdWindowPurposePresetLabel
- PartyAdWindowSelfRoleLabel
- PartyAdWindowTankCurrentValue
- PartyAdWindowTankLabel
- PartyAdWindowTargetsLabel
- PartyAdWindowUseDetectedNeedsLabel
- PartyAdWindowUseSpecCheckLabel
- PeaceOutDisplay
- PieTrackerRowTextTemplate
- RVAPI_ColorDialogEditBoxTemplateLabel
- RVAPI_ColorDialogEditBoxTemplateMetrics
- RVAPI_ColorDialogSliderTemplateEdit
- RVAPI_ColorDialogSliderTemplateLabel
- RVAPI_ColorDialogWindowColorCurrentLabel
- RVAPI_ColorDialogWindowColorNewLabel
- RVAPI_RangeSettingsWindowLabelImportantInformation
- RVAPI_RangeSettingsWindowLabelMapDistancesEnabled
- RVMOD_3DPortraitSettingsWindowLabelEnabled
- RVMOD_ManagerSettingsWindowLabelFadeInOutDelay
- RVMOD_ManagerSettingsWindowLabelUseGlobalScale
- RVMOD_ManagerSettingsWindowLabelZoomInOutDelay
- RVMOD_PlayerStatusSettingsAPBarColorCaption
- RVMOD_PlayerStatusSettingsHPBarColorCaption
- RVMOD_PlayerStatusSettingsLabelEnable
- RVMOD_SquaredDistancesSettingsWindowLabelAnchorPoints
- RVMOD_SquaredDistancesSettingsWindowLabelColor
- RVMOD_SquaredDistancesSettingsWindowLabelFonts
- RVMOD_SquaredDistancesSettingsWindowLabelImportantInformation
- RVMOD_SquaredDistancesSettingsWindowLabelLayers
- RVMOD_SquaredDistancesSettingsWindowLabelUseGlobalScale
- RVMOD_TargetsFrameRowTemplateLabelTemplates
- RVMOD_TargetsFrameRowTemplateLabelTypes
- RVMOD_TargetsTCTitle
- RolodexRowHeader
- RolodexRowName
- ScenarioGroupMiddleBarGroupToggleButtonText
- ScenarioGroupMiddleBarInstructions
- SingleSGroupTemplateName
- StatdollOptionsBackgroundAlphaLabel
- StatdollOptionsBackgroundAlphaValue
- StatdollOptionsButtonGroupVersion
- StatdollOptionsCheckboxTemplateLabel
- StatdollOptionsComboboxTemplateLabel
- StatdollOptionsScaleFactorLabel
- StatdollOptionsScaleFactorValue
- StatdollOptionsTitle
- TOLSettingsWindowPhraseEditWindowPhraseInstructionLabel
- TOLSettingsWindowPhraseEditWindowTitleLabel
- TOLSettingsWindowSkillEditWindowTitleLabel
- TRollAutoRollTitleLabel
- TTitanUICheckBoxLabel
- TaxPayerRowTextTemplate
- TidyChatOptionsTitleLabel
- TidyChatOptionsVersionLabel
- TokenOptionTemplateText
- TomeTracker_Button_SagaListingRowEntryName
- WARCommanderConfigWindowHeaderDefault
- WARCommanderConfigWindowLabelDefault
- WARCommanderConfig_ListenCheckBoxLabelTemplate
- WARCommander_ConfigWindowUpdatesCheckBoxLabel
- WARCommander_ConfigWindowUpdatesTitleCheckBoxLabel
- WARCommander_ConfigWindowWaypointCheckBoxLabel
- WSCTOptionsColorPickerWindowCustomColorText
- WSCTOptionsEventWindowLabel
- WSCTOptionsFrameWindowLabel
- WSCTOptionsProfileWindowCustomLabel
- WSCTOptionsProfileWindowLabel
- WargamesGemsScoreDisplayText
- WargamesGemsScoreText
- WargamesGemsTimeText
- nLootLinkGUIFilterCraftingLevelHeader
- nLootLinkGUIFilterCraftingLevelSeparator
- nLootLinkGUIFilterRankHeader
- nLootLinkGUIFilterRankSeparator
- nLootLinkGUIFilterRenownHeader
- nLootLinkGUIFilterRenownSeparator
- nLootLinkOptionsLabelRarity
- wbLeadHelperConfigTabLfgIconsLabel
- wbLeadHelperConfigTabMessageEndLabel
- wbLeadHelperConfigTabMessageLabel
- wbLeadHelperConfigTabMessageStartLabel
- wbLeadHelperConfigTabMessageTextColorLabel
- zMM_Default_CraftingLabel
- zMM_Default_RarityLabel
- zMailModLogHeaderItem
- zMailModLogHeaderMoney
- zMailModLogHeaderSender
- zMailModLogHeaderSubject
- zMailModLogHeaderType
- zMailModLogRowTemplateItemCount
- zMailModLogRowTemplateItemData
- zMailModLogRowTemplateSender
- zMailModLogRowTemplateSubject
- zMailModLogRowTemplateType

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
