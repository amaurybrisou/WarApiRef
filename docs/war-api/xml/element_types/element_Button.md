# Button

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdjustTheTip, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, ArmorGraphicToggle, Assist, Atlas |
| Files seen in | Code/Assist/AssistConfiguration.xml, Code/CombatLog/CombatLogSnapshotWindow.xml, Code/CombatLog/CombatLogStatsWindow.xml, Code/Core/ChooseIconDialog.xml, Code/Core/Common.xml, Code/Core/ConfigDialog.xml, Code/Core/ConfigurationWindow.xml, Code/Core/Groups/EffectFilterDialog.xml |
| Namespaces detected | Button |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdjustTheTip: AdjustTheTipMenuItemSlider, AdvancedPetAssist: APAOptionsClose, AdvancedPetAssist: APAOptionsTabsAutoRecall, AdvancedPetAssist: APAOptionsTabsControls, AdvancedPetAssist: APAOptionsTabsFollowTarget, AdvancedPetAssist: APAOptionsTabsGeneral |
| XML usage count | 2412 |
| XML attribute usage count | 2412 |
| Lua usage count | 1 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Button is an interactive XML control used for click and mouse-driven callbacks. It usually binds input events to Lua handler functions.

## Common Attributes

- backgroundtexture
- drawchildrenfirst
- font
- handleinput
- highlighttexture
- id
- inherits
- layer
- name
- textalign
- textureScale
- texturescale

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)

## Common Handler Functions

- DevPadWindow.HideNewWindow


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | DevPadWindow.HideNewWindow | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Deathblow.OnTabLBU, , Vectors.Settings.ButtonClicked, APAGui.OnTabButtonUp, CCTV.TOGGLE_TYPE, BankArkel.PackTab, DK_Config.OnPrevNextCombo, Enemy.CombatLogUI_StatsWindow_SortColumnClick, TortallDPSDetail.GenericSort, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick, InterfaceCore.ReloadUI, Sequencer.ToggleReset, AdvancedRenownTraining.OnLButtonUpTab, AuraConfig.OnConfigTabSelected, EA_Window_OpenPartyManage.ToggleWarbandVisibility, EA_Window_OpenPartyWorld.OnLButtonUpSortButton, Ges.OptionsSoundsPlay, GroupIconsSG.OptionsCHKclick, QueueQueuer_GUI.MapButton_OnLButtonUp, TortallDPSMeter.ToggleDetail, AggroMeter.OnTabLBU, AutoBandWindow.OnLButtonUpTab, ClosetGoblinCharacterWindow.EquipmentLButtonUp, DevPadWindow.HideSaveWindow, EA_Window_OpenParty.ToggleFullWindow, Emojii.SelectTab, FastFriends.UI.OnCharOverrideChanged, FastFriends.UI.OnModeChanged, FrameManager.OnLButtonUp, ObjectInspector.CloseWindow, RandomMountUI.OnTabClick, RoR_SoR.KeepClaimDialog, Statdoll.OnNubLBU, TTitan.UI.OnLButtonUpEntryRow, TomeWindow.DK_PageTopList, WarBoard.OnLayoutModeButton, WarBoard.OnOptionsButton, Warbuilder.Decrement, Warbuilder.Increment, AdvancedRenownTraining.Hide, AdvancedRenownTraining.LinkWindowClose, AuctionStats.ToggleCheckBox, AuraSettings.OnClose, AutoBand.HideCopyLink, BustedGUI.SelectNextError, BustedGUI.SelectPrevError, CMapWindow.ToggleMapPinFilter, CMapWindow.ToggleMapPinGutter, CaVESWindowOptions.Cancel, CrusherConfig.OnClose, DK_Config.ToggleMenu, DPSMeter.OnLButtonUpTab, DPSMeter.ToggleOverview, Deathblow.LinkDeath, Deathblow.LinkKill, Enemy.GroupsUI_EffectFilterDialog_Hide, Enemy.IntercomUI_IntercomDialog_Hide, Enemy.MarksUI_MarkConfigDialog_Hide, Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide, Enemy.UI_ConfigDialog_Hide, Enemy.UI_TextEntryDialog_Close, Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide, Enemy.UnitFramesUI_UnitFramePartDialog_Hide, FastFriends.UI.Close, GDes.OptionsClose, Ges.OptionsClose, HopperConfig.OnClose, ItemRack.Sets.OnTabSelect, KillTracker.OpenTome, MapPin.SetupCancel, Megaphone.HideWindow, MiracleGrow2.ContextItem, MotionConfig.OnClose, NBSetup_Save.Show, ObjectInspector.ClearObjects, ObjectInspector.DepthMinus, ObjectInspector.DepthPlus, ObjectInspector.InspectObject, PieTracker.CloseWindow, PotionBarSettings.OnAboutClose, PotionBarSettings.OnCancelButton, PureConfig.OnClose, RVMOD_Manager.OnLButtonUpToggleButton, RaidMeter.OnTabLBU, ReliquaryHunterOptionsWindow.Cancel, Res.CloseOptions, RvRStatsRvRTab.OnTabSelectRvR, ScenarioGroupWindow.Hide, SocialWindow.ToggleOfflineMembersShowing, SpamBayes.ToggleAbout, SpamBayes.ToggleHelps, SpamBayes.ToggleOptions, Statdoll.Options.onCancel, Statdoll.hide, TalismanGenie.Toggle, TastyButtonsOptions.OnEditLButtonUp, TastyButtonsOptions.OnMenuButtonLUp, TastyButtonsOptions.ToggleButtonInfo, TaxPayer.CloseWindow, TidyChat.Options.OnClose, UiModVersionMismatchWindow.OnCancelButton, UiModWindow.OnAdvancedCancelButton, UiModWindow.OnCancelButton, WCDBConfig.OnClose, WCDPConfig.OnClose, WHMGui.HideOptionsWindow, WTes.CloseWindow, Warbuilder.Respecialize, Warbuilder.TogglePreset, WbLeadHelperMessage.MessageDialogHide, alertMod.OnCancelButton, duel_Decline, nLootLinkOptions.toggleShowing, nLootLinkUtil.toggleShowing, wbLeadHelperConfigWindow.OnLButtonUpTab, zMailModOptions.SelectOption, AGTSettingsWindow.CloseSettingsWindow, APAGui.Hide, AdvancedRenownTraining.DeletePreset, AdvancedRenownTraining.ExportCancelButtonPressed, AdvancedRenownTraining.ExportToLink, AdvancedRenownTraining.ExportToWardrobe, AdvancedRenownTraining.ImportCancelButtonPressed, AdvancedRenownTraining.ImportNameInputCancelButtonPressed, AdvancedRenownTraining.ImportNameInputOkButtonPressed, AdvancedRenownTraining.ImportOkButtonPressed, AdvancedRenownTraining.OnExportButtonPressed, AdvancedRenownTraining.OnLoadPressed, AdvancedRenownTraining.PurchaseAdvances, AdvancedRenownTraining.Respecialize, AdvancedRenownTraining.SavePreset, AdvancedRenownTraining.ShowWardrobeImport, AdvancedRenownTraining.TogglePresets, AdvancedRenownTraining.ToggleSaveSettings, AggroMeter.Close, Atlas.Configuration.Close, AtlasMap.OnDynamicDirectionButtonClicked, AtlasMap.OnTabLBU, AtlasMap.Toggle, AuctionStats.CloseOptionsWindow, AuctionStats.ExitOnComplete, AuctionStats.OnLButtonUpItem, AuctionStats.StartUpdate, AuctionStats.StopUpdate, AuctionWindow.ChangeSearchResultsSorting, AuctionWindow.Hide, AuctionWindow.OnLButtonUpBuy, AuctionWindow.OnLButtonUpCancel, AuctionWindow.OnLButtonUpRefresh, AuctionWindow.OnLButtonUpTab, AuctionWindow.OnSearchResultsRowSelected, AuctionWindowSearchControls.OnLButtonUpClear, AuctionWindowSearchControls.OnLButtonUpSearch, AuctionWindowSellControls.Clear, AuctionWindowSellControls.Create, AuctionWindowSellControls.ItemSlotLButtonUp, AuraConfig.OnAbilityIconLButtonUp, AuraConfig.OnActivationAlertTextTestButton, AuraConfig.OnClose, AuraConfig.OnDeactivationAlertTextTestButton, AuraConfig.OnLButtonUpTextureTintColor, AuraConfig.OnLButtonUpTimerTintColor, AuraConfig.OnTextureChangeButton, AuraSettings.ChangeSorting, AuraSettings.OnAddAura, AuraSettings.OnShare, AuraShares.ChangeSorting, AuraShares.OnClose, AuraShares.OnCloseAuraSharesImportExportWindow, AuraShares.OnImportAura, AuraShares.OnImportExportOkButton, AuraTexture.OnClose, AuraTexture.OnIconLButtonUp, AuraTexture.OnRaceTabSelected, AutoBand.cmd_toggle_gui, AutoBandWindow.Hide, AutoBandWindowConfig.OnAltCheckCheckBox, AutoBandWindowConfig.OnBWSorcAsMDPSCheckBox, AutoBandWindowConfig.OnBackfillCheckBox, AutoBandWindowConfig.OnClearTemplateButton, AutoBandWindowConfig.OnCommonRaceNamesCheckBox, AutoBandWindowConfig.OnDpsWeightingCheckBox, AutoBandWindowConfig.OnExcludeRealmHealerAltSpecCheckBox, AutoBandWindowConfig.OnGuildPriorityCheckBox, AutoBandWindowConfig.OnKickCheckBox, AutoBandWindowConfig.OnKickIgnoreCheckBox, AutoBandWindowConfig.OnKickLowRankCheckBox, AutoBandWindowConfig.OnKickRvRZonesCheckBox, AutoBandWindowConfig.OnKickStealthersCheckBox, AutoBandWindowConfig.OnKickToofarCheckBox, AutoBandWindowConfig.OnPressHealerRankMinusButton, AutoBandWindowConfig.OnPressHealerRankPlusButton, AutoBandWindowConfig.OnPressMaxMdpsMinusButton, AutoBandWindowConfig.OnPressMaxMdpsPlusButton, AutoBandWindowConfig.OnPressMaxRdpsMinusButton, AutoBandWindowConfig.OnPressMaxRdpsPlusButton, AutoBandWindowConfig.OnPressRankMinusButton, AutoBandWindowConfig.OnPressRankPlusButton, AutoBandWindowConfig.OnPressTimeMinusButton, AutoBandWindowConfig.OnPressTimePlusButton, AutoBandWindowConfig.OnResetConfig, AutoBandWindowConfig.OnRestrictRaceCheckBox, AutoBandWindowTemplate.OnApplyTemplate, AutoBandWindowTemplate.OnDefaultCheckBox, AutoBandWindowTemplate.OnDeleteTemplate, AutoBandWindowTemplate.OnOrganize, AutoBandWindowTemplate.OnOrganizeSelectedTemplate, AutoBandWindowTemplate.OnPressDpsMinusButton, AutoBandWindowTemplate.OnPressDpsPlusButton, AutoBandWindowTemplate.OnPressHealersMinusButton, AutoBandWindowTemplate.OnPressHealersPlusButton, AutoBandWindowTemplate.OnPressTanksMinusButton, AutoBandWindowTemplate.OnPressTanksPlusButton, AutoBandWindowTemplate.OnRightClickTemplateMenuCheckBox, AutoBandWindowTemplate.OnSaveTemplate, AutoBandWindowTemplateDeleteConfirm.Hide, AutoBandWindowTemplateDeleteConfirm.OnCancel, AutoBandWindowTemplateDeleteConfirm.OnOK, AutoBandWindowTemplateSave.Hide, AutoBandWindowTemplateSave.OnCancel, AutoBandWindowTemplateSave.OnOK, AutoBandWindowTools.OnAutoFormSearchCheckBox, AutoBandWindowTools.OnAutoPartyNoteCheckBox, AutoBandWindowTools.OnAutoSearchCheckBox, AutoBandWindowTools.OnDiscordReqCheckBox, AutoBandWindowTools.OnFormWarbandButton, AutoBandWindowTools.OnKickButton, AutoBandWindowTools.OnKickOfflineCheckBox, AutoBandWindowTools.OnKickRankCheckBox, AutoBandWindowTools.OnKickZoneCheckBox, AutoBandWindowTools.OnNoMicCheckBox, AutoBandWindowTools.OnNotifyBuffsCheckBox, AutoBandWindowTools.OnOrganizeButton, AutoBandWindowTools.OnOrganizeRangeButton, AutoBandWindowTools.OnPrintButton, AutoBandWindowTools.OnPrintRoleCheckBox, AutoBandWindowTools.OnResetTools, AutoBandWindowTools.OnRightClickOrganizeCheckBox, AutoBandWindowTools.OnSearchRoleButton, AutoBandWindowTools.OnStatsBreakdownButton, AutoBandWindowTools.OnStatsButton, AutoSalvage.OptionsWindow.CheckBoxToggle, AutoSalvage.OptionsWindow.WindowClose, AutoSalvage.ToggleEnabled, AutoSalvage.ToggleTurbo, BankArkel.PackClose, BlackBookWindow.OnLButtonUpClose, BlackBookWindow.OnSortList, BloodyMess.ApplyButton_OnLButtonUp, BuddyBind.GrabName, BuddyBind.OnLButtonRawDeviceInput, BuffHead.Setup.AdvancedCompression.OnCloseLUp, BuffHead.Setup.AdvancedCompression.OnNewLClick, BuffHead.Setup.AdvancedCompressionItem.OnAddEffectLClick, BuffHead.Setup.AdvancedCompressionItem.OnApplyLClick, BuffHead.Setup.AdvancedCompressionItem.OnCloseLUp, BuffHead.Setup.AdvancedCompressionItem.OnCreateLClick, BuffHead.Setup.AdvancedCompressionItemEffect.OnApplyLClick, BuffHead.Setup.AdvancedCompressionItemEffect.OnCloseLUp, BuffHead.Setup.AdvancedCompressionItemEffect.OnCreateLClick, BuffHead.Setup.AdvancedContainers.OnCloseLUp, BuffHead.Setup.AdvancedContainers.OnNewLClick, BuffHead.Setup.AdvancedContainersItem.OnApplyLClick, BuffHead.Setup.AdvancedContainersItem.OnCloseLUp, BuffHead.Setup.AdvancedContainersItem.OnCreateLClick, BuffHead.Setup.AdvancedContainersItem.Properties.OnCloseLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysIgnoreLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysShowLUp, BuffHead.Setup.Container.OnCloseLUp, BuffHead.Setup.Display.OnCloseLUp, BuffHead.Setup.EffectCache.OnCloseLUp, BuffHead.Setup.EffectCache.OnRefreshLClick, BuffHead.Setup.EffectCache.OnResetLClick, BuffHead.Setup.EffectCache.OnSortButtonLUp, BuffHead.Setup.EffectCacheTable.OnCloseLUp, BuffHead.Setup.Filter.OnAddLClick, BuffHead.Setup.Filter.OnCloseLUp, BuffHead.Setup.General.OnAlwaysIgnoreLClick, BuffHead.Setup.General.OnAlwaysShowLClick, BuffHead.Setup.General.OnCloseLUp, BuffHead.Setup.Layout.Manager.OnCloseLUp, BuffHead.Setup.Layout.Manager.OnImportLClick, BuffHead.Setup.Layout.Manager.OnLayoutsActivateLClick, BuffHead.Setup.Layout.Manager.OnLayoutsDeleteLClick, BuffHead.Setup.Layout.Manager.OnLayoutsSaveCurrentLayoutLClick, BuffHead.Setup.Layout.OnApplyLClick, BuffHead.Setup.Layout.OnCloseLUp, BuffHead.Setup.Layout.OnLayerLClick, BuffHead.Setup.Layout.OnManagerLClick, BuffHead.Setup.Layout.Properties.OnCloseLUp, BuffHead.Setup.Layout.Properties.OnCoreSizeAutoSizeLClick, BuffHead.Setup.OnCloseLUp, BuffHead.Setup.Performance.OnCloseLUp, BuffHead.Setup.PriorityEffects.OnCloseLUp, BuffHead.Setup.PriorityEffects.OnNewLClick, BuffHead.Setup.PriorityEffectsItem.OnApplyLClick, BuffHead.Setup.PriorityEffectsItem.OnCloseLUp, BuffHead.Setup.PriorityEffectsItem.OnCreateLClick, BuffHead.Setup.Trackers.OnAlwaysIgnoreLClick, BuffHead.Setup.Trackers.OnAlwaysShowLClick, BuffHead.Setup.Trackers.OnCloseLUp, CCTV.Close, CCTV.ICONCHANGE, CDown.RestartTracker, CDownSettings.CloseOptionswindow, CDownSettings.ToggleTestmode, CMapWindow.OnRallyCallLButtonUp, CMapWindow.OnScenarioQueueLButtonUp, CMapWindow.ToggleFilterMenu, CMapWindow.ToggleRankedLeaderboard, CMapWindow.ToggleScenarioGroupWindow, CMapWindow.ToggleScenarioSummaryWindow, CMapWindow.ToggleWorldMapWindow, CMapWindow.ZoomIn, CMapWindow.ZoomOut, CaVESWindowOptions.Save, CallingSetup.Hide, CalljoinGUI.OnCancelButton, CalljoinGUI.OnExecuteButton, CastSequence.Builder.OnCloseLUp, CastSequence.Builder.OnNewLClick, CastSequence.Builder.OnSortButtonLUp, CastSequence.FindAbility.OnCloseLUp, CastSequence.SequenceBuilder.OnApplyLClick, CastSequence.SequenceBuilder.OnCloseLUp, CastSequence.SequenceBuilder.OnCreateLClick, CastSequence.SequenceBuilder.OnFindAbilityLUp, CastSequence.SequenceBuilder.OnNextButtonLUp, CastSequence.SequenceBuilder.OnPreviousButtonLUp, CastSequence.Setup.OnCloseLUp, ChattyCathy.ApplyOptions, ChattyCathy.CancelOptions, ChattyCathy.SaveOptions, Cheeseboard.Close, Cheeseboard.Show, Cheeseboard.ToggleMakeDefault, ClosetGoblinCharacterWindow.Hide, ClosetGoblinCharacterWindow.HotbarChangeToggled1, ClosetGoblinCharacterWindow.HotbarChangeToggled2, ClosetGoblinCharacterWindow.HotbarChangeToggled3, ClosetGoblinCharacterWindow.HotbarChangeToggled4, ClosetGoblinCharacterWindow.HotbarChangeToggled5, ClosetGoblinCharacterWindow.HotbarPageDownProxy, ClosetGoblinCharacterWindow.HotbarPageUpProxy, ClosetGoblinCharacterWindow.OnClickDeleteSetButton, ClosetGoblinCharacterWindow.OnClickNewSetButton, ClosetGoblinCharacterWindow.OnClickSortNameButton, ClosetGoblinCharacterWindow.OnClickSortSetButton, ClosetGoblinCharacterWindow.OnClickSortTacticsButton, ClosetGoblinCharacterWindow.OnClickZoneConfigButton, ClosetGoblinCharacterWindow.ShowCloak, ClosetGoblinCharacterWindow.ShowCloakHeraldry, ClosetGoblinCharacterWindow.ShowHelm, ClosetGoblinCharacterWindow.UseItemRackToggled, ClosetGoblinZoneWindow.Hide, ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton, ClosetGoblinZoneWindow.OnClickSortZoneButton, CrusherConfig.OnApply, CrusherConfig.OnRevert, DAoCBuff.CloseMessageWindow, DAoCBuffSettings.AddEntry, DAoCBuffSettings.AddFilter, DAoCBuffSettings.AddFrame, DAoCBuffSettings.AddList, DAoCBuffSettings.ChangeFrameName, DAoCBuffSettings.ChangeGlobalBorder, DAoCBuffSettings.ChangeGlobalFont, DAoCBuffSettings.ChangeGlobalGlass, DAoCBuffSettings.ChangeGlobalRefresh, DAoCBuffSettings.ChangeGlobalSize, DAoCBuffSettings.ClearLeft, DAoCBuffSettings.ClearRight, DAoCBuffSettings.CloseOptionswindow, DAoCBuffSettings.CopyLeftToRight, DAoCBuffSettings.CopyRightToLeft, DAoCBuffSettings.EditFilterOnLButtonUp, DAoCBuffSettings.FilterNameChanged, DAoCBuffSettings.FilterSettings.AddClassTable, DAoCBuffSettings.FilterSettings.AddCondition, DAoCBuffSettings.FilterSettings.ChangeG4RecursiveConditions, DAoCBuffSettings.FilterSettings.Close, DAoCBuffSettings.FilterSettings.RemoveCondition, DAoCBuffSettings.ImExport.Open, DAoCBuffSettings.MoveFilterDown, DAoCBuffSettings.MoveFilterUp, DAoCBuffSettings.MoveLeftToRight, DAoCBuffSettings.MoveRightToLeft, DAoCBuffSettings.OpenOptionswindow, DAoCBuffSettings.RemoveFilter, DAoCBuffSettings.RemoveFrame, DAoCBuffSettings.RemoveLeft, DAoCBuffSettings.RemoveList, DAoCBuffSettings.RemoveRight, DAoCBuffSettings.RestartTracker, DPSMeter.OnClose, DPSMeter.Reset, DammazKronTTip.DestroyWindow, DebugWindow.ClearTextLog, DebugWindow.CopyToggle, DebugWindow.Hide, DebugWindow.HideOptions, DebugWindow.ScrollToBottom, DebugWindow.ToggleLogging, DebugWindow.ToggleOptions, DeepSleep.Settings.Close, DetauntHelperConfig.OnTabSelect1, DetauntHelperConfig.OnTabSelect2, DetauntHelperConfig.OnTabSelect3, DetauntHelperConfig.UpdateConfig, DetauntHelperMonitor.DestroyWindow, DetauntTarget.ChangeSorting, DetauntTarget.OnApply, DetauntTarget.OnDeleteTarget, DetauntTarget.OnReset, DetauntTargetInfo.PackDB, DevPadWindow.CancelRename, DevPadWindow.ConfirmRename, DevPadWindow.CreateNewFile, DevPadWindow.DeleteFile, DevPadWindow.Execute, DevPadWindow.FileOnLButtonUp, DevPadWindow.Hide, DevPadWindow.HideConfirmLoadWindow, DevPadWindow.HideDeleteWindow, DevPadWindow.HideLoadProject, DevPadWindow.HideNewWindow, DevPadWindow.LoadFile, DevPadWindow.LoadProjectShow, DevPadWindow.SaveFile, DevPadWindow.SaveLoadFile, DevPadWindow.Toggle, DevPadWindow.ToggleSave, DevPadWindow.UndoClick, DuffTimer.Options.Apply, DuffTimer.Options.OnLButtonUp, DuffTimer.Options.OnLButtonUpTab, DuffTimer.Options.OnOptionsButton, DuffTimer.Options.OnRevertButton, DuffTimer.Options.OnSaveButton, DuffTimer.Options.OnTestButton, Dye.Hide, Dye.Reset, EA_Window_Backpack.ChangeSorting, EA_Window_Macro.DetailIconLButtonUp, EA_Window_Macro.Hide, EA_Window_Macro.HideMacroIconSelectionWindow, EA_Window_Macro.IconLButtonUp, EA_Window_Macro.OnSave, EA_Window_Macro.SelectionIconLButtonUp, EA_Window_OpenParty.OnClickTab, EA_Window_OpenPartyManage.FormWarband, EA_Window_OpenPartyNearby.JoinPartySpecified, EA_Window_OpenPartyNearby.OnLButtonUpSortButton, EA_Window_OpenPartyNearby.RequestUpdateFullList, EA_Window_OpenPartyWorld.JoinPartySpecified, EA_Window_OpenPartyWorld.Search, EA_Window_OpenPartyWorld.SetInterests, EA_Window_OverheadMap.ToggleCurrentEvents, EZCraftX.OnLButtonUp_Apply, EZCraftX.OnLButtonUp_ButtonAwesome, EZCraftX.OnLButtonUp_ButtonGood, EZCraftX.OnLButtonUp_ButtonRegular, EZCraftX.OnLButtonUp_Cancel, EZCraftX.OnLButtonUp_Close, EZCraftX.OnLButtonUp_ContextMenu_Item, EZCraftX.OnLButtonUp_CraftingResult_Close, EZCraftX.OnLButtonUp_Defaults, EZCraftX.OnLButtonUp_useCraftingResult, EZCraftX.OnLButtonUp_usePowerPreview, Emojii.ToggleEmojii, Emojii.UI_ChooseIconDialog_Hide, Emojii.UI_ChooseIconDialog_OnListRowLButtonUp, Enemy.AssistUI_ConfigDialog_MarkNewTargetEditTemplate, Enemy.AssistUI_ConfigDialog_OnMarkNewTargetChanged, Enemy.AssistUI_ConfigDialog_OnNewTargetSoundChanged, Enemy.AssistUI_ConfigDialog_OnTargetOnNotifyClickChanged, Enemy.CombatLogUI_SnapshotWindow_Hide, Enemy.CombatLogUI_StatsWindow_Hide, Enemy.CombatLogUI_StatsWindow_SessionAdd, Enemy.CombatLogUI_StatsWindow_SessionDelete, Enemy.CombatLogUI_StatsWindow_SessionRename, Enemy.ConfigurationWindow_OnButtonClick, Enemy.GroupsUI_EffectFilterDialog_Ok, Enemy.IntercomUI_ChooseChannelDialog_Hide, Enemy.IntercomUI_ChooseChannelDialog_OnOkButton, Enemy.IntercomUI_IntercomDialog_OnCreateButton, Enemy.IntercomUI_IntercomDialog_OnCreateButton2, Enemy.IntercomUI_IntercomDialog_OnCreateButton3, Enemy.IntercomUI_IntercomDialog_OnCreateButton4, Enemy.IntercomUI_IntercomDialog_OnCreateButton5, Enemy.IntercomUI_IntercomDialog_OnCreateButton7, Enemy.IntercomUI_IntercomDialog_OnInviteButton, Enemy.IntercomUI_IntercomJoinDialog_Hide, Enemy.IntercomUI_IntercomJoinDialog_OnOkButton, Enemy.KillSpamUI_KillSpamAreaStatsDialog_Hide, Enemy.MarksUI_MarkConfigDialog_Ok, Enemy.OnLeaveButton, Enemy.ScenarioAlerterUI_ConfigDialog_OnEnabledChanged, Enemy.ScenarioInfoToggle, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SwitchToStandard, Enemy.UI_ChooseIconDialog_Hide, Enemy.UI_ChooseIconDialog_OnListRowLButtonUp, Enemy.UI_ConfigDialog_Ok, Enemy.UI_ConfigDialog_Reset, Enemy.UI_ConfigDialog_ResetAll, Enemy.UI_TextEntryDialog_Ok, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsAdd, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsDelete, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsDown, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsEdit, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsEnable, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsExport, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsImport, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsReset, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsUp, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsAdd, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsDelete, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsDown, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsEdit, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsEnable, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsExport, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsImport, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsReset, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsUp, Enemy.UnitFramesUI_ConfigDialog_OnEnabledChanged, Enemy.UnitFramesUI_ConfigDialog_OnSortingEnabledChanged, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsAdd, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsDelete, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsDown, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsEdit, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsEnable, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsExport, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsImport, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsReset, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_ChooseIcon, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersAdd, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersDelete, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersDown, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersEdit, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_Ok, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype1Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype2Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype3Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype1Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype2Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype3Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier1Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier2Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier3Changed, Enemy.UnitFramesUI_UnitFramePartDialog_Ok, Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype1Changed, Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype2Changed, Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype3Changed, Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged, EveryBodyGuard.Settings.Close, FastFriends.UI.OnRestoreDefaults, FastFriends.UI.OnToggleFriends, FastFriends.UI.OnToggleIgnore, FastFriends.UI.OnToggleSyncMaster, FastInteract.OptionsClose, GDes.EnableButtonCheckToggle, GDes.EnableExtraFeatureCheckToggle, GDes.EnableIconHealthCheckToggle, GDes.EnableModCheckToggle, GDes.EnableSoundCheckToggle, GDes.ExtrasFeatureOffsetDown, GDes.ExtrasFeatureOffsetLeft, GDes.ExtrasFeatureOffsetRight, GDes.ExtrasFeatureOffsetUp, GDes.ExtrasIconsOffsetDown, GDes.ExtrasIconsOffsetLeft, GDes.ExtrasIconsOffsetRight, GDes.ExtrasIconsOffsetUp, GDes.OptionsSave, GDes.OptionsSelectTab, GDes.OptionsSoundsPlay, GDes.Toggle, Ges.EnableButtonCheckToggle, Ges.EnableExtraFeatureCheckToggle, Ges.EnableModCheckToggle, Ges.EnableMonitorClickCheckToggle, Ges.EnableMonitorHealthCheckToggle, Ges.EnableMonitorIconCheckToggle, Ges.EnableSoundCheckToggle, Ges.EnableTrackerHealthCheckToggle, Ges.EnableTrackerIconCheckToggle, Ges.ExtrasFeatureOffsetDown, Ges.ExtrasFeatureOffsetLeft, Ges.ExtrasFeatureOffsetRight, Ges.ExtrasFeatureOffsetUp, Ges.ExtrasTrackerOffsetDown, Ges.ExtrasTrackerOffsetLeft, Ges.ExtrasTrackerOffsetRight, Ges.ExtrasTrackerOffsetUp, Ges.OptionsSave, Ges.OptionsSelectTab, Ges.Toggle, GetStats.CloseWindow, GroupIconsSG.SlashHandler, GroupRangeSetup.General.OnCloseLUp, GroupRangeSetup.OnCloseLUp, GroupRangeSetup.OnGeneralSetupLUp, GroupRangeSetup.OnStyleSetupLUp, GroupRangeSetup.Style.OnCloseLUp, GroupSpotter.Settings.CloseWindow, GuildWarden.ScrollDown, GuildWarden.ScrollUp, GuildWarden.filterReset, GuildWardenWin.Toggle, HealGridGui.CloseGui, HealGridGui.MenuItemSelected, HealGridGuiColorSelect.CloseWindow, HealGridGuiColorSelect.SelectColor, HealGridGuiColorSelect.SelectColorGroup, HealGridGuiSpellList.ListSelectRow, HealGridGuiTabMouseClick.EditCommit, HealGridGuiTabMouseClick.ListRemoveRow, HealGridGuiTabMouseClick.ListSelectRow, HealGridGuiTabMouseClick.ListUnselectRow, HealGridGuiTabMouseClick.OnSortList, HealGridGuiTabRangeScan.SetFriendlySpell, HealGridGuiTabRangeScan.SetHostileSpell, HealGridGuiTabRangeScan.SetResurrectSpell, HealGridGuiTabSpellTrack.EditCommit, HealGridGuiTabSpellTrack.ListRemoveRow, HealGridGuiTabSpellTrack.ListSelectRow, HealGridGuiTabSpellTrack.ListUnselectRow, HealGridGuiTabSpellTrack.OnSortList, HealGridGuiTabSpellTrack.SwitchTab, HealGridGuiTabSpellTrack.TabColorChangeSignColor, HealGridGuiTabSpellTrack.TabGeneralLookupSpellName, HealGridGuiUtility.ColorSelectEditForegroundColor, HopperConfig.OnApply, HopperConfig.OnRevert, ItemRack.Sets.ButtonDelete, ItemRack.Sets.ButtonNew, ItemRack.Sets.ButtonSave, JunkDump.JunkBackpack, JunkDumpOptions.Done, KeyBar.IndexSelectDown, KeyBar.IndexSelectUp, KeyBar.LButtonUp, KeyBarSettings.CloseHelpWindow, KeyBarSettings.CloseSettingsWindow, KeyBarSettings.OpenHelpWindow, KeyBarSettings.SaveSettingsWindow, Keyset.Setup.Load.OnCloseLUp, Keyset.Setup.Load.OnDeleteLClick, Keyset.Setup.Load.OnLoadLClick, Keyset.Setup.Save.OnCloseLUp, Keyset.Setup.Save.OnSaveLClick, KeysetMonsterPlay.Setup.Profile.OnCloseLUp, Killer.HideScoreDetailsWindow, Killer.HideSettingsWindow, Killer.OnSettingsCheckboxChanged, Killer.ToggleSettingsWindow, KwestorGui.CloseGui, KwestorGui.SwitchTab, KwestorGuiTabArea.ListAddArea, KwestorGuiTabArea.ListRemoveArea, KwestorGuiTabArea.ListSelectRow, KwestorGuiTabArea.ListToggleRvR, KwestorGuiTabArea.OnSortList, KwestorTracker.LeftButtonClick, LPET.AddProfileOnButtonUp, LPET.HideMenu, LPET.LoadProfileOnButtonUp, LPET.RemoveProfileOnButtonUp, LPET.RenameProfileOnButtonUp, LPET.SaveProfileOnButtonUp, LPET.TabOnLButtonUp, LibAddonButton.Manager.Advanced.OnAnimationAddFrameLClick, LibAddonButton.Manager.Advanced.OnApplyLClick, LibAddonButton.Manager.Advanced.OnCloseLUp, LibAddonButton.Manager.CustomItem.OnApplyLClick, LibAddonButton.Manager.CustomItem.OnCloseLUp, LibAddonButton.Manager.CustomItem.OnCreateLClick, LibAddonButton.Manager.CustomItem.OnMenuManageLClick, LibAddonButton.Manager.MacroHelp.OnCloseLUp, LibAddonButton.Manager.OnCloseLUp, LibAddonButton.Manager.OnCreateLClick, LibAddonButton.Manager.OnNewLClick, LibAddonButton.Manager.OnOptionLClick, LibAddonButton.Menu.OnLButtonUp, LibGroup.Setup.OnCloseLUp, LibWBTogglerManager.ToggleManager, LootAlert.Clicky, LootAlert.Close, MBuffGui.OnCloseLUp, MBuffSetup.SmartBuff.OnCloseLUp, MailWindowTabMessage.OnClose, MailWindowTabMessage.OnLButtonDeleteMessage, MailWindowTabMessage.OnLButtonUpReplyButton, MailWindowTabMessage.OnLButtonUpTakeAll, MapMonster.Editor.Close, MapMonster.Editor.OnClickLeft, MapMonster.Editor.OnClickMiddle, MapMonster.Editor.OnClickRight, MapMonster.Editor.OnLockChanged, MapMonster.Editor.OnPrivateChanged, MapMonster.OnMouseLClickFilterButton, MapMonster.PinTypeEditor.Close, MapMonster.PinTypeEditor.CloseChooser, MapMonster.PinTypeEditor.LeftButton, MapMonster.PinTypeEditor.MiddleButton, MapMonster.PinTypeEditor.OnClickChangeButton, MapMonster.PinTypeEditor.OnLockChanged, MapMonster.PinTypeEditor.OnPrivateChanged, MapMonster.PinTypeEditor.RightButton, MapMonster.ToggleFilterChoice, MapMonster_Calibrate.Hide, MapPin.SetupAccept, MapPin.ToggleCheckbox, MapPin.UI_ChooseIconDialog_Hide, MapPin.UI_ChooseIconDialog_OnListRowLButtonUp, MassMailWindow.OnLButtonUpSendButton, MassMailWindow.OnSlotLButtonUp, MassMailWindow.ToggleCraftingList, MassMailWindow.ToggleRarityList, MassRefine.CancelButton, MassRefine.OkayButton, MassRefine.OnClose, Megaphone.HighlightLeaderToggle, Megaphone.ShowLeaderToggle, Megaphone.TestAlert, Minmap.GroupQueue, Minmap.OnScenarioQueueLButtonUp, Minmap.SoloQueue, MiracleGrow.menuShow, MiracleGrow.onHHoverEnd, MiracleGrow2.ConfigHelp, MiracleGrow2.ConfigSoundTest, MiracleGrow2.LayoutNextBar, MiracleGrow2.LayoutPrevBar, MiracleGrow2.NextVis, MiracleGrow2.PresetLayout, MiracleGrow2.PresetSettings, MiracleGrow2.PrevVis, MiracleGrow2.onClickNutrient, MiracleGrow2.onClickSeed, MiracleGrow2.onClickSoil, MiracleGrow2.onClickWater, MiracleGrow2.onRepeat, MiracleGrowLight.harvestEnd, Motion.OnLButtonUp_ContextMenu_Item, Motion.OnLButtonUp_ContextMenu_Recipe, MotionConfig.OnApply, MotionConfig.OnRevert, NBSBActionBar.OnNext, NBSBActionBar.OnPrevious, NBSBChecks.OnSetCheck, NBSBCore.OnClose, NBSBCore.OnDel, NBSBCore.OnLoad, NBSBCore.OnReset, NBSBCore.OnSave, NBSetup_Save.OnOkayButton, NemesisKill, NemesisPopup, ObjectInspector.Toggle, Obsidian.Setup.Castbar.OnCloseLUp, Obsidian.Setup.EffectTracker.OnCloseLUp, Obsidian.Setup.OnCastbarSetupLUp, Obsidian.Setup.OnCloseLUp, Obsidian.Setup.OnEffectTrackerSetupLUp, PP.ItemSlotLMouse, PP.OnClose, PP.SelectDye, PP.ShowPaperDoll, PP.ToggleWindow, PartyAdWindow.Hide, PartyAdWindow.OnAdvertise, PartyAdWindow.OnClearNote, PartyAdWindow.OnMinusDps, PartyAdWindow.OnMinusHealer, PartyAdWindow.OnMinusTank, PartyAdWindow.OnPlusDps, PartyAdWindow.OnPlusHealer, PartyAdWindow.OnPlusTank, PartyAdWindow.OnToggleDetectedNeeds, PartyAdWindow.OnToggleSpecCheck, PartyAdWindow.OnUpdateNote, Phantom.CloseSettingsWindow, PieTracker.EnableGuildCheckToggle, PieTracker.EnableModCheckToggle, PieTracker.ReSortList, PieTracker.ResetAllCounts, PieTracker.SaveOptions, PotionBarFloating.LButtonUp, PotionBarSettings.AutohideCheck, PotionBarSettings.MoveDown, PotionBarSettings.MoveUp, PotionBarSettings.OnAboutButton, PotionBarSettings.OnDontSplitNameCheck, PotionBarSettings.OnResetButton, PotionBarSettings.OnSaveButton, PotionBarSettings.OnUseCheck, PotionBarSettings.Show1Check, PotionBarSettings.ShowLastCheck, PureConfig.OnApply, PureConfig.OnRevert, PureTargetUnitFrame.OnLButtonUp_Sigil, QTS.ChangeSorting, QTS.OnClose, QueueQueuer_GUI.BlacklistAllButton_OnLButtonUp, QueueQueuer_GUI.BlacklistNoneButton_OnLButtonUp, QueueQueuer_GUI.Close_OnLButtonUp, QueueQueuer_GUI.JoinButton_OnLButtonUp, QueueQueuer_GUI.LeaveButton_OnLButtonUp, QueueQueuer_GUI.OnLButtonUpTab, QueueQueuer_GUI.QueuerCheckButton_OnLButtonUp, QuickWarReport.OnConfirmAccept, QuickWarReport.OnConfirmDecline, RVAPI_ColorDialog.OnLButtonUpButtonCancel, RVAPI_ColorDialog.OnLButtonUpButtonOK, RVAPI_Range.OnCheckBoxMapDistancesEnabled, RVMOD_3DPortrait.OnCheckBoxEnabled, RVMOD_Manager.OnClickReloadUI, RVMOD_Manager.OnClickTabGeneral, RVMOD_Manager.OnClickTabSettings, RVMOD_PlayerStatus.SetEnable, RVMOD_SquaredDistances.OnCheckBoxUseGlobalScale, RVMOD_SquaredDistances.OnLButtonUpBottom, RVMOD_SquaredDistances.OnLButtonUpLeft, RVMOD_SquaredDistances.OnLButtonUpRight, RVMOD_SquaredDistances.OnLButtonUpTop, RVMOD_Targets.OnDeleteRow, RVMOD_Targets.OnToggleRowSettings, RaidMeter.Close, RandomMountUI.Hide, RandomMountUI.OnAddCustomMount, RandomMountUI.OnDropSlotLButtonUp, RealmStatus.HideHistoryWindow, ReferWindow.OnAccept, ReferWindow.OnDecline, ReliquaryHunterOptionsWindow.Save, RememberIt.CloseSettingsWindow, Res.EnableButtonCheckToggle, Res.EnableCastCheckToggle, Res.EnableModCheckToggle, Res.SaveOptions, Res.Toggle, RoR_SoR.CloseSetOffsetWindow, RoR_SoR.CloseSetOpacityWindow, RoR_SoR.CloseSetScaleWindow, RoR_SoR.KeepUnClaimDialog, RoR_SoR.Toggle, RoR_debolster.CloseWindow, Rotation.OnStart, Rotation.OnStop, SOR.CloseNotes, SOR.CloseOptions, SOR.ResetItems, SOROptions.GetSettings, SOROptions.Save, SOROptions.timer1, SOROptions.timer2, SOROptions.timer3, ScenarioGroupWindow.ClaimMainAssist, ScenarioGroupWindow.CloseSetOpacityWindow, ScenarioGroupWindow.HideUngroupedPlayersWindow, ScenarioGroupWindow.JoinGroup, ScenarioGroupWindow.LeaveGroup, ScenarioGroupWindow.ShowUngroupedPlayers, ScenarioRefresh.Refresh, ScenarioStats.CallSettingsWindow, ScenarioStats.SettingsWindowCancel, ScenarioStats.SettingsWindowClose, ScenarioStats.WarningWindowButton, ScenarioStats.WindowClose, ScenarioSummaryWindow.OnLeaveNowClicked, ScenarioSummaryWindow.ToggleShowing, Sequencer.ToggleShowing, SessionRPs.CloseHSWindow, SessionRPs.WarningWindowButton, Shinies.OnLButtonUp_Close, ShiniesAuctionsUI.OnLButtonUp_Results_SortButton, ShiniesAutoUI.OnLButtonUp_AutoDelete, ShiniesBrowseUI.OnLButtonUp_Results_SortButton, ShiniesBrowseUI.OnLButtonUp_Searches_SortButton, ShiniesConfigPrice.OnLButtonUp_DecreasePriority, ShiniesConfigPrice.OnLButtonUp_EnableModule, ShiniesConfigPrice.OnLButtonUp_IncreasePriority, ShiniesPostUI.OnLButtonUp_SortButton, SiegeChatSettings.CloseWindow, SiegeChatSettings.OnToggleChannel, SiegeChatSettings.SaveOptions, SiegeChatSettings.ShowWindow, SocialWindow.Hide, SocialWindow.MemberAdd, SocialWindow.OnLButonTab, SocialWindowBuddyList.OnLButtonUpTab, SocialWindowBuddyList.ShowSocial, SocialWindowBuddyList.ToggleFilterMenu, SocialWindowBuddyList.ToggleSocialWindows, SocialWindowBuddyListTabFriends.OnLButtonUpPlayerRow, SocialWindowTabFriends.AddFriend, SocialWindowTabFriends.OnCancelDescription, SocialWindowTabFriends.OnDescriptionAccept, SocialWindowTabFriends.OnLButtonUpPlayerRow, SocialWindowTabFriends.OnSortPlayerList, SocialWindowTabFriends.RemoveFriend, SocialWindowTabFriends.SwitchFriendType, SocialWindowTabFriends.ToggleFilterMenu, SocialWindowTabIgnore.AddIgnore, SocialWindowTabIgnore.OnLButtonUpPlayerRow, SocialWindowTabIgnore.OnSortPlayerList, SocialWindowTabIgnore.RemoveIgnore, SocialWindowTabOptions.OnPressAFKButton, SocialWindowTabOptions.OnPressAdvisorButton, SocialWindowTabOptions.OnPressAnonymousButton, SocialWindowTabOptions.OnPressDisableBuddyListButton, SocialWindowTabOptions.OnPressHiddenButton, SocialWindowTabOptions.OnPressPrivatePartyButton, SocialWindowTabSearch.OnLButtonUpPlayerRow, SocialWindowTabSearch.OnPressSearchPlayerButton, SocialWindowTabSearch.OnSortPlayerList, Soloq.hideOverviewWindow, SpamBayes.ClearLog, SpamBayes.ToggleControlCenter, Squared.CloseImportExport, Squared.ExportSettings, Squared.ImportSettings, Statdoll.Options.onApply, StatdollWnd.onLBMagicWnd, StatdollWnd.onLBMeleeWnd, StatdollWnd.onLBRangeWnd, TOLSettingsUI.BeginEditPhrases, TOLSettingsUI.CloseLButtonUp, TOLSettingsUI.PhraseDeleteLButtonUp, TOLSettingsUI.PhraseSaveLButtonUp, TOLSettingsUI.SkillDeleteLButtonUp, TOLSettingsUI.SkillSaveLButtonUp, TTitan.TTButton.OnLClick, TTitan.UI.Hide, TTitan.UI.OnLButtonUpFAQButton, TTitan.UI.OnLButtonUpMinusButton, TTitan.UI.OnLButtonUpPlusButton, TacticSetNames.SettingsWindow.CloseSettingsWindow, TacticSetNames.SettingsWindow.OpenSelectFont, TalismanGenie.ContainerCheckToggle, TalismanGenie.CurioCheckToggle, TalismanGenie.EssenceCheckToggle, TalismanGenie.GoNext, TalismanGenie.GoPrev, TalismanGenie.GoldDustCheckToggle, TalismanGenie.Reset, TalismanGenie.Search, TastyButtonsOptions.AddSelectedButton_All, TastyButtonsOptions.AddStateButton, TastyButtonsOptions.BlockButton, TastyButtonsOptions.ButtonDelProfile, TastyButtonsOptions.ButtonLoadProfile, TastyButtonsOptions.ButtonSaveProfile, TastyButtonsOptions.Button_Set_Texture, TastyButtonsOptions.ClearCreateView, TastyButtonsOptions.ClearEditComplex, TastyButtonsOptions.ClearEditRange, TastyButtonsOptions.ClearEditView, TastyButtonsOptions.ClearGroupView, TastyButtonsOptions.ClearMiscView, TastyButtonsOptions.ClearSelectedButton_All, TastyButtonsOptions.ClearStateButton, TastyButtonsOptions.ClearStateView, TastyButtonsOptions.CopyActionBars, TastyButtonsOptions.CreateButton, TastyButtonsOptions.DeleteButton, TastyButtonsOptions.DeleteStateButton, TastyButtonsOptions.GeneralWindowHide, TastyButtonsOptions.GroupButton, TastyButtonsOptions.Hide, TastyButtonsOptions.ModifyStateButton, TastyButtonsOptions.OnButtonSelected, TastyButtonsOptions.OnTabSelect, TastyButtonsOptions.SaveAdvOptions, TastyButtonsOptions.ScaleButton, TastyButtonsOptions.SetABCTint, TastyButtonsOptions.SetHotKeyFont, TastyButtonsOptions.SetOORTint, TastyButtonsOptions.SetPosButton, TastyButtonsOptions.SetTimerFont, TastyButtonsOptions.Set_ModeManual, TastyButtonsOptions.Set_ModeSelect, TastyButtonsOptions.ToggleActionBarsButton, TastyButtonsOptions.ToggleAddon, TastyButtonsOptions.ToggleChangeFontColor, TastyButtonsOptions.ToggleProfileWindow, TastyButtonsOptions.UnGroupButton, TaxPayer.EnableGuildCheckToggle, TaxPayer.EnableModCheckToggle, TaxPayer.ReSortList, TaxPayer.ResetAllCounts, TaxPayer.SaveOptions, TexturedButtons.Setup.Actionbar.OnCloseLUp, TexturedButtons.Setup.AdvancedTextures.OnCloseLUp, TexturedButtons.Setup.Cooldown.OnCloseLUp, TexturedButtons.Setup.Fonts.OnCloseLUp, TexturedButtons.Setup.Misc.OnCloseLUp, TexturedButtons.Setup.OnCloseLUp, TexturedButtons.Setup.Textures.OnCloseLUp, TexturedButtons.Setup.Tint.OnCloseLUp, TidyChat.Copy.CopyNext, TidyChat.Copy.CopyPrev, TidyChat.Copy.OnClose, TidyChat.LootRoll.OnClose, TidyChat.Options.OnApply, TidyChat.Options.OnTabLBU, TidyChat.Options.Reset, TidyQueue.OnJoinSelected, TidyQueue.OnSelectAll, TidyQueue.OnSelectNone, TidyRoll.CustomAutoRoll.AddById, TidyRoll.CustomAutoRoll.OnApply, TidyRoll.CustomAutoRoll.OnClose, TidyRoll.CustomAutoRoll.OnDeleteButton, TidyRollOptions.OnTabLBU, TokenMachine.OpenOption, TokenMachine.ToggleSettings, TomeTracker.Journal.Hide, TomeTracker.Journal.OnLButtonUpTab, TomeWindow.DK_CallPageLastGrudge, TomeWindow.DK_OnBookmark, TomeWindow.DK_OnClickSortButton, TomeWindow.DK_OnSelectProfilMini, TortallDPSMeter.CycleScale, TortallDPSMeter.DamageTab, TortallDPSMeter.HealingTab, TortallDPSMeter.Reset, TortallDPSMeter.Toggle, TortallDPSMeter.ToggleGroup, TurretRange.Setup.Display.OnCloseLUp, TurretRange.Setup.Distance.OnApplyLClick, TurretRange.Setup.Distance.OnCloseLUp, TurretRange.Setup.Distance.OnCreateLClick, TurretRange.Setup.Distances.OnAddLClick, TurretRange.Setup.Distances.OnCloseLUp, TurretRange.Setup.General.OnCloseLUp, TurretRange.Setup.OnCloseLUp, TurretRange.Setup.OnDisplaySetupLUp, TurretRange.Setup.OnDistancesSetupLUp, TurretRange.Setup.OnGeneralSetupLUp, Twister.CloseSettingsWindow, UiModVersionMismatchWindow.OnClickModListSortButton, UiModVersionMismatchWindow.OnReEnableButton, UiModWindow.OnAdvancedButton, UiModWindow.OnAdvancedOkayButton, UiModWindow.OnClickModListSortButton, UiModWindow.OnClickModRow, UiModWindow.OnDebugWindowButton, UiModWindow.OnDisableAllButton, UiModWindow.OnEnableAllButton, UiModWindow.OnOkayButton, UiModWindow.OnReEnableAutoDisabled, UiModWindow.OnToggleModEnabled, Vectors.Settings.CancelImport, Vectors.Settings.Close, Vectors.Settings.WindowHiderClosedClicked, Vectors.Settings.WindowHiderOpenClicked, WARCommander.AllClear, WARCommander.ReportPosition, WARCommander.Target, WARCommander.ToggleVisibility, WARCommander.cmd_ress, WARCommander.cmd_spot, WARCommander.cmd_spotgroup, WARCommander.cmd_spotwb, WARCommanderConfig.Close, WARCommanderConfig.OnListenCheckBoxChanged, WARCommanderConfig.OnUpdatesCheckBoxChanged, WARCommanderConfig.OnUpdatesTitleCheckBoxChanged, WARCommanderConfig.OnWaypointCheckBoxChanged, WBStutterLess.OnLButtonUp, WCDBConfig.OnApply, WCDBConfig.OnRevert, WCDPConfig.OnApply, WCDPConfig.OnRevert, WHMGui.HideDetailsWindow, WHMGui.OnOptionsResetButton, WHMGui.ToggleOptionsWindow, WSCT.ColorAcceptButtonOnButtonUp, WSCT.ColorHideMenu, WSCT.ColorOnButtonUp, WSCT.HideMenu, WSCT.LoadProfileOnButtonUp, WSCT.OnButtonUp, WSCT.TabOnLButtonUp, WSCT.TestOnButtonUp, WTes.Backpack.CheckToggle, WTes.GlyphPos.CheckToggle, WTes.GlyphPos.Reset, WTes.MoraleClick.CheckToggle, WTes.SaveOptions, WTes.ScenStartPos.CheckToggle, WTes.ScenStartPos.Reset, WTes.SiegeChat.CheckToggle, WarBoard.MoveModLeft, WarBoard.MoveModRight, WarBoard.Options.MinusHorPad, WarBoard.Options.MinusVertPad, WarBoard.Options.OnLButtonUpTab, WarBoard.Options.OpenUiModWindow, WarBoard.Options.PlusHorPad, WarBoard.Options.PlusVertPad, WarBoard.Options.SetRowAlign, WarBoard_FPS.OnOptionsButton, WarBoard_FPSOptions.OnToggleAvg, WarBoard_Loc.CloseSetWidthWindow, Warbuilder.AbilityLinkClose, Warbuilder.Close, Warbuilder.ImportCurrent, Warbuilder.OnTabSelect, Warbuilder.Reset, Warbuilder.ShareLink, Warbuilder.ToggleHotBar, WbLeadHelperMessage.OnOk, XpStatus.OnCancelNewQuota, XpStatus.OnCancelSetOpacity, XpStatus.OnSetNewQuota, ZonePOP.AddZone, alertMod.OnOkayButton, alertMod.OnSetDefaults, duel_Accept, duel_Cancel, emotes.OnSortemotes, emotes.Toggleemotes, followTheLeader.OnLButtonUp, minesweep.Close, minesweep.LButtonUp, nLootLinkGUI.changeSorting, nLootLinkGUI.reset, nLootLinkGUI.search, nLootLinkOptions.onToggleRarityConfirm, snt_buttons.toggle_setup, snt_castbar.set1, snt_castbar.set2, snt_castbar.set3, snt_castbar.set4, snt_castbar.set5, snt_castbar.set6, snt_castbar.set7, snt_castbar.set8, snt_castbar.toggle_setup, snt_info.B0, snt_info.B1, snt_info.B2, snt_info.B3, snt_info.F0, snt_info.F1, snt_info.F2, snt_info.F3, snt_info.F4, snt_info.F5, snt_info.F6, snt_info.F7, snt_info.I0, snt_info.I1, snt_info.I2, snt_info.I3, snt_info.toggle_setup, snt_panel.dec, snt_panel.del, snt_panel.inc, snt_panel.new, snt_panel.toggle_setup, wargames.gems.FieldLButtonUp, wargames.gems.Hide, wargames.pairs.CardLButtonUp, wargames.pairs.Hide, wargames.ttt.FieldLButtonUp, wargames.ttt.Hide, wbLeadHelper.hideWbLeadHelperWindowName, wbLeadHelper.toggleColoredChat, wbLeadHelperChooseIcon.Hide, wbLeadHelperChooseIcon.OnListRowLButtonUp, wbLeadHelperConfigTab.ChooseIconMessageEnd, wbLeadHelperConfigTab.ChooseIconMessageStart, wbLeadHelperConfigTab.OnLfgIconsCheckBoxUp, wbLeadHelperConfigTab.OnPreview, wbLeadHelperConfigTab.OnReset, wbLeadHelperConfigTab.OnSave, wbLeadHelperConfigWindow.Hide, wbLeadHelperMessagesTab.ListDelete, wbLeadHelperMessagesTab.ListDown, wbLeadHelperMessagesTab.ListEnable, wbLeadHelperMessagesTab.ListUp, wbLeadHelperMessagesTab.MessageAdd, wbLeadHelperMessagesTab.MessageClone, wbLeadHelperMessagesTab.MessageEdit, wbLeadHelperMessagesTab.OnPreviewMessages, wbLeadHelperMessagesTab.OnResetMessages, wbLeadHelperMessagesTab.OnSaveMessages, zMailMod.SelectMail, zMailModAuction.OpenSelected, zMailModAuction.SelectAll, zMailModAuction.SelectMoney, zMailModInbox.OpenSelected, zMailModInbox.SelectAll, zMailModLog.Hide, zMailModOptions.Hide, zMailModOptions.Save | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | AutoBandWindowConfig.OnSettingMouseOver, Warbuilder.overtooltip, WARCommander.ChatButtonMouseOver, BankArkel.PackTabMover, FrameManager.OnMouseOver, Sequencer.OnMouseOver, QueueQueuer_GUI.MapButton_OnMouseOver, MapMonster.PinTypeEditor.MouseOverDescription, SocialWindowBuddyList.OnMouseOverTooltipElement, AutoBandWindowConfig.OnHealerRankReqMouseOver, AutoBandWindowConfig.OnRankReqMouseOver, Enemy.ConfigurationWindow_ShowTooltip, MiracleGrow2.ContextHover, MiracleGrow2.onHHover, WarBoard.OnMouseOver, WarBoard.OnMouseOverBottom, , ActionBars.OnMouseoverHotbarPageDown, ActionBars.OnMouseoverHotbarPageUp, AnywhereTrainer.OnMouseOver, AtlasMap.OnDynamicDirectionButtonMouseOver, AuctionStats.OnMouseOverItem, AuctionWindow.OnMouseOverResultsIcon, AuctionWindowSellControls.ItemSlotMouseOver, AuraConfig.OnIconMouseOver, AutoBand.OnIconMouseOver, AutoBandWindowConfig.OnBackfillMouseOver, AutoBandWindowConfig.OnCommonRaceNamesMouseOver, AutoBandWindowConfig.OnExcludeRealmHealerAltSpecMouseOver, AutoBandWindowConfig.OnKickLowRankMouseOver, AutoBandWindowTemplate.OnApplyButtonMouseOver, AutoBandWindowTools.OnDiscordReqMouseOver, AutoBandWindowTools.OnNoMicMouseOver, AutoBandWindowTools.OnStatsBreakdownButtonMouseOver, AutoBandWindowTools.OnStatsButtonMouseOver, BankArkel.PackIconOver, BlackBookWindow.OnMouseOverSortButton, BuffHead.Setup.EffectCache.OnSortButtonMouseOver, BuffHead.Setup.Layout.OnLayerMouseOver, CG_ItemRack.EquipmentMouseOver, CMapWindow.OnMouseOverFilterMenuButton, CMapWindow.OnMouseoverRallyCall, CMapWindow.OnMouseoverRankedLeaderboard, CMapWindow.OnMouseoverScenarioGroupBtn, CMapWindow.OnMouseoverScenarioQueue, CMapWindow.OnMouseoverScenarioSummaryBtn, CMapWindow.OnMouseoverWorldMapBtn, CMapWindow.OnMouseoverZoomInBtn, CMapWindow.OnMouseoverZoomOutBtn, CastSequence.Builder.OnSortButtonMouseOver, CastSequence.SequenceBuilder.OnFindAbilityMouseOver, ClosetGoblinCharacterWindow.EquipmentMouseOver, ClosetGoblinCharacterWindow.ShowCloakOptions, ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly, ClosetGoblinCharacterWindow.ShowShowCloakOnly, ClosetGoblinCharacterWindow.ShowShowHelm, ClosetGoblinCharacterWindow.ShowShowHelmOnly, DetauntTarget.OnMouseOverSortButton, DuffTimer.Options.OnTestButtonToolTip, EA_Window_Macro.DetailIconMouseOver, EA_Window_Macro.IconMouseOver, EA_Window_Macro.SelectionIconMouseOver, EA_Window_OpenPartyNearby.UpdatePartyTooltipFromButton, EA_Window_OpenPartyNearby.UpdatePartyTooltipFromWindow, EA_Window_OpenPartyWorld.UpdatePartyTooltipFromButton, EA_Window_OpenPartyWorld.UpdatePartyTooltipFromWindow, EA_Window_OverheadMap.OnMouseOverCurrentEvents, EZCraftX.OnMouseOver_ButtonAwesome, EZCraftX.OnMouseOver_ButtonGood, EZCraftX.OnMouseOver_ButtonRegular, EZCraftX.OnMouseOver_ContextMenu_Item, EZCraftX.OnMouseOver_useCraftingResult, EZCraftX.OnMouseOver_usePowerPreview, Emojii.OnMouseOverStart, Emojii.OnMouseOverStart2, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseOver, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseOver, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseOver, GDes.ToggleButtonMouseOver, Ges.ToggleButtonMouseOver, ItemRack.EquipmentMouseOver, ItemRack.Tactics.TacticsSetMenuTooltip, JunkDump.OnMouseOverButton, KeyBar.MouseOver, KwestorTracker.MouseOverQuest, LPET.OnMouseOver, LibAddonButton.Manager.OnNewMouseOver, LibAddonButton.Menu.OnMouseOver, MapMonster.OnMouseOverFilterButton, MapMonster.ToggleFiltersSubMenu, MassMailWindow.OnSlotMouseOver, Minmap.OnMouseoverScenarioQueue, Minmap.OnMouseoverScenarioSummaryBtn, MiracleGrow.onHHover, MiracleGrow2.ConfigSoundTestTip, MiracleGrow2.onHoverNutrient, MiracleGrow2.onHoverRepeat, MiracleGrow2.onHoverSeed, MiracleGrow2.onHoverSoil, MiracleGrow2.onHoverWater, MiracleGrowLight.harvestStart, MiracleGrowLight.onHover, MoraleSet.MenuTooltip, Motion.OnMouseOver_ContextMenu_Item, NBSBChecks.OnOverCheck, PP.ItemSlotMouseOver, PieTracker.MouseOverSortButton, PotionBarFloating.MouseOver, PotionBarSettings.OnMouseoverAutohide, PotionBarSettings.OnMouseoverDontSplitName, PotionBarSettings.OnMouseoverMoveDown, PotionBarSettings.OnMouseoverMoveUp, PotionBarSettings.OnMouseoverShow1, PotionBarSettings.OnMouseoverShowLast, PotionBarSettings.OnMouseoverUse, PureTargetUnitFrame.OnMouseOver_Sigil, QueueQueuer_GUI.OnMouseOverTab, RandomMountUI.OnDropSlotMouseOver, Res.ToggleButtonMouseOver, RoR_SoR.MainTooltip, ScenarioGroupWindow.OnJoinGroupMouseOver, ScenarioGroupWindow.OnLeaveGroupMouseOver, ScenarioGroupWindow.OnMouseoverVisibleButton, SocialWindow.OnMouseOverTab, SocialWindowBuddyList.OnMouseOverTab, SocialWindowBuddyListTabFriends.OnMouseOverPlayerRow, SocialWindowTabFriends.OnMouseOverPlayerRow, SocialWindowTabFriends.OnMouseOverSortButton, SocialWindowTabIgnore.OnMouseOverPlayerRow, SocialWindowTabIgnore.OnMouseOverSortButton, SocialWindowTabSearch.OnMouseOverPlayerRow, SocialWindowTabSearch.OnMouseOverSortButton, TTitan.TTButton.OnMouseover, TTitan.UI.OnMouseOverFAQButton, TTitan.UI.TomeClearNewEntriesMOver, TastyButtonsOptions.OnMenuButtonMouseOver, TastyButtonsOptions.OnMouseOver, TastyButtonsOptions.OnMouseOverButtonSelectWindow, TaxPayer.MouseOverSortButton, TomeTracker.Journal.OnMouseOverTab, TomeWindow.DK_OnMouseoverBookmark, TomeWindow.DK_OnRollOverSortButton, UiModVersionMismatchWindow.OnMouseOverModListSortButton, UiModWindow.OnMouseOverModListSortButton, WBStutterLess.OnMouseOverButton, WSCT.OnMouseOver, emotes.OnMouseOverFilteremotesButton, emotes.OnMouseOverSortButton, followTheLeader.OnMouseOver, nLootLinkGUI.itemMouseOver, wargames.gems.FieldMouseOver, wargames.pairs.CardMouseOver, wargames.ttt.FieldMouseOver, wbLeadHelperChooseIcon.OnMouseOverCreateTooltip, wbLeadHelperConfigTab.OnMouseOverApplyButton, wbLeadHelperConfigTab.OnMouseOverLfgIconsLabel, wbLeadHelperConfigTab.OnMouseOverResetButton, wbLeadHelperConfigTab.OnMouseOverSaveButton, wbLeadHelperMessagesTab.OnMouseOverApplyButton, wbLeadHelperMessagesTab.OnMouseOverResetButton, wbLeadHelperMessagesTab.OnMouseOverSaveButton, zMailModLog.OnItemMouseOver | `` |  |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Warbuilder.OnTabLBU, Warbuilder.DnD, TTitan.UI.SplashTextLClick, Warbuilder.ChangeLvl, ClosetGoblinCharacterWindow.EquipmentLButtonDown, FrameManager.OnLButtonUp, FrameManager.OnLButtonDown, MiracleGrow2.onHClick, QuickWarReport.OnConfirmNoop, AuctionWindowSearchControls.OnLButtonDownCareerSelection, AuctionWindowSearchControls.OnLButtonDownRestrictionSelection, AuctionWindowSearchControls.OnLButtonDownStatisticSelection, AuctionWindowSellControls.ItemSlotLButtonDown, BuffHead.Setup.Layout.BeginResize, CMapWindow.OnResizeBeginLO, CMapWindow.OnResizeBeginLU, CMapWindow.OnResizeBeginRO, CMapWindow.OnResizeBeginRU, CallingSetup.ChooseBind, CallingSetup.ChooseGroup, CallingSetup.ChooseMacro, CallingSetup.ChoosePrios, CallingSetup.ChooseShow, CallingSetup.ChooseTargeting, CallingSetup.ChooseTutorial, CallingSetup.OnTutNextClick, CallingSetup.OnTutPrevClick, CallingSetup.StartBinding1, CrusherConfig.OnResizeBegin, DaemonAssist.HideWindow, DaemonAssist.OnToggleButtonDown, DebugWindow.OnResizeBegin, DevPad.OnResizeBegin, EA_Window_Macro.DetailIconLButtonDown, EA_Window_Macro.SelectionIconLButtonDown, GetStats.RefreshStats, GuildWardenWin.CF1, GuildWardenWin.CF10, GuildWardenWin.CF11, GuildWardenWin.CF12, GuildWardenWin.CF2, GuildWardenWin.CF3, GuildWardenWin.CF4, GuildWardenWin.CF5, GuildWardenWin.CF6, GuildWardenWin.CF7, GuildWardenWin.CF8, GuildWardenWin.CF9, GuildWardenWin.Toggle, HopperConfig.OnResizeBegin, MotionConfig.OnResizeBegin, ObjectInspector.OnResizeBegin, PotionBarFloating.LButtonDown, PureConfig.OnResizeBegin, QuickWarReport.OnConfirmDecline, ScenarioGroupWindow.ToggleMainGroupVisibility, ScenarioGroupWindow.ToggleVisibility, Sequencer.AbilityCursorSwap, Sequencer.Clear, Sequencer.ReLoad, Sequencer.Save, SocialWindowBuddyList.OnLButtonDownTab, SpamBayes.OnResize, TTitan.UI.TomeClearNewEntriesLClick, WCDBConfig.OnResizeBegin, WCDPConfig.OnResizeBegin, Warbuilder.DeleteData, Warbuilder.HotBarLoad, Warbuilder.HotBarSaveData, Warbuilder.HotbarDelete, Warbuilder.LoadData, Warbuilder.SaveData, Warbuilder.ShowCareer, nLootLinkGUI.itemMouseDownL, wargames.gems.FieldLButtonDown, wargames.pairs.CardLButtonDown, wargames.ttt.FieldLButtonDown | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.CombatLogUI_StatsWindow_SortColumnRClick, TortallDPSDetail.ShowColumnMenu, QueueQueuer_GUI.MapButton_OnRButtonUp, , MiracleGrow2.onRClick, CMapWindow.OnScenarioQueueRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnRClick, MiracleGrowLight.switchMode, AuctionStats.OnRButtonUpItem, AuctionWindowSellControls.ItemSlotRButtonUp, AutoBand.HandleIconRightClick, AutoBandWindowTools.OnStatsBreakdownButtonRightClick, AutoBandWindowTools.OnStatsButtonRightClick, ClosetGoblinCharacterWindow.EquipmentRButtonUp, DuffTimer.Options.OnRButtonUpTab, EA_Window_OpenPartyNearby.OnRButtonUpGroupLeaderLine, EA_Window_OpenPartyWorld.OnRButtonUpGroupLeaderLine, Emojii.RunEmojii, Emojii.UI_ChooseIconDialog_OnListRowRButtonUp, GDes.OptionsShow, Ges.OptionsShow, HealGridGuiTabSpellTrack.TabColorChangeBackgroundColor, HealGridGuiUtility.ColorSelectEditBackgroundColor, JunkDumpOptions.Show, KeyBar.RButtonUp, MapMonster.OnMouseRightClickFilter, MapMonster.ToggleFiltersMainMenu, MapPin.UI_ChooseIconDialog_OnListRowRButtonUp, Minmap.ToggleFilterMenu, Minmap.ToggleScenarioGroupWindow, PP.ItemSlotRMouse, PieTracker.RowClicked, PotionBarFloating.RButtonUp, Res.ShowOptions, ResHelp.RemoveChar, RoR_SoR.OnTabRBU, SocialWindowBuddyListTabFriends.OnRButtonUpPlayerRow, SocialWindowTabFriends.OnRButtonUpPlayerRow, SocialWindowTabSearch.OnRButtonUpPlayerRow, TTitan.TTButton.OnRClick, TastyButtonsOptions.OnButtonSelectListRUp, TastyButtonsOptions.OnMenuButtonRUp, TastyButtonsOptions.OnRButtonUp, TaxPayer.RowClicked, WARCommanderConfig.ShowConfigWindow, emotes.ToggleFilteremotes, followTheLeader.OnRButtonUp, minesweep.RButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | AutoBandWindowConfig.OnSettingMouseOverEnd, FrameManager.OnMouseOverEnd, ClosetGoblinCharacterWindow.HideCloakOptions, AutoBandWindowConfig.OnHealerRankReqMouseOverEnd, AutoBandWindowConfig.OnRankReqMouseOverEnd, ClosetGoblinCharacterWindow.HideShowHelm, WarBoard.OnMouseOverEnd, WarBoard.OnMouseOverEndBottom, AutoBandWindowConfig.OnBackfillMouseOverEnd, AutoBandWindowConfig.OnCommonRaceNamesMouseOverEnd, AutoBandWindowConfig.OnExcludeRealmHealerAltSpecMouseOverEnd, AutoBandWindowConfig.OnKickLowRankMouseOverEnd, AutoBandWindowTemplate.OnApplyButtonMouseOverEnd, AutoBandWindowTools.OnDiscordReqMouseOverEnd, AutoBandWindowTools.OnNoMicMouseOverEnd, AutoBandWindowTools.OnStatsBreakdownButtonMouseOverEnd, AutoBandWindowTools.OnStatsButtonMouseOverEnd, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd, KwestorTracker.MouseOverQuestEnd, LibAddonButton.Menu.OnMouseOut, MiracleGrow.onHHoverEnd, MiracleGrowLight.harvestEnd, RandomMountUI.OnDropSlotMouseOverEnd, SocialWindow.OnMouseOverEndTab, SocialWindowTabFriends.OnMouseOverPlayerRowEnd, SocialWindowTabIgnore.OnMouseOverPlayerRowEnd, SocialWindowTabSearch.OnMouseOverPlayerRowEnd | `` |  |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | FrameManager.OnRButtonDown, ClosetGoblinCharacterWindow.EquipmentRButtonDown, CG_ItemRack.EquipmentRButtonDown, EA_Window_Macro.DetailIconRButtonDown, EA_Window_Macro.IconRButtonDown, EA_Window_Macro.SelectionIconRButtonDown, ItemRack.EquipmentRButtonDown, MassMailWindow.OnSlotRButtonDown, Sequencer.AbilityCursorClear, wargames.gems.FieldRButtonDown, wargames.pairs.CardRButtonDown, wargames.ttt.FieldRButtonDown | `flags, x, y` | MEDIUM |
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_LabelCheckButton.Initialize, EA_GenericCheckButton.Initialize, IraConfig.HelpBtnInit | `` |  |
| [OnMouseDrag](../handlers/handler_OnMouseDrag.md) | input | EA_Window_Macro.IconMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag, Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag | `flags, x, y` | MEDIUM |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | QueueQueuer_GUI.MapButton_OnUpdate | `elapsed` | LOW |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | Minmap.AnswerCall, followTheLeader.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseLeave](../handlers/handler_OnMouseLeave.md) | custom | AtlasMap.OnDynamicDirectionButtonMouseLeave, AutoBand.OnIconMouseLeave | `` |  |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | input | CMapWindow.MWheelWholeZoom, MassMailWindow.OnMouseWheel | `x, y, delta` | MEDIUM |
| OnEscapeProcessed | custom | BuddyBind.OnExitBindingMode | `flags, x, y, bForce` | LOW |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `BroadcastEvent`, `ButtonGetCheckButtonFlag`, `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonGetText`, `ButtonSetCheckButtonFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetStayDownFlag`, `ButtonSetText`, `ComboBoxAddMenuItem`, `ComboBoxClearMenuItems`, `ComboBoxExternalOpenMenu`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `CreateWindow`, `CreateWindowFromTemplate`, `Cursor.Clear`, `Cursor.IconOnCursor`, `Cursor.PickUp`, `DestroyWindow`, `DoesWindowExist`, `DynamicImageSetTexture`, `DynamicImageSetTextureSlice`, `GameData.Player.GetRenownRefundCost`, `LabelGetText`, `LabelGetTextDimensions`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `ListBoxSetDisplayOrder`, `Player.GetMoney`, `RegisterEventHandler`, `ScrollWindowUpdateScrollRect`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `SystemData.ActiveWindow.name:find`, `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxGetText:len`, `TextEditBoxSelectAll`, `TextEditBoxSetText`, `UnregisterEventHandler`, `WindowAddAnchor`, `WindowAssignFocus`, `WindowClearAnchors`, `WindowGetAlpha`, `WindowGetAnchorCount`, `WindowGetDimensions`, `WindowGetHandleInput`, `WindowGetId`, `WindowGetParent`, `WindowGetParent:find`, `WindowGetScale`, `WindowGetScreenPosition`, `WindowGetShowing`, `WindowHiderShowClicked`, `WindowRegisterCoreEventHandler`, `WindowRegisterEventHandler`, `WindowSetAlpha`, `WindowSetDimensions`, `WindowSetGameActionData`, `WindowSetHandleInput`, `WindowSetId`, `WindowSetMoving`, `WindowSetScale`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`, `WindowUtils.RemoveFromOpenList`, `WindowUtils.ToggleShowing`

**OnLButtonUp** handlers call: `BroadcastEvent`, `ButtonGetCheckButtonFlag`, `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonGetText`, `ButtonSetCheckButtonFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetStayDownFlag`, `ButtonSetText`, `ComboBoxAddMenuItem`, `ComboBoxClearMenuItems`, `ComboBoxExternalOpenMenu`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `CreateWindow`, `CreateWindowFromTemplate`, `Cursor.Clear`, `Cursor.IconOnCursor`, `Cursor.PickUp`, `DestroyWindow`, `DoesWindowExist`, `DynamicImageSetTexture`, `DynamicImageSetTextureSlice`, `GameData.Player.GetRenownRefundCost`, `LabelGetText`, `LabelGetTextDimensions`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `ListBoxSetDisplayOrder`, `Player.GetMoney`, `RegisterEventHandler`, `ScrollWindowUpdateScrollRect`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `SystemData.ActiveWindow.name:find`, `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxGetText:len`, `TextEditBoxSelectAll`, `TextEditBoxSetText`, `UnregisterEventHandler`, `WindowAddAnchor`, `WindowAssignFocus`, `WindowClearAnchors`, `WindowGetAlpha`, `WindowGetAnchorCount`, `WindowGetDimensions`, `WindowGetHandleInput`, `WindowGetId`, `WindowGetParent`, `WindowGetParent:find`, `WindowGetScale`, `WindowGetScreenPosition`, `WindowGetShowing`, `WindowHiderShowClicked`, `WindowRegisterCoreEventHandler`, `WindowRegisterEventHandler`, `WindowSetAlpha`, `WindowSetDimensions`, `WindowSetGameActionData`, `WindowSetHandleInput`, `WindowSetId`, `WindowSetMoving`, `WindowSetScale`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`, `WindowUtils.RemoveFromOpenList`, `WindowUtils.ToggleShowing`

**OnMouseOver** handlers call: `ButtonGetDisabledFlag`, `ComboBoxGetSelectedMenuItem`, `ComboBoxIsMenuOpen`, `LabelGetText`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowGetShowing`, `WindowSetAlpha`, `WindowSetShowing`, `WindowUtils.OnMouseOverButton`

**OnLButtonDown** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonSetPressedFlag`, `CreateWindow`, `Cursor.Clear`, `Cursor.IconOnCursor`, `Cursor.PickUp`, `DoesWindowExist`, `DynamicImageSetTexture`, `ListBoxGetDataIndex`, `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowGetId`, `WindowGetParent`, `WindowGetShowing`, `WindowSetMoving`, `WindowSetParent`, `WindowSetShowing`, `WindowSetTintColor`, `WindowUtils.BeginResize`

**OnRButtonUp** handlers call: `DynamicImageSetTexture`, `DynamicImageSetTextureSlice`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetId`, `WindowGetParent`, `WindowGetShowing`, `WindowSetAlpha`, `WindowSetLayer`, `WindowSetShowing`, `WindowSetTintColor`, `WindowUtils.ToggleShowing`

**OnMouseOverEnd** handlers call: `WindowSetShowing`

**OnRButtonDown** handlers call: `WindowGetId`

**OnInitialize** handlers call: `CreateWindow`, `RegisterEventHandler`

**OnMouseDrag** handlers call: `Cursor.IconOnCursor`, `Cursor.PickUp`, `WindowGetId`

**OnEscapeProcessed** handlers call: `ButtonSetCheckButtonFlag`, `ButtonSetPressedFlag`, `ButtonSetStayDownFlag`, `WindowUnregisterCoreEventHandler`, `WindowUnregisterEventHandler`

## Common Inherits

- DefaultButton
- EA_Button_Default
- EA_Button_DefaultCheckBox
- EA_Button_DefaultMinus
- EA_Button_DefaultResizeable
- EA_Button_DefaultText
- EA_Button_DefaultWindowClose
- EA_Button_ListSort
- EA_Button_Tab
- EA_Window_MacroIconButton
- MacroIconSelectionWindowIconButton
- Template_Mastery_Button

## Common Parent Elements

- [Windows](element_Windows.md) — 2407× (HIGH)
- [Window](element_Window.md) — 8× (MEDIUM)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 1954× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 1540× (HIGH)
- [Size](element_Size.md) — 1023× (HIGH)
- [Windows](element_Windows.md) — 131× (HIGH)
- [TexSlices](element_TexSlices.md) — 129× (HIGH)
- [TextColors](element_TextColors.md) — 85× (HIGH)
- [TextOffset](element_TextOffset.md) — 75× (HIGH)
- [TexCoords](element_TexCoords.md) — 47× (HIGH)
- [ResizeImages](element_ResizeImages.md) — 41× (HIGH)
- [OverlayOffset](element_OverlayOffset.md) — 31× (HIGH)
- [Color](element_Color.md) — 27× (HIGH)
- [OverlaySize](element_OverlaySize.md) — 27× (HIGH)
- [OverlayTexCoords](element_OverlayTexCoords.md) — 24× (HIGH)
- [TexDims](element_TexDims.md) — 18× (HIGH)
- [TintColor](element_TintColor.md) — 12× (HIGH)
- [Sounds](element_Sounds.md) — 9× (MEDIUM)
- [Text](element_Text.md) — 3× (MEDIUM)
- [Anchor](element_Anchor.md) — 2× (LOW)
- [OverlayTexSlices](element_OverlayTexSlices.md) — 2× (LOW)
- [AnimatedImages](element_AnimatedImages.md) — 1× (LOW)
- [EventHandler](element_EventHandler.md) — 1× (LOW)
- [Eventhandlers](element_Eventhandlers.md) — 1× (LOW)

## Common Template Bases

- AbilitiesWindowButtonDef
- ActionButtonInput
- Aggro_Button_Template
- AuctionWindowContextMenuItem
- AuctionWindowSortButton
- AuctionWindowTabButton
- AuctionWindowTextButton
- AuraIconButton
- AuraSharesSortButton
- AuraTabButtonTemplate
- AuraWindowButton
- AuraWindowSortButton
- BlackBookWindowSortButton
- BlackBookWindow_Button_EntryRow
- BuffHeadLayoutBottomLeftButton
- BuffHeadLayoutBottomRightButton
- BuffHeadLayoutCornerResizeButton
- BuffHeadLayoutHorizontalButton
- BuffHeadLayoutResizeButton
- BuffHeadLayoutTopLeftButton
- BuffHeadLayoutTopRightButton
- BuffHeadLayoutVerticalButton
- BuffHeadSetupEffectCacheSortTemplate
- CG_ItemRackEquipmentButton
- CastSequenceBuilderSortTemplate
- CastSequenceButtonIcon
- CharacterWindowDefaultButton
- ClosetGoblinDefaultButton
- ClosetGoblinEquipmentButton
- CoreWindowResizeButton
- CoreWindowResizeButtonBottomRight
- CraftingWillardMinus
- CraftingWillardPlus
- DTC_TARGETS_SortButton
- DefaultButton
- DefaultButtonSmall
- DefaultIconButton
- EA_Button_BottomTab
- EA_Button_Default
- EA_Button_DefaultBigLeftArrow
- EA_Button_DefaultBigRightArrow
- EA_Button_DefaultChatScrollToBottom
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultIcon
- EA_Button_DefaultIconFrame
- EA_Button_DefaultIconFrame_Large
- EA_Button_DefaultIconFrame_Small
- EA_Button_DefaultLeftArrow
- EA_Button_DefaultListRow
- EA_Button_DefaultListSort
- EA_Button_DefaultMenuButton
- EA_Button_DefaultMinus
- EA_Button_DefaultPlus
- EA_Button_DefaultResizableComboBoxSelected
- EA_Button_DefaultResizableSmallComboBoxSelected
- EA_Button_DefaultResizeable
- EA_Button_DefaultRightArrow
- EA_Button_DefaultSmallSquare
- EA_Button_DefaultText
- EA_Button_DefaultTimerFrame
- EA_Button_DefaultTimerFrame_Small
- EA_Button_DefaultToggleCircle
- EA_Button_DefaultUp
- EA_Button_DefaultWindowClose
- EA_Button_HUDMenuButton
- EA_Button_ListSort
- EA_Button_OpenPartyTab
- EA_Button_ResizeIconFrame
- EA_Button_ResizeIconFrame_NoNormalImage
- EA_Button_SocialWindowRowTemplate
- EA_Button_Tab
- EA_CheckButtonButton
- EA_FilterMenuButtonTemplate
- EA_ScrollBar_ChatDownArrowButton
- EA_ScrollBar_ChatUpArrowButton
- EA_ScrollBar_DefaultDownArrowButton
- EA_ScrollBar_DefaultUpArrowButton
- EA_Template_OpenPartyNearbySortButton
- EA_Templates_Color_Picker_Button
- EA_Templates_MailWindowTabButton
- EA_Templates_SpecializationChange
- EA_Window_MacroDetailIconButton
- EA_Window_MacroIconButton
- EZCraftX_Button_DefaultMenuButton
- EmojiiChooseIconDialogList_IconButton
- EnemyChooseIconDialogList_IconButton
- FriendsCommandButtonTemplate
- FriendsSortButtonTemplate
- GDesTabButton_Template
- GDes_SoundButton
- GesTabButtonTemplate
- Ges_SoundButton
- HGG_SpellListSortButtonTemplate
- HGG_TabMouseClickSortButtonTemplate
- HGG_TabSpellTrackSortButtonTemplate
- HGG_TabSpellTrackTabTemplate
- HelpRowLabel
- IgnoresSortButtonTemplate
- IraConfigHelpBtn
- ItemRackEquipmentButton
- ItemWindowSlotButton
- JoinScenarioGroupButton
- KeyBarButton
- KwestorGui_TabAreaSortButtonTemplate
- KwestorGui_TabTemplate
- LPETTabButton
- MacroIconSelectionWindowIconButton
- MapFilterContextMenuChoice
- MapMonsterEditorWindowButtonDefault
- MapMonsterPinTypeEditorWindowButtonDefault
- MapPinButton
- MapPinButton_NoBG
- MapPinChooseIconDialogList_IconButton
- MapPinNumberButtonTemplate
- MinusButton
- MiracleGrow2Button
- MiracleGrow2HarvestRptButton
- MiracleGrow2Repeat
- MiracleGrow2SoundButton
- MiracleGrowButton
- MiracleGrowLightButton
- ModWindowSortButton
- Motion_Button_DefaultMenuButton
- NBSBActionButtonInput
- NBSBChecksAddButton
- PieTrackerSortButtonTemplate
- PlusButton
- PotionBarWindowButton
- QuickTacticSwitchWindowSortButton
- Raid_Button_Template
- SearchSortButtonTemplate
- SequencerButtonIcon
- SettingsWindowTabbedButton
- ShiniesAuctionsUI_SortButton
- ShiniesBrowseUI_ResultsSortButton
- ShiniesBrowseUI_SearchesSortButton
- ShiniesConfigPrice_DecreaseButton
- ShiniesConfigPrice_IncreaseButton
- ShiniesPostUI_SortButton
- Shinies_Button_DefaultListSort
- Shinies_Button_ListSort
- Shinies_Default_Button_ClearMediumFont
- Shinies_IconButton
- Shinies_IconButton_Overlay
- SocialWindowBuddyListTabTemplate
- SocialWindowDefaultLargeSquareMinus
- SocialWindowDefaultLargeSquarePlus
- SocialWindowTabTemplate
- SortButton
- SpamBayesTemplateButtonClose
- SpamBayesTemplateButtonText
- TChatButton
- TChatTabButton
- TOKQuestTextButton
- TOKTextButton
- TRollButton
- TRollItemButton
- TTitanUITextButton
- TabButton
- TabButtonTemplate
- TacticsSetMenuButton
- TastyButtonCustomRectInput
- TastyButtonCustomRoundInput
- TastyButtonRoundInput
- TastyButtonsOptionsWindowTabTemplate
- TaxPayerSortButtonTemplate
- Template_Ability_Button
- Template_CoreMoraleButton
- Template_CoreTacticButton
- Template_Mastery_Button
- Template_Morale_Button
- Template_Tactic_Button
- TestMoraleTemplate
- TestTacticTemplate
- TomePairingElves
- TomePairingEvC
- TomePairingGvD
- TomeTracker_Button_SagaListingRow
- TomeTracker_JournalTabButton
- WSCTButtonTemplate
- WSCTTabButton
- WarbuilderTabTemplate
- WargamesGemsItemButtonTemplate
- WargamesPairsItemButtonTemplate
- WargamesTTTItemButtonTemplate
- XpStatusTabButton
- nLootLinkSortingHeaderTemplate
- wbLeadHelperChooseIconDialogList_IconButton
- zMM_Default_Mini_Button
- zMailModOptionsCheckbox
- zMailModOptionsRadioButton
- zMailModParchmentButton


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 94% | EA_Button_DefaultMenuButton, EA_Button_DefaultWindowClose, EA_Button_Tab, EA_Button_DefaultResizeable, ... |
| `id` | optional | 21% | 1, 2, 3, 4, ... |
| `handleinput` | optional | 19% | false, true |
| `textalign` | optional | 18% | left, center, leftcenter, bottomright, ... |
| `layer` | optional | 16% | overlay, popup, default, secondary, ... |
| `font` | optional | 11% | font_default_text, font_clear_default, font_chat_text, font_clear_medium_bold, ... |
| `backgroundtexture` | optional | 7% | shared_01, EA_Training_Specialization, EA_Abilities01_d5, EA_HUD_01, ... |
| `highlighttexture` | optional | 6% | shared_01, EA_Training_Specialization, EA_Abilities01_d5, EA_HUD_01, ... |
| `texturescale` | optional | 3% | 0.8, 2.0, 1, 0.74, ... |
| `drawchildrenfirst` | optional | 3% | true, false |
| `textureScale` | optional | 2% | 0.62, 0.87, 0.85, 0.75, ... |
| `autoresize` | optional | 2% | false, true |
| `textAutoFitMinScale` | optional | 1% | 1.0, 0.5, .5, 0.2 |
| `warnOnTextCropped` | optional | 1% | false |
| `overlaytexture` | optional | 1% | shared_01, EA_HUD_01, EA_VictoryPoints01_32b, spambayes_tint_scrollbar |
| `alpha` | optional | 1% | 0.0, 0, 1 |
| `draganddrop` | optional | 1% | true, false |
| `overlayhighlighttexture` | optional | 1% | shared_01, EA_HUD_01, spambayes_tint_scrollbar |
| `savesettings` | optional | 0% | false, true |
| `gameactionbutton` | optional | 0% | right, left |
| `texture` | optional | 0% | GWicon |
| `wordwrap` | optional | 0% | true |
| `maxchars` | optional | 0% | 256, 150, 80 |
| `popable` | optional | 0% | false, true |
| `sticky` | optional | 0% | false |
| `mirrorTexCoords` | optional | 0% | true |
| `autoresizewidth` | optional | 0% | true, false |
| `movable` | optional | 0% | false |
| `localscriptvars` | optional | 0% | true |
| `virtual` | optional | 0% | true |
| `warnOnTextAutoFit` | optional | 0% | false |
| `overlayhightlighttexture` | optional | 0% | EA_VictoryPoints01_32b |
| `text` | optional | 0% | Save |
| `Layer` | optional | 0% | popup |
| `hilighttexture` | optional | 0% |  |
| `parent` | optional | 0% | Root |
| `scale` | optional | 0% | .5 |
| `visible` | optional | 0% | true |
| `w` | optional | 0% | false |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 1954 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 1540 times as an unnamed child.

### [Size](element_Size.md)

Observed 1023 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 131 times as an unnamed child.

### [TexSlices](element_TexSlices.md)

Observed 129 times as an unnamed child.

### [TextColors](element_TextColors.md)

Observed 85 times as an unnamed child.

### [TextOffset](element_TextOffset.md)

Observed 75 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5, 6, 4, 0 |
| `y` | **required** | 5, 12, 2, 3 |
### [TexCoords](element_TexCoords.md)

Observed 47 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 113 |
| `y` | optional | 0, 163, 51, 90 |
### [ResizeImages](element_ResizeImages.md)

Observed 41 times as an unnamed child.

### [OverlayOffset](element_OverlayOffset.md)

Observed 31 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 93, 193, 148, 223 |
| `y` | **required** | 0, 1, 7, 9 |
### [Color](element_Color.md)

Observed 27 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 255, 0, 73, 225 |
| `g` | **required** | 25, 253, 218, 255 |
| `r` | **required** | 255, 253, 155, 245 |
| `a` | optional | 255, 0.8, 0, 128 |
### [OverlaySize](element_OverlaySize.md)

Observed 27 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 27, 14, 18, 11 |
| `y` | **required** | 28, 15, 13, 19 |
### [OverlayTexCoords](element_OverlayTexCoords.md)

Observed 24 times as an unnamed child.

### [TexDims](element_TexDims.md)

Observed 18 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 64, 256, 32, 48 |
| `y` | **required** | 64, 256, 32, 48 |
### [TintColor](element_TintColor.md)

Observed 12 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.8, 175 |
### [Sounds](element_Sounds.md)

Observed 9 times as an unnamed child.

### [Text](element_Text.md)

Observed 3 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `text` | optional | Killer Settings, Changes apply immediately. Zone K/D history uses 0 for unlimited saved zone leaderboards., Display, Personal |
| `alignment` | optional | left |
### [Anchor](element_Anchor.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, topright, bottomright |
| `relativePoint` | **required** | center, topleft, topright, bottomright |
| `relativeTo` | **required** | Root, $parent, $parentSliderLabel, $parentSliderValueLabel |
| `relativeto` | optional | $parentTitle, $parentPrint, $parentAlert, $parentHyperlink |
| `layer` | optional | secondary, overlay |
| `parent` | optional | Root, $parent |
| `relateiveTo` | optional | $parentDevPadCode, $parentObjectEditBox |
| `textalign` | optional | center |
| `handleinput` | optional | false |
| `relative` | optional | $parent |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
| `input` | optional | numbers |
| `realtivePoint` | optional | center |
| `realtiveTo` | optional | $parentBackground |
### [OverlayTexSlices](element_OverlayTexSlices.md)

Observed 2 times as an unnamed child.

### [AnimatedImages](element_AnimatedImages.md)

Observed 1 times as an unnamed child.

### [EventHandler](element_EventHandler.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `event` | **required** | OnShown, OnHidden, OnLButtonUp, OnSelChanged |
| `function` | **required** | APAGui.OnShown, APAGui.OnHidden, APAGui.Hide, APAGui.OnTabButtonUp |
### [Eventhandlers](element_Eventhandlers.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [Button](element_Button.md)
- [Anchor](element_Anchor.md) (structural, 2×, LOW)
  - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
    - (cycle)
- [Anchors](element_Anchors.md) (structural, 1954×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
      - (cycle)
- [AnimatedImages](element_AnimatedImages.md) (structural, 1×, LOW)
  - [Normal](element_Normal.md) (structural, 1×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
- [Color](element_Color.md) (structural, 27×, HIGH)
- [EventHandler](element_EventHandler.md) (structural, 1×, LOW)
- [EventHandlers](element_EventHandlers.md) (structural, 1540×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
- [Eventhandlers](element_Eventhandlers.md) (structural, 1×, LOW)
  - [EventHandler](element_EventHandler.md) (structural, 2×, HIGH)
- [OverlayOffset](element_OverlayOffset.md) (structural, 31×, HIGH)
- [OverlaySize](element_OverlaySize.md) (structural, 27×, HIGH)
- [OverlayTexCoords](element_OverlayTexCoords.md) (structural, 24×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 20×, HIGH)
  - [Normal](element_Normal.md) (structural, 24×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 24×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 24×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 24×, HIGH)
- [OverlayTexSlices](element_OverlayTexSlices.md) (structural, 2×, LOW)
  - [Normal](element_Normal.md) (structural, 2×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 2×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 2×, HIGH)
- [ResizeImages](element_ResizeImages.md) (structural, 41×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 23×, HIGH)
  - [Normal](element_Normal.md) (structural, 25×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 41×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 26×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 27×, HIGH)
- [Size](element_Size.md) (structural, 1023×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
- [Sounds](element_Sounds.md) (structural, 9×, MEDIUM)
  - [Sound](element_Sound.md) (structural, 20×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 47×, HIGH)
  - [Bottom](element_Bottom.md) (structural, 2×, LOW)
  - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
  - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
  - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
  - [Left](element_Left.md) (structural, 33×, HIGH)
  - [Middle](element_Middle.md) (structural, 41×, HIGH)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
  - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
  - [Normal](element_Normal.md) (structural, 47×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
  - [Right](element_Right.md) (structural, 33×, HIGH)
  - [Top](element_Top.md) (structural, 2×, LOW)
  - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
  - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
- [TexDims](element_TexDims.md) (structural, 18×, HIGH)
- [TexSlices](element_TexSlices.md) (structural, 129×, HIGH)
  - [Bottom](element_Bottom.md) (structural, 2×, LOW)
  - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
  - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
  - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
  - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
  - [Left](element_Left.md) (structural, 3×, MEDIUM)
  - [Middle](element_Middle.md) (structural, 10×, HIGH)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
  - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
  - [Normal](element_Normal.md) (structural, 120×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
  - [Right](element_Right.md) (structural, 3×, MEDIUM)
  - [Top](element_Top.md) (structural, 2×, LOW)
  - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
  - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
- [Text](element_Text.md) (structural, 3×, MEDIUM)
- [TextColors](element_TextColors.md) (structural, 85×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 76×, HIGH)
  - [DisabledPressed](element_DisabledPressed.md) (structural, 26×, HIGH)
  - [Normal](element_Normal.md) (structural, 79×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 83×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 76×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 83×, HIGH)
- [TextOffset](element_TextOffset.md) (structural, 75×, HIGH)
- [TintColor](element_TintColor.md) (structural, 12×, HIGH)
- [Windows](element_Windows.md) (structural, 131×, HIGH)
  - [ActionButtonGroup](element_ActionButtonGroup.md) (named, 3×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 3×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
  - [AnimatedImage](element_AnimatedImage.md) (named, 39×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 38×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [AnimFrames](element_AnimFrames.md) (structural, 19×, HIGH)
      - [AnimFrame](element_AnimFrame.md) (structural, 222×, HIGH)
    - [Size](element_Size.md) (structural, 32×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 1×, LOW)
    - [Windows](element_Windows.md) (structural, 2×, MEDIUM)
      - (cycle)
  - [Button](element_Button.md) (named, 2407×, HIGH)
    - (cycle)
  - [CircleImage](element_CircleImage.md) (named, 56×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 54×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 51×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 37×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 15×, HIGH)
    - [Windows](element_Windows.md) (structural, 2×, LOW)
      - (cycle)
  - [ColorPicker](element_ColorPicker.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [ColorSize](element_ColorSize.md) (structural, 1×, HIGH)
    - [ColorSpacing](element_ColorSpacing.md) (structural, 1×, HIGH)
    - [ColorTexCoords](element_ColorTexCoords.md) (structural, 1×, HIGH)
    - [ColorTexDims](element_ColorTexDims.md) (structural, 1×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
  - [ComboBox](element_ComboBox.md) (named, 689×, HIGH)
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
  - [CooldownDisplay](element_CooldownDisplay.md) (named, 10×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 4×, MEDIUM)
      - (cycle)
  - [DynamicImage](element_DynamicImage.md) (named, 1288×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 1187×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 1×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 155×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 965×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sounds](element_Sounds.md) (structural, 1×, LOW)
      - [Sound](element_Sound.md) (structural, 20×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 76×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 156×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 374×, HIGH)
    - [Windows](element_Windows.md) (structural, 15×, HIGH)
      - (cycle)
  - [EditBox](element_EditBox.md) (named, 416×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 380×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 2×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 269×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 355×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 63×, HIGH)
    - [Windows](element_Windows.md) (structural, 3×, MEDIUM)
      - (cycle)
  - [FullResizeImage](element_FullResizeImage.md) (named, 450×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 2×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 409×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 2×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 5×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 113×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 30×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
      - [Middle](element_Middle.md) (structural, 30×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 19×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
    - [TexSlices](element_TexSlices.md) (structural, 11×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
      - [Left](element_Left.md) (structural, 3×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 10×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
      - [Normal](element_Normal.md) (structural, 120×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
      - [Right](element_Right.md) (structural, 3×, MEDIUM)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 118×, HIGH)
    - [Windows](element_Windows.md) (structural, 4×, MEDIUM)
      - (cycle)
  - [HorizontalResizeImage](element_HorizontalResizeImage.md) (named, 90×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 48×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 49×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 43×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
      - [Middle](element_Middle.md) (structural, 30×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 36×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexSlices](element_TexSlices.md) (structural, 7×, MEDIUM)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
      - [Left](element_Left.md) (structural, 3×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 10×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
      - [Normal](element_Normal.md) (structural, 120×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
      - [Right](element_Right.md) (structural, 3×, MEDIUM)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 20×, HIGH)
    - [Windows](element_Windows.md) (structural, 3×, MEDIUM)
      - (cycle)
  - [Label](element_Label.md) (named, 4814×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 2×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 4631×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 1898×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 330×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 4173×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sounds](element_Sounds.md) (structural, 1×, LOW)
      - [Sound](element_Sound.md) (structural, 20×, HIGH)
    - [Text](element_Text.md) (structural, 69×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 30×, HIGH)
    - [Windows](element_Windows.md) (structural, 14×, HIGH)
      - (cycle)
    - [color](element_color.md) (structural, 1×, LOW)
  - [ListBox](element_ListBox.md) (named, 113×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 110×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 16×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [ListData](element_ListData.md) (structural, 110×, HIGH)
      - [ListColumns](element_ListColumns.md) (structural, 74×, HIGH)
        - [ListColumn](element_ListColumn.md) (structural, 192×, HIGH)
    - [Size](element_Size.md) (structural, 76×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [LogDisplay](element_LogDisplay.md) (named, 3×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 3×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 3×, HIGH)
      - (cycle)
  - [MapDisplay](element_MapDisplay.md) (named, 8×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 8×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 8×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 3×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
  - [NifDisplay](element_NifDisplay.md) (named, 50×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Rotation](element_Rotation.md) (structural, 4×, MEDIUM)
    - [Size](element_Size.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Translation](element_Translation.md) (structural, 50×, HIGH)
  - [PageWindow](element_PageWindow.md) (named, 4×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 4×, HIGH)
      - (cycle)
  - [ScrollWindow](element_ScrollWindow.md) (named, 61×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 43×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 21×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 53×, HIGH)
      - (cycle)
  - [SliderBar](element_SliderBar.md) (named, 225×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 219×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 205×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 186×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 6×, MEDIUM)
      - (cycle)
  - [StatusBar](element_StatusBar.md) (named, 32×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 32×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 19×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
  - [VerticalResizeImage](element_VerticalResizeImage.md) (named, 27×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 14×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 8×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
      - [Middle](element_Middle.md) (structural, 30×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 5×, MEDIUM)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexSlices](element_TexSlices.md) (structural, 3×, MEDIUM)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
      - [Left](element_Left.md) (structural, 3×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 10×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
      - [Normal](element_Normal.md) (structural, 120×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
      - [Right](element_Right.md) (structural, 3×, MEDIUM)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 4×, MEDIUM)
  - [VerticalScrollbar](element_VerticalScrollbar.md) (named, 62×, HIGH)
    - [ActiveZoneOffset](element_ActiveZoneOffset.md) (structural, 2×, LOW)
    - [Anchors](element_Anchors.md) (structural, 56×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [DownOffset](element_DownOffset.md) (structural, 2×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 7×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 47×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [ThumbOffset](element_ThumbOffset.md) (structural, 2×, LOW)
    - [UpOffset](element_UpOffset.md) (structural, 2×, LOW)
  - [Window](element_Window.md) (named, 3695×, HIGH)
    - [Button](element_Button.md) (named, 8×, MEDIUM)
      - (cycle)
    - [ComboBox](element_ComboBox.md) (named, 4×, MEDIUM)
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
    - [DynamicImage](element_DynamicImage.md) (named, 1×, LOW)
      - [Anchor](element_Anchor.md) (structural, 1×, LOW)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
      - [Anchors](element_Anchors.md) (structural, 1187×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 1×, LOW)
      - [EventHandlers](element_EventHandlers.md) (structural, 155×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Size](element_Size.md) (structural, 965×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Sounds](element_Sounds.md) (structural, 1×, LOW)
        - [Sound](element_Sound.md) (structural, 20×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 76×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
        - [Left](element_Left.md) (structural, 33×, HIGH)
        - [Middle](element_Middle.md) (structural, 41×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
        - [Normal](element_Normal.md) (structural, 47×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
        - [Right](element_Right.md) (structural, 33×, HIGH)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
      - [TexDims](element_TexDims.md) (structural, 156×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 374×, HIGH)
      - [Windows](element_Windows.md) (structural, 15×, HIGH)
        - (cycle)
    - [FullResizeImage](element_FullResizeImage.md) (named, 9×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 2×, LOW)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
      - [Anchors](element_Anchors.md) (structural, 409×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 2×, LOW)
      - [EventHandlers](element_EventHandlers.md) (structural, 5×, MEDIUM)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Size](element_Size.md) (structural, 113×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Sizes](element_Sizes.md) (structural, 30×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
        - [Middle](element_Middle.md) (structural, 30×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 19×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
        - [Left](element_Left.md) (structural, 33×, HIGH)
        - [Middle](element_Middle.md) (structural, 41×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
        - [Normal](element_Normal.md) (structural, 47×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
        - [Right](element_Right.md) (structural, 33×, HIGH)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
      - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
      - [TexSlices](element_TexSlices.md) (structural, 11×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
        - [Left](element_Left.md) (structural, 3×, MEDIUM)
        - [Middle](element_Middle.md) (structural, 10×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
        - [Normal](element_Normal.md) (structural, 120×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
        - [Right](element_Right.md) (structural, 3×, MEDIUM)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 118×, HIGH)
      - [Windows](element_Windows.md) (structural, 4×, MEDIUM)
        - (cycle)
    - [Label](element_Label.md) (named, 15×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 2×, LOW)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
      - [Anchors](element_Anchors.md) (structural, 4631×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 1898×, HIGH)
      - [EventHandlers](element_EventHandlers.md) (structural, 330×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Size](element_Size.md) (structural, 4173×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Sounds](element_Sounds.md) (structural, 1×, LOW)
        - [Sound](element_Sound.md) (structural, 20×, HIGH)
      - [Text](element_Text.md) (structural, 69×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 30×, HIGH)
      - [Windows](element_Windows.md) (structural, 14×, HIGH)
        - (cycle)
      - [color](element_color.md) (structural, 1×, LOW)
    - [SliderBar](element_SliderBar.md) (named, 2×, LOW)
      - [Anchors](element_Anchors.md) (structural, 219×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [EventHandlers](element_EventHandlers.md) (structural, 205×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Size](element_Size.md) (structural, 186×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Windows](element_Windows.md) (structural, 6×, MEDIUM)
        - (cycle)
    - [Window](element_Window.md) (named, 5×, MEDIUM)
      - (cycle)
    - [Anchor](element_Anchor.md) (structural, 16×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 2735×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 773×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 1752×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sounds](element_Sounds.md) (structural, 2×, LOW)
      - [Sound](element_Sound.md) (structural, 20×, HIGH)
    - [Visual](element_Visual.md) (structural, 1×, LOW)
      - [Color](element_Color.md) (structural, 1×, HIGH)
    - [Windows](element_Windows.md) (structural, 1498×, HIGH)
      - (cycle)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowSetShowing` | ui | 462 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `WindowGetId` | ui | 286 | OnLButtonDown, OnLButtonUp, OnMouseDrag, OnMouseOver, OnRButtonDown, OnRButtonUp |
| `ButtonSetPressedFlag` | ui | 265 | OnEscapeProcessed, OnLButtonDown, OnLButtonUp |
| `ComboBoxGetSelectedMenuItem` | ui | 142 | OnLButtonUp, OnMouseOver |
| `TextEditBoxGetText` | ui | 114 | OnLButtonDown, OnLButtonUp |
| `WindowGetShowing` | ui | 101 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `LabelSetText` | ui | 90 | OnLButtonUp, OnMouseOver |
| `ButtonGetPressedFlag` | ui | 87 | OnLButtonDown, OnLButtonUp |
| `SystemData.ActiveWindow.name:match` | data | 78 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ComboBoxSetSelectedMenuItem` | ui | 73 | OnLButtonUp |
| `ButtonGetDisabledFlag` | ui | 72 | OnLButtonDown, OnLButtonUp, OnMouseOver |
| `TextEditBoxSetText` | ui | 70 | OnLButtonDown, OnLButtonUp |
| `WindowGetParent` | ui | 52 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `WindowAddAnchor` | ui | 31 | OnLButtonDown, OnLButtonUp |
| `WindowSetDimensions` | ui | 31 | OnLButtonUp |
| `WindowSetTintColor` | ui | 31 | OnLButtonDown, OnLButtonUp, OnRButtonUp |
| `DoesWindowExist` | ui | 27 | OnLButtonDown, OnLButtonUp |
| `LabelSetTextColor` | ui | 27 | OnLButtonUp, OnMouseOver |
| `ComboBoxClearMenuItems` | ui | 24 | OnLButtonUp |
| `DestroyWindow` | ui | 24 | OnLButtonUp |
| `ButtonSetDisabledFlag` | ui | 22 | OnLButtonUp |
| `ListBoxGetDataIndex` | ui | 22 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `WindowClearAnchors` | ui | 22 | OnLButtonDown, OnLButtonUp |
| `SliderBarGetCurrentPosition` | ui | 20 | OnLButtonUp |
| `WindowAssignFocus` | ui | 20 | OnLButtonUp |
| `WindowSetAlpha` | ui | 20 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `SliderBarSetCurrentPosition` | ui | 19 | OnLButtonUp |
| `WindowUtils.ToggleShowing` | ui | 16 | OnLButtonUp, OnRButtonUp |
| `WindowUtils.BeginResize` | ui | 13 | OnLButtonDown |
| `Cursor.Clear` | ui | 12 | OnLButtonDown, OnLButtonUp |
| `Cursor.IconOnCursor` | data | 12 | OnLButtonDown, OnLButtonUp, OnMouseDrag |
| `DynamicImageSetTexture` | ui | 12 | OnLButtonDown, OnLButtonUp, OnRButtonUp |
| `WindowSetScale` | ui | 12 | OnLButtonUp |
| `ButtonSetText` | ui | 11 | OnLButtonUp |
| `CreateWindow` | ui | 11 | OnInitialize, OnLButtonDown, OnLButtonUp |
| `ListBoxSetDisplayOrder` | ui | 11 | OnLButtonUp |
| `ButtonGetText` | ui | 9 | OnLButtonUp |
| `SystemData.MouseOverWindow.name:match` | data | 9 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ButtonSetCheckButtonFlag` | ui | 8 | OnEscapeProcessed, OnLButtonUp |
| `ComboBoxAddMenuItem` | ui | 8 | OnLButtonUp |
| `LabelGetText` | ui | 8 | OnLButtonUp, OnMouseOver |
| `WindowSetLayer` | ui | 8 | OnRButtonUp |
| `WindowSetHandleInput` | ui | 7 | OnLButtonUp |
| `ButtonGetCheckButtonFlag` | ui | 6 | OnLButtonUp |
| `ComboBoxGetSelectedText` | ui | 6 | OnLButtonUp |
| `Cursor.PickUp` | data | 6 | OnLButtonDown, OnLButtonUp, OnMouseDrag |
| `WindowGetAlpha` | ui | 6 | OnLButtonUp |
| `WindowGetDimensions` | ui | 6 | OnLButtonUp, OnMouseOver |
| `RegisterEventHandler` | event | 5 | OnInitialize, OnLButtonUp |
| `ScrollWindowUpdateScrollRect` | ui | 4 | OnLButtonUp |
| `SystemData.ActiveWindow.name:find` | data | 4 | OnLButtonUp |
| `WindowGetHandleInput` | ui | 4 | OnLButtonUp |
| `BroadcastEvent` | event | 3 | OnLButtonUp |
| `CreateWindowFromTemplate` | ui | 3 | OnLButtonUp |
| `GameData.Player.GetRenownRefundCost` | data | 3 | OnLButtonUp |
| `Player.GetMoney` | data | 3 | OnLButtonUp |
| `TextEditBoxGetText:len` | ui | 3 | OnLButtonUp |
| `UnregisterEventHandler` | event | 3 | OnLButtonUp |
| `WindowGetScreenPosition` | ui | 3 | OnLButtonUp, OnMouseOver |
| `WindowRegisterEventHandler` | event | 3 | OnLButtonUp |
| `WindowSetId` | ui | 3 | OnLButtonUp |
| `WindowStartAlphaAnimation` | ui | 3 | OnLButtonUp |
| `WindowUnregisterEventHandler` | ui | 3 | OnEscapeProcessed |
| `ButtonSetStayDownFlag` | ui | 2 | OnEscapeProcessed, OnLButtonUp |
| `DynamicImageSetTextureSlice` | ui | 2 | OnLButtonUp, OnRButtonUp |
| `TextEditBoxSelectAll` | ui | 2 | OnLButtonUp |
| `WindowGetParent:find` | ui | 2 | OnLButtonUp |
| `WindowHiderShowClicked` | ui | 2 | OnLButtonUp |
| `WindowSetGameActionData` | ui | 2 | OnLButtonUp |
| `WindowSetMoving` | ui | 2 | OnLButtonDown, OnLButtonUp |
| `ComboBoxExternalOpenMenu` | ui | 1 | OnLButtonUp |
| `ComboBoxIsMenuOpen` | ui | 1 | OnMouseOver |
| `LabelGetTextDimensions` | ui | 1 | OnLButtonUp |
| `WindowGetAnchorCount` | ui | 1 | OnLButtonUp |
| `WindowGetScale` | ui | 1 | OnLButtonUp |
| `WindowRegisterCoreEventHandler` | event | 1 | OnLButtonUp |
| `WindowSetParent` | ui | 1 | OnLButtonDown |
| `WindowUnregisterCoreEventHandler` | ui | 1 | OnEscapeProcessed |
| `WindowUtils.OnMouseOverButton` | ui | 1 | OnMouseOver |
| `WindowUtils.RemoveFromOpenList` | ui | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnEscapeProcessed

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | boolean | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
| 3 | `bForce` | unknown | unknown |
### OnInitialize

Confidence: HIGH

### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseDrag

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseLeave

Confidence: LOW

### OnMouseOver

Confidence: HIGH

### OnMouseOverEnd

Confidence: HIGH

### OnMouseWheel

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `x` | number | mouse_x |
| 1 | `y` | number | mouse_y |
| 2 | `delta` | number | wheel_delta |
### OnRButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnUpdate

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
## Lua Functions Manipulating This Type

- DaemonAssist.UpdateWindow
- ScenarioRefresh.Initialize
- RvRStatsRvRTab.Init
- TacticSetNames.local.SetLabels
- TastyButtons.Button_State_TypeSwap
- nLootLinkGUIIntegration.initButtons
- CMapWindow.DeactivateRallyCall
- SocialWindowTabSearch.Initialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- TortallDPSMeter.Initialize
- DammazKron.local.InitializeBookmark
- Enemy.GroupsUI_EffectFilterDialog_Open
- ScenarioStats.SettingsWindowClose
- TalismanGenie.Initialise
- Tortall_DPS.switchTab
- Enemy.IntercomInitialize
- GuildWardenWin.CF10
- Enemy.UI_TextEntryDialog_Open
- GetStats.OnInitialize
- MapMonster.local.FilterButtonState
- AdjustTheTip.CreateSliderMenuItem
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Emojii.UI_ChooseIconDialog_OnListRowLButtonUp
- KeyBar.MoveWindow
- PotionBar.local.UpdateButton
- TortallDPSMeter.Toggle
- mmv.SetNewLayout
- CMapWindow.OnMouseoverRallyCall
- DAoCBuff.ShowMessageWindow
- HealGridGuiTabSpellTrack.Initialize
- MarkBuff.CreateContextBuffItem
- TTitan.UI.ToggleHideCompleted
- zmm.UpdateLayout
- CMapWindow.UpdateScenarioQueueButton
- Dye.Initialize
- MapMonster.local.CreateFilterMenuEntry
- BuddyBind.init
- DetauntTarget.InitUI
- alertMod.SetLabels
- QuickWarReport.local.PrepareConfirmWindowChrome
- EA_Window_OpenPartyWorld.HideAllSortButtonArrows
- EA_Window_OpenPartyWorld.Initialize
- Enemy.IntercomUI_ChooseChannelDialog_Open
- GuildWardenWin.CF9
- AdvancedRenownTrainer.SetImportExportLabels
- GuildWardenWin.CF11
- MapMonster.ShutdownPins
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- MapMonster.FilterButtonState
- Sequencer.OnInitialize
- TastyButtons.Load
- Enemy.UI_ConfigDialog_Open
- GuardRange.UpdateStateMachine
- FixGit.BypassIsBeingThrown
- KeyBar.Update
- KwestorTracker.DrawQuest
- WSCT.ColorOnInitialize
- Enemy.ScenarioInfoInitialize
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- RandomMountUI.OnInitialize
- SocialWindow.ShowFriendsListContextWindow
- Enemy.IntercomUI_IntercomDialog_Open
- MapMonster.CreateFilterMenuEntry
- Phantom.Initialize
- TacticSetNames.SetLabels
- snt_info.tdps_module.entry_point
- BuddyBind.GrabName
- GuardList.UpdateStateMachine
- TTitan.TTButton.Initialize
- TidyQueueHooks.SetupHooks
- GuildWardenWin.CF8
- LibAddonButton.Menu.AddMenuItem
- CDown.SetTMBL
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- UiModWindow.UpdateEnableDisableAllButtons
- MapMonster.local.CleanEditorWindow
- Sequencer.ToggleReset
- EA_SummoningAcceptPrompt.UpdateButtons
- XpStatus.InitializeQuotaWindow
- Emojii.UI_ChooseIconDialog_OnListRowRButtonUp
- RvRStatsRvRTab.Show
- CMapWindow.UpdateScenarioButtons
- GuildWardenWin.CF2
- TastyButtons.UnLoad
- WARCommander.Init
- HealGridGuiTabMouseClick.Initialize
- ScenarioRefresh.OnShown
- TTitan.UI.TomeClearNewEntriesCheckForNew
- BuddyBind.OnRawDeviceInput
- CDown.local.SetTMBL
- HealGridGuiSpellList.Initialize
- NBSetup_Save.Initialize
- GuildWardenWin.CF6
- BloodyMess.OnInitialize
- CMapWindow.OnRallyCallLButtonUp
- AdvancedRenownTrainer.local.SetImportExportLabels
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Tortall_DPS.SetGroupAnchor
- UiModVersionMismatchWindow.Initialize
- MapPin.RButtonUp
- ClosetGoblinCharacterWindow.UpdateActionBarSettings
- MiracleGrow2.InitConfig
- TalismanGenie.Reset
- TortallDPSMeter.ToggleGroup
- ChattyCathy.CC_ModelWindow
- TastyButtons.Button_State_TypeToggle
- APAGui.OnShown
- BlackBookWindow.Initialize
- CDown.SetLabels
- ClosetGoblinCharacterWindow.OnShow
- GuildWardenWin.WinSetup
- ScenarioStats.CallSettingsWindow
- Sequencer.Load
- BankArkel.CreatePackWin
- CaVESWindow.Initialize
- Emojii.SelectTab
- GuildWardenWin.CF5
- PotionBarSettings.OnAboutShown
- EA_Window_Macro.UpdateDetails
- AggroMeter.OnTabLBU
- CMapWindow.ActivateRallyCall
- Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- MapMonster.PinTypeEditor.Initialize
- SocialWindow.ShowIgnoreListContextWindow
- DammazKron.InitializeBookmark
- SiegeChatSettings.WindowInit
- BustedGUI.UpdateNextPrevButtonStatus
- CCTV.Update
- Enemy.MarksUI_MarkConfigDialog_Open
- GuildWardenWin.CF4
- TidyRollOptions.Initialize
- ClosetGoblinCharacterWindow.ShowHelm
- TTitan.UI.TomeClearNewEntriesLClick
- Tortall_DPS.local.SetGroupAnchor
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- DammazKron.local.InitializeNFO
- nRarity.Border:new
- ClosetGoblinCharacterWindow.UseItemRackToggled
- MapMonster.InitializeMapPins
- SOROptions.Initialize
- Sequencer.Clear
- SessionRPs.WarningWindow
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- SocialWindowTabFriends.Initialize
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- Phantom.PopulateWindow
- ClosetGoblinCharacterWindow.HideShowHelm
- UiModWindow.InitAdvancedWindow
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- RaidMeter.OnTabLBU
- ClosetGoblinCharacterWindow.ShowShowHelm
- ClosetGoblinCharacterWindow.ShowCloakOptions
- AdvancedRenownTrainer.local.SetLabels
- EZCraftX.createContextMenu_ItemWindow
- Killer.Initialize
- ReliquaryHunterOptionsWindow.Initialize
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- GuildWardenWin.CF3
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- Enemy.CombatLogUI_StatsWindow_Open
- EA_Window_OpenPartyNearby.HideAllSortButtonArrows
- ScenarioGroupWindow.HideLeaveGroupElements
- TortallDPSMeter.CycleScale
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- SocialWindow.ToggleOfflineMembersShowing
- MarkBuff.local.CreateContextBuffItem
- ClosetGoblinCharacterWindow.OnInitialize
- DammazKron.InitializeNFO
- AuctionWindow.Initialize
- MapMonster.local.OpenTypeViewerWindow
- CMapWindow.ToggleScenarioButtons
- EA_Window_OpenPartyNearby.Initialize
- WbLeadHelperMessage.MessageDialogOpen
- JunkDump.Initialize
- MapPin.EditMarker
- EA_Window_Macro.SelectionIconLButtonDown
- MacroIcons.UpdateSelectedIcon
- CCTV.UpdateSettings
- CDown.local.SetLabels
- KwestorGuiTabArea.Initialize
- PotionBarSettings.OnShown
- RandomMountUI.OnAddCustomMount
- Squared.InitImportExport
- GuildWardenWin.CF7
- RoR_SoR.Restack
- TortallDPSMeter.DamageUpdate
- EA_Window_Macro.Initialize
- MapMonster.CleanEditorWindow
- SocialWindowTabIgnore.Initialize
- UiModWindow.Initialize
- RvRStatsRvRTab.UpdateMode
- ClosetGoblinCharacterWindow.HideCloakOptions
- CallingSetup.UpdateBindingTexts
- AdvancedRenownTrainer.SetLabels
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- MapMonster.OpenTypeViewerWindow
- TastyButtons.CheckInitialSetting
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- GuildWardenWin.CF12
- DPSMeter.InitializeWindow
- ClosetGoblinCharacterWindow.UpdateSetRow
- ScenarioGroupWindow.Initialize
- AggroMeter.Initialize
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- Emojii.OnInitialize
- ScenarioRefresh.OnHidden
- MapPin.local.EditMarker
- TTitan.UI.ToggleButton
- LibWBTogglerManager.Initialize
- CallingSetup.UpdateMacros
- DAoCBuffSettings.SetLabels
- GuildWardenWin.CF1
- TTitan.UI.Initialize
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- AuctionWindow.UpdateBuyButton
- PotionBar.UpdateButton
- RandomMountUI.OnDropSlotLButtonUp
- ScenarioStats.Warning03
- ClosetGoblinZoneWindow.OnInitialize
- CMapWindow.Initialize
- DPSMeter.ChangeTab
- Tortall_DPS.local.switchTab
- MacroIcons.NewIcon
- BuddyBind.update
- ClosetGoblinCharacterWindow.ShowCloak
- AuctionWindow.SwitchTabs
- ScenarioGroupWindow.HideJoinGroupElements
- QuickWarReport.PrepareConfirmWindowChrome


## Binding Resolution

- Total handler declarations: 1920
- Resolved to Lua functions: 1916 (99%)

## .mod Lifecycle: Startup Windows

This element type is instantiated as a startup window by the following .mod addon(s):

| Frame Name | Addon | Hook | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| QueueQueuer_GUI_MapButton | Queue Queuer | OnInitialize | exact | HIGH |
| RVMOD_ManagerToggleButtonWindow | RVMOD_Manager | OnInitialize | exact | HIGH |
| TortallDPSToggle | Tortall_DPS | OnInitialize | exact | HIGH |
## Seen In

- AdjustTheTip
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- ArmorGraphicToggle
- Assist
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoSalvage
- BankArkel
- BlackBook
- Bloody Mess
- BuddyBind
- BuffHead
- Busted
- CCTV
- CDown
- CM_ClosetGoblin
- CMap
- CaVES
- Calling
- CastSequence
- ChattyCathy
- Cheeseboard
- CleanUnitFrames
- CraftingWillard
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- Dascore
- Deathblow
- Deathblow2
- DeepSleep
- DetauntHelper
- Duel
- DuffTimer
- Dye Preview
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FozAuction
- GDes
- Ges
- GetStats
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardList
- GuardRange
- GuildWarden
- HealGrid
- Hopper
- ItemRack
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- LootAlert
- LoyalPet
- ManualScenarioRefresh
- MapMonster
- MapPin
- MarkBuff
- Mass Refine
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleSet
- Motion
- NerfedButtons
- ObjectInspector
- Obsidian
- PartyAd
- PartyCast
- Phantom
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RaidMeter
- RandomMount
- RealmStatus
- Refer
- ReliquaryHunter
- RememberIt
- Res
- ResHelp
- RoR_SoR
- RoR_debolster
- Rotation
- RvRStats
- RvRStatsTab
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_INFO
- SNT_PANEL
- SOR
- ScenarioStats
- Sequencer
- SessionRPs
- Shinies
- SocialWindow 2.0
- Soloq
- SpamBayes
- Squared
- Statdoll
- Statdoll Remix
- TacticSetNames
- TalismanGenie
- TastyButtons
- TaxPayer
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- TokenMachine
- Tome Titan
- TomeTracker
- Tortall_DPS
- TurretRange
- Twister
- TwisterSet
- Vectors
- WARCommander
- WBStutterLess
- WSCT
- WTes
- WaaaghBar
- WarBoard
- WarBoard_FPS
- WarBoard_Loc
- WarBoard_TDPS
- Warbuilder
- Wargames
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- ZonePOP
- alertMod
- bigger_MacroWindow
- emotes
- followTheLeader
- minesweep
- nLootLink
- nRarity
- wbLeadHelper
- zMailMod

## Examples

- AdjustTheTip: AdjustTheTipMenuItemSlider -> Button AdjustTheTipMenuItemSlider
- AdvancedPetAssist: APAOptionsClose -> Button APAOptionsClose
- AdvancedPetAssist: APAOptionsTabsAutoRecall -> Button APAOptionsTabsAutoRecall
- AdvancedPetAssist: APAOptionsTabsControls -> Button APAOptionsTabsControls
- AdvancedPetAssist: APAOptionsTabsFollowTarget -> Button APAOptionsTabsFollowTarget
- AdvancedPetAssist: APAOptionsTabsGeneral -> Button APAOptionsTabsGeneral

## Related APIs

- [Anchor](element_Anchor.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [Color](element_Color.md) (HIGH 100/100) - XML Element Type
- [EA_Button_BottomTab](../../globals/constants/constant_EA_Button_BottomTab.md) (HIGH 100/100) - Constant
- [EA_Button_Default](../../globals/constants/constant_EA_Button_Default.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultBigLeftArrow](../../globals/constants/constant_EA_Button_DefaultBigLeftArrow.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultBigRightArrow](../../globals/constants/constant_EA_Button_DefaultBigRightArrow.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultCheckBox](../../globals/constants/constant_EA_Button_DefaultCheckBox.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultDown](../../globals/constants/constant_EA_Button_DefaultDown.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultIconFrame](../../globals/constants/constant_EA_Button_DefaultIconFrame.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultIconFrame_Large](../../globals/constants/constant_EA_Button_DefaultIconFrame_Large.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultIconFrame_Small](../../globals/constants/constant_EA_Button_DefaultIconFrame_Small.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultLeftArrow](../../globals/constants/constant_EA_Button_DefaultLeftArrow.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultListRow](../../globals/constants/constant_EA_Button_DefaultListRow.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultListSort](../../globals/constants/constant_EA_Button_DefaultListSort.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultMenuButton](../../globals/constants/constant_EA_Button_DefaultMenuButton.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultMinus](../../globals/constants/constant_EA_Button_DefaultMinus.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultPlus](../../globals/constants/constant_EA_Button_DefaultPlus.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultResizableComboBoxSelected](../../globals/constants/constant_EA_Button_DefaultResizableComboBoxSelected.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultResizeable](../../globals/constants/constant_EA_Button_DefaultResizeable.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultRightArrow](../../globals/constants/constant_EA_Button_DefaultRightArrow.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultSmallSquare](../../globals/constants/constant_EA_Button_DefaultSmallSquare.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultText](../../globals/constants/constant_EA_Button_DefaultText.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultTimerFrame](../../globals/constants/constant_EA_Button_DefaultTimerFrame.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultToggleCircle](../../globals/constants/constant_EA_Button_DefaultToggleCircle.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultUp](../../globals/constants/constant_EA_Button_DefaultUp.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultWindowClose](../../globals/constants/constant_EA_Button_DefaultWindowClose.md) (HIGH 100/100) - Constant
- [EA_Button_HUDMenuButton](../../globals/constants/constant_EA_Button_HUDMenuButton.md) (HIGH 100/100) - Constant
- [EA_Button_ListSort](../../globals/constants/constant_EA_Button_ListSort.md) (HIGH 100/100) - Constant
- [EA_Button_ResizeIconFrame_NoNormalImage](../../globals/constants/constant_EA_Button_ResizeIconFrame_NoNormalImage.md) (HIGH 100/100) - Constant
- [EA_Button_Tab](../../globals/constants/constant_EA_Button_Tab.md) (HIGH 100/100) - Constant
- [EA_FilterMenuButtonTemplate](../../globals/constants/constant_EA_FilterMenuButtonTemplate.md) (HIGH 100/100) - Constant
- [EA_Window_MacroIconButton](../../globals/constants/constant_EA_Window_MacroIconButton.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [OverlayOffset](element_OverlayOffset.md) (HIGH 100/100) - XML Element Type
- [OverlaySize](element_OverlaySize.md) (HIGH 100/100) - XML Element Type
- [OverlayTexCoords](element_OverlayTexCoords.md) (HIGH 100/100) - XML Element Type
- [ResizeImages](element_ResizeImages.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Sounds](element_Sounds.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexDims](element_TexDims.md) (HIGH 100/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (HIGH 100/100) - XML Element Type
- [Text](element_Text.md) (HIGH 100/100) - XML Element Type
- [TextColors](element_TextColors.md) (HIGH 100/100) - XML Element Type
- [TextOffset](element_TextOffset.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_Button_DefaultChatScrollToBottom](../../globals/constants/constant_EA_Button_DefaultChatScrollToBottom.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultIcon](../../globals/constants/constant_EA_Button_DefaultIcon.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultResizableSmallComboBoxSelected](../../globals/constants/constant_EA_Button_DefaultResizableSmallComboBoxSelected.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultTimerFrame_Small](../../globals/constants/constant_EA_Button_DefaultTimerFrame_Small.md) (HIGH 90/100) - Constant
- [EA_Button_OpenPartyTab](../../globals/constants/constant_EA_Button_OpenPartyTab.md) (HIGH 90/100) - Constant
- [EA_Button_ResizeIconFrame](../../globals/constants/constant_EA_Button_ResizeIconFrame.md) (HIGH 90/100) - Constant
- [EA_Button_SocialWindowRowTemplate](../../globals/constants/constant_EA_Button_SocialWindowRowTemplate.md) (HIGH 90/100) - Constant
- [EA_CheckButtonButton](../../globals/constants/constant_EA_CheckButtonButton.md) (HIGH 90/100) - Constant
- [EA_ScrollBar_ChatDownArrowButton](../../globals/constants/constant_EA_ScrollBar_ChatDownArrowButton.md) (HIGH 90/100) - Constant
- [EA_ScrollBar_ChatUpArrowButton](../../globals/constants/constant_EA_ScrollBar_ChatUpArrowButton.md) (HIGH 90/100) - Constant
- [EA_ScrollBar_DefaultDownArrowButton](../../globals/constants/constant_EA_ScrollBar_DefaultDownArrowButton.md) (HIGH 90/100) - Constant
- [EA_ScrollBar_DefaultUpArrowButton](../../globals/constants/constant_EA_ScrollBar_DefaultUpArrowButton.md) (HIGH 90/100) - Constant
- [EA_Template_OpenPartyNearbySortButton](../../globals/constants/constant_EA_Template_OpenPartyNearbySortButton.md) (HIGH 90/100) - Constant
- [EA_Templates_Color_Picker_Button](../../globals/constants/constant_EA_Templates_Color_Picker_Button.md) (HIGH 90/100) - Constant
- [EA_Templates_MailWindowTabButton](../../globals/constants/constant_EA_Templates_MailWindowTabButton.md) (HIGH 90/100) - Constant
- [EA_Templates_SpecializationChange](../../globals/constants/constant_EA_Templates_SpecializationChange.md) (HIGH 90/100) - Constant
- [EA_Window_MacroDetailIconButton](../../globals/constants/constant_EA_Window_MacroDetailIconButton.md) (HIGH 90/100) - Constant
- [AnimatedImages](element_AnimatedImages.md) (MEDIUM 45/100) - XML Element Type
- [EventHandler](element_EventHandler.md) (MEDIUM 45/100) - XML Element Type
- [Eventhandlers](element_Eventhandlers.md) (MEDIUM 45/100) - XML Element Type
- [OverlayTexSlices](element_OverlayTexSlices.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [AnimatedImage](element_AnimatedImage.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Pressed](element_Pressed.md) (MEDIUM 45/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 88/100) - XML Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md) (HIGH 76/100) - XML Event
- [OnMouseLeave](../handlers/handler_OnMouseLeave.md) (HIGH 76/100) - XML Event
