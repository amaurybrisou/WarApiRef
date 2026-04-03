# EA_LabelCheckButton

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
| Addons seen in | Bloody Mess, BuffHead, CDown, Calling, CastSequence, DAoCBuff, DammazKron, DetauntHelper |
| Files seen in | Bloody Mess.xml, CDownSettingsTabs.xml, Calling.xml, Conf/DK_Config.xml, FastInteract.xml, Gui.xml, Gui/HealGridGuiTabBattlegroup.xml, Gui/HealGridGuiTabGroup.xml |
| Namespaces detected | EA_LabelCheckButton |
| Source kinds | xml_attributes |
| Example locations | BlacklistCheckBox, BloodyMessOptionsEnableCheckBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsAlwaysShowEnable, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsPermanent, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputEnableRemovable, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputShowTooltips |
| XML usage count | 233 |
| XML attribute usage count | 233 |
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

Engine-supplied XML constant or template class referenced by 35 addons.

## Seen In

- Bloody Mess
- BuffHead
- CDown
- Calling
- CastSequence
- DAoCBuff
- DammazKron
- DetauntHelper
- EA_OpenPartyWindow
- EA_UiDebugTools
- EA_UiModWindow
- EveryBodyGuard
- FastInteract
- GroupRange
- GroupSpotter
- HealGrid
- KeyBar
- Keyset
- KeysetMonsterPlay
- LibAddonButton
- LibGroup
- MarkBuff
- Minmap
- NerfedButtons
- Obsidian
- Queue Queuer
- ReliquaryHunter
- SpamBayes
- TastyButtons
- TexturedButtons
- TurretRange
- Vectors
- WSCT
- WarBoard
- nLootLink

## Used By

- BlacklistCheckBox
- BloodyMessOptionsEnableCheckBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsAlwaysShowEnable
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsPermanent
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputEnableRemovable
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputShowTooltips
- BuffHeadSetupDisplayWindowEnableSort
- BuffHeadSetupEffectCacheWindowEnableCaching
- BuffHeadSetupFilterRowTemplateEnable
- BuffHeadSetupLayoutManagerWindowExport
- BuffHeadSetupLayoutManagerWindowImport
- BuffHeadSetupLayoutManagerWindowLayouts
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarEnable
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarForegroundStretch
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarOrientationReverse
- BuffHeadSetupLayoutWindowLayerWindowLayer0Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer1Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer2Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer3Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer4Checkbox
- BuffHeadSetupPerformanceWindowEnableFading
- BuffHeadSetupPerformanceWindowEnablePriorityUpdates
- BuffHeadSetupPerformanceWindowEnableSync
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeFriendly
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeGroup
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeHostile
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeSelf
- BuffHeadSetupPriorityEffectsWindowSortFirst
- BuffHeadSetupTrackersWindowOnTargetChangeClearAlwaysShow
- BuffHeadSetupTrackersWindowOnTargetChangeClearBuffs
- BuffHeadSetupTrackersWindowOnTargetChangeClearDebuffs
- BuffHeadSetupTrackersWindowTrackerEnable
- BuffHeadSetupTrackersWindowTrackerPermanentEffects
- CDownNLayoutSettingsTab_ScrollChild_NBorderCheckBox
- CDownNLayoutSettingsTab_ScrollChild_NFadingCheckBox
- CDownNLayoutSettingsTab_ScrollChild_NGlassCheckBox
- CDownNLayoutSettingsTab_ScrollChild_NHPTECheckBox
- CDownNLayoutSettingsTab_ScrollChild_NTimerBelowCheckBox
- CDownSLayoutSettingsTab_ScrollChild_SBEndCheckBox
- CDownSLayoutSettingsTab_ScrollChild_SBackCheckBox
- CDownSLayoutSettingsTab_ScrollChild_SBorderCheckBox
- CDownSLayoutSettingsTab_ScrollChild_SCDBelowCheckBox
- CDownSLayoutSettingsTab_ScrollChild_SGlassCheckBox
- CDownSLayoutSettingsTab_ScrollChild_SHPTECheckBox
- CDownSLayoutSettingsTab_ScrollChild_SShowNameCheckBox
- CallingSetupGroupInviteOnRequestGA
- CallingSetupGroupInviteOnRequestL
- CallingSetupGroupInviteOnRequestPW
- CallingSetupGroupShowInvites
- CallingSetupGroupSupressInvitesInGroup
- CallingSetupShowCallerIcon
- CallingSetupShowCallingList
- CallingSetupShowKillsList
- CallingSetupShowNotificationColor
- CallingSetupShowTargetIcon
- CallingSetupTargetingEvaluateRvrKills
- CallingSetupTargetingShowTargetMarker
- CallingSetupTargetingWithdrawNPC
- CastSequenceSequenceBuilderWindowEnabledCheckbox
- CastSequenceSequenceBuilderWindowResetInCombatChangedCheckbox
- CastSequenceSequenceBuilderWindowResetOnTargetChangedCheckbox
- CastSequenceSequenceBuilderWindowResetTimeoutCheckbox
- CastSequenceSequenceBuilderWindowShowAttachedCheckbox
- CastSequenceSequenceBuilderWindowSlotHelpCheckbox
- DAoCBuffFrameSettingsTab_ScrollChild_ActiveCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_AdvancedFiltersCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_BuffsBelowCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_GlassCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_HPTECheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_HideLongtimeCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_LongtoPermaCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_SelfEffectsCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_ShowBorderCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_StaticCondenseCheckBox
- DAoCBuffFrameSettings_G4Filter_G4NotResultCheckBox
- DAoCBuffGeneralSettingsTab_ScrollChild_KillBuffsCheckBox
- DAoCBuffGeneralSettingsTab_ScrollChild_ToggleTestModeCheckBox
- DAoCBuffListManagerTab_ScrollChild_RightClickCopyCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_ColorBCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_ColorCCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_CombatHideCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_CondenseCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_DeleteCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnableClasstableCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnableFilterCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_EndaCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_InvertResultCheckBox
- DKconfigWindowDeathBlowsAlert
- DKconfigWindowDeathBlowsPrint
- DKconfigWindowDeathBlowsSound
- DKconfigWindowDeathsAlert
- DKconfigWindowDeathsPrint
- DKconfigWindowDeathsSound
- DKconfigWindowKillsAlert
- DKconfigWindowKillsPrint
- DKconfigWindowKillsSound
- DKconfigWindowMiscHTScore
- DKconfigWindowMiscHyperlink
- DKconfigWindowMiscHyperlinkSC
- DKconfigWindowMiscMoth
- DTC_BarSettingsTemplateBarPropertiesShowDPS
- DTC_BarSettingsTemplateBarPropertiesShowIcon
- DTC_BarSettingsTemplateBarPropertiesShowLabel
- DTC_BarSettingsTemplateMiscellaneousDetauntRefresh
- DTC_BarSettingsTemplateSoundNotificationsSoundOn
- DTC_TARGETS_TemplateMonitorOff
- DTC_TARGETS_TemplateMonitorOn
- DTC_TARGETS_TemplateScenario
- EA_LabelCheckButtonSmallCopy
- EA_Window_OpenPartyManageSocketDefAutoLootRvR
- EA_Window_OpenPartyManageSocketDefNeedOnUseButton
- EA_Window_OpenPartyNearbySocketDefOpenPartyFlag
- EA_Window_OpenPartyWorldSocketDefHideFullPartiesCheck
- EA_Window_OpenPartyWorldSocketDefSearchableCheck
- EveryBodyGuard_Templates_CheckBoxCheckBox
- FastInteractWindowCheckHealer
- FastInteractWindowCheckQuestAccept
- FastInteractWindowCheckQuestComplete
- FastInteractWindowCheckRalleyMaster
- FastInteractWindowCheckVendorJD
- FastInteractWindowCheckVendorSkip
- GroupRangeSetupStyleGroupBoxWindowAutoSize
- GroupRangeSetupStyleGroupBoxWindowMovable
- GroupRangeSetupStylePointerReverseWindowAutoScale
- GroupRangeSetupStylePointerReverseWindowHideInRange
- GroupRangeSetupStylePointerWindowAutoPosition
- GroupRangeSetupStylePointerWindowAutoScale
- GroupRangeSetupStylePointerWindowHideInRange
- GroupRangeSetupStyleSimpleTextWindowAutoPosition
- GroupRangeSetupStyleWindowUseStyle
- GroupSpotterSettingsWindow_LinkCheckBox
- GroupSpotterSettingsWindow_ModeCheckBox0
- GroupSpotterSettingsWindow_ModeCheckBox1
- GroupSpotterSettingsWindow_ModeCheckBox2
- GroupSpotterSettingsWindow_SpotterCheckBox
- HGG_TabBattlegroupWindowTemplateHideWARWarbandFrames
- HGG_TabBattlegroupWindowTemplatePreviewGrid
- HGG_TabBattlegroupWindowTemplateShowHUD
- HGG_TabBattlegroupWindowTemplateStayHidden
- HGG_TabGroupWindowTemplateHideInBattlegroup
- HGG_TabGroupWindowTemplateHideInScenariogroup
- HGG_TabGroupWindowTemplateHideWARGroupFrame
- HGG_TabGroupWindowTemplatePreviewGrid
- HGG_TabGroupWindowTemplateShowAvatar
- HGG_TabGroupWindowTemplateShowHUD
- HGG_TabGroupWindowTemplateShowPets
- HGG_TabGroupWindowTemplateStayHidden
- HGG_TabHUDWindowTemplateStayHidden
- HGG_TabMouseClickWindowTemplateCheckButtonModifierKeyAlt
- HGG_TabMouseClickWindowTemplateCheckButtonModifierKeyAny
- HGG_TabMouseClickWindowTemplateCheckButtonModifierKeyCtrl
- HGG_TabMouseClickWindowTemplateCheckButtonModifierKeyShift
- HGG_TabScenariogroupWindowTemplateCollapseGroups
- HGG_TabScenariogroupWindowTemplateHideWARGroupFrame
- HGG_TabScenariogroupWindowTemplatePreviewGrid
- HGG_TabScenariogroupWindowTemplateShowHUD
- HGG_TabScenariogroupWindowTemplateShowUngroupedPlayers
- HGG_TabScenariogroupWindowTemplateStayHidden
- HGG_TabTooltipWindowTemplateShowTooltip
- KeyBarCheckBox
- KeysetMonsterPlayProfileWindowEnable
- KeysetSaveWindowSaveAll
- LibAddonButtonManagerAdvancedWindowItemTextureAnimation
- LibAddonButtonManagerAdvancedWindowItemTextureSimple
- LibAddonButtonManagerWindowAutoSort
- LibAddonButtonManagerWindowShowButton
- LibGroupSetupWindowEnableGroupDistance
- MarkBuffSetupWindowManualSelect
- MinmapPinMenu_WorldZoomCheckBox
- NBSB_ParamRowTemplateboolean
- ObsidianSetupCastbarWindowElementGeneralAccurate
- ObsidianSetupCastbarWindowElementGlobalCooldownEnable
- ObsidianSetupCastbarWindowElementGlobalCooldownReverse
- ObsidianSetupCastbarWindowElementIconEnable
- ObsidianSetupCastbarWindowElementLatencyEnable
- ObsidianSetupCastbarWindowElementLatencyEnableText
- ObsidianSetupCastbarWindowElementRangeEnable
- ObsidianSetupEffectTrackerWindowElementTrackerElementGeneralEnable
- ObsidianSetupEffectTrackerWindowElementTrackerElementIconEnable
- OptionCheckBox
- ReliquaryHunterOptionsWindowEnableWorldMapWindowBackgroundOpacity
- ReliquaryHunterOptionsWindowEnableWorldMapWindowOpacity
- SpamBayesOptionsIsAddOnEnabled
- TastyButtonsOptionsWindowCreateViewCheckCustomTexture
- TastyButtonsOptionsWindowCreateViewCheckDelAll
- TastyButtonsOptionsWindowCreateViewCheckisButtonRound
- TastyButtonsOptionsWindowEditViewCheckCustomTexture
- TastyButtonsOptionsWindowEditViewCheckisButtonRound
- TastyButtonsOptionsWindowMiscViewCheckABCMode
- TastyButtonsOptionsWindowMiscViewCheckAPHandler
- TastyButtonsOptionsWindowMiscViewCheckAutoMode
- TastyButtonsOptionsWindowMiscViewCheckGlobalCooldown
- TastyButtonsOptionsWindowMiscViewCheckHideToolTip
- TastyButtonsOptionsWindowMiscViewCheckResourceHandler
- TastyButtonsOptionsWindowMiscViewCheckShowEmpty
- TastyButtonsOptionsWindowMiscViewCheckShowHotkeys
- TastyButtonsOptionsWindowMiscViewCheckTimeIndicator
- TastyButtonsOptionsWindowMiscViewCheckTintIconOnly
- TexturedButtonsSetupActionbarWindowEnable
- TexturedButtonsSetupActionbarWindowHideBackground
- TexturedButtonsSetupActionbarWindowHideEmpty
- TexturedButtonsSetupAdvancedTexturesWindowDisabled
- TexturedButtonsSetupAdvancedTexturesWindowUseCustom
- TexturedButtonsSetupAdvancedTexturesWindowUsePreset
- TexturedButtonsSetupCooldownWindowEnable
- TexturedButtonsSetupCooldownWindowEnableButtonTinting
- TexturedButtonsSetupCooldownWindowHideFlash
- TexturedButtonsSetupCooldownWindowRemoveCooldownS
- TexturedButtonsSetupCooldownWindowShowGlobalCooldownText
- TexturedButtonsSetupFontsWindowEnable
- TexturedButtonsSetupFontsWindowHideCooldown
- TexturedButtonsSetupFontsWindowHideKeybind
- TexturedButtonsSetupMiscWindowEnableCooldownPulse
- TexturedButtonsSetupMiscWindowHideActive
- TexturedButtonsSetupMiscWindowHideGlow
- TexturedButtonsSetupMiscWindowHideQuicklock
- TexturedButtonsSetupMiscWindowMovableQuicklock
- TexturedButtonsSetupMiscWindowSaveQuicklock
- TexturedButtonsSetupTexturesWindowDisabled
- TexturedButtonsSetupTexturesWindowUseCustom
- TexturedButtonsSetupTexturesWindowUsePreset
- TexturedButtonsSetupTintWindowEnable
- TurretRangeSetupDisplayWindowCircleInvert
- TurretRangeSetupDistanceWindowHide
- TurretRangeSetupGeneralWindowEnable
- UiModAdvancedWindowShowInAddOnsListCheck
- Vectors_Templates_CheckBoxCheckBox
- WSCTCheckBox
- WarBoardButtonCheckBox
- WarBoardButtonCheckBoxDefault
- WarBoardButtonCheckBoxEnabled
- nLootLinkGUIFilterSetItem
- nLootLinkGUIFilterUsable

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
