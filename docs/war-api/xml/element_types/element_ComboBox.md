# ComboBox

- Category: XML Element Type
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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, AutoBand, BankArkel, BuffHead, Busted, CCTV |
| Files seen in | Code/Assist/AssistConfiguration.xml, Code/CombatLog/CombatLogStatsWindow.xml, Code/Core/ConfigDialog.xml, Code/Core/ConfigurationWindow.xml, Code/Core/Groups/EffectFilterDialog.xml, Code/Intercom/ChooseChannelDialog.xml, Code/Intercom/IntercomJoinDialog.xml, Code/UnitFrames/ClickCastingDialog.xml |
| Namespaces detected | ComboBox |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAComboAttackBind, AdvancedPetAssist: APAComboAutoReattack, AdvancedPetAssist: APAComboAutoReattackDelay, AdvancedPetAssist: APAComboCastDelay, AdvancedPetAssist: APAComboCastOnAcquire, AdvancedPetAssist: APAComboCombatExitDelay |
| XML usage count | 694 |
| XML attribute usage count | 694 |
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

ComboBox is an interactive XML control. It commonly appears under ScrollWindow and Window. It is typically used to organize structural children such as Anchors, EventHandlers, MenuButtonOffset and bind XML events like OnLButtonUp, OnMouseOver, OnMouseOverEnd to Lua.

## Common Attributes

- autoresize
- id
- inherits
- layer
- maxvisibleitems
- menubackground
- menuitembutton
- name
- scrollbar
- selectedbutton
- show
- warnOnTextCropped

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnSelChanged](../handlers/handler_OnSelChanged.md)

## Common Handler Functions

- APAGui.OnComboChanged
- EA_Window_OpenPartyWorld.UpdateComboBoxDisable
- LPET.QuickHealthOnSelChanged
- LPET.QuickModeOnSelChanged
- LPET.QuickPriorityOnSelChanged
- WHMGui.OnOptionsComboChanged
- DK_Config.OnUpdateSoundCombo
- DetauntBarSettings.OnUpdateCombo
- MapPin.ComboBoxUpdate
- BustedGUI.SelectAddonView
- CDownSettings.GrowHorizontalChanged
- CDownSettings.GrowLeftChanged
- CDownSettings.GrowUpChanged
- CDownSettings.TimerFontChanged
- DaemonAssist.OnBindingComboChanged
- DuffTimer.Options.OnSelChanged
- Dye.OnSelChanged
- EA_Window_OpenPartyWorld.PopulateSpecificComboBoxes
- Enemy.ConfigurationWindow_OnChange
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Killer.OnSettingsComboChanged
- Megaphone.SaveSettings
- PP.UpdateDyeList
- TalismanGenie.CurioChange
- TalismanGenie.EssenceChange
- TalismanGenie.FragmentChange
- TalismanGenie.GoldDustChange
- TalismanGenie.TalismanStats
- TaxPayer.UpdateSelection
- Warbuilder.ComboBoxUpdate
- AdvancedRenownTraining.SelectedItemChanged
- AuctionWindowSearchControls.OnChangeCategory
- AuctionWindowSearchControls.OnChangeMaxRank
- AuctionWindowSearchControls.OnChangeMinRank
- AuctionWindowSearchControls.OnCheckItemSlot
- AuctionWindowSearchControls.OnCheckItemType
- AuraConfig.OnActivationSoundComboChanged
- AuraConfig.OnDeactivationSoundComboChanged
- AuraConfig.OnTriggerTypeSelChanged
- AutoBandWindowConfig.OnSelChangedComboBox
- AutoBandWindowTemplate.OnSelChangedComboBox
- BankArkel.PackCombo
- BuffHead.Setup.AdvancedContainersItem.OnPositionChanged
- BuffHead.Setup.AdvancedContainersItem.OnTargetTypeChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsBuffsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsDebuffsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnElementChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthHorizontalChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthVerticalChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnLayerChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnLayoutChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionAnchorChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionPlacementChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertiesChanged
- BuffHead.Setup.Container.OnContainerAlwaysShowAnchorChanged
- BuffHead.Setup.Container.OnContainerAlwaysShowPlacementChanged
- BuffHead.Setup.Container.OnContainerBuffsAnchorChanged
- BuffHead.Setup.Container.OnContainerBuffsPlacementChanged
- BuffHead.Setup.Container.OnContainerDebuffsAnchorChanged
- BuffHead.Setup.Container.OnContainerDebuffsPlacementChanged
- BuffHead.Setup.Container.OnSizeContainerTypeChanged
- BuffHead.Setup.Display.OnLayerChanged
- BuffHead.Setup.Display.OnSortByChanged
- BuffHead.Setup.Display.OnSortDirectionChanged
- BuffHead.Setup.General.OnCompressionChanged
- BuffHead.Setup.Layout.Manager.OnExportLayoutChanged
- BuffHead.Setup.Layout.Manager.OnLayoutsLayoutChanged
- BuffHead.Setup.Layout.Properties.OnDurationFormatChanged
- BuffHead.Setup.Layout.Properties.OnElementChanged
- BuffHead.Setup.Layout.Properties.OnFontAlignmentChanged
- BuffHead.Setup.Layout.Properties.OnIconBorderColorTypeChanged
- BuffHead.Setup.Layout.Properties.OnLayerLayerChanged
- BuffHead.Setup.Layout.Properties.OnPropertiesChanged
- BuffHead.Setup.Layout.Properties.OnStatusBarForegroundColorTypeChanged
- BuffHead.Setup.Layout.Properties.OnStatusBarOrientationChanged
- BuffHead.Setup.Performance.OnEffectAnchoringChanged
- BuffHead.Setup.PriorityEffects.OnAnimationChanged
- BuffHead.Setup.Trackers.OnTrackerBuffChanged
- BuffHead.Setup.Trackers.OnTrackerDebuffChanged
- BuffHead.Setup.Trackers.OnTrackerTypeChanged
- CCTV.MenuSelect
- CCTV.MenuSelect2
- CCTV.MenuSelect3
- CCTV.MenuSelect4
- CDownSettings.ABNameFontChanged
- CDownSettings.CSettingChanged
- CDownSettings.CountChanged
- CDownSettings.LayoutChanged
- CDownSettings.MaxCDChanged
- CDownSettings.MinCDChanged
- CDownSettings.OrderChanged
- CDownSettings.RefreshChanged
- CDownSettings.RowChanged
- CDownSettings.SMaxCountChanged
- CallingSetup.OnLanguageChanged
- CallingSetup.OnPrioSelectionChanged
- CallingSetup.OnPrioTypeChanged
- CastSequence.Setup.OnAdvancementTypeChanged
- ChattyCathy.UpdateEntrySetup
- Cheeseboard.OnClassComboChanged
- Cheeseboard.OnTotalComboChanged
- CrusherConfig_Profiles_OnActiveProfileChanged
- CrusherConfig_Profiles_OnCopyProfileChanged
- DAoCBuffSettings.BufforderChanged
- DAoCBuffSettings.CountChanged
- DAoCBuffSettings.DivideChanged
- DAoCBuffSettings.FilterSettings.ClassTableChanged
- DAoCBuffSettings.FilterSettings.ConditionChanged
- DAoCBuffSettings.FilterSettings.ConditionTypeChanged
- DAoCBuffSettings.FilterSettings.FilterPropertyChanged
- DAoCBuffSettings.FilterSettings.G1FilterChanged
- DAoCBuffSettings.FilterSettings.G2ListChanged
- DAoCBuffSettings.FilterSettings.G4HistoryBrowserChanged
- DAoCBuffSettings.FilterSettings.G4UseandChanged
- DAoCBuffSettings.FilterSettings.G5DurationChanged
- DAoCBuffSettings.FilterSettings.G5FilterChanged
- DAoCBuffSettings.FilterSettings.TextureChanged
- DAoCBuffSettings.FilterSettings.TextureTypeChanged
- DAoCBuffSettings.FilterSettings.UseandChanged
- DAoCBuffSettings.FontChanged
- DAoCBuffSettings.GrowHorizontalChanged
- DAoCBuffSettings.GrowLeftChanged
- DAoCBuffSettings.GrowUpChanged
- DAoCBuffSettings.LeftListChanged
- DAoCBuffSettings.ManagerModeChanged
- DAoCBuffSettings.PermabuffsChanged
- DAoCBuffSettings.RefreshChanged
- DAoCBuffSettings.RemoveListChanged
- DAoCBuffSettings.RightListChanged
- DAoCBuffSettings.RowChanged
- DAoCBuffSettings.StickChanged
- DAoCBuffSettings.TargetChanged
- DAoCBuffSettings.TypeChanged
- DPSMeter.AbilityChanged
- DPSMeter.ValueChanged
- DeepSleep.Settings.ComboBoxChanged
- DetauntBarSettings.OnDetauntReadySelChanged
- DetauntBarSettings.OnDetauntSelChanged
- DetauntBarSettings.OnUpdateFontColorCombo
- DetauntBarSettings.OnUpdateFontCombo
- EA_Window_OpenPartyLootRollOptions.OnOptionChange
- EA_Window_OpenPartyManage.OnLeaderMarkSelChange
- EA_Window_OpenPartyManage.OnLootModeSelChange
- EA_Window_OpenPartyManage.OnLootThresholdSelChange
- EA_Window_OpenPartyManage.OnMasterLooterSelChange
- Enemy.AssistUI_ConfigDialog_OnNewTargetSoundIdSelChanged
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- Enemy.CombatLogUI_StatsWindow_OnTypeSelChanged
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- Enemy.UnitFramesUI_ConfigDialog_OnSorting1Changed
- Enemy.UnitFramesUI_ConfigDialog_OnSorting2Changed
- Enemy.UnitFramesUI_ConfigDialog_OnSorting3Changed
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- Enemy.UnitFramesUI_UnitFramePartDialog_OnTypeSelChanged
- EveryBodyGuard.Settings.ComboBoxChanged
- GDes.ExtrasFeatureSelected
- GDes.ExtrasPositionSelected
- Ges.ColourTabIconSelected
- Ges.ExtrasFeatureSelected
- Ges.ExtrasIconSelected
- Ges.ExtrasPositionSelected
- GroupRangeSetup.Style.GroupBox.OnColorStyleChanged
- GroupRangeSetup.Style.GroupBox.OnGrowthChanged
- GroupRangeSetup.Style.GroupBox.OnLayerChanged
- GroupRangeSetup.Style.OnStyleChanged
- GroupRangeSetup.Style.Pointer.OnColorStyleChanged
- GroupRangeSetup.Style.Pointer.OnLayerChanged
- GroupRangeSetup.Style.PointerReverse.OnColorStyleChanged
- GroupRangeSetup.Style.PointerReverse.OnLayerChanged
- GroupRangeSetup.Style.SimpleText.OnColorStyleChanged
- GroupRangeSetup.Style.SimpleText.OnLayerChanged
- HealGridGuiTabBattlegroup.ChangeShowDebuffFlags
- HealGridGuiTabBattlegroup.ChangeSkin
- HealGridGuiTabBattlegroup.GridGroupingChanged
- HealGridGuiTabBattlegroup.GridGrowthChanged
- HealGridGuiTabBattlegroup.GridOrientationChanged
- HealGridGuiTabBattlegroup.ShowActionPointsBarChanged
- HealGridGuiTabBattlegroup.ShowCareerLineIconChanged
- HealGridGuiTabBattlegroup.ShowHealthPointsBarChanged
- HealGridGuiTabBattlegroup.ShowMoraleLevelBarChanged
- HealGridGuiTabGroup.ChangeShowDebuffFlags
- HealGridGuiTabGroup.ChangeSkin
- HealGridGuiTabGroup.GridGroupingChanged
- HealGridGuiTabGroup.GridGrowthChanged
- HealGridGuiTabGroup.GridOrientationChanged
- HealGridGuiTabGroup.ShowActionPointsBarChanged
- HealGridGuiTabGroup.ShowCareerLineIconChanged
- HealGridGuiTabGroup.ShowHealthPointsBarChanged
- HealGridGuiTabGroup.ShowMoraleLevelBarChanged
- HealGridGuiTabHUD.ChangeShowFriendlyTargetDebuffFlags
- HealGridGuiTabHUD.ChangeShowHostileTargetDebuffFlags
- HealGridGuiTabHUD.ChangeShowPlayerDebuffFlags
- HealGridGuiTabHUD.ChangeSkin
- HealGridGuiTabHUD.ShowActionPointsBarChanged
- HealGridGuiTabHUD.ShowCareerLineIconChanged
- HealGridGuiTabHUD.ShowHealthPointsBarChanged
- HealGridGuiTabHUD.ShowMoraleLevelBarChanged
- HealGridGuiTabHUDBuffs.ChangeFriendlyTargetBuffsFilter
- HealGridGuiTabHUDBuffs.ChangeFriendlyTargetDebuffsFilter
- HealGridGuiTabHUDBuffs.ChangeHostileTargetBuffsFilter
- HealGridGuiTabHUDBuffs.ChangeHostileTargetDebuffsFilter
- HealGridGuiTabHUDBuffs.ChangePlayerBuffsFilter
- HealGridGuiTabHUDBuffs.ChangePlayerDebuffsFilter
- HealGridGuiTabMouseClick.ActionChanged
- HealGridGuiTabScenariogroup.ChangeShowDebuffFlags
- HealGridGuiTabScenariogroup.ChangeSkin
- HealGridGuiTabScenariogroup.GridGroupingChanged
- HealGridGuiTabScenariogroup.GridGrowthChanged
- HealGridGuiTabScenariogroup.GridOrientationChanged
- HealGridGuiTabScenariogroup.ShowActionPointsBarChanged
- HealGridGuiTabScenariogroup.ShowCareerLineIconChanged
- HealGridGuiTabScenariogroup.ShowHealthPointsBarChanged
- HealGridGuiTabScenariogroup.ShowMoraleLevelBarChanged
- HopperConfig_Profiles_OnActiveProfileChanged
- HopperConfig_Profiles_OnCopyProfileChanged
- ItemRack.Sets.Select
- ItemRack.Tactics.OnSetMenuSelectionChanged
- KeysetMonsterPlay.Setup.Profile.OnDefaultProfileChanged
- KeysetMonsterPlay.Setup.Profile.OnMonsterProfileChanged
- LPET.AttackRangeCheckComboOnSelChanged
- LPET.AutoAttackComboOnSelChanged
- LPET.AutoDefendComboOnSelChanged
- LPET.AutoSwitchComboOnSelChanged
- LPET.BiteComboOnSelChanged
- LPET.BiteHealthComboOnSelChanged
- LPET.BitePriorityComboOnSelChanged
- LPET.ClawSweepComboOnSelChanged
- LPET.ClawSweepHealthComboOnSelChanged
- LPET.ClawSweepPriorityComboOnSelChanged
- LPET.ComboOnSelChanged
- LPET.CoruscatingEnergyComboOnSelChanged
- LPET.CoruscatingEnergyHealthComboOnSelChanged
- LPET.CoruscatingEnergyPriorityComboOnSelChanged
- LPET.DaemonicConsumptionComboOnSelChanged
- LPET.DaemonicConsumptionHealthComboOnSelChanged
- LPET.DaemonicConsumptionPriorityComboOnSelChanged
- LPET.DaemonicFireComboOnSelChanged
- LPET.DaemonicFireHealthComboOnSelChanged
- LPET.DaemonicFirePriorityComboOnSelChanged
- LPET.DeathFromAboveComboOnSelChanged
- LPET.DeathFromAboveHealthComboOnSelChanged
- LPET.DeathFromAbovePriorityComboOnSelChanged
- LPET.DefendRangeCheckComboOnSelChanged
- LPET.FangAndClawComboOnSelChanged
- LPET.FangAndClawHealthComboOnSelChanged
- LPET.FangAndClawPriorityComboOnSelChanged
- LPET.FlameOfTzeentchComboOnSelChanged
- LPET.FlameOfTzeentchHealthComboOnSelChanged
- LPET.FlameOfTzeentchPriorityComboOnSelChanged
- LPET.FlamesOfChangeComboOnSelChanged
- LPET.FlamesOfChangeHealthComboOnSelChanged
- LPET.FlamesOfChangePriorityComboOnSelChanged
- LPET.FlamethrowerComboOnSelChanged
- LPET.FlamethrowerHealthComboOnSelChanged
- LPET.FlamethrowerPriorityComboOnSelChanged
- LPET.GoopShootinComboOnSelChanged
- LPET.GoopShootinHealthComboOnSelChanged
- LPET.GoopShootinPriorityComboOnSelChanged
- LPET.GoreComboOnSelChanged
- LPET.GoreHealthComboOnSelChanged
- LPET.GorePriorityComboOnSelChanged
- LPET.GutRipperComboOnSelChanged
- LPET.GutRipperHealthComboOnSelChanged
- LPET.GutRipperPriorityComboOnSelChanged
- LPET.HeadButtComboOnSelChanged
- LPET.HeadButtHealthComboOnSelChanged
- LPET.HeadButtPriorityComboOnSelChanged
- LPET.HighExplosiveGrenadeComboOnSelChanged
- LPET.HighExplosiveGrenadeHealthComboOnSelChanged
- LPET.HighExplosiveGrenadePriorityComboOnSelChanged
- LPET.LegTearComboOnSelChanged
- LPET.LegTearHealthComboOnSelChanged
- LPET.LegTearPriorityComboOnSelChanged
- LPET.LionsRoarComboOnSelChanged
- LPET.LionsRoarHealthComboOnSelChanged
- LPET.LionsRoarPriorityComboOnSelChanged
- LPET.MachineGunComboOnSelChanged
- LPET.MachineGunHealthComboOnSelChanged
- LPET.MachineGunPriorityComboOnSelChanged
- LPET.MaulComboOnSelChanged
- LPET.MaulHealthComboOnSelChanged
- LPET.MaulPriorityComboOnSelChanged
- LPET.PenetratingRoundComboOnSelChanged
- LPET.PenetratingRoundHealthComboOnSelChanged
- LPET.PenetratingRoundPriorityComboOnSelChanged
- LPET.PetAttackComboOnSelChanged
- LPET.PetFollowComboOnSelChanged
- LPET.PoisonedSpineComboOnSelChanged
- LPET.PoisonedSpineHealthComboOnSelChanged
- LPET.PoisonedSpinePriorityComboOnSelChanged
- LPET.ProfilesComboOnSelChanged
- LPET.QuickSettingsComboOnSelChanged
- LPET.SelfFollowComboOnSelChanged
- LPET.ShockGrenadeComboOnSelChanged
- LPET.ShockGrenadeHealthComboOnSelChanged
- LPET.ShockGrenadePriorityComboOnSelChanged
- LPET.ShredComboOnSelChanged
- LPET.ShredHealthComboOnSelChanged
- LPET.ShredPriorityComboOnSelChanged
- LPET.SpineFlingComboOnSelChanged
- LPET.SpineFlingHealthComboOnSelChanged
- LPET.SpineFlingPriorityComboOnSelChanged
- LPET.SporeCloudComboOnSelChanged
- LPET.SporeCloudHealthComboOnSelChanged
- LPET.SporeCloudPriorityComboOnSelChanged
- LPET.SquigSquealComboOnSelChanged
- LPET.SquigSquealHealthComboOnSelChanged
- LPET.SquigSquealPriorityComboOnSelChanged
- LPET.SteamVentComboOnSelChanged
- LPET.SteamVentHealthComboOnSelChanged
- LPET.SteamVentPriorityComboOnSelChanged
- LPET.SwitchRangeCheckComboOnSelChanged
- LPET.TerrifyingRoarComboOnSelChanged
- LPET.TerrifyingRoarHealthComboOnSelChanged
- LPET.TerrifyingRoarPriorityComboOnSelChanged
- LPET.WarpingEnergyComboOnSelChanged
- LPET.WarpingEnergyHealthComboOnSelChanged
- LPET.WarpingEnergyPriorityComboOnSelChanged
- LibAddonButton.Manager.Advanced.OnItemChanged
- LibAddonButton.Manager.CustomItem.OnTypeChanged
- MBuffSetup.SmartBuff.OnClassTypeChanged
- Map.OnSelChanged
- MapMonster.Editor.OnPinTypeChange
- MapMonster.Editor.OnSubTypeChange
- MapMonster.Editor.OnZoneNameChange
- MapMonster.PinTypeEditor.OnSubTypeChange
- MiracleGrow2.LayoutPlotArrChanged
- MiracleGrow2.LayoutShowChanged
- MoraleSet.OnSetMenuSelectionChanged
- MotionConfig_Profiles_OnActiveProfileChanged
- MotionConfig_Profiles_OnCopyProfileChanged
- NBSBCore.OnAddonSelected
- NBSBParam.OnParamChange
- Obsidian.Setup.Castbar.OnElementChanged
- Obsidian.Setup.Castbar.OnFillUpdatePriorityChanged
- Obsidian.Setup.Castbar.OnGlobalCooldownPositionChanged
- Obsidian.Setup.Castbar.OnIconPositionChanged
- Obsidian.Setup.Castbar.OnNameAlignmentChanged
- Obsidian.Setup.Castbar.OnTimerAlignmentChanged
- Obsidian.Setup.EffectTracker.OnBarElementChanged
- Obsidian.Setup.EffectTracker.OnElementChanged
- Obsidian.Setup.EffectTracker.OnFillUpdatePriorityChanged
- Obsidian.Setup.EffectTracker.OnNameAlignmentChanged
- Obsidian.Setup.EffectTracker.OnTimerAlignmentChanged
- Obsidian.Setup.EffectTracker.OnTrackerElementChanged
- Obsidian.Setup.EffectTracker.OnTrackerFillColorChanged
- Obsidian.Setup.EffectTracker.OnTrackerIconPositionChanged
- Obsidian.Setup.EffectTracker.OnTrackerPositionChanged
- PartyAdWindow.OnSelChangedPurposePreset
- PartyAdWindow.OnSelChangedSelfRole
- PotionBarSettings.ActivatorComboSelChanged
- PotionBarSettings.BuildComboSelChanged
- PotionBarSettings.ComboMethod
- PotionBarSettings.InfoTextBRComboSelChanged
- PotionBarSettings.InfoTextTRComboSelChanged
- PotionBarSettings.QuickActionsSelChanged
- PureConfig_Profiles_OnActiveProfileChanged
- PureConfig_Profiles_OnCopyProfileChanged
- RVMOD_Manager.OnSelChangedSortBy
- RVMOD_SquaredDistances.OnComboBoxAnchorPointsChange
- RVMOD_SquaredDistances.OnComboBoxFontsChange
- RVMOD_SquaredDistances.OnComboBoxLayerChange
- RVMOD_Targets.OnComboBoxConditionClick
- RVMOD_Targets.OnComboBoxTemplateChange
- RaidMeter.MenuSelect
- Res.OptionsIconBackgroundSelected
- Res.OptionsIconForegroundSelected
- Res.OptionsIconHighlightSelected
- ScenarioStats.CharacterSelectionChanged
- ScenarioStats.ScenarioSelectionChanged
- ShiniesBrowseUI.OnSelChanged_Criteria_ItemSlotCombo
- ShiniesBrowseUI.OnSelChanged_Criteria_ItemTypeCombo
- ShiniesBrowseUI.OnSelChanged_Criteria_ModifierCombo
- SocialWindowTabSearch.OnFilterSelChanged
- TOLSettingsUI.OnEventSelChanged
- TOLSettingsUI.OnPhraseSelChanged
- TOLSettingsUI.OnSkillSelChanged
- TastyButtonsOptions.OnComboButtonSelectionChanged
- TastyButtonsOptions.OnComboFontChanged
- TastyButtonsOptions.OnComboStateChanged
- TexturedButtons.Setup.Actionbar.OnBarChanged
- TexturedButtons.Setup.Actionbar.OnSelectorChanged
- TexturedButtons.Setup.AdvancedTextures.OnCustomChanged
- TexturedButtons.Setup.AdvancedTextures.OnCustomTextureChanged
- TexturedButtons.Setup.AdvancedTextures.OnCustomTextureSliceChanged
- TexturedButtons.Setup.AdvancedTextures.OnPresetChanged
- TexturedButtons.Setup.AdvancedTextures.OnSlotTypeChanged
- TexturedButtons.Setup.Misc.OnActionButtonPickUpModifierChanged
- TexturedButtons.Setup.Misc.OnCustomGlowChanged
- TexturedButtons.Setup.Textures.OnCustomChanged
- TexturedButtons.Setup.Textures.OnCustomTextureChanged
- TexturedButtons.Setup.Textures.OnCustomTextureSliceChanged
- TexturedButtons.Setup.Textures.OnPresetChanged
- TexturedButtons.Setup.Tint.OnTintTypeChanged
- TidyChat.Options.UpdateGroupTabs
- TidyRoll.CustomAutoRoll.OnChoiceChange
- TokenMachine.UpdateOption
- TurretRange.Setup.Display.OnCircleModeChanged
- TurretRange.Setup.Display.OnCircleTypeChanged
- TurretRange.Setup.Display.OnDistanceLayoutChanged
- TurretRange.Setup.Display.OnDistanceTypeChanged
- TurretRange.Setup.Display.OnElementChanged
- TurretRange.Setup.Display.OnGraphicTypeChanged
- TwisterSet.OnSetMenuSelectionChanged
- UiModWindow.OnCategoryComboSelChanged
- Vectors.Settings.ComboBoxChanged
- WARCommanderConfig.OnRVRAlertsChange
- WARCommanderConfig.OnRVRCommandsChange
- WARCommanderConfig.OnScenAlertsChange
- WARCommanderConfig.OnScenCommandsChange
- WCDBConfig_Filters_OnActiveFilterChanged
- WCDBConfig_Profiles_OnActiveProfileChanged
- WCDBConfig_Profiles_OnCopyProfileChanged
- WCDPConfig_Filters_OnActiveFilterChanged
- WCDPConfig_Profiles_OnActiveProfileChanged
- WCDPConfig_Profiles_OnCopyProfileChanged
- WSCT.ComboOnSelChanged
- WSCT.EventOnSelChanged
- WSCT.FrameComboOnSelChanged
- WSCT.FrameOnSelChanged
- WarBoard_FPSOptions.OnCmbChange
- Warbuilder.ChangeSelect
- Warbuilder.LinkLevel
- Warbuilder.MenuSelect
- ZonePOP.OnFilterSelChanged
- nLootLinkGUI.onCareerSelected
- nLootLinkGUI.onCategorySelected
- nLootLinkGUI.onRaritySelected
- nLootLinkGUI.onSlotSelected
- nLootLinkGUI.onStatSelected
- nLootLinkGUI.onTypeSelected
- nLootLinkOptions.onRaritySelected
- wbLeadHelperConfigTab.OnChanged
- LPET.OnMouseOver
- Enemy.ConfigurationWindow_ShowTooltip
- AutoBandWindowConfig.OnSettingMouseOver
- DevPadWindow.OnMouseOverCode
- EA_Window_OpenPartyManage.OnMouseoverLeaderMark
- EA_Window_OpenPartyManage.OnMouseoverLootModeCombo
- EA_Window_OpenPartyManage.OnMouseoverLootThresholdCombo
- EZCraftX.OnMouseOver_Dropdown
- MoraleSet.OnMouseOverSetMenu
- PotionBarSettings.OnMouseoverActivator
- PotionBarSettings.OnMouseoverBuild
- PotionBarSettings.OnMouseoverInfoTextBR
- PotionBarSettings.OnMouseoverInfoTextTR
- PotionBarSettings.OnMouseoverMethod
- PotionBarSettings.OnMouseoverQuickActions
- TwisterSet.OnMouseover
- WSCT.FrameComboOnMouseOver
- WSCT.OnMouseOver
- WbLeadHelperMessage.OnMouseOverSubmenuComboBox
- WbLeadHelperMessage.OnMouseOverTypeComboBox
- wbLeadHelperConfigTab.OnMouseOverMessageTextColorLabel
- MoraleSet.OnLButtonUpSetMenu
- TwisterSet.OnLClick
- alertMod.OnSliderChanged
- EA_Window_OpenPartyLootRollOptions.OnRClickComboBox
- MoraleSet.OnRButtonUpSetMenu
- TwisterSet.OnRClick
- AutoBandWindowConfig.OnSettingMouseOverEnd


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnSelChanged](../handlers/handler_OnSelChanged.md) | data | APAGui.OnComboChanged, , EA_Window_OpenPartyWorld.UpdateComboBoxDisable, LPET.QuickHealthOnSelChanged, LPET.QuickModeOnSelChanged, LPET.QuickPriorityOnSelChanged, WHMGui.OnOptionsComboChanged, DK_Config.OnUpdateSoundCombo, DetauntBarSettings.OnUpdateCombo, MapPin.ComboBoxUpdate, BustedGUI.SelectAddonView, CDownSettings.GrowHorizontalChanged, CDownSettings.GrowLeftChanged, CDownSettings.GrowUpChanged, CDownSettings.TimerFontChanged, DaemonAssist.OnBindingComboChanged, DuffTimer.Options.OnSelChanged, Dye.OnSelChanged, EA_Window_OpenPartyWorld.PopulateSpecificComboBoxes, Enemy.ConfigurationWindow_OnChange, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, Killer.OnSettingsComboChanged, Megaphone.SaveSettings, PP.UpdateDyeList, TalismanGenie.CurioChange, TalismanGenie.EssenceChange, TalismanGenie.FragmentChange, TalismanGenie.GoldDustChange, TalismanGenie.TalismanStats, TaxPayer.UpdateSelection, Warbuilder.ComboBoxUpdate, AdvancedRenownTraining.SelectedItemChanged, AuctionWindowSearchControls.OnChangeCategory, AuctionWindowSearchControls.OnChangeMaxRank, AuctionWindowSearchControls.OnChangeMinRank, AuctionWindowSearchControls.OnCheckItemSlot, AuctionWindowSearchControls.OnCheckItemType, AuraConfig.OnActivationSoundComboChanged, AuraConfig.OnDeactivationSoundComboChanged, AuraConfig.OnTriggerTypeSelChanged, AutoBandWindowConfig.OnSelChangedComboBox, AutoBandWindowTemplate.OnSelChangedComboBox, BankArkel.PackCombo, BuffHead.Setup.AdvancedContainersItem.OnPositionChanged, BuffHead.Setup.AdvancedContainersItem.OnTargetTypeChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsBuffsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsDebuffsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnElementChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthHorizontalChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthVerticalChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnLayerChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnLayoutChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionAnchorChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionPlacementChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertiesChanged, BuffHead.Setup.Container.OnContainerAlwaysShowAnchorChanged, BuffHead.Setup.Container.OnContainerAlwaysShowPlacementChanged, BuffHead.Setup.Container.OnContainerBuffsAnchorChanged, BuffHead.Setup.Container.OnContainerBuffsPlacementChanged, BuffHead.Setup.Container.OnContainerDebuffsAnchorChanged, BuffHead.Setup.Container.OnContainerDebuffsPlacementChanged, BuffHead.Setup.Container.OnSizeContainerTypeChanged, BuffHead.Setup.Display.OnLayerChanged, BuffHead.Setup.Display.OnSortByChanged, BuffHead.Setup.Display.OnSortDirectionChanged, BuffHead.Setup.General.OnCompressionChanged, BuffHead.Setup.Layout.Manager.OnExportLayoutChanged, BuffHead.Setup.Layout.Manager.OnLayoutsLayoutChanged, BuffHead.Setup.Layout.Properties.OnDurationFormatChanged, BuffHead.Setup.Layout.Properties.OnElementChanged, BuffHead.Setup.Layout.Properties.OnFontAlignmentChanged, BuffHead.Setup.Layout.Properties.OnIconBorderColorTypeChanged, BuffHead.Setup.Layout.Properties.OnLayerLayerChanged, BuffHead.Setup.Layout.Properties.OnPropertiesChanged, BuffHead.Setup.Layout.Properties.OnStatusBarForegroundColorTypeChanged, BuffHead.Setup.Layout.Properties.OnStatusBarOrientationChanged, BuffHead.Setup.Performance.OnEffectAnchoringChanged, BuffHead.Setup.PriorityEffects.OnAnimationChanged, BuffHead.Setup.Trackers.OnTrackerBuffChanged, BuffHead.Setup.Trackers.OnTrackerDebuffChanged, BuffHead.Setup.Trackers.OnTrackerTypeChanged, CCTV.MenuSelect, CCTV.MenuSelect2, CCTV.MenuSelect3, CCTV.MenuSelect4, CDownSettings.ABNameFontChanged, CDownSettings.CSettingChanged, CDownSettings.CountChanged, CDownSettings.LayoutChanged, CDownSettings.MaxCDChanged, CDownSettings.MinCDChanged, CDownSettings.OrderChanged, CDownSettings.RefreshChanged, CDownSettings.RowChanged, CDownSettings.SMaxCountChanged, CallingSetup.OnLanguageChanged, CallingSetup.OnPrioSelectionChanged, CallingSetup.OnPrioTypeChanged, CastSequence.Setup.OnAdvancementTypeChanged, ChattyCathy.UpdateEntrySetup, Cheeseboard.OnClassComboChanged, Cheeseboard.OnTotalComboChanged, CrusherConfig_Profiles_OnActiveProfileChanged, CrusherConfig_Profiles_OnCopyProfileChanged, DAoCBuffSettings.BufforderChanged, DAoCBuffSettings.CountChanged, DAoCBuffSettings.DivideChanged, DAoCBuffSettings.FilterSettings.ClassTableChanged, DAoCBuffSettings.FilterSettings.ConditionChanged, DAoCBuffSettings.FilterSettings.ConditionTypeChanged, DAoCBuffSettings.FilterSettings.FilterPropertyChanged, DAoCBuffSettings.FilterSettings.G1FilterChanged, DAoCBuffSettings.FilterSettings.G2ListChanged, DAoCBuffSettings.FilterSettings.G4HistoryBrowserChanged, DAoCBuffSettings.FilterSettings.G4UseandChanged, DAoCBuffSettings.FilterSettings.G5DurationChanged, DAoCBuffSettings.FilterSettings.G5FilterChanged, DAoCBuffSettings.FilterSettings.TextureChanged, DAoCBuffSettings.FilterSettings.TextureTypeChanged, DAoCBuffSettings.FilterSettings.UseandChanged, DAoCBuffSettings.FontChanged, DAoCBuffSettings.GrowHorizontalChanged, DAoCBuffSettings.GrowLeftChanged, DAoCBuffSettings.GrowUpChanged, DAoCBuffSettings.LeftListChanged, DAoCBuffSettings.ManagerModeChanged, DAoCBuffSettings.PermabuffsChanged, DAoCBuffSettings.RefreshChanged, DAoCBuffSettings.RemoveListChanged, DAoCBuffSettings.RightListChanged, DAoCBuffSettings.RowChanged, DAoCBuffSettings.StickChanged, DAoCBuffSettings.TargetChanged, DAoCBuffSettings.TypeChanged, DPSMeter.AbilityChanged, DPSMeter.ValueChanged, DeepSleep.Settings.ComboBoxChanged, DetauntBarSettings.OnDetauntReadySelChanged, DetauntBarSettings.OnDetauntSelChanged, DetauntBarSettings.OnUpdateFontColorCombo, DetauntBarSettings.OnUpdateFontCombo, EA_Window_OpenPartyLootRollOptions.OnOptionChange, EA_Window_OpenPartyManage.OnLeaderMarkSelChange, EA_Window_OpenPartyManage.OnLootModeSelChange, EA_Window_OpenPartyManage.OnLootThresholdSelChange, EA_Window_OpenPartyManage.OnMasterLooterSelChange, Enemy.AssistUI_ConfigDialog_OnNewTargetSoundIdSelChanged, Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged, Enemy.CombatLogUI_StatsWindow_OnTypeSelChanged, Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged, Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged, Enemy.UI_ConfigDialog_OnSectionSelChanged, Enemy.UnitFramesUI_ConfigDialog_OnSorting1Changed, Enemy.UnitFramesUI_ConfigDialog_OnSorting2Changed, Enemy.UnitFramesUI_ConfigDialog_OnSorting3Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged, Enemy.UnitFramesUI_UnitFramePartDialog_OnTypeSelChanged, EveryBodyGuard.Settings.ComboBoxChanged, GDes.ExtrasFeatureSelected, GDes.ExtrasPositionSelected, Ges.ColourTabIconSelected, Ges.ExtrasFeatureSelected, Ges.ExtrasIconSelected, Ges.ExtrasPositionSelected, GroupRangeSetup.Style.GroupBox.OnColorStyleChanged, GroupRangeSetup.Style.GroupBox.OnGrowthChanged, GroupRangeSetup.Style.GroupBox.OnLayerChanged, GroupRangeSetup.Style.OnStyleChanged, GroupRangeSetup.Style.Pointer.OnColorStyleChanged, GroupRangeSetup.Style.Pointer.OnLayerChanged, GroupRangeSetup.Style.PointerReverse.OnColorStyleChanged, GroupRangeSetup.Style.PointerReverse.OnLayerChanged, GroupRangeSetup.Style.SimpleText.OnColorStyleChanged, GroupRangeSetup.Style.SimpleText.OnLayerChanged, HealGridGuiTabBattlegroup.ChangeShowDebuffFlags, HealGridGuiTabBattlegroup.ChangeSkin, HealGridGuiTabBattlegroup.GridGroupingChanged, HealGridGuiTabBattlegroup.GridGrowthChanged, HealGridGuiTabBattlegroup.GridOrientationChanged, HealGridGuiTabBattlegroup.ShowActionPointsBarChanged, HealGridGuiTabBattlegroup.ShowCareerLineIconChanged, HealGridGuiTabBattlegroup.ShowHealthPointsBarChanged, HealGridGuiTabBattlegroup.ShowMoraleLevelBarChanged, HealGridGuiTabGroup.ChangeShowDebuffFlags, HealGridGuiTabGroup.ChangeSkin, HealGridGuiTabGroup.GridGroupingChanged, HealGridGuiTabGroup.GridGrowthChanged, HealGridGuiTabGroup.GridOrientationChanged, HealGridGuiTabGroup.ShowActionPointsBarChanged, HealGridGuiTabGroup.ShowCareerLineIconChanged, HealGridGuiTabGroup.ShowHealthPointsBarChanged, HealGridGuiTabGroup.ShowMoraleLevelBarChanged, HealGridGuiTabHUD.ChangeShowFriendlyTargetDebuffFlags, HealGridGuiTabHUD.ChangeShowHostileTargetDebuffFlags, HealGridGuiTabHUD.ChangeShowPlayerDebuffFlags, HealGridGuiTabHUD.ChangeSkin, HealGridGuiTabHUD.ShowActionPointsBarChanged, HealGridGuiTabHUD.ShowCareerLineIconChanged, HealGridGuiTabHUD.ShowHealthPointsBarChanged, HealGridGuiTabHUD.ShowMoraleLevelBarChanged, HealGridGuiTabHUDBuffs.ChangeFriendlyTargetBuffsFilter, HealGridGuiTabHUDBuffs.ChangeFriendlyTargetDebuffsFilter, HealGridGuiTabHUDBuffs.ChangeHostileTargetBuffsFilter, HealGridGuiTabHUDBuffs.ChangeHostileTargetDebuffsFilter, HealGridGuiTabHUDBuffs.ChangePlayerBuffsFilter, HealGridGuiTabHUDBuffs.ChangePlayerDebuffsFilter, HealGridGuiTabMouseClick.ActionChanged, HealGridGuiTabScenariogroup.ChangeShowDebuffFlags, HealGridGuiTabScenariogroup.ChangeSkin, HealGridGuiTabScenariogroup.GridGroupingChanged, HealGridGuiTabScenariogroup.GridGrowthChanged, HealGridGuiTabScenariogroup.GridOrientationChanged, HealGridGuiTabScenariogroup.ShowActionPointsBarChanged, HealGridGuiTabScenariogroup.ShowCareerLineIconChanged, HealGridGuiTabScenariogroup.ShowHealthPointsBarChanged, HealGridGuiTabScenariogroup.ShowMoraleLevelBarChanged, HopperConfig_Profiles_OnActiveProfileChanged, HopperConfig_Profiles_OnCopyProfileChanged, ItemRack.Sets.Select, ItemRack.Tactics.OnSetMenuSelectionChanged, KeysetMonsterPlay.Setup.Profile.OnDefaultProfileChanged, KeysetMonsterPlay.Setup.Profile.OnMonsterProfileChanged, LPET.AttackRangeCheckComboOnSelChanged, LPET.AutoAttackComboOnSelChanged, LPET.AutoDefendComboOnSelChanged, LPET.AutoSwitchComboOnSelChanged, LPET.BiteComboOnSelChanged, LPET.BiteHealthComboOnSelChanged, LPET.BitePriorityComboOnSelChanged, LPET.ClawSweepComboOnSelChanged, LPET.ClawSweepHealthComboOnSelChanged, LPET.ClawSweepPriorityComboOnSelChanged, LPET.ComboOnSelChanged, LPET.CoruscatingEnergyComboOnSelChanged, LPET.CoruscatingEnergyHealthComboOnSelChanged, LPET.CoruscatingEnergyPriorityComboOnSelChanged, LPET.DaemonicConsumptionComboOnSelChanged, LPET.DaemonicConsumptionHealthComboOnSelChanged, LPET.DaemonicConsumptionPriorityComboOnSelChanged, LPET.DaemonicFireComboOnSelChanged, LPET.DaemonicFireHealthComboOnSelChanged, LPET.DaemonicFirePriorityComboOnSelChanged, LPET.DeathFromAboveComboOnSelChanged, LPET.DeathFromAboveHealthComboOnSelChanged, LPET.DeathFromAbovePriorityComboOnSelChanged, LPET.DefendRangeCheckComboOnSelChanged, LPET.FangAndClawComboOnSelChanged, LPET.FangAndClawHealthComboOnSelChanged, LPET.FangAndClawPriorityComboOnSelChanged, LPET.FlameOfTzeentchComboOnSelChanged, LPET.FlameOfTzeentchHealthComboOnSelChanged, LPET.FlameOfTzeentchPriorityComboOnSelChanged, LPET.FlamesOfChangeComboOnSelChanged, LPET.FlamesOfChangeHealthComboOnSelChanged, LPET.FlamesOfChangePriorityComboOnSelChanged, LPET.FlamethrowerComboOnSelChanged, LPET.FlamethrowerHealthComboOnSelChanged, LPET.FlamethrowerPriorityComboOnSelChanged, LPET.GoopShootinComboOnSelChanged, LPET.GoopShootinHealthComboOnSelChanged, LPET.GoopShootinPriorityComboOnSelChanged, LPET.GoreComboOnSelChanged, LPET.GoreHealthComboOnSelChanged, LPET.GorePriorityComboOnSelChanged, LPET.GutRipperComboOnSelChanged, LPET.GutRipperHealthComboOnSelChanged, LPET.GutRipperPriorityComboOnSelChanged, LPET.HeadButtComboOnSelChanged, LPET.HeadButtHealthComboOnSelChanged, LPET.HeadButtPriorityComboOnSelChanged, LPET.HighExplosiveGrenadeComboOnSelChanged, LPET.HighExplosiveGrenadeHealthComboOnSelChanged, LPET.HighExplosiveGrenadePriorityComboOnSelChanged, LPET.LegTearComboOnSelChanged, LPET.LegTearHealthComboOnSelChanged, LPET.LegTearPriorityComboOnSelChanged, LPET.LionsRoarComboOnSelChanged, LPET.LionsRoarHealthComboOnSelChanged, LPET.LionsRoarPriorityComboOnSelChanged, LPET.MachineGunComboOnSelChanged, LPET.MachineGunHealthComboOnSelChanged, LPET.MachineGunPriorityComboOnSelChanged, LPET.MaulComboOnSelChanged, LPET.MaulHealthComboOnSelChanged, LPET.MaulPriorityComboOnSelChanged, LPET.PenetratingRoundComboOnSelChanged, LPET.PenetratingRoundHealthComboOnSelChanged, LPET.PenetratingRoundPriorityComboOnSelChanged, LPET.PetAttackComboOnSelChanged, LPET.PetFollowComboOnSelChanged, LPET.PoisonedSpineComboOnSelChanged, LPET.PoisonedSpineHealthComboOnSelChanged, LPET.PoisonedSpinePriorityComboOnSelChanged, LPET.ProfilesComboOnSelChanged, LPET.QuickSettingsComboOnSelChanged, LPET.SelfFollowComboOnSelChanged, LPET.ShockGrenadeComboOnSelChanged, LPET.ShockGrenadeHealthComboOnSelChanged, LPET.ShockGrenadePriorityComboOnSelChanged, LPET.ShredComboOnSelChanged, LPET.ShredHealthComboOnSelChanged, LPET.ShredPriorityComboOnSelChanged, LPET.SpineFlingComboOnSelChanged, LPET.SpineFlingHealthComboOnSelChanged, LPET.SpineFlingPriorityComboOnSelChanged, LPET.SporeCloudComboOnSelChanged, LPET.SporeCloudHealthComboOnSelChanged, LPET.SporeCloudPriorityComboOnSelChanged, LPET.SquigSquealComboOnSelChanged, LPET.SquigSquealHealthComboOnSelChanged, LPET.SquigSquealPriorityComboOnSelChanged, LPET.SteamVentComboOnSelChanged, LPET.SteamVentHealthComboOnSelChanged, LPET.SteamVentPriorityComboOnSelChanged, LPET.SwitchRangeCheckComboOnSelChanged, LPET.TerrifyingRoarComboOnSelChanged, LPET.TerrifyingRoarHealthComboOnSelChanged, LPET.TerrifyingRoarPriorityComboOnSelChanged, LPET.WarpingEnergyComboOnSelChanged, LPET.WarpingEnergyHealthComboOnSelChanged, LPET.WarpingEnergyPriorityComboOnSelChanged, LibAddonButton.Manager.Advanced.OnItemChanged, LibAddonButton.Manager.CustomItem.OnTypeChanged, MBuffSetup.SmartBuff.OnClassTypeChanged, Map.OnSelChanged, MapMonster.Editor.OnPinTypeChange, MapMonster.Editor.OnSubTypeChange, MapMonster.Editor.OnZoneNameChange, MapMonster.PinTypeEditor.OnSubTypeChange, MiracleGrow2.LayoutPlotArrChanged, MiracleGrow2.LayoutShowChanged, MoraleSet.OnSetMenuSelectionChanged, MotionConfig_Profiles_OnActiveProfileChanged, MotionConfig_Profiles_OnCopyProfileChanged, NBSBCore.OnAddonSelected, NBSBParam.OnParamChange, Obsidian.Setup.Castbar.OnElementChanged, Obsidian.Setup.Castbar.OnFillUpdatePriorityChanged, Obsidian.Setup.Castbar.OnGlobalCooldownPositionChanged, Obsidian.Setup.Castbar.OnIconPositionChanged, Obsidian.Setup.Castbar.OnNameAlignmentChanged, Obsidian.Setup.Castbar.OnTimerAlignmentChanged, Obsidian.Setup.EffectTracker.OnBarElementChanged, Obsidian.Setup.EffectTracker.OnElementChanged, Obsidian.Setup.EffectTracker.OnFillUpdatePriorityChanged, Obsidian.Setup.EffectTracker.OnNameAlignmentChanged, Obsidian.Setup.EffectTracker.OnTimerAlignmentChanged, Obsidian.Setup.EffectTracker.OnTrackerElementChanged, Obsidian.Setup.EffectTracker.OnTrackerFillColorChanged, Obsidian.Setup.EffectTracker.OnTrackerIconPositionChanged, Obsidian.Setup.EffectTracker.OnTrackerPositionChanged, PartyAdWindow.OnSelChangedPurposePreset, PartyAdWindow.OnSelChangedSelfRole, PotionBarSettings.ActivatorComboSelChanged, PotionBarSettings.BuildComboSelChanged, PotionBarSettings.ComboMethod, PotionBarSettings.InfoTextBRComboSelChanged, PotionBarSettings.InfoTextTRComboSelChanged, PotionBarSettings.QuickActionsSelChanged, PureConfig_Profiles_OnActiveProfileChanged, PureConfig_Profiles_OnCopyProfileChanged, RVMOD_Manager.OnSelChangedSortBy, RVMOD_SquaredDistances.OnComboBoxAnchorPointsChange, RVMOD_SquaredDistances.OnComboBoxFontsChange, RVMOD_SquaredDistances.OnComboBoxLayerChange, RVMOD_Targets.OnComboBoxConditionClick, RVMOD_Targets.OnComboBoxTemplateChange, RaidMeter.MenuSelect, Res.OptionsIconBackgroundSelected, Res.OptionsIconForegroundSelected, Res.OptionsIconHighlightSelected, ScenarioStats.CharacterSelectionChanged, ScenarioStats.ScenarioSelectionChanged, ShiniesBrowseUI.OnSelChanged_Criteria_ItemSlotCombo, ShiniesBrowseUI.OnSelChanged_Criteria_ItemTypeCombo, ShiniesBrowseUI.OnSelChanged_Criteria_ModifierCombo, SocialWindowTabSearch.OnFilterSelChanged, TOLSettingsUI.OnEventSelChanged, TOLSettingsUI.OnPhraseSelChanged, TOLSettingsUI.OnSkillSelChanged, TastyButtonsOptions.OnComboButtonSelectionChanged, TastyButtonsOptions.OnComboFontChanged, TastyButtonsOptions.OnComboStateChanged, TexturedButtons.Setup.Actionbar.OnBarChanged, TexturedButtons.Setup.Actionbar.OnSelectorChanged, TexturedButtons.Setup.AdvancedTextures.OnCustomChanged, TexturedButtons.Setup.AdvancedTextures.OnCustomTextureChanged, TexturedButtons.Setup.AdvancedTextures.OnCustomTextureSliceChanged, TexturedButtons.Setup.AdvancedTextures.OnPresetChanged, TexturedButtons.Setup.AdvancedTextures.OnSlotTypeChanged, TexturedButtons.Setup.Misc.OnActionButtonPickUpModifierChanged, TexturedButtons.Setup.Misc.OnCustomGlowChanged, TexturedButtons.Setup.Textures.OnCustomChanged, TexturedButtons.Setup.Textures.OnCustomTextureChanged, TexturedButtons.Setup.Textures.OnCustomTextureSliceChanged, TexturedButtons.Setup.Textures.OnPresetChanged, TexturedButtons.Setup.Tint.OnTintTypeChanged, TidyChat.Options.UpdateGroupTabs, TidyRoll.CustomAutoRoll.OnChoiceChange, TokenMachine.UpdateOption, TurretRange.Setup.Display.OnCircleModeChanged, TurretRange.Setup.Display.OnCircleTypeChanged, TurretRange.Setup.Display.OnDistanceLayoutChanged, TurretRange.Setup.Display.OnDistanceTypeChanged, TurretRange.Setup.Display.OnElementChanged, TurretRange.Setup.Display.OnGraphicTypeChanged, TwisterSet.OnSetMenuSelectionChanged, UiModWindow.OnCategoryComboSelChanged, Vectors.Settings.ComboBoxChanged, WARCommanderConfig.OnRVRAlertsChange, WARCommanderConfig.OnRVRCommandsChange, WARCommanderConfig.OnScenAlertsChange, WARCommanderConfig.OnScenCommandsChange, WCDBConfig_Filters_OnActiveFilterChanged, WCDBConfig_Profiles_OnActiveProfileChanged, WCDBConfig_Profiles_OnCopyProfileChanged, WCDPConfig_Filters_OnActiveFilterChanged, WCDPConfig_Profiles_OnActiveProfileChanged, WCDPConfig_Profiles_OnCopyProfileChanged, WSCT.ComboOnSelChanged, WSCT.EventOnSelChanged, WSCT.FrameComboOnSelChanged, WSCT.FrameOnSelChanged, WarBoard_FPSOptions.OnCmbChange, Warbuilder.ChangeSelect, Warbuilder.LinkLevel, Warbuilder.MenuSelect, ZonePOP.OnFilterSelChanged, nLootLinkGUI.onCareerSelected, nLootLinkGUI.onCategorySelected, nLootLinkGUI.onRaritySelected, nLootLinkGUI.onSlotSelected, nLootLinkGUI.onStatSelected, nLootLinkGUI.onTypeSelected, nLootLinkOptions.onRaritySelected, wbLeadHelperConfigTab.OnChanged | `index` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | LPET.OnMouseOver, , Enemy.ConfigurationWindow_ShowTooltip, AutoBandWindowConfig.OnSettingMouseOver, DevPadWindow.OnMouseOverCode, EA_Window_OpenPartyManage.OnMouseoverLeaderMark, EA_Window_OpenPartyManage.OnMouseoverLootModeCombo, EA_Window_OpenPartyManage.OnMouseoverLootThresholdCombo, EZCraftX.OnMouseOver_Dropdown, MoraleSet.OnMouseOverSetMenu, PotionBarSettings.OnMouseoverActivator, PotionBarSettings.OnMouseoverBuild, PotionBarSettings.OnMouseoverInfoTextBR, PotionBarSettings.OnMouseoverInfoTextTR, PotionBarSettings.OnMouseoverMethod, PotionBarSettings.OnMouseoverQuickActions, TwisterSet.OnMouseover, WSCT.FrameComboOnMouseOver, WSCT.OnMouseOver, WbLeadHelperMessage.OnMouseOverSubmenuComboBox, WbLeadHelperMessage.OnMouseOverTypeComboBox, wbLeadHelperConfigTab.OnMouseOverMessageTextColorLabel | `` |  |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | MoraleSet.OnLButtonUpSetMenu, TOLSettingsUI.OnEventSelChanged, TwisterSet.OnLClick, alertMod.OnSliderChanged | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | EA_Window_OpenPartyLootRollOptions.OnRClickComboBox, MoraleSet.OnRButtonUpSetMenu, TOLSettingsUI.OnEventSelChanged, TwisterSet.OnRClick | `flags, x, y` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | AutoBandWindowConfig.OnSettingMouseOverEnd | `` |  |

### Per-Event Lua API Calls

**OnSelChanged** handlers call: `ButtonGetPressedFlag`, `ButtonGetText`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetText`, `ButtonSetTextColor`, `ComboBoxAddMenuItem`, `ComboBoxClearMenuItems`, `ComboBoxGetDisabledFlag`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetDisabledFlag`, `ComboBoxSetSelectedMenuItem`, `CreateWindowFromTemplate`, `DestroyWindow`, `DoesWindowExist`, `DynamicImageSetTexture`, `DynamicImageSetTextureSlice`, `LabelGetTextDimensions`, `LabelSetFont`, `LabelSetText`, `LabelSetTextColor`, `ListBoxSetDisplayOrder`, `ScrollWindowUpdateScrollRect`, `SliderBarGetCurrentPosition`, `SystemData.ActiveWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxSetText`, `TextEditBoxSetTextColor`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetShowing`, `WindowRegisterEventHandler`, `WindowSetDimensions`, `WindowSetHandleInput`, `WindowSetShowing`

**OnMouseOver** handlers call: `ComboBoxGetSelectedText`, `WindowGetParent`

**OnLButtonUp** handlers call: `ComboBoxSetSelectedMenuItem`, `WindowSetShowing`

**OnRButtonUp** handlers call: `ButtonGetText`, `ButtonSetText`, `ButtonSetTextColor`, `ComboBoxGetSelectedMenuItem`, `ComboBoxSetDisabledFlag`, `ComboBoxSetSelectedMenuItem`, `WindowGetId`, `WindowGetParent`, `WindowSetShowing`

## Common Inherits

- APA_ComboBox
- Aura_ComboBox_DefaultResizableLarge
- Aura_ComboBox_DefaultResizableTiny
- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizableLarge
- EA_ComboBox_DefaultResizableMedium
- EA_ComboBox_DefaultResizableSmall
- EA_ComboBox_DefaultResizableTiny
- EA_ComboBox_DefaultResizable_Fixed
- EA_ComboBox_DefaultResizable_Fixed_Small
- IraConfigCombo
- Tiny_ComboBox

## Common Parent Elements

- [Windows](element_Windows.md) — 689× (HIGH)
- [Window](element_Window.md) — 4× (MEDIUM)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 643× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 534× (HIGH)
- [Size](element_Size.md) — 186× (HIGH)
- [MenuButtonOffset](element_MenuButtonOffset.md) — 33× (HIGH)
- [OverlayOffset](element_OverlayOffset.md) — 4× (MEDIUM)
- [TextOffset](element_TextOffset.md) — 1× (LOW)

## Common Template Bases

- APA_ComboBox
- APA_ComboBoxWide
- Aura_ComboBox_DefaultResizable
- Aura_ComboBox_DefaultResizableLarge
- Aura_ComboBox_DefaultResizableTiny
- CBWithCenterAlignedText
- CaVES_ComboBox_DefaultResizable
- Crusher_ComboBox_DefaultResizable
- DaemonAssist_ComboBoxWide
- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizable3
- EA_ComboBox_DefaultResizableLarge
- EA_ComboBox_DefaultResizableMedium
- EA_ComboBox_DefaultResizableMediumLarge
- EA_ComboBox_DefaultResizableSmall
- EA_ComboBox_DefaultResizableTiny
- EA_ComboBox_DefaultResizable_Fixed
- EA_ComboBox_DefaultResizable_Fixed_Small
- EA_Template_AutoRollComboBox
- EA_Template_OpenParty_ComboBox
- Hopper_ComboBox_DefaultResizable
- IraConfigCombo
- ItemRackTacticsSetComboBox
- MapMonsterPinTypeEditor_ComboBox_DefaultResizable
- MoraleSetComboBox
- Motion_ComboBox_DefaultResizable
- PartyAd_PurposePresetComboBoxTemplate
- PartyAd_SelfRoleComboBoxTemplate
- Pure_ComboBox_DefaultResizable
- Shinies_ComboBox_DefaultResizableLarge
- Small_ComboBox
- Tiny_ComboBox
- TooltipComboBox
- TwisterSetComboBox
- WCDB_ComboBox_DefaultResizable
- WCDP_ComboBox_DefaultResizable
- Zonepop_ComboBox_DefaultResizable


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 95% | APA_ComboBox, APA_ComboBoxWide, EA_ComboBox_DefaultResizable, Aura_ComboBox_DefaultResizableLarge, ... |
| `layer` | optional | 40% | secondary, popup, primary, default, ... |
| `maxvisibleitems` | optional | 9% | 10, 13, 5, 8, ... |
| `menubackground` | optional | 4% | EA_Window_ComboBoxMenuBackground, EA_Window_DefaultFrame, TacticsComboBoxBackground, MoraleComboBoxBackground, ... |
| `menuitembutton` | optional | 4% | EA_Button_DefaultMenuButtonSmall, EA_Button_DefaultMenuButton, Aura_Button_DefaultMenuButton, Aura_Button_DefaultMenuButtonLarge, ... |
| `selectedbutton` | optional | 4% | APA_ComboBoxButton, APA_ComboBoxButtonWide, Aura_ComboBox_DefaultResizableComboBoxSelected, Aura_ComboBox_DefaultResizableComboBoxSelectedLarge, ... |
| `scrollbar` | optional | 3% | EA_ScrollBar_DefaultVerticalChain, EA_ScrollBar_ChatVertical |
| `id` | optional | 2% | 1, 2, 3, 4, ... |
| `show` | optional | 1% | false |
| `warnOnTextCropped` | optional | 1% | false |
| `autoresize` | optional | 1% | true |
| `handleinput` | optional | 0% | true |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 643 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 534 times as an unnamed child.

### [Size](element_Size.md)

Observed 186 times as an unnamed child.

### [MenuButtonOffset](element_MenuButtonOffset.md)

Observed 33 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5, 12, 0 |
| `y` | **required** | 5, 0 |
### [OverlayOffset](element_OverlayOffset.md)

Observed 4 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 93, 193, 148, 223 |
| `y` | **required** | 0, 1, 7, 9 |
### [TextOffset](element_TextOffset.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5, 6, 4, 0 |
| `y` | **required** | 5, 12, 2, 3 |
## Recursive Hierarchy

- Root: [ComboBox](element_ComboBox.md)
- [Anchors](element_Anchors.md) (structural, 643×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
      - (cycle)
- [EventHandlers](element_EventHandlers.md) (structural, 534×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
- [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 33×, HIGH)
- [OverlayOffset](element_OverlayOffset.md) (structural, 4×, MEDIUM)
- [Size](element_Size.md) (structural, 186×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
- [TextOffset](element_TextOffset.md) (structural, 1×, LOW)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `ComboBoxGetSelectedMenuItem` | ui | 250 | OnRButtonUp, OnSelChanged |
| `WindowGetParent` | ui | 232 | OnMouseOver, OnRButtonUp, OnSelChanged |
| `WindowSetShowing` | ui | 74 | OnLButtonUp, OnRButtonUp, OnSelChanged |
| `LabelSetText` | ui | 62 | OnSelChanged |
| `ComboBoxSetSelectedMenuItem` | ui | 42 | OnLButtonUp, OnRButtonUp, OnSelChanged |
| `TextEditBoxGetText` | ui | 22 | OnSelChanged |
| `ButtonSetDisabledFlag` | ui | 13 | OnSelChanged |
| `ButtonGetPressedFlag` | ui | 10 | OnSelChanged |
| `ButtonSetPressedFlag` | ui | 9 | OnSelChanged |
| `WindowGetId` | ui | 9 | OnRButtonUp, OnSelChanged |
| `ComboBoxGetSelectedText` | ui | 8 | OnMouseOver, OnSelChanged |
| `ComboBoxSetDisabledFlag` | ui | 8 | OnRButtonUp, OnSelChanged |
| `SliderBarGetCurrentPosition` | ui | 8 | OnSelChanged |
| `ButtonSetText` | ui | 7 | OnRButtonUp, OnSelChanged |
| `WindowSetHandleInput` | ui | 6 | OnSelChanged |
| `DynamicImageSetTexture` | ui | 5 | OnSelChanged |
| `LabelSetTextColor` | ui | 4 | OnSelChanged |
| `WindowAddAnchor` | ui | 4 | OnSelChanged |
| `ComboBoxAddMenuItem` | ui | 3 | OnSelChanged |
| `ComboBoxClearMenuItems` | ui | 3 | OnSelChanged |
| `ComboBoxGetDisabledFlag` | ui | 3 | OnSelChanged |
| `DoesWindowExist` | ui | 3 | OnSelChanged |
| `LabelSetFont` | ui | 3 | OnSelChanged |
| `WindowClearAnchors` | ui | 3 | OnSelChanged |
| `WindowRegisterEventHandler` | event | 3 | OnSelChanged |
| `ButtonGetText` | ui | 2 | OnRButtonUp, OnSelChanged |
| `ButtonSetTextColor` | ui | 2 | OnRButtonUp, OnSelChanged |
| `DynamicImageSetTextureSlice` | ui | 2 | OnSelChanged |
| `ListBoxSetDisplayOrder` | ui | 2 | OnSelChanged |
| `TextEditBoxSetText` | ui | 2 | OnSelChanged |
| `TextEditBoxSetTextColor` | ui | 2 | OnSelChanged |
| `CreateWindowFromTemplate` | ui | 1 | OnSelChanged |
| `DestroyWindow` | ui | 1 | OnSelChanged |
| `LabelGetTextDimensions` | ui | 1 | OnSelChanged |
| `ScrollWindowUpdateScrollRect` | ui | 1 | OnSelChanged |
| `SystemData.ActiveWindow.name:match` | data | 1 | OnSelChanged |
| `WindowGetDimensions` | ui | 1 | OnSelChanged |
| `WindowGetShowing` | ui | 1 | OnSelChanged |
| `WindowSetDimensions` | ui | 1 | OnSelChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseOver

Confidence: HIGH

### OnMouseOverEnd

Confidence: LOW

### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnSelChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `index` | number | selected_index |
## Lua Functions Manipulating This Type

- DaemonAssist.PopulateBindingCombos
- Enemy.CombatLogUI_StatsWindow_Open
- Enemy.IntercomUI_IntercomJoinDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- ChattyCathy.UpdateEntrySetup
- Enemy.CombatLogUI_StatsWindow_UpdateSessionsList
- Enemy.GroupsUI_EffectFilterDialog_Open
- GuildWardenWin.WinSetup
- TalismanGenie.CurioCheckToggle
- TwisterSet.OnSetMenuSelectionChanged
- CCTV.Initialize
- ChattyCathy.CC_ModelWindow
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- MapPin.EditMarker
- TalismanGenie.Reset
- TwisterSet.OnRClick
- ZonePOP.UpdateZoneList
- CallingSetup.Initialize
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- MapPin.local.EditMarker
- TalismanGenie.GoldDustCheckToggle
- UiModWindow.UpdateCategoryComboBox
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- TwisterSet.Initialize
- Enemy.IntercomUI_ChooseChannelDialog_Open
- TalismanGenie.EssenceCheckToggle
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- BustedGUI.Initialize
- BustedGUI.NewErrorRecorded
- Enemy.UI_ConfigDialog_Open
- Enemy.IntercomUI_IntercomJoinDialog_AddGroup
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- ScenarioStats.CreateDropdownList
- AdvancedRenownTrainer.GeneratePresetByLinkData
- BankArkel.SetupCombos
- Enemy.UnitFramesUI_ConfigDialog_Import
- MapPin.RButtonUp
- alertMod.SetLabels


## Binding Resolution

- Total handler declarations: 676
- Resolved to Lua functions: 675 (99%)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- AutoBand
- BankArkel
- BuffHead
- Busted
- CCTV
- CDown
- CaVES
- Calling
- CastSequence
- ChattyCathy
- Cheeseboard
- Crafting Info Tooltip
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- DeepSleep
- DetauntHelper
- DuffTimer
- Dye Preview
- EA_OpenPartyWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- Enemy
- EveryBodyGuard
- FozAuction
- GDes
- Ges
- GroupRange
- GuildWarden
- HealGrid
- Hopper
- ItemRack
- Keyset
- KeysetMonsterPlay
- Killer
- LibAddonButton
- LoyalPet
- Map
- MapMonster
- MapPin
- MarkBuff
- MegaphonePlusPlus
- Miracle Grow Remix
- MoraleSet
- Motion
- NerfedButtons
- Obsidian
- OverheadFonts
- PartyAd
- Pocket Palette
- PotionBar
- Pure
- RVMOD_Manager
- RVMOD_SquaredDistances
- RVMOD_Targets
- RaidMeter
- Res
- SOR
- ScenarioStats
- Shinies
- SocialWindow 2.0
- Statdoll Remix
- TalismanGenie
- TastyButtons
- TaxPayer
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TokenMachine
- TurretRange
- TwisterSet
- Vectors
- WARCommander
- WBStutterLess
- WSCT
- WarBoard
- WarBoard_FPS
- Warbuilder
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- ZonePOP
- alertMod
- nLootLink
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAComboAttackBind -> ComboBox APAComboAttackBind
- AdvancedPetAssist: APAComboAutoReattack -> ComboBox APAComboAutoReattack
- AdvancedPetAssist: APAComboAutoReattackDelay -> ComboBox APAComboAutoReattackDelay
- AdvancedPetAssist: APAComboCastDelay -> ComboBox APAComboCastDelay
- AdvancedPetAssist: APAComboCastOnAcquire -> ComboBox APAComboCastOnAcquire
- AdvancedPetAssist: APAComboCombatExitDelay -> ComboBox APAComboCombatExitDelay

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EA_ComboBox_DefaultResizable](../../globals/constants/constant_EA_ComboBox_DefaultResizable.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizableLarge](../../globals/constants/constant_EA_ComboBox_DefaultResizableLarge.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizableMedium](../../globals/constants/constant_EA_ComboBox_DefaultResizableMedium.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizableSmall](../../globals/constants/constant_EA_ComboBox_DefaultResizableSmall.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizableTiny](../../globals/constants/constant_EA_ComboBox_DefaultResizableTiny.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizable_Fixed](../../globals/constants/constant_EA_ComboBox_DefaultResizable_Fixed.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [MenuButtonOffset](element_MenuButtonOffset.md) (HIGH 100/100) - XML Element Type
- [OverlayOffset](element_OverlayOffset.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TextOffset](element_TextOffset.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_ComboBox_DefaultResizable3](../../globals/constants/constant_EA_ComboBox_DefaultResizable3.md) (HIGH 90/100) - Constant
- [EA_ComboBox_DefaultResizableMediumLarge](../../globals/constants/constant_EA_ComboBox_DefaultResizableMediumLarge.md) (HIGH 90/100) - Constant
- [EA_ComboBox_DefaultResizable_Fixed_Small](../../globals/constants/constant_EA_ComboBox_DefaultResizable_Fixed_Small.md) (HIGH 90/100) - Constant
- [EA_Template_AutoRollComboBox](../../globals/constants/constant_EA_Template_AutoRollComboBox.md) (HIGH 90/100) - Constant
- [EA_Template_OpenParty_ComboBox](../../globals/constants/constant_EA_Template_OpenParty_ComboBox.md) (HIGH 90/100) - Constant

## Used With

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnSelChanged](../handlers/handler_OnSelChanged.md) (HIGH 88/100) - XML Event
