# OnLButtonUp

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, BagOMatic, BankArkel |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/AnywhereTrainerAdditions/AnywhereTrainerAdditions.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0` |
| Namespaces detected | OnLButtonUp |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnLButtonUp, AdvancedPetAssist: APAInstantOnlyHUD.OnLButtonUp, AdvancedPetAssist: APAKitingHUD.OnLButtonUp, AdvancedPetAssist: APAOptionsClose.OnLButtonUp, AdvancedPetAssist: APAOptionsTabsAutoRecall.OnLButtonUp, AdvancedPetAssist: APAOptionsTabsControls.OnLButtonUp |
| XML usage count | 633 |
| XML attribute usage count | 633 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 1 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed as an engine-supplied UI event hook used by 29 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Aura
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- LibGroup
- MiracleGrowLight
- MoraleCircle
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- followTheLeader

## Registrars And Handlers

- APAGui.Hide
- APAGui.OnFollowTargetHUDClicked
- APAGui.OnInstantOnlyHUDClicked
- APAGui.OnKitingHUDClicked
- APAGui.OnTabButtonUp
- AdvancedRenownTraining.DeletePreset
- AdvancedRenownTraining.ExportCancelButtonPressed
- AdvancedRenownTraining.ExportToLink
- AdvancedRenownTraining.ExportToWardrobe
- AdvancedRenownTraining.Hide
- AdvancedRenownTraining.ImportCancelButtonPressed
- AdvancedRenownTraining.ImportNameInputCancelButtonPressed
- AdvancedRenownTraining.ImportNameInputOkButtonPressed
- AdvancedRenownTraining.ImportOkButtonPressed
- AdvancedRenownTraining.LinkWindowClose
- AdvancedRenownTraining.OnExportButtonPressed
- AdvancedRenownTraining.OnLButtonUpTab
- AdvancedRenownTraining.OnLoadPressed
- AdvancedRenownTraining.PurchaseAdvances
- AdvancedRenownTraining.Respecialize
- AdvancedRenownTraining.SavePreset
- AdvancedRenownTraining.ShowWardrobeImport
- AdvancedRenownTraining.TogglePresets
- AdvancedRenownTraining.ToggleSaveSettings
- AggroMeter.Close
- AggroMeter.OnTabLBU
- AggroMeter.SelectChar
- AnywhereTrainer.OnLButtonUp
- AuraConfig.OnAbilityIconLButtonUp
- AuraConfig.OnActivationAlertTextTestButton
- AuraConfig.OnCircledImageChanged
- AuraConfig.OnClose
- AuraConfig.OnConfigTabSelected
- AuraConfig.OnDeactivationAlertTextTestButton
- AuraConfig.OnDisplayTimerChanged
- AuraConfig.OnLButtonUpTextureTintColor
- AuraConfig.OnLButtonUpTimerTintColor
- AuraConfig.OnMirrorImageChanged
- AuraConfig.OnTextureChangeButton
- AuraSettings.ChangeSorting
- AuraSettings.ConfigChange_EnableAddon
- AuraSettings.ConfigChange_EnableDebugging
- AuraSettings.OnAddAura
- AuraSettings.OnClose
- AuraSettings.OnLButtonUpAuraList
- AuraSettings.OnShare
- AuraShares.ChangeSorting
- AuraShares.OnClose
- AuraShares.OnCloseAuraSharesImportExportWindow
- AuraShares.OnFilterCharactersAurasToggle
- AuraShares.OnFilterSameNameToggle
- AuraShares.OnImportAura
- AuraShares.OnImportExportOkButton
- AuraShares.OnLButtonUpAuraList
- AuraTexture.OnClose
- AuraTexture.OnIconLButtonUp
- AuraTexture.OnRaceTabSelected
- BagOMatic.wnd_on_lbutton_up
- BankArkel.PackClose
- BankArkel.PackTab
- BuffHead.Setup.AdvancedCompression.OnCloseLUp
- BuffHead.Setup.AdvancedCompression.OnNewLClick
- BuffHead.Setup.AdvancedCompression.OnRowLUp
- BuffHead.Setup.AdvancedCompressionItem.OnAddEffectLClick
- BuffHead.Setup.AdvancedCompressionItem.OnApplyLClick
- BuffHead.Setup.AdvancedCompressionItem.OnCloseLUp
- BuffHead.Setup.AdvancedCompressionItem.OnCreateLClick
- BuffHead.Setup.AdvancedCompressionItem.OnRowLUp
- BuffHead.Setup.AdvancedCompressionItemEffect.OnApplyLClick
- BuffHead.Setup.AdvancedCompressionItemEffect.OnCloseLUp
- BuffHead.Setup.AdvancedCompressionItemEffect.OnCreateLClick
- BuffHead.Setup.AdvancedContainers.OnCloseLUp
- BuffHead.Setup.AdvancedContainers.OnNewLClick
- BuffHead.Setup.AdvancedContainers.OnRowLUp
- BuffHead.Setup.AdvancedContainersItem.OnApplyLClick
- BuffHead.Setup.AdvancedContainersItem.OnCloseLUp
- BuffHead.Setup.AdvancedContainersItem.OnCreateLClick
- BuffHead.Setup.AdvancedContainersItem.OnRowLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnCloseLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysIgnoreLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysShowEnableLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysShowLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsPermanentLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputEnableRemovableLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputShowTooltipsLUp
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleLClick
- BuffHead.Setup.Container.OnCloseLUp
- BuffHead.Setup.Container.OnContainerClick
- BuffHead.Setup.Display.OnCloseLUp
- BuffHead.Setup.Display.OnEnableSortLUp
- BuffHead.Setup.EffectCache.OnCloseLUp
- BuffHead.Setup.EffectCache.OnEnableCachingLUp
- BuffHead.Setup.EffectCache.OnRefreshLClick
- BuffHead.Setup.EffectCache.OnResetLClick
- BuffHead.Setup.EffectCache.OnRowLUp
- BuffHead.Setup.EffectCache.OnSortButtonLUp
- BuffHead.Setup.EffectCacheTable.OnCloseLUp
- BuffHead.Setup.Filter.OnAddLClick
- BuffHead.Setup.Filter.OnCloseLUp
- BuffHead.Setup.Filter.OnEnableFilter
- BuffHead.Setup.Filter.OnRowLUp
- BuffHead.Setup.General.OnAlwaysIgnoreLClick
- BuffHead.Setup.General.OnAlwaysShowLClick
- BuffHead.Setup.General.OnCloseLUp
- BuffHead.Setup.Layout.Manager.OnCloseLUp
- BuffHead.Setup.Layout.Manager.OnExportLUp
- BuffHead.Setup.Layout.Manager.OnImportLClick
- BuffHead.Setup.Layout.Manager.OnImportLUp
- BuffHead.Setup.Layout.Manager.OnLayoutsActivateLClick
- BuffHead.Setup.Layout.Manager.OnLayoutsDeleteLClick
- BuffHead.Setup.Layout.Manager.OnLayoutsLUp
- BuffHead.Setup.Layout.Manager.OnLayoutsSaveCurrentLayoutLClick
- BuffHead.Setup.Layout.OnApplyLClick
- BuffHead.Setup.Layout.OnCloseLUp
- BuffHead.Setup.Layout.OnControlFrameLButtonUp
- BuffHead.Setup.Layout.OnLayerLClick
- BuffHead.Setup.Layout.OnLayersChanged
- BuffHead.Setup.Layout.OnManagerLClick
- BuffHead.Setup.Layout.Properties.OnCloseLUp
- BuffHead.Setup.Layout.Properties.OnCoreSizeAutoSizeLClick
- BuffHead.Setup.Layout.Properties.OnFontColorLUp
- BuffHead.Setup.Layout.Properties.OnFontFontLUp
- BuffHead.Setup.Layout.Properties.OnIconBorderColorLUp
- BuffHead.Setup.Layout.Properties.OnPropertyTitleLClick
- BuffHead.Setup.Layout.Properties.OnStatusBarBackgroundColorLUp
- BuffHead.Setup.Layout.Properties.OnStatusBarEnableLUp
- BuffHead.Setup.Layout.Properties.OnStatusBarForegroundColorLUp
- BuffHead.Setup.Layout.Properties.OnStatusBarForegroundStretchLUp
- BuffHead.Setup.Layout.Properties.OnStatusBarForegroundTextureLUp
- BuffHead.Setup.Layout.Properties.OnStatusBarOrientationReverseLUp
- BuffHead.Setup.OnCloseLUp
- BuffHead.Setup.OnRowLUp
- BuffHead.Setup.Performance.OnCloseLUp
- BuffHead.Setup.Performance.OnEnableFadingLUp
- BuffHead.Setup.Performance.OnEnablePriorityLUp
- BuffHead.Setup.Performance.OnEnableSyncLUp
- BuffHead.Setup.PriorityEffects.OnCloseLUp
- BuffHead.Setup.PriorityEffects.OnNewLClick
- BuffHead.Setup.PriorityEffects.OnRowLUp
- BuffHead.Setup.PriorityEffects.OnSortFirstLUp
- BuffHead.Setup.PriorityEffectsItem.OnApplyLClick
- BuffHead.Setup.PriorityEffectsItem.OnCloseLUp
- BuffHead.Setup.PriorityEffectsItem.OnCreateLClick
- BuffHead.Setup.PriorityEffectsItem.OnRowLUp
- BuffHead.Setup.PriorityEffectsItem.OnTargetTypeLUp
- BuffHead.Setup.SelectTexture.OnTextureRowLUp
- BuffHead.Setup.Trackers.OnAlwaysIgnoreLClick
- BuffHead.Setup.Trackers.OnAlwaysShowLClick
- BuffHead.Setup.Trackers.OnCloseLUp
- BuffHead.Setup.Trackers.OnTargetChangeClearAlwaysShowLUp
- BuffHead.Setup.Trackers.OnTargetChangeClearBuffsLUp
- BuffHead.Setup.Trackers.OnTargetChangeClearDebuffsLUp
- BuffHead.Setup.Trackers.OnTrackerEnableLUp
- BuffHead.Setup.Trackers.OnTrackerPermanentLUp
- ClosetGoblinCharacterWindow.EquipmentLButtonUp
- ClosetGoblinCharacterWindow.Hide
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- ClosetGoblinCharacterWindow.HotbarPageDownProxy
- ClosetGoblinCharacterWindow.HotbarPageUpProxy
- ClosetGoblinCharacterWindow.OnClickDeleteSetButton
- ClosetGoblinCharacterWindow.OnClickNewSetButton
- ClosetGoblinCharacterWindow.OnClickSetRow
- ClosetGoblinCharacterWindow.OnClickSortNameButton
- ClosetGoblinCharacterWindow.OnClickSortSetButton
- ClosetGoblinCharacterWindow.OnClickSortTacticsButton
- ClosetGoblinCharacterWindow.OnClickZoneConfigButton
- ClosetGoblinCharacterWindow.ShowCloak
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- ClosetGoblinCharacterWindow.ShowHelm
- ClosetGoblinCharacterWindow.UseItemRackToggled
- ClosetGoblinOptionWindow.OnLButtonUp
- ClosetGoblinZoneWindow.Hide
- ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton
- ClosetGoblinZoneWindow.OnClickSortZoneButton
- ClosetGoblinZoneWindow.OnClickZoneRow
- DAoCBuff.CloseMessageWindow
- DAoCBuffFrame.OnLButtonUp
- DAoCBuffSettings.AddEntry
- DAoCBuffSettings.AddFilter
- DAoCBuffSettings.AddFrame
- DAoCBuffSettings.AddList
- DAoCBuffSettings.ChangeFrameName
- DAoCBuffSettings.ChangeGlobalBorder
- DAoCBuffSettings.ChangeGlobalFont
- DAoCBuffSettings.ChangeGlobalGlass
- DAoCBuffSettings.ChangeGlobalRefresh
- DAoCBuffSettings.ChangeGlobalSize
- DAoCBuffSettings.ClearLeft
- DAoCBuffSettings.ClearRight
- DAoCBuffSettings.CloseOptionswindow
- DAoCBuffSettings.CopyLeftToRight
- DAoCBuffSettings.CopyRightToLeft
- DAoCBuffSettings.EditFilterOnLButtonUp
- DAoCBuffSettings.FilterNameChanged
- DAoCBuffSettings.FilterRowOnLButtonUp
- DAoCBuffSettings.FilterSettings.AddClassTable
- DAoCBuffSettings.FilterSettings.AddCondition
- DAoCBuffSettings.FilterSettings.ChangeG4RecursiveConditions
- DAoCBuffSettings.FilterSettings.Close
- DAoCBuffSettings.FilterSettings.RemoveCondition
- DAoCBuffSettings.FrameSettingsRowItemOnLButtonUp
- DAoCBuffSettings.ImExport.Open
- DAoCBuffSettings.MoveFilterDown
- DAoCBuffSettings.MoveFilterUp
- DAoCBuffSettings.MoveLeftToRight
- DAoCBuffSettings.MoveRightToLeft
- DAoCBuffSettings.OpenOptionswindow
- DAoCBuffSettings.RemoveFilter
- DAoCBuffSettings.RemoveFrame
- DAoCBuffSettings.RemoveLeft
- DAoCBuffSettings.RemoveList
- DAoCBuffSettings.RemoveRight
- DAoCBuffSettings.RestartTracker
- DAoCBuffSettings.ShowGeneralOptions
- DAoCBuffSettings.ShowListManager
- DAoCBuffSettings.ToggleCheckBox
- EA_LabelCheckButton.Toggle
- EA_Window_Backpack.ChangeSorting
- EA_Window_Macro.DetailIconLButtonUp
- EA_Window_Macro.Hide
- EA_Window_Macro.HideMacroIconSelectionWindow
- EA_Window_Macro.IconLButtonUp
- EA_Window_Macro.OnSave
- EA_Window_Macro.SelectionIconLButtonUp
- Enemy.AssistUI_ConfigDialog_MarkNewTargetEditTemplate
- Enemy.AssistUI_ConfigDialog_OnMarkNewTargetChanged
- Enemy.AssistUI_ConfigDialog_OnNewTargetSoundChanged
- Enemy.AssistUI_ConfigDialog_OnTargetOnNotifyClickChanged
- Enemy.CombatLogUI_EpsWindow_OnInDpsLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnInHpsLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutDpsAoeLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutDpsLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutHpsAoeLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutHpsLButtonUp
- Enemy.CombatLogUI_IDSAnchor_OnLButtonUp
- Enemy.CombatLogUI_SnapshotWindow_Hide
- Enemy.CombatLogUI_StatsWindow_Hide
- Enemy.CombatLogUI_StatsWindow_OnEpsAoeCurClick
- Enemy.CombatLogUI_StatsWindow_OnEpsAoeMaxClick
- Enemy.CombatLogUI_StatsWindow_OnEpsCurClick
- Enemy.CombatLogUI_StatsWindow_OnEpsMaxClick
- Enemy.CombatLogUI_StatsWindow_SessionAdd
- Enemy.CombatLogUI_StatsWindow_SessionDelete
- Enemy.CombatLogUI_StatsWindow_SessionRename
- Enemy.CombatLogUI_StatsWindow_SortColumnClick
- Enemy.ConfigurationWindow_OnBoolClick
- Enemy.ConfigurationWindow_OnButtonClick
- Enemy.GroupsUI_EffectFilterDialog_Hide
- Enemy.GroupsUI_EffectFilterDialog_Ok
- Enemy.Guard_GuardIndicator_OnLButtonUp
- Enemy.IntercomUI_ChooseChannelDialog_Hide
- Enemy.IntercomUI_ChooseChannelDialog_OnOkButton
- Enemy.IntercomUI_IntercomDialog_Hide
- Enemy.IntercomUI_IntercomDialog_OnCreateButton
- Enemy.IntercomUI_IntercomDialog_OnCreateButton2
- Enemy.IntercomUI_IntercomDialog_OnCreateButton3
- Enemy.IntercomUI_IntercomDialog_OnCreateButton4
- Enemy.IntercomUI_IntercomDialog_OnCreateButton5
- Enemy.IntercomUI_IntercomDialog_OnCreateButton7
- Enemy.IntercomUI_IntercomDialog_OnInviteButton
- Enemy.IntercomUI_IntercomJoinDialog_Hide
- Enemy.IntercomUI_IntercomJoinDialog_OnOkButton
- Enemy.KillSpamUI_KillSpamAreaStatsDialog_Hide
- Enemy.KillSpamUI_KillSpamDialog_OnLButtonUp
- Enemy.KillSpamUI_KillSpamDialog_OnRowLButtonUp
- Enemy.KillSpamUI_KillSpamDialog_OnTotalLButtonUp
- Enemy.MarksUI_EnemyMarkIcon_OnLButtonUp
- Enemy.MarksUI_EnemyMarksWindow_OnAddLButtonUp
- Enemy.MarksUI_EnemyMarksWindow_OnLButtonUp
- Enemy.MarksUI_MarkConfigDialog_Hide
- Enemy.MarksUI_MarkConfigDialog_Ok
- Enemy.OnLeaveButton
- Enemy.ScenarioAlerterUI_ConfigDialog_OnEnabledChanged
- Enemy.ScenarioInfoToggle
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_SwitchToStandard
- Enemy.StopwatchButtonClick
- Enemy.TimerUI_OnLButtonUp
- Enemy.UI_ChooseIconDialog_Hide
- Enemy.UI_ChooseIconDialog_OnListRowLButtonUp
- Enemy.UI_ConfigDialog_Hide
- Enemy.UI_ConfigDialog_Ok
- Enemy.UI_ConfigDialog_Reset
- Enemy.UI_ConfigDialog_ResetAll
- Enemy.UI_Icon_OnLButtonUp
- Enemy.UI_TextEntryDialog_Close
- Enemy.UI_TextEntryDialog_Ok
- Enemy.UnitFramesUI_Anchor_OnLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsAdd
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsDelete
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsDown
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsEdit
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsEnable
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsExport
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsImport
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsReset
- Enemy.UnitFramesUI_ConfigDialog_ClickCastingsUp
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsAdd
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsDelete
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsDown
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsEdit
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsEnable
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsExport
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsImport
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsReset
- Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsUp
- Enemy.UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_OnEnabledChanged
- Enemy.UnitFramesUI_ConfigDialog_OnSortingEnabledChanged
- Enemy.UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsAdd
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsDelete
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsDown
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsEdit
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsEnable
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsExport
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsImport
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsReset
- Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsUp
- Enemy.UnitFramesUI_EffectsIndicatorDialog_ChooseIcon
- Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersAdd
- Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersDelete
- Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersDown
- Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersEdit
- Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersUp
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Ok
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype1Changed
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype2Changed
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype3Changed
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionLButtonUp
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype1Changed
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype2Changed
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype3Changed
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier1Changed
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier2Changed
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier3Changed
- Enemy.UnitFramesUI_UnitFramePartDialog_Hide
- Enemy.UnitFramesUI_UnitFramePartDialog_Ok
- Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype1Changed
- Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype2Changed
- Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype3Changed
- Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_UnitFrame_OnLButtonUp
- FrameManager.OnLButtonUp
- Killer.HideScoreDetailsWindow
- Killer.HideSettingsWindow
- Killer.OnAllTimeSummaryLButtonUp
- Killer.OnRecentSummaryLButtonUp
- Killer.OnSettingsCheckboxChanged
- Killer.ToggleSettingsWindow
- LibGroup.Setup.OnCloseLUp
- LibGroup.Setup.OnEnableGroupDistanceLUp
- LibWBTogglerManager.ToggleManager
- Map.OnClickMap
- MiracleGrowLight.harvestEnd
- MoraleCircle.Click
- PP.ItemSlotLMouse
- PP.OnClose
- PP.PersistantToggle
- PP.SelectDye
- PP.ShowPaperDoll
- PP.ToggleWindow
- PotionBarFloating.ActivatorLButtonUp
- PotionBarFloating.LButtonUp
- PotionBarSettings.AutohideCheck
- PotionBarSettings.MoveDown
- PotionBarSettings.MoveUp
- PotionBarSettings.OnAboutButton
- PotionBarSettings.OnAboutClose
- PotionBarSettings.OnCancelButton
- PotionBarSettings.OnDontSplitNameCheck
- PotionBarSettings.OnResetButton
- PotionBarSettings.OnSaveButton
- PotionBarSettings.OnUseCheck
- PotionBarSettings.Show1Check
- PotionBarSettings.ShowLastCheck
- RoR_SoR.CloseSetOffsetWindow
- RoR_SoR.CloseSetOpacityWindow
- RoR_SoR.CloseSetScaleWindow
- RoR_SoR.KeepClaimDialog
- RoR_SoR.KeepUnClaimDialog
- RoR_SoR.Toggle
- ScenarioSummaryWindow.OnLeaveNowClicked
- Shinies.OnLButtonUp_Close
- ShiniesAuctionsUI.OnLButtonUp_Results_SortButton
- ShiniesAutoUI.OnLButtonUp_AutoDelete
- ShiniesAutoUI.OnLButtonUp_AutoListRow
- ShiniesBrowseUI.OnLButtonUp_Results_ListItem
- ShiniesBrowseUI.OnLButtonUp_Results_SortButton
- ShiniesBrowseUI.OnLButtonUp_Searches
- ShiniesBrowseUI.OnLButtonUp_Searches_SortButton
- ShiniesConfigPrice.OnLButtonUp_DecreasePriority
- ShiniesConfigPrice.OnLButtonUp_EnableModule
- ShiniesConfigPrice.OnLButtonUp_IncreasePriority
- ShiniesConfigUI.OnLButtonUp_ListItem
- ShiniesPostUI.OnLButtonUp_SortButton
- TexturedButtons.Setup.Actionbar.OnCloseLUp
- TexturedButtons.Setup.Actionbar.OnEnableLUp
- TexturedButtons.Setup.Actionbar.OnHideBackgroundLUp
- TexturedButtons.Setup.Actionbar.OnHideEmptyLUp
- TexturedButtons.Setup.AdvancedTextures.OnCloseLUp
- TexturedButtons.Setup.AdvancedTextures.OnDisabledLUp
- TexturedButtons.Setup.AdvancedTextures.OnUseCustomLUp
- TexturedButtons.Setup.AdvancedTextures.OnUsePresetLUp
- TexturedButtons.Setup.Cooldown.OnCloseLUp
- TexturedButtons.Setup.Cooldown.OnCooldownColorLUp
- TexturedButtons.Setup.Cooldown.OnEnableButtonTintingLUp
- TexturedButtons.Setup.Cooldown.OnEnableLUp
- TexturedButtons.Setup.Cooldown.OnHideFlashLUp
- TexturedButtons.Setup.Cooldown.OnRemoveCooldownSLUp
- TexturedButtons.Setup.Cooldown.OnShowGlobalCooldownTextLUp
- TexturedButtons.Setup.Fonts.OnCloseLUp
- TexturedButtons.Setup.Fonts.OnCooldownColorLUp
- TexturedButtons.Setup.Fonts.OnCooldownFontLUp
- TexturedButtons.Setup.Fonts.OnEnableLUp
- TexturedButtons.Setup.Fonts.OnHideCooldownLUp
- TexturedButtons.Setup.Fonts.OnHideKeybindLUp
- TexturedButtons.Setup.Fonts.OnKeybindColorLUp
- TexturedButtons.Setup.Fonts.OnKeybindFontLUp
- TexturedButtons.Setup.Misc.OnCloseLUp
- TexturedButtons.Setup.Misc.OnEnableCooldownPulseLUp
- TexturedButtons.Setup.Misc.OnHideActiveLUp
- TexturedButtons.Setup.Misc.OnHideGlowLUp
- TexturedButtons.Setup.Misc.OnHideQuicklockLUp
- TexturedButtons.Setup.Misc.OnMovableQuicklockLUp
- TexturedButtons.Setup.Misc.OnSaveQuicklockLUp
- TexturedButtons.Setup.OnCloseLUp
- TexturedButtons.Setup.OnRowLUp
- TexturedButtons.Setup.Textures.OnCloseLUp
- TexturedButtons.Setup.Textures.OnDisabledLUp
- TexturedButtons.Setup.Textures.OnUseCustomLUp
- TexturedButtons.Setup.Textures.OnUsePresetLUp
- TexturedButtons.Setup.Tint.OnCloseLUp
- TexturedButtons.Setup.Tint.OnEnableLUp
- TidyChat.Copy.CopyNext
- TidyChat.Copy.CopyPrev
- TidyChat.Copy.OnClose
- TidyChat.LootRoll.OnClose
- TidyChat.Options.OnApply
- TidyChat.Options.OnCheckboxLBU
- TidyChat.Options.OnClose
- TidyChat.Options.OnTabLBU
- TidyChat.Options.Reset
- TidyRoll.CustomAutoRoll.AddById
- TidyRoll.CustomAutoRoll.OnApply
- TidyRoll.CustomAutoRoll.OnClose
- TidyRoll.CustomAutoRoll.OnDeleteButton
- TidyRoll.CustomAutoRoll.OnListLbuttonUp
- TidyRoll.CustomAutoRoll.ToggleOptions
- TidyRollOptions.OnApply
- TidyRollOptions.OnCheckboxLBU
- TidyRollOptions.OnClose
- TidyRollOptions.OnRadioLBU
- TidyRollOptions.OnTabLBU
- TidyRollOptions.Reset
- TurretRange.OnBoxLUp
- TurretRange.Setup.Display.OnCircleColorLUp
- TurretRange.Setup.Display.OnCircleInvertLUp
- TurretRange.Setup.Display.OnCloseLUp
- TurretRange.Setup.Display.OnDistanceColorLUp
- TurretRange.Setup.Display.OnDistanceFontLUp
- TurretRange.Setup.Display.OnGraphicColorLUp
- TurretRange.Setup.Distance.OnApplyLClick
- TurretRange.Setup.Distance.OnCloseLUp
- TurretRange.Setup.Distance.OnCreateLClick
- TurretRange.Setup.Distances.OnAddLClick
- TurretRange.Setup.Distances.OnCloseLUp
- TurretRange.Setup.Distances.OnRowLUp
- TurretRange.Setup.General.OnCloseLUp
- TurretRange.Setup.General.OnEnableLUp
- TurretRange.Setup.OnCloseLUp
- TurretRange.Setup.OnDisplaySetupLUp
- TurretRange.Setup.OnDistancesSetupLUp
- TurretRange.Setup.OnGeneralSetupLUp
- WHMGui.HideDetailsWindow
- WHMGui.HideOptionsWindow
- WHMGui.OnHealerNameMouseUp
- WHMGui.OnOptionsResetButton
- WHMGui.ToggleOptionsWindow
- WSCT.CheckBoxOnLButtonUp
- WSCT.ColorAcceptButtonOnButtonUp
- WSCT.ColorHideMenu
- WSCT.ColorOnButtonUp
- WSCT.HideMenu
- WSCT.LoadProfileOnButtonUp
- WSCT.OnButtonUp
- WSCT.OnLButtonUpColorPicker
- WSCT.TabOnLButtonUp
- WSCT.TestOnButtonUp
- WarBoard.MoveModLeft
- WarBoard.MoveModRight
- WarBoard.OnLayoutModeButton
- WarBoard.OnOptionsButton
- WarBoard.Options.DisableCornerButtons
- WarBoard.Options.MinusHorPad
- WarBoard.Options.MinusVertPad
- WarBoard.Options.OnLButtonUpTab
- WarBoard.Options.OpenUiModWindow
- WarBoard.Options.PlusHorPad
- WarBoard.Options.PlusVertPad
- WarBoard.Options.SetBarDefault
- WarBoard.Options.SetRowAlign
- WarBoard.Options.ToggleBar
- WindowRegisterCoreEventHandler
- core
- followTheLeader.OnLButtonUp
- none

## Examples

- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRoll.CustomAutoRoll.ToggleOptions
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.OnApply
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.OnClose
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.Reset
- AdvancedPetAssist: APAFollowTargetHUD -> APAFollowTargetHUD.OnLButtonUp -> APAGui.OnFollowTargetHUDClicked
- AdvancedPetAssist: APAInstantOnlyHUD -> APAInstantOnlyHUD.OnLButtonUp -> APAGui.OnInstantOnlyHUDClicked

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](../../globals/functions/global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](../../globals/functions/global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](../../globals/functions/global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 98/100) - Global Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 80/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining](../../globals/tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- none
