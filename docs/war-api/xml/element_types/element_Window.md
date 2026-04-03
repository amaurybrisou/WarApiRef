# Window

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, AutoMark, BankArkel |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/AnywhereTrainerAdditions/AnywhereTrainerAdditions.xml:0`, `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0` |
| Namespaces detected | Window |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAInstantOnlyHUD, AdvancedPetAssist: APAKitingHUD, AdvancedPetAssist: APAOptions, AdvancedPetAssist: APAOptionsBackground, AdvancedPetAssist: APAOptionsContent |
| XML usage count | 1007 |
| XML attribute usage count | 1007 |
| Lua usage count | 18 |
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

Observed XML element type instantiated by 32 addons.

## Common Attributes

- alpha
- drawchildrenfirst
- handleinput
- inherits
- layer
- movable
- name
- popable
- savesettings
- scale
- skipinput
- sticky

## Common Handlers

- [OnHidden](../handlers/handler_OnHidden.md)
- [OnInitialize](../handlers/handler_OnInitialize.md)
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOut](../handlers/handler_OnMouseOut.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnShown](../handlers/handler_OnShown.md)

## Common Handler Functions

- EA_LabelCheckButton.Initialize
- DAoCBuffSettings.ToggleCheckBox
- WindowUtils.OnShown
- WindowUtils.OnHidden
- BuffHead.Setup.Layout.OnLayersChanged
- WindowUtils.TrapClick
- BuffHead.Setup.Layout.Properties.OnColorExampleMouseOut
- BuffHead.Setup.Layout.Properties.OnColorExampleMouseOver
- BuffHead.Setup.PriorityEffectsItem.OnTargetTypeLUp
- Killer.HideRowTooltip
- RoR_SoR.BroadCastOption
- RoR_SoR.OnMouseOverStart


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnHidden](../handlers/handler_OnHidden.md) | lifecycle | WindowUtils.OnHidden, APAGui.OnFollowTargetHUDHidden, APAGui.OnHidden, APAGui.OnInstantOnlyHUDHidden, APAGui.OnKitingHUDHidden, AdvancedRenownTraining.OnExportHidden | `function()` | MEDIUM |
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_LabelCheckButton.Initialize, AdvancedRenownTraining.Initialize, AuraConfig.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, EA_Window_DefaultLabelToggleCircle.Initialize, EA_Window_Macro.Initialize | `function()` | MEDIUM |
| [OnKeyEscape](../handlers/handler_OnKeyEscape.md) | custom | Enemy.CombatLogUI_SnapshotWindow_Hide, Enemy.CombatLogUI_StatsWindow_Hide, Enemy.GroupsUI_EffectFilterDialog_Hide, Enemy.IntercomUI_ChooseChannelDialog_Hide, Enemy.IntercomUI_IntercomDialog_Hide, Enemy.IntercomUI_IntercomJoinDialog_Hide | `function(...)` | LOW |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | WindowUtils.TrapClick, Enemy.UnitFramesUI_Anchor_OnLButtonDown, AdvancedRenownTraining.Select, BuffHead.Setup.AdvancedCompression.OnRowLDown, BuffHead.Setup.AdvancedCompressionItem.OnRowLDown, BuffHead.Setup.AdvancedContainers.OnRowLDown | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | DAoCBuffSettings.ToggleCheckBox, BuffHead.Setup.Layout.OnLayersChanged, BuffHead.Setup.PriorityEffectsItem.OnTargetTypeLUp, EA_LabelCheckButton.Toggle, AnywhereTrainer.OnLButtonUp, Enemy.UnitFramesUI_Anchor_OnLButtonUp | `flags, x, y` | MEDIUM |
| [OnMButtonDown](../handlers/handler_OnMButtonDown.md) | input | Enemy.UnitFramesUI_UnitFrame_OnMButtonDown, MoraleCircle.Reset | `flags, x, y` | MEDIUM |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | Enemy.UnitFramesUI_UnitFrame_OnMButtonUp, TidyRollFrame.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOut](../handlers/handler_OnMouseOut.md) | input | Killer.HideRowTooltip | `function()` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOver, RoR_SoR.OnMouseOverStart, AggroMeter.OnMouseOverStart, Enemy.ConfigurationWindow_ShowTooltip, Enemy.UnitFramesUI_Anchor_OnMouseOver, TexturedButtons.Setup.Fonts.OnColorExampleMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOut, Enemy.UnitFramesUI_Anchor_OnMouseOverEnd, TexturedButtons.Setup.Fonts.OnColorExampleMouseOut, BuffHead.Setup.AdvancedCompression.OnRowMouseOut, BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut, BuffHead.Setup.AdvancedContainers.OnRowMouseOut | `function(...)` | LOW |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | input | FrameManager.OnMouseWheel, MoraleCircle.OnMouseWheel, TidyChat.Copy.OnMouseWheel | `function(delta)` | MEDIUM |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | BuffHead.Setup.AdvancedCompression.OnRowRDown, BuffHead.Setup.AdvancedCompressionItem.OnRowRDown, BuffHead.Setup.AdvancedContainers.OnRowRDown, BuffHead.Setup.AdvancedContainersItem.OnRowRDown, BuffHead.Setup.Layout.OnControlFrameRButtonDown, BuffHead.Setup.Layout.OnLayoutWindowRButtonDown | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | RoR_SoR.BroadCastOption, RoR_SoR.POPOption, AggroMeter.OnTabRBU, AnywhereTrainer.OnRButtonUp, AnywhereTrainerAdditions.OnRButtonUp, AuraSettings.OnRButtonUpAuraList | `flags, x, y` | MEDIUM |
| [OnRawDeviceInput](../handlers/handler_OnRawDeviceInput.md) | custom | BuffHead.Setup.Layout.OnRawDeviceInput | `function(...)` | LOW |
| [OnShown](../handlers/handler_OnShown.md) | lifecycle | WindowUtils.OnShown, APAGui.OnFollowTargetHUDShown, APAGui.OnInstantOnlyHUDShown, APAGui.OnKitingHUDShown, APAGui.OnPetTargetHUDShown, APAGui.OnShown | `function()` | MEDIUM |
| [OnShutdown](../handlers/handler_OnShutdown.md) | lifecycle | ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, EA_Window_Macro.Shutdown | `function()` | MEDIUM |
| [OnSizeUpdated](../handlers/handler_OnSizeUpdated.md) | custom | RoR_SoR.OnSizeUpdated | `function(...)` | LOW |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | BuffHead.Setup.Layout.OnUpdate, TidyRoll.OnUpdate | `function(elapsed)` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | DAoCBuffSettings.ToggleCheckBox, BuffHead.Setup.Layout.OnLayersChanged, BuffHead.Setup.PriorityEffectsItem.OnTargetTypeLUp, EA_LabelCheckButton.Toggle, AnywhereTrainer.OnLButtonUp, Enemy.UnitFramesUI_Anchor_OnLButtonUp, WSCT.CheckBoxOnLButtonUp, APAGui.OnFollowTargetHUDClicked, APAGui.OnInstantOnlyHUDClicked, APAGui.OnKitingHUDClicked, AuraConfig.OnCircledImageChanged, AuraConfig.OnDisplayTimerChanged, AuraConfig.OnMirrorImageChanged, AuraSettings.ConfigChange_EnableAddon, AuraSettings.ConfigChange_EnableDebugging, AuraSettings.OnLButtonUpAuraList, AuraShares.OnFilterCharactersAurasToggle, AuraShares.OnFilterSameNameToggle, AuraShares.OnLButtonUpAuraList, BuffHead.Setup.AdvancedCompression.OnRowLUp, BuffHead.Setup.AdvancedCompressionItem.OnRowLUp, BuffHead.Setup.AdvancedContainers.OnRowLUp, BuffHead.Setup.AdvancedContainersItem.OnRowLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysShowEnableLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsPermanentLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputEnableRemovableLUp, BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputShowTooltipsLUp, BuffHead.Setup.Display.OnEnableSortLUp, BuffHead.Setup.EffectCache.OnEnableCachingLUp, BuffHead.Setup.EffectCache.OnRowLUp, BuffHead.Setup.Filter.OnEnableFilter, BuffHead.Setup.Filter.OnRowLUp, BuffHead.Setup.Layout.Manager.OnExportLUp, BuffHead.Setup.Layout.Manager.OnImportLUp, BuffHead.Setup.Layout.Manager.OnLayoutsLUp, BuffHead.Setup.Layout.OnControlFrameLButtonUp, BuffHead.Setup.Layout.Properties.OnFontColorLUp, BuffHead.Setup.Layout.Properties.OnIconBorderColorLUp, BuffHead.Setup.Layout.Properties.OnStatusBarBackgroundColorLUp, BuffHead.Setup.Layout.Properties.OnStatusBarEnableLUp, BuffHead.Setup.Layout.Properties.OnStatusBarForegroundColorLUp, BuffHead.Setup.Layout.Properties.OnStatusBarForegroundStretchLUp, BuffHead.Setup.Layout.Properties.OnStatusBarForegroundTextureLUp, BuffHead.Setup.Layout.Properties.OnStatusBarOrientationReverseLUp, BuffHead.Setup.OnRowLUp, BuffHead.Setup.Performance.OnEnableFadingLUp, BuffHead.Setup.Performance.OnEnablePriorityLUp, BuffHead.Setup.Performance.OnEnableSyncLUp, BuffHead.Setup.PriorityEffects.OnRowLUp, BuffHead.Setup.PriorityEffects.OnSortFirstLUp, BuffHead.Setup.PriorityEffectsItem.OnRowLUp, BuffHead.Setup.SelectTexture.OnTextureRowLUp, BuffHead.Setup.Trackers.OnTargetChangeClearAlwaysShowLUp, BuffHead.Setup.Trackers.OnTargetChangeClearBuffsLUp, BuffHead.Setup.Trackers.OnTargetChangeClearDebuffsLUp, BuffHead.Setup.Trackers.OnTrackerEnableLUp, BuffHead.Setup.Trackers.OnTrackerPermanentLUp, DAoCBuffFrame.OnLButtonUp, DAoCBuffSettings.FilterRowOnLButtonUp, DAoCBuffSettings.FrameSettingsRowItemOnLButtonUp, DAoCBuffSettings.ShowGeneralOptions, DAoCBuffSettings.ShowListManager, Enemy.CombatLogUI_IDSAnchor_OnLButtonUp, Enemy.ConfigurationWindow_OnBoolClick, Enemy.KillSpamUI_KillSpamDialog_OnLButtonUp, Enemy.KillSpamUI_KillSpamDialog_OnRowLButtonUp, Enemy.MarksUI_EnemyMarkIcon_OnLButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnLButtonUp, Enemy.StopwatchButtonClick, Enemy.TimerUI_OnLButtonUp, Enemy.UnitFramesUI_UnitFrame_OnLButtonUp, FrameManager.OnLButtonUp, Killer.OnAllTimeSummaryLButtonUp, Killer.OnRecentSummaryLButtonUp, LibGroup.Setup.OnEnableGroupDistanceLUp, MoraleCircle.Click, PP.PersistantToggle, PotionBarFloating.ActivatorLButtonUp, ShiniesAutoUI.OnLButtonUp_AutoListRow, ShiniesBrowseUI.OnLButtonUp_Results_ListItem, ShiniesBrowseUI.OnLButtonUp_Searches, ShiniesConfigUI.OnLButtonUp_ListItem, TexturedButtons.Setup.Actionbar.OnEnableLUp, TexturedButtons.Setup.Actionbar.OnHideBackgroundLUp, TexturedButtons.Setup.Actionbar.OnHideEmptyLUp, TexturedButtons.Setup.AdvancedTextures.OnDisabledLUp, TexturedButtons.Setup.AdvancedTextures.OnUseCustomLUp, TexturedButtons.Setup.AdvancedTextures.OnUsePresetLUp, TexturedButtons.Setup.Cooldown.OnCooldownColorLUp, TexturedButtons.Setup.Cooldown.OnEnableButtonTintingLUp, TexturedButtons.Setup.Cooldown.OnEnableLUp, TexturedButtons.Setup.Cooldown.OnHideFlashLUp, TexturedButtons.Setup.Cooldown.OnRemoveCooldownSLUp, TexturedButtons.Setup.Cooldown.OnShowGlobalCooldownTextLUp, TexturedButtons.Setup.Fonts.OnCooldownColorLUp, TexturedButtons.Setup.Fonts.OnEnableLUp, TexturedButtons.Setup.Fonts.OnHideCooldownLUp, TexturedButtons.Setup.Fonts.OnHideKeybindLUp, TexturedButtons.Setup.Fonts.OnKeybindColorLUp, TexturedButtons.Setup.Misc.OnEnableCooldownPulseLUp, TexturedButtons.Setup.Misc.OnHideActiveLUp, TexturedButtons.Setup.Misc.OnHideGlowLUp, TexturedButtons.Setup.Misc.OnHideQuicklockLUp, TexturedButtons.Setup.Misc.OnMovableQuicklockLUp, TexturedButtons.Setup.Misc.OnSaveQuicklockLUp, TexturedButtons.Setup.OnRowLUp, TexturedButtons.Setup.Textures.OnDisabledLUp, TexturedButtons.Setup.Textures.OnUseCustomLUp, TexturedButtons.Setup.Textures.OnUsePresetLUp, TexturedButtons.Setup.Tint.OnEnableLUp, TidyChat.Options.OnCheckboxLBU, TidyRollOptions.OnCheckboxLBU, TidyRollOptions.OnRadioLBU, TurretRange.OnBoxLUp, TurretRange.Setup.Display.OnCircleInvertLUp, TurretRange.Setup.Distances.OnRowLUp, TurretRange.Setup.General.OnEnableLUp, WHMGui.OnHealerNameMouseUp, WarBoard.Options.DisableCornerButtons, WarBoard.Options.SetBarDefault, WarBoard.Options.ToggleBar | `flags, x, y` | MEDIUM |
| [OnHidden](../handlers/handler_OnHidden.md) | lifecycle | WindowUtils.OnHidden, APAGui.OnFollowTargetHUDHidden, APAGui.OnHidden, APAGui.OnInstantOnlyHUDHidden, APAGui.OnKitingHUDHidden, AdvancedRenownTraining.OnExportHidden, AdvancedRenownTraining.OnHidden, AdvancedRenownTraining.OnImportHidden, AdvancedRenownTraining.OnImportNameInputHidden, AdvancedRenownTraining.OnLinkHidden, AdvancedRenownTraining.PresetOnHidden, AuraConfig.OnHidden, AuraSettings.OnHidden, AuraShares.OnHidden, BuffHead.Setup.AdvancedCompression.OnHidden, BuffHead.Setup.AdvancedCompressionItem.OnHidden, BuffHead.Setup.AdvancedCompressionItemEffect.OnHidden, BuffHead.Setup.AdvancedContainers.OnHidden, BuffHead.Setup.AdvancedContainersItem.OnHidden, BuffHead.Setup.AdvancedContainersItem.Properties.OnHidden, BuffHead.Setup.Container.OnHidden, BuffHead.Setup.Display.OnHidden, BuffHead.Setup.EffectCache.OnHidden, BuffHead.Setup.EffectCacheTable.OnHidden, BuffHead.Setup.Filter.OnHidden, BuffHead.Setup.General.OnHidden, BuffHead.Setup.Layout.Manager.OnHidden, BuffHead.Setup.Layout.OnHidden, BuffHead.Setup.Layout.Properties.OnHidden, BuffHead.Setup.OnHidden, BuffHead.Setup.Performance.OnHidden, BuffHead.Setup.PriorityEffects.OnHidden, BuffHead.Setup.PriorityEffectsItem.OnHidden, BuffHead.Setup.Trackers.OnHidden, ClosetGoblinCharacterWindow.OnHide, ClosetGoblinZoneWindow.OnHide, EA_Window_Macro.OnHidden, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnHidden, Shinies.OnHidden, ShiniesAuctionsUI.OnHidden, ShiniesAutoUI.OnHidden, ShiniesBrowseUI.OnHidden, ShiniesPostUI.OnHidden, TexturedButtons.Setup.Actionbar.OnHidden, TexturedButtons.Setup.AdvancedTextures.OnHidden, TexturedButtons.Setup.Cooldown.OnHidden, TexturedButtons.Setup.Fonts.OnHidden, TexturedButtons.Setup.Misc.OnHidden, TexturedButtons.Setup.OnHidden, TexturedButtons.Setup.Textures.OnHidden, TexturedButtons.Setup.Tint.OnHidden, TidyChat.Copy.OnHidden, TidyChat.LootRoll.OnHidden, TidyChat.Options.OnHidden, TidyRoll.CustomAutoRoll.OnHidden, TidyRoll.OnEsc, TidyRollOptions.OnHidden, TurretRange.Setup.Display.OnHidden, TurretRange.Setup.Distance.OnHidden, TurretRange.Setup.Distances.OnHidden, TurretRange.Setup.General.OnHidden, TurretRange.Setup.OnHidden, WHMGui.OnWindowHidden, WSCT.OnHidden | `` |  |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOver, RoR_SoR.OnMouseOverStart, AggroMeter.OnMouseOverStart, Enemy.ConfigurationWindow_ShowTooltip, Enemy.UnitFramesUI_Anchor_OnMouseOver, TexturedButtons.Setup.Fonts.OnColorExampleMouseOver, AdvancedRenownTraining.AbilityTooltip, AuraSettings.OnMouseOverAuraList, AuraShares.OnMouseOverAuraList, BuffHead.Setup.AdvancedCompression.OnRowMouseOver, BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOver, BuffHead.Setup.AdvancedContainers.OnRowMouseOver, BuffHead.Setup.AdvancedContainersItem.OnRowMouseOver, BuffHead.Setup.EffectCache.OnRowMouseOver, BuffHead.Setup.Filter.OnRowMouseOver, BuffHead.Setup.Layout.Properties.OnTextureButtonMouseOver, BuffHead.Setup.OnRowMouseOver, BuffHead.Setup.PriorityEffects.OnRowMouseOver, BuffHead.Setup.PriorityEffectsItem.OnRowMouseOver, BuffHead.Setup.SelectTexture.OnTextureRowMouseOver, DAoCBuffFrame.OnMouseOver, Enemy.CombatLogUI_EpsWindow_OnMouseOver, Enemy.CombatLogUI_IDSAnchor_OnMouseOver, Enemy.CombatLogUI_SnapshotWindow_ListRowMouseOver, Enemy.CombatLogUI_TargetDefenseTotalWindow_OnMouseOver, Enemy.CombatLogUI_TargetDefenseWindow_OnMouseOver, Enemy.KillSpamUI_KillSpamAreaStatsDialog_OnRowMouseOver, Enemy.KillSpamUI_KillSpamDialog_OnRowMouseOver, Enemy.KillSpamUI_PlayerKDR_OnMouseOver, Enemy.MarksUI_EnemyMarkIcon_OnMouseOver, Enemy.MarksUI_EnemyMarksWindow_OnMouseOver, Enemy.TimerUI_OnMouseOver, Enemy.UnitFramesUI_UnitFrame_OnMouseOver, FrameManager.OnMouseOver, Killer.OnAllTimeSummaryMouseOver, Killer.OnFeedRowMouseOver, Killer.OnPersonalCounterMouseOver, Killer.OnRecentSummaryMouseOver, MiracleGrowLight.onHover, MoraleCircle.OnMouseOverStart, PotionBarFloating.ActivatorMouseOver, PotionBarFloating.Autoshow, RoR_SoR.ShowPopper, ShiniesAuctionsUI.OnMouseOver_Results_ListItem, ShiniesAutoUI.OnMouseOver_AutoListRow, ShiniesBrowseUI.OnMouseOver_Results_ListItem, ShiniesBrowseUI.OnMouseOver_Searches, ShiniesPostUI.OnMouseOver_Results_ListItem, TexturedButtons.Setup.Cooldown.OnColorExampleMouseOver, TexturedButtons.Setup.OnRowMouseOver, TurretRange.Setup.Distances.OnRowMouseOver, WSCT.CritOnMouseOver, WSCT.OnMouseOver | `` |  |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | RoR_SoR.BroadCastOption, RoR_SoR.POPOption, AggroMeter.OnTabRBU, AnywhereTrainer.OnRButtonUp, AnywhereTrainerAdditions.OnRButtonUp, AuraSettings.OnRButtonUpAuraList, AuraShares.OnRButtonUpAuraList, BuffHead.Setup.AdvancedCompression.OnRowRUp, BuffHead.Setup.AdvancedCompressionItem.OnRowRUp, BuffHead.Setup.AdvancedContainers.OnRowRUp, BuffHead.Setup.AdvancedContainersItem.OnRowRUp, BuffHead.Setup.EffectCache.OnRowRUp, BuffHead.Setup.Filter.OnRowRUp, BuffHead.Setup.Layout.OnControlFrameRButtonUp, BuffHead.Setup.OnRowRUp, BuffHead.Setup.PriorityEffects.OnRowRUp, BuffHead.Setup.PriorityEffectsItem.OnRowRUp, BuffHead.Setup.SelectTexture.OnTextureRowRUp, DAoCBuffFrame.OnRButtonUp, Enemy.CombatLogUI_TargetDefenseTotalWindow_OnRButtonUp, Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp, Enemy.KillSpamUI_KillSpamDialog_OnRButtonUp, Enemy.KillSpamUI_KillSpamDialog_OnRowRButtonUp, Enemy.KillSpamUI_PlayerKDR_OnRButtonUp, Enemy.MarkUI_EnemyMark_OnRButtonUp, Enemy.MarksUI_EnemyMarkIcon_OnRButtonUp, Enemy.StopwatchReset, Enemy.TimerUI_OnRButtonUp, Enemy.UI_Debug_OnRButtonUp, Enemy.UnitFramesUI_UnitFrame_OnRButtonUp, FrameManager.OnRButtonUp, Killer.OnPersonalCounterRButtonUp, MiracleGrowLight.switchBackground, MoraleCircle.RightClick, PotionBarFloating.ActivatorRButtonUp, RoR_SoR.OnTabRBU, ShiniesAuctionsUI.OnRButtonUp_Results_ListItem, ShiniesBrowseUI.OnRButtonUp_Results_ListItem, ShiniesBrowseUI.OnRButtonUp_Searches, ShiniesPostUI.OnRButtonUp_Results_ListItem, TexturedButtons.Setup.OnRowRUp, TurretRange.Setup.Distances.OnRowRUp, WarBoard.OpenLayoutMenu | `flags, x, y` | MEDIUM |
| [OnShown](../handlers/handler_OnShown.md) | lifecycle | WindowUtils.OnShown, APAGui.OnFollowTargetHUDShown, APAGui.OnInstantOnlyHUDShown, APAGui.OnKitingHUDShown, APAGui.OnPetTargetHUDShown, APAGui.OnShown, AdvancedRenownTraining.OnExportShown, AdvancedRenownTraining.OnShown, AdvancedRenownTraining.PresetOnShown, AuraTexture.OnShown, ClosetGoblinCharacterWindow.OnShow, ClosetGoblinZoneWindow.OnShow, EA_Window_Macro.OnShown, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnShown, PP.OnShown, PotionBarSettings.OnAboutShown, Shinies.OnShown, ShiniesAuctionsUI.OnShown, ShiniesAutoUI.OnShown, ShiniesBrowseUI.OnShown, ShiniesPostUI.OnShown, TidyChat.Copy.OnShown, TidyChat.LootRoll.OnShown, TidyChat.Options.OnShown, TidyRoll.CustomAutoRoll.OnShown, TidyRollOptions.OnShown, WHMGui.OnDetailsShown, WHMGui.OnOptionsShown, WHMGui.OnWindowShown, WSCT.OnShown, WarBoard.Options.ShowTopOptions | `` |  |
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_LabelCheckButton.Initialize, AdvancedRenownTraining.Initialize, AuraConfig.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, EA_Window_DefaultLabelToggleCircle.Initialize, EA_Window_Macro.Initialize, WHMGui.OnDetailsInitialize, WHMGui.OnOptionsInitialize, WHMGui.OnWindowInitialize, WSCT.ColorOnInitialize, WSCT.OptionsOnInitialize | `` |  |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | WindowUtils.TrapClick, Enemy.UnitFramesUI_Anchor_OnLButtonDown, AdvancedRenownTraining.Select, BuffHead.Setup.AdvancedCompression.OnRowLDown, BuffHead.Setup.AdvancedCompressionItem.OnRowLDown, BuffHead.Setup.AdvancedContainers.OnRowLDown, BuffHead.Setup.AdvancedContainersItem.OnRowLDown, BuffHead.Setup.EffectCache.OnRowLDown, BuffHead.Setup.Layout.ClearSelection, BuffHead.Setup.Layout.OnControlFrameLButtonDown, BuffHead.Setup.Layout.OnLayoutWindowLButtonDown, BuffHead.Setup.OnRowLDown, BuffHead.Setup.PriorityEffects.OnRowLDown, BuffHead.Setup.PriorityEffectsItem.OnRowLDown, BuffHead.Setup.SelectTexture.OnTextureRowLDown, Enemy.AssistUI_Target_OnLButtonDown, Enemy.CombatLogUI_IDSAnchor_OnLButtonDown, Enemy.GroupIcons_GroupIcon_OnLButtonDown, Enemy.KillSpamUI_KillSpamDialog_OnLButtonDown, Enemy.MarkUI_EnemyMark_OnLButtonDown, Enemy.MarksUI_EnemyMarksWindow_OnLButtonDown, Enemy.UnitFramesUI_UnitFrame_OnLButtonDown, MoraleCircle.Click, PotionBarFloating.ActivatorLButtonDown, TexturedButtons.Setup.OnRowLDown, TurretRange.Setup.Distances.OnRowLDown, WHMGui.OnHealerNameMouseDown | `flags, x, y` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOut, Enemy.UnitFramesUI_Anchor_OnMouseOverEnd, TexturedButtons.Setup.Fonts.OnColorExampleMouseOut, BuffHead.Setup.AdvancedCompression.OnRowMouseOut, BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut, BuffHead.Setup.AdvancedContainers.OnRowMouseOut, BuffHead.Setup.AdvancedContainersItem.OnRowMouseOut, BuffHead.Setup.EffectCache.OnRowMouseOut, BuffHead.Setup.Filter.OnRowMouseOut, BuffHead.Setup.Layout.Properties.OnTextureButtonMouseOut, BuffHead.Setup.OnRowMouseOut, BuffHead.Setup.PriorityEffects.OnRowMouseOut, BuffHead.Setup.PriorityEffectsItem.OnRowMouseOut, BuffHead.Setup.SelectTexture.OnTextureRowMouseOut, DAoCBuffFrame.OnMouseOverEnd, Enemy.CombatLogUI_IDSAnchor_OnMouseOverEnd, Enemy.MarksUI_EnemyMarksWindow_OnMouseOverEnd, FrameManager.OnMouseOverEnd, ShiniesAuctionsUI.OnMouseOverEnd_Results_ListItem, ShiniesAutoUI.OnMouseOverEnd_AutoListRow, ShiniesBrowseUI.OnMouseOverEnd_Results_ListItem, ShiniesBrowseUI.OnMouseOverEnd_Searches, ShiniesPostUI.OnMouseOverEnd_Results_ListItem, TexturedButtons.Setup.Cooldown.OnColorExampleMouseOut, TexturedButtons.Setup.OnRowMouseOut, TurretRange.Setup.Distances.OnRowMouseOut | `` |  |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | BuffHead.Setup.AdvancedCompression.OnRowRDown, BuffHead.Setup.AdvancedCompressionItem.OnRowRDown, BuffHead.Setup.AdvancedContainers.OnRowRDown, BuffHead.Setup.AdvancedContainersItem.OnRowRDown, BuffHead.Setup.Layout.OnControlFrameRButtonDown, BuffHead.Setup.Layout.OnLayoutWindowRButtonDown, BuffHead.Setup.Layout.TrapClick, BuffHead.Setup.OnRowRDown, BuffHead.Setup.PriorityEffects.OnRowRDown, BuffHead.Setup.PriorityEffectsItem.OnRowRDown, BuffHead.Setup.SelectTexture.OnTextureRowRDown, Enemy.UnitFramesUI_UnitFrame_OnRButtonDown, FrameManager.OnRButtonDown, TexturedButtons.Setup.OnRowRDown, TurretRange.Setup.Distances.OnRowRDown | `flags, x, y` | MEDIUM |
| [OnKeyEscape](../handlers/handler_OnKeyEscape.md) | custom | Enemy.CombatLogUI_SnapshotWindow_Hide, Enemy.CombatLogUI_StatsWindow_Hide, Enemy.GroupsUI_EffectFilterDialog_Hide, Enemy.IntercomUI_ChooseChannelDialog_Hide, Enemy.IntercomUI_IntercomDialog_Hide, Enemy.IntercomUI_IntercomJoinDialog_Hide, Enemy.MarksUI_MarkConfigDialog_Hide, Enemy.UI_ChooseIconDialog_Hide, Enemy.UI_ConfigDialog_Hide, Enemy.UI_TextEntryDialog_Close, Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide, Enemy.UnitFramesUI_UnitFramePartDialog_Hide | `wn` | LOW |
| [OnMouseOut](../handlers/handler_OnMouseOut.md) | input | Killer.HideRowTooltip | `` |  |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | input | FrameManager.OnMouseWheel, MoraleCircle.OnMouseWheel, TidyChat.Copy.OnMouseWheel | `x, y, delta` | MEDIUM |
| [OnShutdown](../handlers/handler_OnShutdown.md) | lifecycle | ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, EA_Window_Macro.Shutdown | `` |  |
| [OnMButtonDown](../handlers/handler_OnMButtonDown.md) | input | Enemy.UnitFramesUI_UnitFrame_OnMButtonDown, MoraleCircle.Reset | `flags, x, y` | MEDIUM |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | Enemy.UnitFramesUI_UnitFrame_OnMButtonUp, TidyRollFrame.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | BuffHead.Setup.Layout.OnUpdate, TidyRoll.OnUpdate | `elapsed` | MEDIUM |
| [OnRawDeviceInput](../handlers/handler_OnRawDeviceInput.md) | custom | BuffHead.Setup.Layout.OnRawDeviceInput | `deviceId, itemId, itemDown` | LOW |
| [OnSizeUpdated](../handlers/handler_OnSizeUpdated.md) | custom | RoR_SoR.OnSizeUpdated | `` |  |

### Per-Event Lua API Calls

**OnHidden** handlers call: `BroadcastEvent`, `ButtonSetDisabledFlag`, `ButtonSetText`, `ComboBoxSetDisabledFlag`, `DoesWindowExist`, `TextEditBoxSetText`, `WindowSetShowing`, `WindowUtils.OnHidden`

**OnInitialize** handlers call: `ButtonSetText`, `CreateWindow`, `DoesWindowExist`, `LabelSetText`, `RegisterEventHandler`, `UnregisterEventHandler`, `WindowSetAlpha`, `WindowSetShowing`, `WindowSetTintColor`

**OnKeyEscape** handlers call: `DestroyWindow`, `DoesWindowExist`, `WindowSetShowing`

**OnLButtonDown** handlers call: `BroadcastEvent`, `WindowGetId`, `WindowGetParent`, `WindowSetGameActionData`, `WindowSetTintColor`

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `LabelGetText`, `WindowGetId`, `WindowGetParent`, `WindowSetMovable`, `WindowSetShowing`

**OnMouseOver** handlers call: `DoesWindowExist`, `LabelSetTextColor`, `WindowGetAlpha`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetAlpha`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnMouseOverEnd** handlers call: `DoesWindowExist`, `LabelSetTextColor`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnRButtonDown** handlers call: `WindowGetId`, `WindowSetTintColor`

**OnRButtonUp** handlers call: `WindowGetId`, `WindowGetMovable`, `WindowGetParent`, `WindowSetShowing`

**OnShown** handlers call: `BroadcastEvent`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxSetDisabledFlag`, `DoesWindowExist`, `LabelSetText`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowSetDimensions`, `WindowSetShowing`, `WindowUtils.OnShown`

**OnShutdown** handlers call: `UnregisterEventHandler`

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `LabelGetText`, `WindowGetId`, `WindowGetParent`, `WindowSetMovable`, `WindowSetShowing`

**OnHidden** handlers call: `BroadcastEvent`, `ButtonSetDisabledFlag`, `ButtonSetText`, `ComboBoxSetDisabledFlag`, `DoesWindowExist`, `TextEditBoxSetText`, `WindowSetShowing`, `WindowUtils.OnHidden`

**OnMouseOver** handlers call: `DoesWindowExist`, `LabelSetTextColor`, `WindowGetAlpha`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetAlpha`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnRButtonUp** handlers call: `WindowGetId`, `WindowGetMovable`, `WindowGetParent`, `WindowSetShowing`

**OnShown** handlers call: `BroadcastEvent`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxSetDisabledFlag`, `DoesWindowExist`, `LabelSetText`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowSetDimensions`, `WindowSetShowing`, `WindowUtils.OnShown`

**OnInitialize** handlers call: `ButtonSetText`, `CreateWindow`, `DoesWindowExist`, `LabelSetText`, `RegisterEventHandler`, `UnregisterEventHandler`, `WindowSetAlpha`, `WindowSetShowing`, `WindowSetTintColor`

**OnLButtonDown** handlers call: `BroadcastEvent`, `WindowGetId`, `WindowGetParent`, `WindowSetGameActionData`, `WindowSetTintColor`

**OnMouseOverEnd** handlers call: `DoesWindowExist`, `LabelSetTextColor`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnRButtonDown** handlers call: `WindowGetId`, `WindowSetTintColor`

**OnKeyEscape** handlers call: `DestroyWindow`, `DoesWindowExist`, `WindowSetShowing`

**OnShutdown** handlers call: `UnregisterEventHandler`

## Common Inherits

- Aura_LabelCheckButton
- EA_LabelCheckButton
- EA_TitleBar_Default
- EA_Window_Default
- EA_Window_DefaultBackgroundFrame
- EA_Window_DefaultButtonBottomFrame
- EA_Window_DefaultFrame
- EA_Window_DefaultSeparator
- EnemyScenarioInfoDialog_CareerStatsTemplate
- TChatCheckboxTemplate
- WSCTCheckBox
- WSCTEvent

## Common Parent Elements

- [Windows](element_Windows.md) — 1001× (HIGH)
- [Window](element_Window.md) — 5× (MEDIUM)

## Common Named Child Elements

- [Label](element_Label.md) — 15× (HIGH)
- [Button](element_Button.md) — 8× (MEDIUM)
- [Window](element_Window.md) — 5× (MEDIUM)
- [ComboBox](element_ComboBox.md) — 4× (MEDIUM)
- [FullResizeImage](element_FullResizeImage.md) — 2× (LOW)
- [SliderBar](element_SliderBar.md) — 2× (LOW)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 745× (HIGH)
- [Size](element_Size.md) — 624× (HIGH)
- [Windows](element_Windows.md) — 476× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 280× (HIGH)
- [Anchor](element_Anchor.md) — 10× (HIGH)
- [Visual](element_Visual.md) — 1× (LOW)

## Common Template Bases

- Aggro_Timer_Template
- AnywhereTrainerTabTemplate
- Aura_LabelCheckButton
- Aura_LargeLabelCheckButton
- BuffHeadColorExample
- BuffHeadSetupSelectTextureRowTemplate
- ClosetGoblinActionBarPageSelector
- DAoCBuffCondenseTooltipItem
- DAoCBuffFrameSettings_G1Filter
- DAoCBuffFrameSettings_G2Filter
- DAoCBuffFrameSettings_G4Filter
- DAoCBuffFrameSettings_G5Filter
- DAoCBuff_FakeSettingsRow
- EA_LabelCheckButton
- EA_TitleBar_Default
- EA_Window_CityRating
- EA_Window_ComboBoxMenuBackground
- EA_Window_Default
- EA_Window_DefaultBackgroundFrame
- EA_Window_DefaultButtonBottomFrame
- EA_Window_DefaultContextMenuFrame
- EA_Window_DefaultFrame
- EA_Window_DefaultFrameStatusBar
- EA_Window_DefaultSeparator
- EA_Window_DefaultTooltipBackground
- EA_Window_DefaultVerticalSeparator
- EA_Window_TabSeparatorLeftSide
- EA_Window_TabSeparatorRightSide
- EnemyCombatLogStatsWindow_ListHeaderTemplate
- EnemyMark
- EnemyScenarioInfoDialog_CareerStatsTemplate
- EnemyScenarioInfoDialog_PlayerStatsHeaderTemplate
- EnemyScenarioInfoDialog_StatsRowTemplate
- EnemyScenarioInfoDialog_StatsRowTemplateBig
- EnemyScenarioInfoDialog_StatsRowTemplateBig2
- EnemyScenarioInfoDialog_StatsRowTemplateBig3
- Frame_BG_Temlate
- MiracleGrowLightLine
- MoneyFrame
- OptionsTemplate
- PartyFrameStatusBar
- RoR_SoR_BO_Template
- Shinies_TitleBar_Default
- SliderWindowTemplate
- TChatCheckboxTemplate
- TChatTabLogsTemplate
- TChatTabMiscTemplate
- TChatTabTextEntryTemplate
- TChatTabWindowsTemplate
- TRollOverlay
- TexturedButtonsColorExample
- TooltipBase
- TurretRangeDisplay
- WSCTCheckBox
- WSCTComboBoxTemplate
- WSCTEvent
- WSCTSliderTemplate


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<Window alpha="1.0" ignorealpha="false" name="...">
  <Size/>
  <Visual/>
  <Windows/>
</Window>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 59% | EA_Window_DefaultFrame, EA_TitleBar_Default, EA_Window_Default, EA_Window_TabSeparatorLeftSide, ... |
| `layer` | optional | 31% | secondary, background, popup, overlay, ... |
| `movable` | optional | 20% | true, false |
| `handleinput` | optional | 20% | true, false |
| `savesettings` | optional | 13% | true, false |
| `popable` | optional | 6% | false, true |
| `sticky` | optional | 6% | false, true |
| `scale` | optional | 3% | 1.0 |
| `skipinput` | optional | 1% | false, true |
| `drawchildrenfirst` | optional | 0% | true |
| `alpha` | optional | 0% | 0.85, 1.0, 1, 0.6 |
| `draganddrop` | optional | 0% | true |
| `parent` | optional | 0% | Root, root |
| `id` | optional | 0% | 1 |
| `show` | optional | 0% | false |
| `ignorealpha` | optional | 0% | false |
| `inhserits` | optional | 0% | EA_Window_Default |
| `localscriptvars` | optional | 0% | true |
| `wordwrap` | optional | 0% | false |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 745 times as an unnamed child.

### [Size](element_Size.md)

Observed 624 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 476 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 280 times as an unnamed child.

### [Anchor](element_Anchor.md)

Observed 10 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, bottomright, right |
| `relativePoint` | **required** | center, topleft, bottomright, left |
| `relativeTo` | **required** | Root, $parent, $parentGeneral, $parentFollowTarget |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
### [Visual](element_Visual.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [Window](element_Window.md)
- [Button](element_Button.md) (named, 8×, MEDIUM)
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
    - (cycle)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `RegisterEventHandler` | event | 285 | OnInitialize |
| `TextEditBoxSetText` | ui | 277 | OnHidden, OnShown |
| `LabelSetText` | ui | 248 | OnInitialize, OnShown |
| `ButtonSetText` | ui | 113 | OnHidden, OnInitialize, OnShown |
| `ButtonSetPressedFlag` | ui | 104 | OnLButtonUp, OnShown |
| `ButtonGetPressedFlag` | ui | 85 | OnLButtonUp |
| `WindowSetShowing` | ui | 69 | OnHidden, OnInitialize, OnKeyEscape, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp, OnShown |
| `DoesWindowExist` | ui | 63 | OnHidden, OnInitialize, OnKeyEscape, OnMouseOver, OnMouseOverEnd, OnShown |
| `WindowSetDimensions` | ui | 54 | OnShown |
| `WindowSetTintColor` | ui | 30 | OnInitialize, OnLButtonDown, OnMouseOver, OnMouseOverEnd, OnRButtonDown |
| `ButtonGetDisabledFlag` | ui | 29 | OnLButtonUp |
| `WindowGetId` | ui | 20 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonDown, OnRButtonUp |
| `WindowUtils.OnHidden` | ui | 16 | OnHidden |
| `WindowUtils.OnShown` | ui | 16 | OnShown |
| `WindowGetParent` | ui | 14 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `WindowStartAlphaAnimation` | ui | 11 | OnMouseOver, OnMouseOverEnd |
| `UnregisterEventHandler` | event | 5 | OnInitialize, OnShutdown |
| `CreateWindow` | ui | 4 | OnInitialize |
| `BroadcastEvent` | event | 3 | OnHidden, OnLButtonDown, OnShown |
| `ButtonSetDisabledFlag` | ui | 3 | OnHidden, OnLButtonUp |
| `DestroyWindow` | ui | 3 | OnKeyEscape |
| `WindowSetAlpha` | ui | 3 | OnInitialize, OnMouseOver |
| `ComboBoxSetDisabledFlag` | ui | 2 | OnHidden, OnShown |
| `LabelSetTextColor` | ui | 2 | OnMouseOver, OnMouseOverEnd |
| `WindowGetAlpha` | ui | 2 | OnMouseOver |
| `LabelGetText` | ui | 1 | OnLButtonUp |
| `WindowAddAnchor` | ui | 1 | OnShown |
| `WindowClearAnchors` | ui | 1 | OnShown |
| `WindowGetDimensions` | ui | 1 | OnMouseOver |
| `WindowGetMovable` | ui | 1 | OnRButtonUp |
| `WindowGetScreenPosition` | ui | 1 | OnMouseOver |
| `WindowSetGameActionData` | ui | 1 | OnLButtonDown |
| `WindowSetMovable` | ui | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHidden

Confidence: HIGH

### OnInitialize

Confidence: HIGH

### OnKeyEscape

Confidence: LOW

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
### OnMButtonDown

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
### OnMouseOut

Confidence: HIGH

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
### OnRawDeviceInput

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `deviceId` | string | identifier |
| 1 | `itemId` | string | identifier |
| 2 | `itemDown` | unknown | unknown |
### OnShown

Confidence: HIGH

### OnShutdown

Confidence: HIGH

### OnSizeUpdated

Confidence: LOW

### OnUpdate

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
## Lua Functions Manipulating This Type

- ClosetGoblinCharacterWindow.UseItemRackToggled
- Enemy.CombatLogUI_EpsWindow_Update
- RoR_SoR.OnInitialize
- APAGui.ToggleKitingHUD
- ClosetGoblinCharacterWindow.UpdateSlotIcons
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- Enemy.IntercomUI_IntercomDialog_Hide
- MoraleCircle.OnSetCustomColorEmpty
- PotionBarSettings.OnAboutShown
- APAGui.Hide
- ClosetGoblinCharacterWindow.ShowShowHelm
- Killer.ApplyFeedLayout
- Swift Assist.SetTexLabel
- SwiftAssist.aaLabelColorSet
- AdvancedRenownTraining.OnShown
- ClosetGoblinCharacterWindow.Hide
- Enemy.Timer_Update
- PotionBarSettings.OnDontSplitNameCheck
- RoR_SoR.OnCombat
- AnywhereTrainer.OnLeftClickAuction
- Swift Assist.local.SetSmartLabel
- WHMGui.OnOptionsInitialize
- AdvancedRenownTraining.ChangeTab
- Enemy.UnitFramesUI_EffectsIndicatorDialog_IsOpen
- Enemy.CombatLogUI_TargetDefenseTotalWindow_IsOpen
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Update
- WarBoard.OnOptionsButton
- DAoCBuffFrame.MouseOverUpdate
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- GuardLine.Libguard_Toggle
- GuardLine.OnLayoutEditorFinished
- MoraleCircle.ColorChanger4
- RoR_SoR.SET_CITY
- BankArkel.CreatePackWin
- Enemy.local._OnArchetypeChanged
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- WHMGui.RefreshConfigurationWindow
- Aura.Aura:CreateEditWindows
- BuffHead.RegisterLayoutEditor
- DAoCBuff.CloseMessageWindow
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- RoR_SoR.HidePopper
- TidyRollOptions.Initialize
- AggroMeter.SplitText
- ClosetGoblinZoneWindow.RefreshOption
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Open
- GuardLine.update
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- Enemy.MarksUI_EnemyMarksWindow_IsOpen
- ClosetGoblinCharacterWindow.ShowCloak
- Enemy.UI_TextEntryDialog_Open
- Enemy.UI_ChooseIconDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- WSCT.ColorOnButtonUp
- AutoMark.local.CreateMarker
- BankArkel.SetupCombos
- RoR_SoR.OnChatLogUpdated
- SwiftAssist.OnMacroUpdated
- WarBoard.Options.EnableBoard
- MoraleCircle.OnSetCustomColorFull
- PotionBarSettings.OnResetButton
- ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- APAGui.UpdateInstantOnlyHUD
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- Enemy.IntercomUI_IntercomJoinDialog_Open
- Enemy.MarksUI_MarkConfigDialog_Open
- Killer.ShowTopKillersTooltip
- MoraleCircle.ColorChanger1
- AggroMeter.Initialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide
- Enemy.EnemyGroupIcon:Attach
- ClosetGoblinCharacterWindow.ShowCloakOptions
- ClosetGoblinCharacterWindow.HotbarPageDownProxy
- WSCT.WSCT:ObjectIDAnimation
- BankArkel.PackShow
- Enemy.CombatLogUI_EpsWindow_UpdateLayout
- Killer.ShowRowTooltip
- WSCT.WSCT:DisplayText
- WSCT.OnSetCustomColor
- Enemy.MarksUI_MarkConfigDialog_Hide
- Enemy.SetStatsRow
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- ClosetGoblinZoneWindow.Show
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- GuardLine.GetIDs
- PotionBarSettings.QuickActionsSelChanged
- RoR_SoR.CloseSetOpacityWindow
- APAGui.Show
- APAGui.UpdatePetTargetHUD
- MoraleCircle.ColorChanger3
- PotionBar.Initialize
- PotionBar.LibSlashHandler
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide
- Enemy.CombatLogUI_IDS_OnSettingsChanged
- PotionBar.UpdateButton
- RoR_SoR.ShowPopper
- WarBoard.LoadGeneralSettings
- WHMGui.HideOptionsWindow
- AdvancedRenownTrainer.local.CreateAbilityWindow
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- WhoHealedMe.local.IsMainWindowVisible
- Killer.RefreshPersonalCounter
- PotionBarSettings.OnAlphaSliderChanged
- RoR_SoR.SET_KEEP
- AdvancedPetAssist.local.AnchorInContent
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide
- DAoCBuffSettings.UC
- Enemy.GroupsUI_EffectFilterDialog_Hide
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- Enemy.TimerInitialize
- RoR_SoR.DefaultSettings
- RoR_SoR.Restack
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- DAoCBuffSettings.CloseOptionswindow
- WHMGui.ToggleOptionsWindow
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- PotionBar.Show
- DAoCBuff.DAoCBuffHeadTracker:Create
- Enemy.GroupsInitialize
- Killer.ShowPersonalStatsTooltip
- PP.UpdateDyeFilter
- AdvancedRenownTrainer.CreateAbilityWindow
- AdvancedRenownTrainer.local.CreateTab
- Enemy.CombatLogUI_StatsWindow_UpdateList
- MoraleCircle.init
- ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- Enemy.local.SetStatsRow
- Enemy.CombatLogUI_StatsWindow_Hide
- TurretRange.local.UpdateDisplay
- AggroMeter.OnTabLBU
- BankArkel.BackPackHide
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Hide
- ClosetGoblinCharacterWindow.ShowHelm
- Enemy.StopwatchEnabledChanged
- PotionBarFloating.Scale
- RoR_SoR.SET_CAMPAIGN
- Swift Assist.SetSmartLabel
- DAoCBuff.ShowMessageWindow
- MiracleGrowLight.switchBackground
- Enemy.CombatLogUI_EpsWindow_Initialize
- RoR_SoR.OnSlideWindowOptionsOpacity
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_Hide
- PotionBarSettings.RefreshSingleRightLine
- RoR_SoR.CloseSetScaleWindow
- WSCT.ColorHideMenu
- WSCT.HideMenu
- DAoCBuffSettings.Reactivate
- Enemy.CombatLogUI_EpsWindow_Hide
- Enemy.IntercomUI_IntercomDialog_Open
- Enemy.Stopwatch_Update
- GuardLine.init
- AdvancedPetAssist.ApplyHUDVisibilityFromSettings
- AuraAddon.OnLoad
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.CombatLogUI_EpsWindow_IsOpen
- MoraleCircle.ColorChanger2
- WarBoard.SlashCommand
- APAGui.UpdateFollowTargetHUD
- Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged
- Enemy.GroupsUI_EffectFilterDialog_IsOpen
- Enemy.GroupsUI_EffectFilterDialog_Ok
- Enemy.CombatLogUI_TargetDefenseWindow_Hide
- MoraleCircle.OnSetCustomColor
- PotionBar.local.UpdateButton
- BankArkel.SetCharInfo
- ClosetGoblinCharacterWindow.HotbarPageUpProxy
- Enemy.UI_ChooseIconDialog_IsOpen
- Enemy.IntercomUI_ChooseChannelDialog_Open
- Enemy.CombatLogUI_IDS_Update
- Enemy.CombatLogUI_TargetDefenseWindow_Initialize
- BankArkel.PackImg
- BankArkel.Init
- Enemy.IntercomUI_IntercomJoinDialog_Hide
- RoR_SoR.OnWindowOptionsSetScale
- BuffHead.local.RegisterLayoutEditor
- ClosetGoblinCharacterWindow.UpdateSlotIcon
- DAoCBuffSettings.PopulateSettings
- Enemy.UI_ConfigDialog_Open
- Enemy._OnKeyModifierChanged
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- ClosetGoblinCharacterWindow.OnShow
- ClosetGoblinZoneWindow.UpdateHighlightOnRow
- Enemy.AssistUI_Target_Show
- PotionBarSettings.OnShown
- WSCT.OnHidden
- WSCT.ColorOnInitialize
- WSCT.WSCT:OpenMenu
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- Enemy.UI_ChooseIconDialog_Hide
- RoR_SoR.OnWindowOptionsSetOffset
- RoR_SoR.SET_BO
- PotionBarFloating.ReflowButtons
- DAoCTooltips.CreateCondenseTooltip
- Enemy._Initialize
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_IsOpen
- GuardLine.OffTarget
- PotionBarSettings.OnAboutClose
- Shinies.Searches_UpdateWindowDisplay
- AuraTexture.OnClose
- Enemy.AssistInitialize
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- ClosetGoblinCharacterWindow.UpdateActionBarSettings
- DAoCBuffSettings.Disable
- Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.CombatLogUI_TargetDefenseWindow_Update
- APAGui.ToggleInstantOnlyHUD
- ClosetGoblinCharacterWindow.OnInitialize
- WhoHealedMe.Initialize
- Enemy.CombatLogUI_EpsWindow_Open
- GuardLine.OnLayoutEditorStart
- Killer.Initialize
- PotionBarSettings.OnUseCheckUpdateHighlighting
- Swift Assist.local.SetTexLabel
- ClosetGoblinZoneWindow.Hide
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- Enemy.EnemyMarkTemplate:ToggleMark
- WhoHealedMe.IsMainWindowVisible
- APAGui.HidePetTargetHUD
- AuraTexture.OnLoad
- WHMGui.ShowOptionsWindow
- WSCT.ColorAcceptButtonOnButtonUp
- WSCT.UpdateAnimationOptions
- ClosetGoblinCharacterWindow.HideShowHelm
- ClosetGoblinCharacterWindow.UpdateSetRow
- Enemy.UI_ConfigDialog_Hide
- Enemy.IntercomUI_IntercomJoinDialog_AddGroup
- Enemy.Assist_OnIntercomMessage
- Enemy.EnemyUnitFrame:ApplySettings
- AdvancedRenownTrainer.GeneratePresetByLinkData
- AdvancedRenownTraining.OnHidden
- WHMCommands.CmdConfig
- PotionBarSettings.OnCancelButton
- PotionBarSettings.Show1Check
- PotionBarSettings.OnAboutButton
- PotionBarSettings.UpdateLastCheckBasedOnInfoText
- PotionBarSettings.OnScaleSliderChanged
- PP.UpdateListRow
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- RoR_SoR.Text_Stream_Fetch
- WSCT.OnLButtonUpColorPicker
- WHMCore.ApplyBackgroundFillColor
- AdvancedPetAssist.local.ApplyHUDVisibilityFromSettings
- Enemy.CombatLogUI_StatsWindow_UpdateSessionsList
- ClosetGoblinCharacterWindow.Show
- Enemy.UnitFramesUI_ConfigDialog_Import
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- RoR_SoR.CloseSetOffsetWindow
- RoR_SoR.OnSlideWindowOptionsScale
- Swift Assist.local.WriteLabels
- AuraAddon.Slash
- BankArkel.PackHide
- Enemy.CombatLogUI_StatsWindow_Open
- TurretRange.OnUpdate
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- Enemy.MarksInitialize
- Enemy.MarksUI_EnemyMarksWindow_Hide
- Enemy.local._OnKeyModifierChanged
- Enemy.CombatLogUI_TargetDefenseWindow_IsOpen
- LibWBTogglerManager.Initialize
- AnywhereTrainer.OnLeftClickRenown
- ClosetGoblinCharacterWindow.HideCloakOptions
- Aura.Aura:CreateRuntimeWindows
- Enemy.MarksUI_EnemyMarksWindow_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- Enemy.CombatLogUI_StatsWindow_IsOpen
- PotionBarSettings.OnHidden
- APAGui.ApplyPetTargetHUDLayout
- AdvancedRenownTrainer.CreateTab
- RoR_SoR.OnWindowOptionsSetOpacity
- TurretRange.OnLoadComplete
- Enemy.MarksUI_MarkConfigDialog_IsOpen
- Enemy.UnitFramesUI_UnitFramePartDialog_IsOpen
- PotionBarSettings.AutohideCheck
- WarBoard.Options.OnSlide
- WhoHealedMe.Shutdown
- DAoCTooltips.Init
- DAoCBuffSettings.OpenOptionswindow
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Open
- GuardLine.update2
- Swift Assist.WriteLabels
- DAoCBuffSettings.CreateOptionswindow
- Enemy.IntercomUI_ChooseChannelDialog_OnOkButton
- PotionBarFloating.Alpha
- PotionBarSettings.OnUseCheck
- SwiftAssist.Initialize
- Enemy.IntercomInitialize
- Enemy._OnArchetypeChanged
- Enemy.CombatLogUI_TargetDefenseWindow_Open
- WHMCore.RunSettingEffects
- DAoCTooltips.UpdateCondenseTooltip
- DAoCBuffSettings.Change_Setting
- AutoMark.CreateMarker
- TidyChat.LootRoll.OnRollLinkLButtonUp
- APAGui.OnShown
- AdvancedRenownTraining.Initialize
- BankArkel.BackPackShow
- Enemy.IntercomUI_ChooseChannelDialog_Hide
- Enemy.MarksUI_EnemyMarksWindow_Update
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- Enemy.CombatLogUI_IDS_Initialize
- MoraleCircle.OnSetCustomColorFill
- AdvancedPetAssist.AnchorInContent
- AggroMeter.Close
- AnywhereTrainerAdditions.OnLeftClickAuction
- PotionBar.Hide
- RoR_SoR.SetWindowShow
- RoR_SoR.OnScenario
- ClosetGoblinZoneWindow.OnInitialize
- Enemy._CombatLogUI_IDS_UpdateWindow
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_IsOpen
- TurretRange.UpdateDisplay
- PotionBar.Shutdown
- PotionBarFloating.ShowSettings
- APAGui.UpdateKitingHUD
- DAoCBuffSettings.SetLabels


## Binding Resolution

- Total handler declarations: 536
- Resolved to Lua functions: 536 (100%)

## .mod Lifecycle: Startup Windows

This element type is instantiated as a startup window by the following .mod addon(s):

| Frame Name | Addon | Hook | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| APAFollowTargetHUD | AdvancedPetAssist | OnInitialize | exact | HIGH |
| APAKitingHUD | AdvancedPetAssist | OnInitialize | exact | HIGH |
| APAInstantOnlyHUD | AdvancedPetAssist | OnInitialize | exact | HIGH |
| APAPetTargetHUD | AdvancedPetAssist | OnInitialize | exact | HIGH |
| AdvancedRenownTrainingWindow | AdvancedRenownTrainer | OnInitialize | exact | HIGH |
| AuraTexture | Aura | OnInitialize | exact | HIGH |
| AuraColorPicker | Aura | OnInitialize | exact | HIGH |
| AuraConfig | Aura | OnInitialize | exact | HIGH |
| AuraShares | Aura | OnInitialize | exact | HIGH |
| AuraSettings | Aura | OnInitialize | exact | HIGH |
| BuffHeadSetupSelectTextureWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupSelectColorWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupLayoutManagerWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupAdvancedCompressionWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupAdvancedCompressionItemWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupAdvancedCompressionItemEffectWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupAdvancedContainersWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupAdvancedContainersItemWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupAdvancedContainersItemPropertiesWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupPriorityEffectsWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupPriorityEffectsItemWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupFilterWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupGeneralWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupLayoutWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupLayoutPropertiesWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupPerformanceWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupDisplayWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupContainerWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupTrackersWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupMenuWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupEffectCacheWindow | BuffHead | OnInitialize | exact | HIGH |
| BuffHeadSetupEffectCacheTableWindow | BuffHead | OnInitialize | exact | HIGH |
| ClosetGoblinCharacterWindow | CM_ClosetGoblin | OnInitialize | exact | HIGH |
| ClosetGoblinZoneWindow | CM_ClosetGoblin | OnInitialize | exact | HIGH |
| ClosetGoblinOptionWindow | CM_ClosetGoblin | OnInitialize | exact | HIGH |
| KillerWindow | Killer | OnInitialize | exact | HIGH |
| LibGroupSetupWindow | LibGroup | OnInitialize | exact | HIGH |
| TexturedButtonsSetupSelectColorWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupMenuWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupTexturesWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupAdvancedTexturesWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupCooldownWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupFontsWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupTintWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupMiscWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TexturedButtonsSetupActionbarWindow | TexturedButtons | OnInitialize | exact | HIGH |
| TurretRangeSetupMenuWindow | TurretRange | OnInitialize | exact | HIGH |
| TurretRangeSetupDistancesWindow | TurretRange | OnInitialize | exact | HIGH |
| TurretRangeSetupDistanceWindow | TurretRange | OnInitialize | exact | HIGH |
| TurretRangeSetupDisplayWindow | TurretRange | OnInitialize | exact | HIGH |
| TurretRangeSetupGeneralWindow | TurretRange | OnInitialize | exact | HIGH |
| TurretRangeDisplay | TurretRange | OnInitialize | exact | HIGH |
| WSCTContainer | WSCT | OnInitialize | exact | HIGH |
| WhoHealedMeWindow | WhoHealedMe | OnInitialize | exact | HIGH |
| WhoHealedMeOptions | WhoHealedMe | OnInitialize | exact | HIGH |
| WhoHealedMeDetails | WhoHealedMe | OnInitialize | exact | HIGH |
| MacroIconSelectionWindow | bigger_MacroWindow | OnInitialize | exact | HIGH |
| EA_Window_Macro | bigger_MacroWindow | OnInitialize | exact | HIGH |
## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Aura
- AutoMark
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibGroup
- LibWBToggler
- MiracleGrowLight
- MoraleCircle
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow

## Examples

- AdvancedPetAssist: APAFollowTargetHUD -> Window APAFollowTargetHUD
- AdvancedPetAssist: APAInstantOnlyHUD -> Window APAInstantOnlyHUD
- AdvancedPetAssist: APAKitingHUD -> Window APAKitingHUD
- AdvancedPetAssist: APAOptions -> Window APAOptions
- AdvancedPetAssist: APAOptionsBackground -> Window APAOptionsBackground
- AdvancedPetAssist: APAOptionsContent -> Window APAOptionsContent

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetTextColor](../../window_api/functions/window_ButtonGetTextColor.md) (HIGH 100/100) - Window Function
- [ButtonSetCheckButtonFlag](../../window_api/functions/window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ButtonSetTextColor](../../window_api/functions/window_ButtonSetTextColor.md) (HIGH 100/100) - Window Function
- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetDisabledFlag](../../window_api/functions/window_ComboBoxGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetDisabledFlag](../../window_api/functions/window_ComboBoxSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetRotation](../../window_api/functions/window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](../../window_api/functions/window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [EA_LabelCheckButton](../../globals/constants/constant_EA_LabelCheckButton.md) (HIGH 100/100) - Constant
- [EA_TitleBar_Default](../../globals/constants/constant_EA_TitleBar_Default.md) (HIGH 100/100) - Constant
- [EA_Window_ContextMenu.AddCascadingMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_Default](../../globals/constants/constant_EA_Window_Default.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultBackgroundFrame](../../globals/constants/constant_EA_Window_DefaultBackgroundFrame.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultButtonBottomFrame](../../globals/constants/constant_EA_Window_DefaultButtonBottomFrame.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultContextMenuFrame](../../globals/constants/constant_EA_Window_DefaultContextMenuFrame.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultFrame](../../globals/constants/constant_EA_Window_DefaultFrame.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultFrameStatusBar](../../globals/constants/constant_EA_Window_DefaultFrameStatusBar.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultSeparator](../../globals/constants/constant_EA_Window_DefaultSeparator.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultTooltipBackground](../../globals/constants/constant_EA_Window_DefaultTooltipBackground.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultVerticalSeparator](../../globals/constants/constant_EA_Window_DefaultVerticalSeparator.md) (HIGH 100/100) - Constant
- [EA_Window_TabSeparatorLeftSide](../../globals/constants/constant_EA_Window_TabSeparatorLeftSide.md) (HIGH 100/100) - Constant
- [EA_Window_TabSeparatorRightSide](../../globals/constants/constant_EA_Window_TabSeparatorRightSide.md) (HIGH 100/100) - Constant
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [LabelGetTextColor](../../window_api/functions/window_LabelGetTextColor.md) (HIGH 100/100) - Window Function
- [LabelGetTextDimensions](../../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelGetWordWrap](../../window_api/functions/window_LabelGetWordWrap.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](../../window_api/functions/window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](../../window_api/functions/window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [ScrollWindowSetOffset](../../window_api/functions/window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](../../window_api/functions/window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [SliderBar](element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [StatusBarGetCurrentValue](../../window_api/functions/window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetBackgroundTint](../../window_api/functions/window_StatusBarSetBackgroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetCurrentValue](../../window_api/functions/window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](../../window_api/functions/window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](../../window_api/functions/window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowForceProcessAnchors](../../window_api/functions/window_WindowForceProcessAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetFontAlpha](../../window_api/functions/window_WindowGetFontAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](../../window_api/functions/window_WindowGetLayer.md) (HIGH 100/100) - Window Function
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetPopable](../../window_api/functions/window_WindowGetPopable.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowGetTintColor](../../window_api/functions/window_WindowGetTintColor.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetFontAlpha](../../window_api/functions/window_WindowSetFontAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionTrigger](../../window_api/functions/window_WindowSetGameActionTrigger.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](../../window_api/functions/window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetId](../../window_api/functions/window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetMovable](../../window_api/functions/window_WindowSetMovable.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetPopable](../../window_api/functions/window_WindowSetPopable.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](../../window_api/functions/window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 90/100) - Window Function
- [ButtonSetTexture](../../window_api/functions/window_ButtonSetTexture.md) (HIGH 90/100) - Window Function
- [ButtonSetTextureSlice](../../window_api/functions/window_ButtonSetTextureSlice.md) (HIGH 90/100) - Window Function
- [EA_Window_CityRating](../../globals/constants/constant_EA_Window_CityRating.md) (HIGH 90/100) - Constant
- [EA_Window_ComboBoxMenuBackground](../../globals/constants/constant_EA_Window_ComboBoxMenuBackground.md) (HIGH 90/100) - Constant
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 90/100) - Window Function
- [WindowRestoreDefaultSettings](../../window_api/functions/window_WindowRestoreDefaultSettings.md) (HIGH 90/100) - Window Function
- [WindowSetMoving](../../window_api/functions/window_WindowSetMoving.md) (HIGH 90/100) - Window Function
- [WindowSetResizing](../../window_api/functions/window_WindowSetResizing.md) (HIGH 90/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 90/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 80/100) - Window Function
- [ComboBoxExternalOpenMenu](../../window_api/functions/window_ComboBoxExternalOpenMenu.md) (HIGH 80/100) - Window Function
- [StatusBarGetMaximumValue](../../window_api/functions/window_StatusBarGetMaximumValue.md) (HIGH 80/100) - Window Function
- [TextEditBoxSetMaxChars](../../window_api/functions/window_TextEditBoxSetMaxChars.md) (HIGH 80/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 80/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function
- [Anchor](element_Anchor.md) (MEDIUM 55/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type
- [Visual](element_Visual.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Event
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMButtonDown](../handlers/handler_OnMButtonDown.md) (HIGH 100/100) - XML Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Event
- [OnShutdown](../handlers/handler_OnShutdown.md) (HIGH 100/100) - XML Event
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 100/100) - XML Event
- [OnMouseOut](../handlers/handler_OnMouseOut.md) (HIGH 93/100) - XML Event
- [OnRawDeviceInput](../handlers/handler_OnRawDeviceInput.md) (HIGH 73/100) - XML Event
- [OnSizeUpdated](../handlers/handler_OnSizeUpdated.md) (HIGH 73/100) - XML Event

## Affects

- none
