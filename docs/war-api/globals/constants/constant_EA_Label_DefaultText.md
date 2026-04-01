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
| Addons seen in | AdvancedPetAssist, AutoBand, JunkDump, LoyalPet, MapMonster, MapPin, PeaceOut, RVAPI_ColorDialog |
| Files seen in | `/workspace/AdvancedPetAssist/APAGui.xml:1012`, `/workspace/AdvancedPetAssist/APAGui.xml:1034`, `/workspace/AdvancedPetAssist/APAGui.xml:1056`, `/workspace/AdvancedPetAssist/APAGui.xml:1068`, `/workspace/AdvancedPetAssist/APAGui.xml:1090`, `/workspace/AdvancedPetAssist/APAGui.xml:1102`, `/workspace/AdvancedPetAssist/APAGui.xml:1114`, `/workspace/AdvancedPetAssist/APAGui.xml:1126` |
| Namespaces detected | EA_Label_DefaultText |
| Source kinds | xml_attributes |
| Example locations | APAFollowTargetHUDLabel, APAInstantOnlyHUDLabel, APAKitingHUDLabel, APALabelAttackBind, APALabelAutoReattack, APALabelAutoReattackDelay |
| XML usage count | 152 |
| XML attribute usage count | 152 |
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

Observed engine XML template or inherited constant referenced by 14 addons.

## Seen In

- AdvancedPetAssist
- AutoBand
- JunkDump
- LoyalPet
- MapMonster
- MapPin
- PeaceOut
- RVAPI_ColorDialog
- RVMOD_Manager
- ThinkOutLoud
- TidyChat
- TidyRoll
- WSCT
- wbLeadHelper

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
- AutoBandWindowConfigAutoKickLabel
- AutoBandWindowConfigAutoKickToofarLabel
- AutoBandWindowConfigComboBoxLabel
- AutoBandWindowConfigDefaultLabel
- AutoBandWindowConfigKickTimeLabel
- AutoBandWindowConfigMapIconLabel
- AutoBandWindowConfigRankReqLabel
- AutoBandWindowTemplateComboBoxLabel
- AutoBandWindowTemplateSaveNameLabel
- AutoBandWindowToolsKickLabel
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
- PeaceOutDisplay
- RVAPI_ColorDialogEditBoxTemplateLabel
- RVAPI_ColorDialogEditBoxTemplateMetrics
- RVAPI_ColorDialogSliderTemplateEdit
- RVAPI_ColorDialogSliderTemplateLabel
- RVAPI_ColorDialogWindowColorCurrentLabel
- RVAPI_ColorDialogWindowColorNewLabel
- RVMOD_ManagerSettingsWindowLabelFadeInOutDelay
- RVMOD_ManagerSettingsWindowLabelUseGlobalScale
- RVMOD_ManagerSettingsWindowLabelZoomInOutDelay
- TOLSettingsWindowPhraseEditWindowTitleLabel
- TOLSettingsWindowSkillEditWindowTitleLabel
- TRollAutoRollTitleLabel
- TidyChatOptionsTitleLabel
- TidyChatOptionsVersionLabel
- WSCTOptionsColorPickerWindowCustomColorText
- WSCTOptionsEventWindowLabel
- WSCTOptionsFrameWindowLabel
- WSCTOptionsProfileWindowCustomLabel
- WSCTOptionsProfileWindowLabel
- wbLeadHelperConfigTabLfgIconsLabel
- wbLeadHelperConfigTabMessageEndLabel
- wbLeadHelperConfigTabMessageLabel
- wbLeadHelperConfigTabMessageStartLabel
- wbLeadHelperConfigTabMessageTextColorLabel

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
