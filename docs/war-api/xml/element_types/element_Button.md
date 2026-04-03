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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, BankArkel, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0` |
| Namespaces detected | Button |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAOptionsClose, AdvancedPetAssist: APAOptionsTabsAutoRecall, AdvancedPetAssist: APAOptionsTabsControls, AdvancedPetAssist: APAOptionsTabsFollowTarget, AdvancedPetAssist: APAOptionsTabsGeneral, AdvancedPetAssist: APAOptionsTabsHUD |
| XML usage count | 672 |
| XML attribute usage count | 672 |
| Lua usage count | 9 |
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

Observed XML element type instantiated by 27 addons.

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
- texturescale
- warnOnTextCropped

## Common Handlers

- [OnInitialize](../handlers/handler_OnInitialize.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md)
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

## Common Handler Functions

- APAGui.OnTabButtonUp
- BankArkel.PackTab
- BankArkel.PackTabMover
- Enemy.CombatLogUI_StatsWindow_SortColumnClick
- Enemy.CombatLogUI_StatsWindow_SortColumnRClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick
- AdvancedRenownTraining.OnLButtonUpTab
- AuraConfig.OnConfigTabSelected
- AggroMeter.OnTabLBU
- ClosetGoblinCharacterWindow.EquipmentLButtonDown
- ClosetGoblinCharacterWindow.EquipmentLButtonUp
- ClosetGoblinCharacterWindow.HideCloakOptions


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_GenericCheckButton.Initialize | `function()` | MEDIUM |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | ClosetGoblinCharacterWindow.EquipmentLButtonDown, BuffHead.Setup.Layout.BeginResize, EA_Window_Macro.DetailIconLButtonDown, EA_Window_Macro.SelectionIconLButtonDown, PotionBarFloating.LButtonDown | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | APAGui.OnTabButtonUp, BankArkel.PackTab, Enemy.CombatLogUI_StatsWindow_SortColumnClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick, AdvancedRenownTraining.OnLButtonUpTab, AuraConfig.OnConfigTabSelected | `function(...)` | LOW |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | followTheLeader.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseDrag](../handlers/handler_OnMouseDrag.md) | input | EA_Window_Macro.IconMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag, Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | BankArkel.PackTabMover, Enemy.ConfigurationWindow_ShowTooltip, WarBoard.OnMouseOver, WarBoard.OnMouseOverBottom, ActionBars.OnMouseoverHotbarPageDown, ActionBars.OnMouseoverHotbarPageUp | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | ClosetGoblinCharacterWindow.HideCloakOptions, ClosetGoblinCharacterWindow.HideShowHelm, WarBoard.OnMouseOverEnd, WarBoard.OnMouseOverEndBottom, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd, MiracleGrowLight.harvestEnd | `function(...)` | LOW |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | ClosetGoblinCharacterWindow.EquipmentRButtonDown, CG_ItemRack.EquipmentRButtonDown, EA_Window_Macro.DetailIconRButtonDown, EA_Window_Macro.IconRButtonDown, EA_Window_Macro.SelectionIconRButtonDown | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.CombatLogUI_StatsWindow_SortColumnRClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnRClick, MiracleGrowLight.switchMode, ClosetGoblinCharacterWindow.EquipmentRButtonUp, PP.ItemSlotRMouse, PotionBarFloating.RButtonUp | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | APAGui.OnTabButtonUp, BankArkel.PackTab, Enemy.CombatLogUI_StatsWindow_SortColumnClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick, AdvancedRenownTraining.OnLButtonUpTab, AuraConfig.OnConfigTabSelected, AggroMeter.OnTabLBU, ClosetGoblinCharacterWindow.EquipmentLButtonUp, RoR_SoR.KeepClaimDialog, WarBoard.OnLayoutModeButton, WarBoard.OnOptionsButton, AdvancedRenownTraining.Hide, AdvancedRenownTraining.LinkWindowClose, AuraSettings.OnClose, Enemy.GroupsUI_EffectFilterDialog_Hide, Enemy.IntercomUI_IntercomDialog_Hide, Enemy.MarksUI_MarkConfigDialog_Hide, Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide, Enemy.UI_ConfigDialog_Hide, Enemy.UI_TextEntryDialog_Close, Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide, Enemy.UnitFramesUI_UnitFramePartDialog_Hide, PotionBarSettings.OnAboutClose, PotionBarSettings.OnCancelButton, TidyChat.Options.OnClose, WHMGui.HideOptionsWindow, APAGui.Hide, AdvancedRenownTraining.DeletePreset, AdvancedRenownTraining.ExportCancelButtonPressed, AdvancedRenownTraining.ExportToLink, AdvancedRenownTraining.ExportToWardrobe, AdvancedRenownTraining.ImportCancelButtonPressed, AdvancedRenownTraining.ImportNameInputCancelButtonPressed, AdvancedRenownTraining.ImportNameInputOkButtonPressed, AdvancedRenownTraining.ImportOkButtonPressed, AdvancedRenownTraining.OnExportButtonPressed, AdvancedRenownTraining.OnLoadPressed, AdvancedRenownTraining.PurchaseAdvances, AdvancedRenownTraining.Respecialize, AdvancedRenownTraining.SavePreset, AdvancedRenownTraining.ShowWardrobeImport, AdvancedRenownTraining.TogglePresets, AdvancedRenownTraining.ToggleSaveSettings, AggroMeter.Close, AuraConfig.OnAbilityIconLButtonUp, AuraConfig.OnActivationAlertTextTestButton, AuraConfig.OnClose, AuraConfig.OnDeactivationAlertTextTestButton, AuraConfig.OnLButtonUpTextureTintColor, AuraConfig.OnLButtonUpTimerTintColor, AuraConfig.OnTextureChangeButton, AuraSettings.ChangeSorting, AuraSettings.OnAddAura, AuraSettings.OnShare, AuraShares.ChangeSorting, AuraShares.OnClose, AuraShares.OnCloseAuraSharesImportExportWindow, AuraShares.OnImportAura, AuraShares.OnImportExportOkButton, AuraTexture.OnClose, AuraTexture.OnIconLButtonUp, AuraTexture.OnRaceTabSelected, BankArkel.PackClose, BuffHead.Setup.AdvancedCompression.OnCloseLUp, BuffHead.Setup.AdvancedCompression.OnNewLClick, BuffHead.Setup.AdvancedCompressionItem.OnAddEffectLClick, BuffHead.Setup.AdvancedCompressionItem.OnApplyLClick, BuffHead.Setup.AdvancedCompressionItem.OnCloseLUp, BuffHead.Setup.AdvancedCompressionItem.OnCreateLClick, BuffHead.Setup.AdvancedCompressionItemEffect.OnApplyLClick, BuffHead.Setup.AdvancedCompressionItemEffect.OnCloseLUp, BuffHead.Setup.AdvancedCompressionItemEffect.OnCreateLClick, BuffHead.Setup.AdvancedContainers.OnCloseLUp, BuffHead.Setup.AdvancedContainers.OnNewLClick, BuffHead.Setup.AdvancedContainersItem.OnApplyLClick, BuffHead.Setup.AdvancedContainersItem.OnCloseLUp, BuffHead.Setup.AdvancedContainersItem.OnCreateLClick, BuffHead.Setup.AdvancedContainersItem.Properties.OnCloseLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysIgnoreLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysShowLUp, BuffHead.Setup.Container.OnCloseLUp, BuffHead.Setup.Display.OnCloseLUp, BuffHead.Setup.EffectCache.OnCloseLUp, BuffHead.Setup.EffectCache.OnRefreshLClick, BuffHead.Setup.EffectCache.OnResetLClick, BuffHead.Setup.EffectCache.OnSortButtonLUp, BuffHead.Setup.EffectCacheTable.OnCloseLUp, BuffHead.Setup.Filter.OnAddLClick, BuffHead.Setup.Filter.OnCloseLUp, BuffHead.Setup.General.OnAlwaysIgnoreLClick, BuffHead.Setup.General.OnAlwaysShowLClick, BuffHead.Setup.General.OnCloseLUp, BuffHead.Setup.Layout.Manager.OnCloseLUp, BuffHead.Setup.Layout.Manager.OnImportLClick, BuffHead.Setup.Layout.Manager.OnLayoutsActivateLClick, BuffHead.Setup.Layout.Manager.OnLayoutsDeleteLClick, BuffHead.Setup.Layout.Manager.OnLayoutsSaveCurrentLayoutLClick, BuffHead.Setup.Layout.OnApplyLClick, BuffHead.Setup.Layout.OnCloseLUp, BuffHead.Setup.Layout.OnLayerLClick, BuffHead.Setup.Layout.OnManagerLClick, BuffHead.Setup.Layout.Properties.OnCloseLUp, BuffHead.Setup.Layout.Properties.OnCoreSizeAutoSizeLClick, BuffHead.Setup.OnCloseLUp, BuffHead.Setup.Performance.OnCloseLUp, BuffHead.Setup.PriorityEffects.OnCloseLUp, BuffHead.Setup.PriorityEffects.OnNewLClick, BuffHead.Setup.PriorityEffectsItem.OnApplyLClick, BuffHead.Setup.PriorityEffectsItem.OnCloseLUp, BuffHead.Setup.PriorityEffectsItem.OnCreateLClick, BuffHead.Setup.Trackers.OnAlwaysIgnoreLClick, BuffHead.Setup.Trackers.OnAlwaysShowLClick, BuffHead.Setup.Trackers.OnCloseLUp, ClosetGoblinCharacterWindow.Hide, ClosetGoblinCharacterWindow.HotbarChangeToggled1, ClosetGoblinCharacterWindow.HotbarChangeToggled2, ClosetGoblinCharacterWindow.HotbarChangeToggled3, ClosetGoblinCharacterWindow.HotbarChangeToggled4, ClosetGoblinCharacterWindow.HotbarChangeToggled5, ClosetGoblinCharacterWindow.HotbarPageDownProxy, ClosetGoblinCharacterWindow.HotbarPageUpProxy, ClosetGoblinCharacterWindow.OnClickDeleteSetButton, ClosetGoblinCharacterWindow.OnClickNewSetButton, ClosetGoblinCharacterWindow.OnClickSortNameButton, ClosetGoblinCharacterWindow.OnClickSortSetButton, ClosetGoblinCharacterWindow.OnClickSortTacticsButton, ClosetGoblinCharacterWindow.OnClickZoneConfigButton, ClosetGoblinCharacterWindow.ShowCloak, ClosetGoblinCharacterWindow.ShowCloakHeraldry, ClosetGoblinCharacterWindow.ShowHelm, ClosetGoblinCharacterWindow.UseItemRackToggled, ClosetGoblinZoneWindow.Hide, ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton, ClosetGoblinZoneWindow.OnClickSortZoneButton, DAoCBuff.CloseMessageWindow, DAoCBuffSettings.AddEntry, DAoCBuffSettings.AddFilter, DAoCBuffSettings.AddFrame, DAoCBuffSettings.AddList, DAoCBuffSettings.ChangeFrameName, DAoCBuffSettings.ChangeGlobalBorder, DAoCBuffSettings.ChangeGlobalFont, DAoCBuffSettings.ChangeGlobalGlass, DAoCBuffSettings.ChangeGlobalRefresh, DAoCBuffSettings.ChangeGlobalSize, DAoCBuffSettings.ClearLeft, DAoCBuffSettings.ClearRight, DAoCBuffSettings.CloseOptionswindow, DAoCBuffSettings.CopyLeftToRight, DAoCBuffSettings.CopyRightToLeft, DAoCBuffSettings.EditFilterOnLButtonUp, DAoCBuffSettings.FilterNameChanged, DAoCBuffSettings.FilterSettings.AddClassTable, DAoCBuffSettings.FilterSettings.AddCondition, DAoCBuffSettings.FilterSettings.ChangeG4RecursiveConditions, DAoCBuffSettings.FilterSettings.Close, DAoCBuffSettings.FilterSettings.RemoveCondition, DAoCBuffSettings.ImExport.Open, DAoCBuffSettings.MoveFilterDown, DAoCBuffSettings.MoveFilterUp, DAoCBuffSettings.MoveLeftToRight, DAoCBuffSettings.MoveRightToLeft, DAoCBuffSettings.OpenOptionswindow, DAoCBuffSettings.RemoveFilter, DAoCBuffSettings.RemoveFrame, DAoCBuffSettings.RemoveLeft, DAoCBuffSettings.RemoveList, DAoCBuffSettings.RemoveRight, DAoCBuffSettings.RestartTracker, EA_Window_Backpack.ChangeSorting, EA_Window_Macro.DetailIconLButtonUp, EA_Window_Macro.Hide, EA_Window_Macro.HideMacroIconSelectionWindow, EA_Window_Macro.IconLButtonUp, EA_Window_Macro.OnSave, EA_Window_Macro.SelectionIconLButtonUp, Enemy.AssistUI_ConfigDialog_MarkNewTargetEditTemplate, Enemy.AssistUI_ConfigDialog_OnMarkNewTargetChanged, Enemy.AssistUI_ConfigDialog_OnNewTargetSoundChanged, Enemy.AssistUI_ConfigDialog_OnTargetOnNotifyClickChanged, Enemy.CombatLogUI_SnapshotWindow_Hide, Enemy.CombatLogUI_StatsWindow_Hide, Enemy.CombatLogUI_StatsWindow_SessionAdd, Enemy.CombatLogUI_StatsWindow_SessionDelete, Enemy.CombatLogUI_StatsWindow_SessionRename, Enemy.ConfigurationWindow_OnButtonClick, Enemy.GroupsUI_EffectFilterDialog_Ok, Enemy.IntercomUI_ChooseChannelDialog_Hide, Enemy.IntercomUI_ChooseChannelDialog_OnOkButton, Enemy.IntercomUI_IntercomDialog_OnCreateButton, Enemy.IntercomUI_IntercomDialog_OnCreateButton2, Enemy.IntercomUI_IntercomDialog_OnCreateButton3, Enemy.IntercomUI_IntercomDialog_OnCreateButton4, Enemy.IntercomUI_IntercomDialog_OnCreateButton5, Enemy.IntercomUI_IntercomDialog_OnCreateButton7, Enemy.IntercomUI_IntercomDialog_OnInviteButton, Enemy.IntercomUI_IntercomJoinDialog_Hide, Enemy.IntercomUI_IntercomJoinDialog_OnOkButton, Enemy.KillSpamUI_KillSpamAreaStatsDialog_Hide, Enemy.MarksUI_MarkConfigDialog_Ok, Enemy.OnLeaveButton, Enemy.ScenarioAlerterUI_ConfigDialog_OnEnabledChanged, Enemy.ScenarioInfoToggle, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SwitchToStandard, Enemy.UI_ChooseIconDialog_Hide, Enemy.UI_ChooseIconDialog_OnListRowLButtonUp, Enemy.UI_ConfigDialog_Ok, Enemy.UI_ConfigDialog_Reset, Enemy.UI_ConfigDialog_ResetAll, Enemy.UI_TextEntryDialog_Ok, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsAdd, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsDelete, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsDown, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsEdit, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsEnable, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsExport, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsImport, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsReset, Enemy.UnitFramesUI_ConfigDialog_ClickCastingsUp, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsAdd, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsDelete, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsDown, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsEdit, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsEnable, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsExport, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsImport, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsReset, Enemy.UnitFramesUI_ConfigDialog_EffectsIndicatorsUp, Enemy.UnitFramesUI_ConfigDialog_OnEnabledChanged, Enemy.UnitFramesUI_ConfigDialog_OnSortingEnabledChanged, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsAdd, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsDelete, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsDown, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsEdit, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsEnable, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsExport, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsImport, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsReset, Enemy.UnitFramesUI_ConfigDialog_UnitFramePartsUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_ChooseIcon, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersAdd, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersDelete, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersDown, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersEdit, Enemy.UnitFramesUI_EffectsIndicatorDialog_EffectFiltersUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_Ok, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype1Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype2Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnArchetype3Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype1Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype2Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnArchetype3Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier1Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier2Changed, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnKeyModifier3Changed, Enemy.UnitFramesUI_UnitFramePartDialog_Ok, Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype1Changed, Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype2Changed, Enemy.UnitFramesUI_UnitFramePartDialog_OnArchetype3Changed, Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged, Killer.HideScoreDetailsWindow, Killer.HideSettingsWindow, Killer.OnSettingsCheckboxChanged, Killer.ToggleSettingsWindow, LibGroup.Setup.OnCloseLUp, LibWBTogglerManager.ToggleManager, MiracleGrowLight.harvestEnd, PP.ItemSlotLMouse, PP.OnClose, PP.SelectDye, PP.ShowPaperDoll, PP.ToggleWindow, PotionBarFloating.LButtonUp, PotionBarSettings.AutohideCheck, PotionBarSettings.MoveDown, PotionBarSettings.MoveUp, PotionBarSettings.OnAboutButton, PotionBarSettings.OnDontSplitNameCheck, PotionBarSettings.OnResetButton, PotionBarSettings.OnSaveButton, PotionBarSettings.OnUseCheck, PotionBarSettings.Show1Check, PotionBarSettings.ShowLastCheck, RoR_SoR.CloseSetOffsetWindow, RoR_SoR.CloseSetOpacityWindow, RoR_SoR.CloseSetScaleWindow, RoR_SoR.KeepUnClaimDialog, RoR_SoR.Toggle, ScenarioSummaryWindow.OnLeaveNowClicked, Shinies.OnLButtonUp_Close, ShiniesAuctionsUI.OnLButtonUp_Results_SortButton, ShiniesAutoUI.OnLButtonUp_AutoDelete, ShiniesBrowseUI.OnLButtonUp_Results_SortButton, ShiniesBrowseUI.OnLButtonUp_Searches_SortButton, ShiniesConfigPrice.OnLButtonUp_DecreasePriority, ShiniesConfigPrice.OnLButtonUp_EnableModule, ShiniesConfigPrice.OnLButtonUp_IncreasePriority, ShiniesPostUI.OnLButtonUp_SortButton, TexturedButtons.Setup.Actionbar.OnCloseLUp, TexturedButtons.Setup.AdvancedTextures.OnCloseLUp, TexturedButtons.Setup.Cooldown.OnCloseLUp, TexturedButtons.Setup.Fonts.OnCloseLUp, TexturedButtons.Setup.Misc.OnCloseLUp, TexturedButtons.Setup.OnCloseLUp, TexturedButtons.Setup.Textures.OnCloseLUp, TexturedButtons.Setup.Tint.OnCloseLUp, TidyChat.Copy.CopyNext, TidyChat.Copy.CopyPrev, TidyChat.Copy.OnClose, TidyChat.LootRoll.OnClose, TidyChat.Options.OnApply, TidyChat.Options.OnTabLBU, TidyChat.Options.Reset, TidyRoll.CustomAutoRoll.AddById, TidyRoll.CustomAutoRoll.OnApply, TidyRoll.CustomAutoRoll.OnClose, TidyRoll.CustomAutoRoll.OnDeleteButton, TidyRollOptions.OnTabLBU, TurretRange.Setup.Display.OnCloseLUp, TurretRange.Setup.Distance.OnApplyLClick, TurretRange.Setup.Distance.OnCloseLUp, TurretRange.Setup.Distance.OnCreateLClick, TurretRange.Setup.Distances.OnAddLClick, TurretRange.Setup.Distances.OnCloseLUp, TurretRange.Setup.General.OnCloseLUp, TurretRange.Setup.OnCloseLUp, TurretRange.Setup.OnDisplaySetupLUp, TurretRange.Setup.OnDistancesSetupLUp, TurretRange.Setup.OnGeneralSetupLUp, WHMGui.HideDetailsWindow, WHMGui.OnOptionsResetButton, WHMGui.ToggleOptionsWindow, WSCT.ColorAcceptButtonOnButtonUp, WSCT.ColorHideMenu, WSCT.ColorOnButtonUp, WSCT.HideMenu, WSCT.LoadProfileOnButtonUp, WSCT.OnButtonUp, WSCT.TabOnLButtonUp, WSCT.TestOnButtonUp, WarBoard.MoveModLeft, WarBoard.MoveModRight, WarBoard.Options.MinusHorPad, WarBoard.Options.MinusVertPad, WarBoard.Options.OnLButtonUpTab, WarBoard.Options.OpenUiModWindow, WarBoard.Options.PlusHorPad, WarBoard.Options.PlusVertPad, WarBoard.Options.SetRowAlign, followTheLeader.OnLButtonUp | `flags, x, y` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | BankArkel.PackTabMover, Enemy.ConfigurationWindow_ShowTooltip, WarBoard.OnMouseOver, WarBoard.OnMouseOverBottom, ActionBars.OnMouseoverHotbarPageDown, ActionBars.OnMouseoverHotbarPageUp, AnywhereTrainer.OnMouseOver, AuraConfig.OnIconMouseOver, BankArkel.PackIconOver, BuffHead.Setup.EffectCache.OnSortButtonMouseOver, BuffHead.Setup.Layout.OnLayerMouseOver, CG_ItemRack.EquipmentMouseOver, ClosetGoblinCharacterWindow.EquipmentMouseOver, ClosetGoblinCharacterWindow.ShowCloakOptions, ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly, ClosetGoblinCharacterWindow.ShowShowCloakOnly, ClosetGoblinCharacterWindow.ShowShowHelm, ClosetGoblinCharacterWindow.ShowShowHelmOnly, EA_Window_Macro.DetailIconMouseOver, EA_Window_Macro.IconMouseOver, EA_Window_Macro.SelectionIconMouseOver, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseOver, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseOver, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseOver, FrameManager.OnMouseOver, MiracleGrowLight.harvestStart, MiracleGrowLight.onHover, PP.ItemSlotMouseOver, PotionBarFloating.MouseOver, PotionBarSettings.OnMouseoverAutohide, PotionBarSettings.OnMouseoverDontSplitName, PotionBarSettings.OnMouseoverMoveDown, PotionBarSettings.OnMouseoverMoveUp, PotionBarSettings.OnMouseoverShow1, PotionBarSettings.OnMouseoverShowLast, PotionBarSettings.OnMouseoverUse, RoR_SoR.MainTooltip, WSCT.OnMouseOver, followTheLeader.OnMouseOver | `` |  |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.CombatLogUI_StatsWindow_SortColumnRClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnRClick, MiracleGrowLight.switchMode, ClosetGoblinCharacterWindow.EquipmentRButtonUp, PP.ItemSlotRMouse, PotionBarFloating.RButtonUp, RoR_SoR.OnTabRBU, followTheLeader.OnRButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | ClosetGoblinCharacterWindow.HideCloakOptions, ClosetGoblinCharacterWindow.HideShowHelm, WarBoard.OnMouseOverEnd, WarBoard.OnMouseOverEndBottom, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd, MiracleGrowLight.harvestEnd | `` |  |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | ClosetGoblinCharacterWindow.EquipmentLButtonDown, BuffHead.Setup.Layout.BeginResize, EA_Window_Macro.DetailIconLButtonDown, EA_Window_Macro.SelectionIconLButtonDown, PotionBarFloating.LButtonDown | `flags, x, y` | MEDIUM |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | ClosetGoblinCharacterWindow.EquipmentRButtonDown, CG_ItemRack.EquipmentRButtonDown, EA_Window_Macro.DetailIconRButtonDown, EA_Window_Macro.IconRButtonDown, EA_Window_Macro.SelectionIconRButtonDown | `flags, x, y` | LOW |
| [OnMouseDrag](../handlers/handler_OnMouseDrag.md) | input | EA_Window_Macro.IconMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag, Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag | `flags, x, y` | LOW |
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_GenericCheckButton.Initialize | `` |  |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | followTheLeader.OnMButtonUp | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonDown** handlers call: `WindowGetId`

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonGetText`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxClearMenuItems`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `CreateWindow`, `DestroyWindow`, `DoesWindowExist`, `GameData.Player.GetRenownRefundCost`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `ListBoxSetDisplayOrder`, `RegisterEventHandler`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `SystemData.ActiveWindow.name:find`, `SystemData.ActiveWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxGetText:len`, `TextEditBoxSelectAll`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowAssignFocus`, `WindowClearAnchors`, `WindowGetAnchorCount`, `WindowGetDimensions`, `WindowGetHandleInput`, `WindowGetId`, `WindowGetParent`, `WindowGetScale`, `WindowGetScreenPosition`, `WindowGetShowing`, `WindowSetHandleInput`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnMouseDrag** handlers call: `Cursor.IconOnCursor`, `Cursor.PickUp`

**OnMouseOver** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetShowing`

**OnMouseOverEnd** handlers call: `WindowSetShowing`

**OnRButtonUp** handlers call: `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetId`, `WindowSetLayer`

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonGetText`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxClearMenuItems`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `CreateWindow`, `DestroyWindow`, `DoesWindowExist`, `GameData.Player.GetRenownRefundCost`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `ListBoxSetDisplayOrder`, `RegisterEventHandler`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `SystemData.ActiveWindow.name:find`, `SystemData.ActiveWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxGetText:len`, `TextEditBoxSelectAll`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowAssignFocus`, `WindowClearAnchors`, `WindowGetAnchorCount`, `WindowGetDimensions`, `WindowGetHandleInput`, `WindowGetId`, `WindowGetParent`, `WindowGetScale`, `WindowGetScreenPosition`, `WindowGetShowing`, `WindowSetHandleInput`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnMouseOver** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetShowing`

**OnRButtonUp** handlers call: `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetId`, `WindowSetLayer`

**OnMouseOverEnd** handlers call: `WindowSetShowing`

**OnLButtonDown** handlers call: `WindowGetId`

**OnMouseDrag** handlers call: `Cursor.IconOnCursor`, `Cursor.PickUp`

## Common Inherits

- CG_ItemRackEquipmentButton
- ClosetGoblinEquipmentButton
- DefaultButton
- EA_Button_Default
- EA_Button_DefaultCheckBox
- EA_Button_DefaultListSort
- EA_Button_DefaultResizeable
- EA_Button_DefaultText
- EA_Button_DefaultWindowClose
- EA_Button_Tab
- EA_Window_MacroIconButton
- MacroIconSelectionWindowIconButton

## Common Parent Elements

- [Windows](element_Windows.md) — 664× (HIGH)
- [Window](element_Window.md) — 8× (MEDIUM)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 556× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 418× (HIGH)
- [Size](element_Size.md) — 259× (HIGH)
- [TexSlices](element_TexSlices.md) — 22× (HIGH)
- [TextOffset](element_TextOffset.md) — 18× (HIGH)
- [Windows](element_Windows.md) — 18× (HIGH)
- [TextColors](element_TextColors.md) — 16× (HIGH)
- [ResizeImages](element_ResizeImages.md) — 14× (HIGH)
- [OverlayOffset](element_OverlayOffset.md) — 7× (MEDIUM)
- [OverlaySize](element_OverlaySize.md) — 7× (MEDIUM)
- [OverlayTexCoords](element_OverlayTexCoords.md) — 7× (MEDIUM)
- [TexCoords](element_TexCoords.md) — 7× (MEDIUM)
- [TexDims](element_TexDims.md) — 4× (MEDIUM)
- [Anchor](element_Anchor.md) — 2× (LOW)
- [AnimatedImages](element_AnimatedImages.md) — 1× (LOW)

## Common Template Bases

- AbilitiesWindowButtonDef
- Aggro_Button_Template
- AuraIconButton
- AuraSharesSortButton
- AuraTabButtonTemplate
- AuraWindowButton
- AuraWindowSortButton
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
- CharacterWindowDefaultButton
- ClosetGoblinDefaultButton
- ClosetGoblinEquipmentButton
- DefaultButton
- DefaultIconButton
- EA_Button_BottomTab
- EA_Button_Default
- EA_Button_DefaultBigLeftArrow
- EA_Button_DefaultBigRightArrow
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultIconFrame
- EA_Button_DefaultIconFrame_Large
- EA_Button_DefaultIconFrame_Small
- EA_Button_DefaultListSort
- EA_Button_DefaultMenuButton
- EA_Button_DefaultMinus
- EA_Button_DefaultPlus
- EA_Button_DefaultResizeable
- EA_Button_DefaultSmallSquare
- EA_Button_DefaultText
- EA_Button_DefaultToggleCircle
- EA_Button_DefaultUp
- EA_Button_DefaultWindowClose
- EA_Button_ListSort
- EA_Button_ResizeIconFrame_NoNormalImage
- EA_Button_Tab
- EA_FilterMenuButtonTemplate
- EA_ScrollBar_DefaultDownArrowButton
- EA_ScrollBar_DefaultUpArrowButton
- EA_Templates_Color_Picker_Button
- EA_Window_MacroDetailIconButton
- EA_Window_MacroIconButton
- EnemyChooseIconDialogList_IconButton
- ItemWindowSlotButton
- MacroIconSelectionWindowIconButton
- MiracleGrowLightButton
- PotionBarWindowButton
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
- TChatButton
- TChatTabButton
- TRollButton
- TRollItemButton
- TabButtonTemplate
- WSCTButtonTemplate
- WSCTTabButton


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<Button backgroundtexture="shared_01" font="font_default_text" handleinput="false" highlighttexture="shared_01" name="..." overlayhighlighttexture="shared_01" overlaytexture="shared_01" textalign="left">
  <Size/>
  <TextColors>
    <Normal a="255" b="255" g="255" r="255"/>
    <NormalHighlit a="255" b="63" g="213" r="250"/>
    <Pressed a="255" b="63" g="213" r="250"/>
    <PressedHighlit a="255" b="63" g="213" r="250"/>
    <Disabled a="255" b="92" g="92" r="92"/>
  </TextColors>
  <ResizeImages>
    <Normal def="EA_HorizontalResizeImage_Defa..."/>
    <NormalHighlit def="EA_HorizontalResizeImage_Defa..."/>
    <Pressed def="EA_HorizontalResizeImage_Defa..."/>
    <PressedHighlit def="EA_HorizontalResizeImage_Defa..."/>
    <Disabled def="EA_HorizontalResizeImage_Defa..."/>
  </ResizeImages>
  <OverlaySize x="27" y="28"/>
  <OverlayOffset x="93" y="0"/>
  <OverlayTexCoords>
    <Normal x="0" y="28"/>
    <NormalHighlit x="27" y="28"/>
    <Pressed x="0" y="56"/>
    <PressedHighlit x="0" y="56"/>
    <Disabled x="27" y="56"/>
  </OverlayTexCoords>
  <TextOffset x="5" y="5"/>
</Button>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 97% | EA_Button_DefaultWindowClose, EA_Button_Tab, EA_Button_DefaultResizeable, EA_Button_DefaultCheckBox, ... |
| `id` | optional | 31% | 1, 2, 3, 4, ... |
| `textalign` | optional | 14% | left, center, leftcenter, bottomright, ... |
| `font` | optional | 12% | font_default_text, font_chat_text, font_clear_tiny, font_clear_small_bold, ... |
| `layer` | optional | 10% | overlay, popup, default, secondary, ... |
| `handleinput` | optional | 7% | false, true |
| `backgroundtexture` | optional | 4% | shared_01, EA_Training_Specialization, EA_Abilities01_d5, bpKtxt, ... |
| `highlighttexture` | optional | 3% | shared_01, EA_Training_Specialization, EA_Abilities01_d5, EA_HUD_01, ... |
| `drawchildrenfirst` | optional | 2% | true, false |
| `warnOnTextCropped` | optional | 1% | false |
| `texturescale` | optional | 1% | 2.0, 0.74, 1.171, 0.3, ... |
| `textureScale` | optional | 1% | 0.62, 0.58, 0.7, 0.0, ... |
| `overlayhighlighttexture` | optional | 1% | shared_01 |
| `overlaytexture` | optional | 1% | shared_01 |
| `draganddrop` | optional | 0% | true, false |
| `gameactionbutton` | optional | 0% | right, left |
| `autoresize` | optional | 0% | false |
| `savesettings` | optional | 0% | false, true |
| `movable` | optional | 0% | false |
| `popable` | optional | 0% | false |
| `wordwrap` | optional | 0% | true |
| `Layer` | optional | 0% | popup |
| `scale` | optional | 0% | .5 |
| `text` | optional | 0% | Save |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 556 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 418 times as an unnamed child.

### [Size](element_Size.md)

Observed 259 times as an unnamed child.

### [TexSlices](element_TexSlices.md)

Observed 22 times as an unnamed child.

### [TextOffset](element_TextOffset.md)

Observed 18 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5, 4, 10, 0 |
| `y` | **required** | 5, 2, 3, 0 |
### [Windows](element_Windows.md)

Observed 18 times as an unnamed child.

### [TextColors](element_TextColors.md)

Observed 16 times as an unnamed child.

### [ResizeImages](element_ResizeImages.md)

Observed 14 times as an unnamed child.

### [OverlayOffset](element_OverlayOffset.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 93, 193, 148, 223 |
| `y` | **required** | 0 |
### [OverlaySize](element_OverlaySize.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 27 |
| `y` | **required** | 28 |
### [OverlayTexCoords](element_OverlayTexCoords.md)

Observed 7 times as an unnamed child.

### [TexCoords](element_TexCoords.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 32 |
| `y` | optional | 0, 163, 51, 32 |
### [TexDims](element_TexDims.md)

Observed 4 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 64, 256, 32, 430 |
| `y` | **required** | 64, 256, 32, 430 |
### [Anchor](element_Anchor.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, bottomright, right |
| `relativePoint` | **required** | center, topleft, bottomright, left |
| `relativeTo` | **required** | Root, $parent, $parentGeneral, $parentFollowTarget |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
### [AnimatedImages](element_AnimatedImages.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [Button](element_Button.md)
- [Anchor](element_Anchor.md) (structural, 2×, LOW)
  - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
    - (cycle)
- [Anchors](element_Anchors.md) (structural, 556×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
      - (cycle)
- [AnimatedImages](element_AnimatedImages.md) (structural, 1×, LOW)
  - [Normal](element_Normal.md) (structural, 1×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
- [EventHandlers](element_EventHandlers.md) (structural, 418×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
- [OverlayOffset](element_OverlayOffset.md) (structural, 7×, MEDIUM)
- [OverlaySize](element_OverlaySize.md) (structural, 7×, MEDIUM)
- [OverlayTexCoords](element_OverlayTexCoords.md) (structural, 7×, MEDIUM)
  - [Disabled](element_Disabled.md) (structural, 7×, HIGH)
  - [Normal](element_Normal.md) (structural, 7×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 7×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 7×, HIGH)
- [ResizeImages](element_ResizeImages.md) (structural, 14×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 8×, HIGH)
  - [Normal](element_Normal.md) (structural, 10×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 14×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 10×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 9×, HIGH)
- [Size](element_Size.md) (structural, 259×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 7×, MEDIUM)
  - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
  - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
  - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
  - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
  - [Left](element_Left.md) (structural, 7×, MEDIUM)
  - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
  - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
  - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
  - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
  - [Right](element_Right.md) (structural, 7×, MEDIUM)
  - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
  - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
  - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
- [TexDims](element_TexDims.md) (structural, 4×, MEDIUM)
- [TexSlices](element_TexSlices.md) (structural, 22×, HIGH)
  - [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
  - [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
  - [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
  - [Disabled](element_Disabled.md) (structural, 11×, HIGH)
  - [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
  - [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
  - [Normal](element_Normal.md) (structural, 22×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
  - [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
  - [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
  - [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)
- [TextColors](element_TextColors.md) (structural, 16×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 16×, HIGH)
  - [DisabledPressed](element_DisabledPressed.md) (structural, 4×, MEDIUM)
  - [Normal](element_Normal.md) (structural, 16×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
- [TextOffset](element_TextOffset.md) (structural, 18×, HIGH)
- [Windows](element_Windows.md) (structural, 18×, HIGH)
  - [AnimatedImage](element_AnimatedImage.md) (named, 12×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 12×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [AnimFrames](element_AnimFrames.md) (structural, 2×, MEDIUM)
      - [AnimFrame](element_AnimFrame.md) (structural, 14×, HIGH)
    - [Size](element_Size.md) (structural, 9×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Windows](element_Windows.md) (structural, 2×, MEDIUM)
      - (cycle)
  - [Button](element_Button.md) (named, 664×, HIGH)
    - (cycle)
  - [CircleImage](element_CircleImage.md) (named, 23×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 22×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 22×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 20×, HIGH)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TexDims](element_TexDims.md) (structural, 2×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 6×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [ColorPicker](element_ColorPicker.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [ColorSize](element_ColorSize.md) (structural, 1×, HIGH)
    - [ColorSpacing](element_ColorSpacing.md) (structural, 1×, HIGH)
    - [ColorTexCoords](element_ColorTexCoords.md) (structural, 1×, HIGH)
    - [ColorTexDims](element_ColorTexDims.md) (structural, 1×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [ComboBox](element_ComboBox.md) (named, 233×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 228×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 179×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 7×, MEDIUM)
    - [Size](element_Size.md) (structural, 52×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [CooldownDisplay](element_CooldownDisplay.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [DynamicImage](element_DynamicImage.md) (named, 237×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 216×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 1×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 40×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 190×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 14×, HIGH)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TexDims](element_TexDims.md) (structural, 65×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 25×, HIGH)
    - [Windows](element_Windows.md) (structural, 6×, MEDIUM)
      - (cycle)
  - [EditBox](element_EditBox.md) (named, 151×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 126×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 102×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 127×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 15×, HIGH)
  - [FullResizeImage](element_FullResizeImage.md) (named, 156×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 139×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 1×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 35×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 14×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
      - [Middle](element_Middle.md) (structural, 14×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 9×, MEDIUM)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
    - [TexSlices](element_TexSlices.md) (structural, 5×, MEDIUM)
      - [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 11×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 22×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
      - [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 28×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [HorizontalResizeImage](element_HorizontalResizeImage.md) (named, 19×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 11×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 4×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 8×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
      - [Middle](element_Middle.md) (structural, 14×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 8×, HIGH)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 7×, HIGH)
  - [Label](element_Label.md) (named, 1243×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 1238×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 375×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 95×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 1168×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Text](element_Text.md) (structural, 60×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 14×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [ListBox](element_ListBox.md) (named, 42×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 42×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 10×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [ListData](element_ListData.md) (structural, 42×, HIGH)
      - [ListColumns](element_ListColumns.md) (structural, 25×, HIGH)
        - [ListColumn](element_ListColumn.md) (structural, 42×, HIGH)
    - [Size](element_Size.md) (structural, 31×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [LogDisplay](element_LogDisplay.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Windows](element_Windows.md) (structural, 1×, HIGH)
      - (cycle)
  - [MapDisplay](element_MapDisplay.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [ScrollWindow](element_ScrollWindow.md) (named, 26×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 16×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 12×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Windows](element_Windows.md) (structural, 24×, HIGH)
      - (cycle)
  - [SliderBar](element_SliderBar.md) (named, 83×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 83×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 80×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 76×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [StatusBar](element_StatusBar.md) (named, 9×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 9×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
  - [VerticalResizeImage](element_VerticalResizeImage.md) (named, 4×, MEDIUM)
    - [Sizes](element_Sizes.md) (structural, 1×, LOW)
      - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
      - [Middle](element_Middle.md) (structural, 14×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 3×, HIGH)
  - [VerticalScrollbar](element_VerticalScrollbar.md) (named, 25×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 25×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 23×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [Window](element_Window.md) (named, 1001×, HIGH)
    - [Button](element_Button.md) (named, 8×, MEDIUM)
      - (cycle)
    - [ComboBox](element_ComboBox.md) (named, 4×, MEDIUM)
      - [Anchors](element_Anchors.md) (structural, 228×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [EventHandlers](element_EventHandlers.md) (structural, 179×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 7×, MEDIUM)
      - [Size](element_Size.md) (structural, 52×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [FullResizeImage](element_FullResizeImage.md) (named, 2×, LOW)
      - [Anchors](element_Anchors.md) (structural, 139×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 1×, LOW)
      - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [Size](element_Size.md) (structural, 35×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
      - [Sizes](element_Sizes.md) (structural, 14×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
        - [Middle](element_Middle.md) (structural, 14×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 9×, MEDIUM)
        - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
        - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
        - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
        - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
        - [Left](element_Left.md) (structural, 7×, MEDIUM)
        - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
        - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
        - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
        - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
        - [Right](element_Right.md) (structural, 7×, MEDIUM)
        - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
        - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
        - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
      - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
      - [TexSlices](element_TexSlices.md) (structural, 5×, MEDIUM)
        - [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
        - [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
        - [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
        - [Disabled](element_Disabled.md) (structural, 11×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
        - [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
        - [Normal](element_Normal.md) (structural, 22×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
        - [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
        - [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
        - [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)
      - [TintColor](element_TintColor.md) (structural, 28×, HIGH)
      - [Windows](element_Windows.md) (structural, 1×, LOW)
        - (cycle)
    - [Label](element_Label.md) (named, 15×, HIGH)
      - [Anchors](element_Anchors.md) (structural, 1238×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 375×, HIGH)
      - [EventHandlers](element_EventHandlers.md) (structural, 95×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [Size](element_Size.md) (structural, 1168×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
      - [Text](element_Text.md) (structural, 60×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 14×, HIGH)
      - [Windows](element_Windows.md) (structural, 1×, LOW)
        - (cycle)
    - [SliderBar](element_SliderBar.md) (named, 2×, LOW)
      - [Anchors](element_Anchors.md) (structural, 83×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [EventHandlers](element_EventHandlers.md) (structural, 80×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [Size](element_Size.md) (structural, 76×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Window](element_Window.md) (named, 5×, MEDIUM)
      - (cycle)
    - [Anchor](element_Anchor.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 745×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 280×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 624×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Visual](element_Visual.md) (structural, 1×, LOW)
      - [Color](element_Color.md) (structural, 1×, HIGH)
    - [Windows](element_Windows.md) (structural, 476×, HIGH)
      - (cycle)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowSetShowing` | ui | 114 | OnLButtonUp, OnMouseOver, OnMouseOverEnd |
| `ComboBoxGetSelectedMenuItem` | ui | 56 | OnLButtonUp |
| `WindowGetId` | ui | 44 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ButtonSetPressedFlag` | ui | 42 | OnLButtonUp |
| `ButtonGetDisabledFlag` | ui | 39 | OnLButtonUp |
| `TextEditBoxGetText` | ui | 24 | OnLButtonUp |
| `SystemData.ActiveWindow.name:match` | data | 19 | OnLButtonUp, OnRButtonUp |
| `DoesWindowExist` | ui | 14 | OnLButtonUp |
| `ButtonGetPressedFlag` | ui | 13 | OnLButtonUp |
| `WindowGetShowing` | ui | 9 | OnLButtonUp |
| `LabelSetText` | ui | 8 | OnLButtonUp |
| `TextEditBoxSetText` | ui | 8 | OnLButtonUp |
| `DestroyWindow` | ui | 7 | OnLButtonUp |
| `ListBoxSetDisplayOrder` | ui | 7 | OnLButtonUp |
| `WindowSetLayer` | ui | 7 | OnRButtonUp |
| `WindowAddAnchor` | ui | 6 | OnLButtonUp |
| `WindowClearAnchors` | ui | 6 | OnLButtonUp |
| `WindowGetParent` | ui | 6 | OnLButtonUp, OnMouseOver |
| `ComboBoxSetSelectedMenuItem` | ui | 5 | OnLButtonUp |
| `SliderBarSetCurrentPosition` | ui | 5 | OnLButtonUp |
| `ButtonSetDisabledFlag` | ui | 4 | OnLButtonUp |
| `ComboBoxGetSelectedText` | ui | 4 | OnLButtonUp |
| `SystemData.ActiveWindow.name:find` | data | 4 | OnLButtonUp |
| `SystemData.MouseOverWindow.name:match` | data | 3 | OnMouseOver, OnRButtonUp |
| `TextEditBoxGetText:len` | ui | 3 | OnLButtonUp |
| `WindowSetHandleInput` | ui | 3 | OnLButtonUp |
| `WindowStartAlphaAnimation` | ui | 3 | OnLButtonUp |
| `ButtonSetText` | ui | 2 | OnLButtonUp |
| `ComboBoxClearMenuItems` | ui | 2 | OnLButtonUp |
| `CreateWindow` | ui | 2 | OnLButtonUp |
| `LabelSetTextColor` | ui | 2 | OnLButtonUp |
| `ListBoxGetDataIndex` | ui | 2 | OnLButtonUp |
| `RegisterEventHandler` | event | 2 | OnLButtonUp |
| `SliderBarGetCurrentPosition` | ui | 2 | OnLButtonUp |
| `WindowGetDimensions` | ui | 2 | OnLButtonUp, OnMouseOver |
| `WindowGetScreenPosition` | ui | 2 | OnLButtonUp, OnMouseOver |
| `ButtonGetText` | ui | 1 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 1 | OnMouseDrag |
| `Cursor.PickUp` | data | 1 | OnMouseDrag |
| `GameData.Player.GetRenownRefundCost` | data | 1 | OnLButtonUp |
| `TextEditBoxSelectAll` | ui | 1 | OnLButtonUp |
| `WindowAssignFocus` | ui | 1 | OnLButtonUp |
| `WindowGetAnchorCount` | ui | 1 | OnLButtonUp |
| `WindowGetHandleInput` | ui | 1 | OnLButtonUp |
| `WindowGetScale` | ui | 1 | OnLButtonUp |
| `WindowSetTintColor` | ui | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

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
### OnMouseOver

Confidence: HIGH

### OnMouseOverEnd

Confidence: HIGH

### OnRButtonDown

Confidence: LOW

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
## Lua Functions Manipulating This Type

- ClosetGoblinCharacterWindow.HideShowHelm
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- BankArkel.CreatePackWin
- ClosetGoblinCharacterWindow.OnShow
- Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy.ScenarioInfoInitialize
- ClosetGoblinCharacterWindow.UpdateSetRow
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.MarksUI_MarkConfigDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- PotionBar.local.UpdateButton
- Enemy.CombatLogUI_StatsWindow_Open
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- AggroMeter.Initialize
- ClosetGoblinCharacterWindow.UpdateActionBarSettings
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- Enemy.IntercomUI_ChooseChannelDialog_Open
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- DAoCBuff.ShowMessageWindow
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- ClosetGoblinCharacterWindow.ShowCloak
- ClosetGoblinCharacterWindow.HideCloakOptions
- ClosetGoblinCharacterWindow.ShowCloakOptions
- Enemy.IntercomUI_IntercomDialog_Open
- PotionBarSettings.OnAboutShown
- ClosetGoblinCharacterWindow.ShowShowHelm
- ClosetGoblinCharacterWindow.ShowHelm
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- ClosetGoblinCharacterWindow.OnInitialize
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- Killer.Initialize
- WSCT.ColorOnInitialize
- LibWBTogglerManager.Initialize
- Enemy.UI_ConfigDialog_Open
- Enemy.IntercomInitialize
- DAoCBuffSettings.SetLabels
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- AggroMeter.OnTabLBU
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- RoR_SoR.Restack
- TidyRollOptions.Initialize
- APAGui.OnShown
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- PotionBar.UpdateButton
- ClosetGoblinZoneWindow.OnInitialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- ClosetGoblinCharacterWindow.UseItemRackToggled
- ClosetGoblinCharacterWindow.HotbarChangeToggled3


## Binding Resolution

- Total handler declarations: 500
- Resolved to Lua functions: 494 (98%)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- LibGroup
- MiracleGrowLight
- PartyCast
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

## Examples

- AdvancedPetAssist: APAOptionsClose -> Button APAOptionsClose
- AdvancedPetAssist: APAOptionsTabsAutoRecall -> Button APAOptionsTabsAutoRecall
- AdvancedPetAssist: APAOptionsTabsControls -> Button APAOptionsTabsControls
- AdvancedPetAssist: APAOptionsTabsFollowTarget -> Button APAOptionsTabsFollowTarget
- AdvancedPetAssist: APAOptionsTabsGeneral -> Button APAOptionsTabsGeneral
- AdvancedPetAssist: APAOptionsTabsHUD -> Button APAOptionsTabsHUD

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [OverlayOffset](element_OverlayOffset.md) (HIGH 100/100) - XML Element Type
- [OverlaySize](element_OverlaySize.md) (HIGH 100/100) - XML Element Type
- [OverlayTexCoords](element_OverlayTexCoords.md) (HIGH 100/100) - XML Element Type
- [ResizeImages](element_ResizeImages.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexDims](element_TexDims.md) (HIGH 100/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (HIGH 100/100) - XML Element Type
- [TextColors](element_TextColors.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextOffset](element_TextOffset.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 98/100) - Global Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 80/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [Anchor](element_Anchor.md) (MEDIUM 55/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type
- [AnimatedImages](element_AnimatedImages.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [Anchor](element_Anchor.md) (MEDIUM 55/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type

## Triggered By

- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event

## Affects

- none
