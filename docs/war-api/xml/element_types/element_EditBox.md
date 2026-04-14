# EditBox

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
| Addons seen in | AdvancedRenownTrainer, Aura, AutoBand, BuffHead, Busted, CastSequence, CraftingWillard, Crusher |
| Files seen in | Code/Core/Common.xml, Code/Core/ConfigurationWindow.xml, Code/Core/Groups/EffectFilterDialog.xml, Code/Intercom/ChooseChannelDialog.xml, Code/UnitFrames/ClickCastingDialog.xml, Code/UnitFrames/EffectsIndicatorDialog.xml, Code/UnitFrames/UnitFramePartDialog.xml, Configuration/Config.xml |
| Namespaces detected | EditBox |
| Source kinds | xml_frames |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox, AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput, Aura: AuraConfigActivationAlertTextText |
| XML usage count | 416 |
| XML attribute usage count | 416 |
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

EditBox is an interactive XML control. It commonly appears under Window. It is typically used to organize structural children such as Anchors, Color, EventHandlers and bind XML events like OnEnterPressed, OnFocusLost, OnKeyEnter to Lua.

## Common Attributes

- font
- handleinput
- history
- inherits
- input
- layer
- maxChars
- maxchars
- name
- scrolling
- taborder
- warnOnTextCropped

## Common Handlers

- [OnEnterPressed](../handlers/handler_OnEnterPressed.md)
- [OnFocusLost](../handlers/handler_OnFocusLost.md)
- [OnKeyEnter](../handlers/handler_OnKeyEnter.md)
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md)
- OnKeyTab
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- OnRawDeviceInput
- [OnShown](../handlers/handler_OnShown.md)
- [OnTextChanged](../handlers/handler_OnTextChanged.md)
- [OnUpdate](../handlers/handler_OnUpdate.md)

## Common Handler Functions

- Enemy.ConfigurationWindow_OnChange
- zMailModSend.FixRecipient
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Killer.OnSettingsEditChanged
- MiracleGrow2.LayoutBarCChanged
- MiracleGrow2.LayoutProgDimChanged
- BuffHead.Setup.SelectColor.OnTintChanged
- MapPin.TimeChanged
- MiracleGrow2.ConfigThrobCChanged
- Obsidian.Setup.SelectColor.OnTintChanged
- ShiniesAutoUI.OnPriceChange
- ShiniesPostUI.OnPriceChange
- TexturedButtons.Setup.SelectColor.OnTintChanged
- TexturedButtons.Setup.Tint.OnTintChanged
- TurretRange.Setup.Display.OnTintChanged
- TurretRange.Setup.Distance.OnTintChanged
- Vectors.Settings.FloatEditBoxChanged
- MassMailWindow.FixRecipient
- MiracleGrow2.LayoutDimChanged
- NBSBParam.OnParamChange
- wbLeadHelperConfigTab.OnChanged
- zMailModSend.OnTextChanged
- AuraConfig.OnTextureOffsetXChanged
- AuraConfig.OnTextureOffsetYChanged
- AuraConfig.OnTimerOffsetXChanged
- AuraConfig.OnTimerOffsetYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeColumnsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeRowsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetYChanged
- BuffHead.Setup.Display.OnOffsetXChanged
- BuffHead.Setup.Display.OnOffsetYChanged
- BuffHead.Setup.EffectCache.OnSearchChanged
- BuffHead.Setup.Layout.Properties.OnCoreSizeSizeHeightChanged
- BuffHead.Setup.Layout.Properties.OnCoreSizeSizeWidthChanged
- BuffHead.Setup.Layout.Properties.OnOffsetXChanged
- BuffHead.Setup.Layout.Properties.OnOffsetYChanged
- BuffHead.Setup.Layout.Properties.OnSizeScaleChanged
- BuffHead.Setup.Layout.Properties.OnSizeSizeHeightChanged
- BuffHead.Setup.Layout.Properties.OnSizeSizeWidthChanged
- BuffHead.Setup.PriorityEffectsItem.OnAbilityIdChanged
- CastSequence.FindAbility.OnAbilityIdChanged
- CastSequence.SequenceBuilder.OnPageChanged
- CastSequence.SequenceBuilder.OnSlotChanged
- CastSequence.Setup.OnCastModifierChanged
- DebugWindow.AutoSender
- DebugWindow.PreventType
- DevPadWindow.OnCodeChanged
- GDes.ExtrasFeatureOffsetXKeyEntry
- GDes.ExtrasFeatureOffsetYKeyEntry
- GDes.ExtrasIconsOffsetXKeyEntry
- GDes.ExtrasIconsOffsetYKeyEntry
- Ges.ExtrasFeatureOffsetXKeyEntry
- Ges.ExtrasFeatureOffsetYKeyEntry
- Ges.ExtrasTrackerOffsetXKeyEntry
- Ges.ExtrasTrackerOffsetYKeyEntry
- GroupRangeSetup.General.OnOffsetXChanged
- GroupRangeSetup.General.OnOffsetYChanged
- GroupRangeSetup.Style.Pointer.OnOffsetXChanged
- GroupRangeSetup.Style.Pointer.OnOffsetYChanged
- GroupRangeSetup.Style.PointerReverse.OnOffsetXChanged
- GroupRangeSetup.Style.PointerReverse.OnOffsetYChanged
- GroupRangeSetup.Style.SimpleText.OnOffsetXChanged
- GroupRangeSetup.Style.SimpleText.OnOffsetYChanged
- LibAddonButton.Manager.Advanced.OnAnimationFpsChanged
- LibAddonButton.Manager.Advanced.OnHeightChanged
- LibAddonButton.Manager.Advanced.OnSimpleTextureChanged
- LibAddonButton.Manager.Advanced.OnWidthChanged
- MapMonster.Editor.OnLabelChange
- MapMonster.Editor.OnNoteChange
- MapMonster.Editor.OnPosTextChanged
- MapMonster.PinTypeEditor.OnPinTypeTextChange
- MapMonster.PinTypeEditor.OnRadiusTextChange
- MapMonster.PinTypeEditor.OnSubTypeTextChange
- MassMailWindow.OnTextChanged
- Megaphone.SaveSettings
- MiracleGrow2.ConfigSoundChanged
- Obsidian.Setup.Castbar.OnBackgroundBorderSizeChanged
- Obsidian.Setup.Castbar.OnGeneralSizeHeightChanged
- Obsidian.Setup.Castbar.OnGeneralSizeWidthChanged
- Obsidian.Setup.Castbar.OnGlobalCooldownOffsetXChanged
- Obsidian.Setup.Castbar.OnGlobalCooldownOffsetYChanged
- Obsidian.Setup.Castbar.OnGlobalCooldownSizeHeightChanged
- Obsidian.Setup.Castbar.OnGlobalCooldownSizeWidthChanged
- Obsidian.Setup.Castbar.OnIconOffsetXChanged
- Obsidian.Setup.Castbar.OnIconOffsetYChanged
- Obsidian.Setup.Castbar.OnNameOffsetXChanged
- Obsidian.Setup.Castbar.OnNameOffsetYChanged
- Obsidian.Setup.Castbar.OnTimerOffsetXChanged
- Obsidian.Setup.Castbar.OnTimerOffsetYChanged
- Obsidian.Setup.EffectTracker.OnGeneralFixedDurationChanged
- Obsidian.Setup.EffectTracker.OnGeneralMaximumDurationChanged
- Obsidian.Setup.EffectTracker.OnGeneralSizeHeightChanged
- Obsidian.Setup.EffectTracker.OnGeneralSizeWidthChanged
- Obsidian.Setup.EffectTracker.OnGeneralSpacingChanged
- Obsidian.Setup.EffectTracker.OnNameOffsetXChanged
- Obsidian.Setup.EffectTracker.OnNameOffsetYChanged
- Obsidian.Setup.EffectTracker.OnTimerOffsetXChanged
- Obsidian.Setup.EffectTracker.OnTimerOffsetYChanged
- Obsidian.Setup.EffectTracker.OnTrackerIconOffsetXChanged
- Obsidian.Setup.EffectTracker.OnTrackerIconOffsetYChanged
- Obsidian.Setup.EffectTracker.OnTrackerOffsetXChanged
- Obsidian.Setup.EffectTracker.OnTrackerOffsetYChanged
- PP.UpdateDyeFilter
- PartyAdWindow.OnPurposeTextChanged
- RVAPI_ColorDialog.OnTextChangedEdit
- ShiniesConfigGeneral.OnTextChanged_UIScale
- ShiniesPostUI.OnStackChange
- ShiniesPostUI.OnStackSizeChange
- TexturedButtons.Setup.Actionbar.OnPaddingXChanged
- TexturedButtons.Setup.Actionbar.OnPaddingYChanged
- TexturedButtons.Setup.Actionbar.OnSpacingXChanged
- TexturedButtons.Setup.Actionbar.OnSpacingYChanged
- TurretRange.Setup.Display.OnDistanceOffsetXChanged
- TurretRange.Setup.Display.OnDistanceOffsetYChanged
- TurretRange.Setup.Display.OnGraphicLimitChanged
- Twister.OnButtonTextChanged
- XpStatus.OnQuotaTextChanged
- zMailModOptions.OnTextChanged
- nLootLinkGUI.search
- WarBoard_FPSOptions.OnEnterKeyPressed
- snt_panel.text_input
- EA_ChatWindow.OnRename
- MapPin.SendCommand
- MapPin.SendText
- ObjectInspector.InspectObject
- zMailModSend.FixComplete
- AuctionWindowSearchControls.OnLButtonUpSearch
- DebugWindow.TextSend
- DevPadWindow.ConfirmRename
- DevPadWindow.CreateNewFile
- DevPadWindow.SaveFile
- DuffTimer.Options.OnEditBoxChanged
- MapMonster.Editor.OnPosEnterKey
- MassMailWindow.FixComplete
- MassRefine.OnKeyEnter
- RandomMountUI.OnMinLevelChanged
- SocialWindowTabFriends.AddFriend
- SocialWindowTabFriends.OnDescriptionAccept
- SocialWindowTabIgnore.AddIgnore
- Twister.OnButtonTextEnter
- XpStatus.OnKeyEnter
- EA_ChatWindow.OnCancelRename
- DebugWindow.TextClear
- DevPadWindow.OnKeyEscape
- MassRefine.OnKeyEscape
- SocialWindowTabFriends.OnCancelDescription
- XpStatus.OnKeyEscape
- Enemy.ConfigurationWindow_ShowTooltip
- MapMonster.PinTypeEditor.MouseOverDescription
- EZCraftX.OnMouseOver_Text3Chars
- Enemy.GroupsUI_EffectFilterDialog_OnAbilityIdsMouseOver
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig2CommandMouseOver
- LPET.OnMouseOver
- WbLeadHelperMessage.OnMouseOverLabelEditBox
- WbLeadHelperMessage.OnMouseOverMessageEditBox
- wbLeadHelperConfigTab.OnMouseOverMessageEndLabel
- wbLeadHelperConfigTab.OnMouseOverMessageStartLabel
- DebugWindow.OnKeyTab
- zMailModSend.OnKeyTab
- DevPadWindow.OnKeyTab
- MassMailWindow.OnKeyTab
- FastFriends.UI.OnMinLevelCommit
- PartyAdWindow.OnPurposeFocusLost
- zMailModSend.OnRawDeviceInput
- MassMailWindow.OnRawDeviceInput
- DebugWindow.OnShowFocus
- DevPadWindow.OnShown
- EZCraftX.OnUpdate


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnTextChanged](../handlers/handler_OnTextChanged.md) | data | Enemy.ConfigurationWindow_OnChange, zMailModSend.FixRecipient, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, Killer.OnSettingsEditChanged, MiracleGrow2.LayoutBarCChanged, MiracleGrow2.LayoutProgDimChanged, BuffHead.Setup.SelectColor.OnTintChanged, MapPin.TimeChanged, MiracleGrow2.ConfigThrobCChanged, Obsidian.Setup.SelectColor.OnTintChanged, ShiniesAutoUI.OnPriceChange, ShiniesPostUI.OnPriceChange, TexturedButtons.Setup.SelectColor.OnTintChanged, TexturedButtons.Setup.Tint.OnTintChanged, TurretRange.Setup.Display.OnTintChanged, TurretRange.Setup.Distance.OnTintChanged, Vectors.Settings.FloatEditBoxChanged, MassMailWindow.FixRecipient, MiracleGrow2.LayoutDimChanged, NBSBParam.OnParamChange, wbLeadHelperConfigTab.OnChanged, zMailModSend.OnTextChanged, , AuraConfig.OnTextureOffsetXChanged, AuraConfig.OnTextureOffsetYChanged, AuraConfig.OnTimerOffsetXChanged, AuraConfig.OnTimerOffsetYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeColumnsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeRowsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetYChanged, BuffHead.Setup.Display.OnOffsetXChanged, BuffHead.Setup.Display.OnOffsetYChanged, BuffHead.Setup.EffectCache.OnSearchChanged, BuffHead.Setup.Layout.Properties.OnCoreSizeSizeHeightChanged, BuffHead.Setup.Layout.Properties.OnCoreSizeSizeWidthChanged, BuffHead.Setup.Layout.Properties.OnOffsetXChanged, BuffHead.Setup.Layout.Properties.OnOffsetYChanged, BuffHead.Setup.Layout.Properties.OnSizeScaleChanged, BuffHead.Setup.Layout.Properties.OnSizeSizeHeightChanged, BuffHead.Setup.Layout.Properties.OnSizeSizeWidthChanged, BuffHead.Setup.PriorityEffectsItem.OnAbilityIdChanged, CastSequence.FindAbility.OnAbilityIdChanged, CastSequence.SequenceBuilder.OnPageChanged, CastSequence.SequenceBuilder.OnSlotChanged, CastSequence.Setup.OnCastModifierChanged, DebugWindow.AutoSender, DebugWindow.PreventType, DevPadWindow.OnCodeChanged, GDes.ExtrasFeatureOffsetXKeyEntry, GDes.ExtrasFeatureOffsetYKeyEntry, GDes.ExtrasIconsOffsetXKeyEntry, GDes.ExtrasIconsOffsetYKeyEntry, Ges.ExtrasFeatureOffsetXKeyEntry, Ges.ExtrasFeatureOffsetYKeyEntry, Ges.ExtrasTrackerOffsetXKeyEntry, Ges.ExtrasTrackerOffsetYKeyEntry, GroupRangeSetup.General.OnOffsetXChanged, GroupRangeSetup.General.OnOffsetYChanged, GroupRangeSetup.Style.Pointer.OnOffsetXChanged, GroupRangeSetup.Style.Pointer.OnOffsetYChanged, GroupRangeSetup.Style.PointerReverse.OnOffsetXChanged, GroupRangeSetup.Style.PointerReverse.OnOffsetYChanged, GroupRangeSetup.Style.SimpleText.OnOffsetXChanged, GroupRangeSetup.Style.SimpleText.OnOffsetYChanged, LibAddonButton.Manager.Advanced.OnAnimationFpsChanged, LibAddonButton.Manager.Advanced.OnHeightChanged, LibAddonButton.Manager.Advanced.OnSimpleTextureChanged, LibAddonButton.Manager.Advanced.OnWidthChanged, MapMonster.Editor.OnLabelChange, MapMonster.Editor.OnNoteChange, MapMonster.Editor.OnPosTextChanged, MapMonster.PinTypeEditor.OnPinTypeTextChange, MapMonster.PinTypeEditor.OnRadiusTextChange, MapMonster.PinTypeEditor.OnSubTypeTextChange, MassMailWindow.OnTextChanged, Megaphone.SaveSettings, MiracleGrow2.ConfigSoundChanged, Obsidian.Setup.Castbar.OnBackgroundBorderSizeChanged, Obsidian.Setup.Castbar.OnGeneralSizeHeightChanged, Obsidian.Setup.Castbar.OnGeneralSizeWidthChanged, Obsidian.Setup.Castbar.OnGlobalCooldownOffsetXChanged, Obsidian.Setup.Castbar.OnGlobalCooldownOffsetYChanged, Obsidian.Setup.Castbar.OnGlobalCooldownSizeHeightChanged, Obsidian.Setup.Castbar.OnGlobalCooldownSizeWidthChanged, Obsidian.Setup.Castbar.OnIconOffsetXChanged, Obsidian.Setup.Castbar.OnIconOffsetYChanged, Obsidian.Setup.Castbar.OnNameOffsetXChanged, Obsidian.Setup.Castbar.OnNameOffsetYChanged, Obsidian.Setup.Castbar.OnTimerOffsetXChanged, Obsidian.Setup.Castbar.OnTimerOffsetYChanged, Obsidian.Setup.EffectTracker.OnGeneralFixedDurationChanged, Obsidian.Setup.EffectTracker.OnGeneralMaximumDurationChanged, Obsidian.Setup.EffectTracker.OnGeneralSizeHeightChanged, Obsidian.Setup.EffectTracker.OnGeneralSizeWidthChanged, Obsidian.Setup.EffectTracker.OnGeneralSpacingChanged, Obsidian.Setup.EffectTracker.OnNameOffsetXChanged, Obsidian.Setup.EffectTracker.OnNameOffsetYChanged, Obsidian.Setup.EffectTracker.OnTimerOffsetXChanged, Obsidian.Setup.EffectTracker.OnTimerOffsetYChanged, Obsidian.Setup.EffectTracker.OnTrackerIconOffsetXChanged, Obsidian.Setup.EffectTracker.OnTrackerIconOffsetYChanged, Obsidian.Setup.EffectTracker.OnTrackerOffsetXChanged, Obsidian.Setup.EffectTracker.OnTrackerOffsetYChanged, PP.UpdateDyeFilter, PartyAdWindow.OnPurposeTextChanged, RVAPI_ColorDialog.OnTextChangedEdit, ShiniesConfigGeneral.OnTextChanged_UIScale, ShiniesPostUI.OnStackChange, ShiniesPostUI.OnStackSizeChange, TexturedButtons.Setup.Actionbar.OnPaddingXChanged, TexturedButtons.Setup.Actionbar.OnPaddingYChanged, TexturedButtons.Setup.Actionbar.OnSpacingXChanged, TexturedButtons.Setup.Actionbar.OnSpacingYChanged, TurretRange.Setup.Display.OnDistanceOffsetXChanged, TurretRange.Setup.Display.OnDistanceOffsetYChanged, TurretRange.Setup.Display.OnGraphicLimitChanged, Twister.OnButtonTextChanged, XpStatus.OnQuotaTextChanged, zMailModOptions.OnTextChanged | `text` | MEDIUM |
| [OnKeyEnter](../handlers/handler_OnKeyEnter.md) | custom | , nLootLinkGUI.search, WarBoard_FPSOptions.OnEnterKeyPressed, snt_panel.text_input, EA_ChatWindow.OnRename, MapPin.SendCommand, MapPin.SendText, ObjectInspector.InspectObject, zMailModSend.FixComplete, AuctionWindowSearchControls.OnLButtonUpSearch, DebugWindow.TextSend, DevPadWindow.ConfirmRename, DevPadWindow.CreateNewFile, DevPadWindow.SaveFile, DuffTimer.Options.OnEditBoxChanged, Ges.ExtrasFeatureOffsetXKeyEntry, Ges.ExtrasFeatureOffsetYKeyEntry, Ges.ExtrasTrackerOffsetXKeyEntry, Ges.ExtrasTrackerOffsetYKeyEntry, MapMonster.Editor.OnPosEnterKey, MassMailWindow.FixComplete, MassRefine.OnKeyEnter, RandomMountUI.OnMinLevelChanged, SocialWindowTabFriends.AddFriend, SocialWindowTabFriends.OnDescriptionAccept, SocialWindowTabIgnore.AddIgnore, Twister.OnButtonTextEnter, XpStatus.OnKeyEnter | `Command` | LOW |
| [OnKeyEscape](../handlers/handler_OnKeyEscape.md) | custom | , EA_ChatWindow.OnCancelRename, zMailModSend.FixRecipient, DebugWindow.TextClear, DevPadWindow.OnKeyEscape, DuffTimer.Options.OnEditBoxChanged, Ges.ExtrasFeatureOffsetXKeyEntry, Ges.ExtrasFeatureOffsetYKeyEntry, Ges.ExtrasTrackerOffsetXKeyEntry, Ges.ExtrasTrackerOffsetYKeyEntry, MassMailWindow.FixRecipient, MassRefine.OnKeyEscape, SocialWindowTabFriends.OnCancelDescription, XpStatus.OnKeyEscape | `` |  |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ConfigurationWindow_ShowTooltip, MapMonster.PinTypeEditor.MouseOverDescription, EZCraftX.OnMouseOver_Text3Chars, Enemy.GroupsUI_EffectFilterDialog_OnAbilityIdsMouseOver, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig2CommandMouseOver, LPET.OnMouseOver, WbLeadHelperMessage.OnMouseOverLabelEditBox, WbLeadHelperMessage.OnMouseOverMessageEditBox, wbLeadHelperConfigTab.OnMouseOverMessageEndLabel, wbLeadHelperConfigTab.OnMouseOverMessageStartLabel | `` |  |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | zMailModSend.FixRecipient, MassMailWindow.FixRecipient | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | zMailModSend.FixRecipient, MassMailWindow.FixRecipient | `flags, x, y` | MEDIUM |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | zMailModSend.FixRecipient, MassMailWindow.FixRecipient | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | zMailModSend.FixRecipient, MassMailWindow.FixRecipient | `flags, x, y` | MEDIUM |
| OnKeyTab | custom | DebugWindow.OnKeyTab, zMailModSend.OnKeyTab, DevPadWindow.OnKeyTab, Ges.ExtrasFeatureOffsetXKeyEntry, Ges.ExtrasFeatureOffsetYKeyEntry, Ges.ExtrasTrackerOffsetXKeyEntry, Ges.ExtrasTrackerOffsetYKeyEntry, MassMailWindow.OnKeyTab | `` |  |
| [OnFocusLost](../handlers/handler_OnFocusLost.md) | custom | FastFriends.UI.OnMinLevelCommit, Megaphone.SaveSettings, PartyAdWindow.OnPurposeFocusLost | `` |  |
| OnRawDeviceInput | custom | zMailModSend.OnRawDeviceInput, MassMailWindow.OnRawDeviceInput | `deviceId, itemId, itemDown` | LOW |
| [OnEnterPressed](../handlers/handler_OnEnterPressed.md) | input | FastFriends.UI.OnMinLevelCommit, Megaphone.SaveSettings | `` |  |
| [OnShown](../handlers/handler_OnShown.md) | lifecycle | DebugWindow.OnShowFocus, DevPadWindow.OnShown | `` |  |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | EZCraftX.OnUpdate | `elapsed` | MEDIUM |

### Per-Event Lua API Calls

**OnTextChanged** handlers call: `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ComboBoxGetSelectedMenuItem`, `LabelSetText`, `LabelSetTextColor`, `RegisterEventHandler`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `SystemData.ActiveWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxGetText:gsub`, `TextEditBoxSetText`, `TextEditBoxSetTextColor`, `WindowGetId`, `WindowGetParent`, `WindowSetDimensions`, `WindowSetGameActionTrigger`, `WindowSetShowing`, `WindowSetTintColor`

**OnKeyEnter** handlers call: `ComboBoxAddMenuItem`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `LabelGetText`, `ListBoxSetDisplayOrder`, `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowAssignFocus`, `WindowSetShowing`, `WindowStartAlphaAnimation`, `WindowStartScaleAnimation`

**OnKeyEscape** handlers call: `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowAssignFocus`, `WindowGetShowing`, `WindowSetShowing`

**OnMouseOver** handlers call: `WindowGetId`, `WindowGetParent`

**OnLButtonDown** handlers call: `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowSetShowing`

**OnLButtonUp** handlers call: `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowSetShowing`

**OnRButtonDown** handlers call: `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowSetShowing`

**OnRButtonUp** handlers call: `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowSetShowing`

**OnKeyTab** handlers call: `TextEditBoxInsertText`, `TextEditBoxSelectAll`, `WindowAssignFocus`

**OnFocusLost** handlers call: `ButtonGetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `TextEditBoxGetText`, `TextEditBoxSetText`

**OnRawDeviceInput** handlers call: `TextEditBoxSetText`

**OnEnterPressed** handlers call: `ButtonGetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `TextEditBoxGetText`, `TextEditBoxSetText`

**OnShown** handlers call: `WindowAssignFocus`, `WindowGetShowing`

## Common Inherits

- Aura_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline
- EA_EditBox_NoFrame
- IraConfigNumEdit
- IraConfigNumSpin
- MapMonsterEditorWindowEditBoxCoord
- MapMonsterPinTypeEditorWindowEditBoxDefault
- RVAPI_ColorDialogEditBoxTemplate
- Shinies_BrassCoin_EditBox_DefaultFrame
- Shinies_GoldCoin_EditBox_DefaultFrame
- Shinies_SilverCoin_EditBox_DefaultFrame

## Common Parent Elements

- [Windows](element_Windows.md) — 416× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 380× (HIGH)
- [Size](element_Size.md) — 355× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 269× (HIGH)
- [TextOffset](element_TextOffset.md) — 63× (HIGH)
- [Windows](element_Windows.md) — 3× (MEDIUM)
- [Color](element_Color.md) — 2× (LOW)

## Common Template Bases

- Aura_EditBox_DefaultFrame
- CraftingWillardEditbox
- EA_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline
- EA_EditBox_NoFrame
- IraConfigNumEdit
- IraConfigNumSpin
- MapMonsterEditorWindowEditBoxCoord
- MapMonsterEditorWindowEditBoxDefault
- MapMonsterPinTypeEditorWindowEditBoxDefault
- RVAPI_ColorDialogEditBoxTemplate
- Shinies_BrassCoin_EditBox_DefaultFrame
- Shinies_GoldCoin_EditBox_DefaultFrame
- Shinies_SilverCoin_EditBox_DefaultFrame
- zMailModOptionsEditbox


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 95% | EA_EditBox_DefaultFrame, Aura_EditBox_DefaultFrame, EA_EditBox_DefaultFrame_Multiline, CraftingWillardEditbox, ... |
| `maxchars` | optional | 68% | 76, 20, 5, 3, ... |
| `input` | optional | 29% | nospaces, numbers, textandnumbers |
| `layer` | optional | 25% | secondary, default, popup, 2, ... |
| `taborder` | optional | 24% | 1, 2, 3, 4, ... |
| `scrolling` | optional | 14% | vert, none, horz, horiz |
| `history` | optional | 12% | 15, 0, 5, 30, ... |
| `font` | optional | 11% | font_chat_text, font_clear_small, font_clear_default, font_clear_small_bold, ... |
| `maxChars` | optional | 11% | 3, 300, 4, 100, ... |
| `warnOnTextCropped` | optional | 8% | false |
| `handleinput` | optional | 4% | true, false |
| `background` | optional | 4% | EA_FullResizeImage_TanBorder, EA_Window_DefaultFrame, DefaultWindowInnerBackground |
| `id` | optional | 2% | 1, 3, 5, 2, ... |
| `textalign` | optional | 2% | leftcenter, right, center, topright |
| `scrollbar` | optional | 2% | EA_ScrollBar_DefaultVerticalChain, CopyScrollBar, $parentDevPadCodeScrollBar, $parentObjectScrollbar |
| `align` | optional | 1% | center, rightcenter |
| `alpha` | optional | 1% | 0.97 |
| `autoHideScrollBar` | optional | 0% | true |
| `linecount` | optional | 0% | 2, 9 |
| `maxchar` | optional | 0% | 79000 |
| `numeric` | optional | 0% | true |
| `wordwrap` | optional | 0% | false |
| `sticky` | optional | 0% | false |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 380 times as an unnamed child.

### [Size](element_Size.md)

Observed 355 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 269 times as an unnamed child.

### [TextOffset](element_TextOffset.md)

Observed 63 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5, 6, 4, 0 |
| `y` | **required** | 5, 12, 2, 3 |
### [Windows](element_Windows.md)

Observed 3 times as an unnamed child.

### [Color](element_Color.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 255, 0, 73, 225 |
| `g` | **required** | 25, 253, 218, 255 |
| `r` | **required** | 255, 253, 155, 245 |
| `a` | optional | 255, 0.8, 0, 128 |
## Recursive Hierarchy

- Root: [EditBox](element_EditBox.md)
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
| `TextEditBoxGetText` | ui | 498 | OnEnterPressed, OnFocusLost, OnKeyEnter, OnKeyEscape, OnLButtonDown, OnLButtonUp, OnRButtonDown, OnRButtonUp, OnTextChanged |
| `TextEditBoxSetText` | ui | 97 | OnEnterPressed, OnFocusLost, OnKeyEnter, OnKeyEscape, OnLButtonDown, OnLButtonUp, OnRButtonDown, OnRButtonUp, OnRawDeviceInput, OnTextChanged |
| `WindowSetShowing` | ui | 90 | OnKeyEnter, OnKeyEscape, OnLButtonDown, OnLButtonUp, OnRButtonDown, OnRButtonUp, OnTextChanged |
| `ComboBoxGetSelectedMenuItem` | ui | 59 | OnEnterPressed, OnFocusLost, OnKeyEnter, OnTextChanged |
| `SliderBarSetCurrentPosition` | ui | 55 | OnTextChanged |
| `SliderBarGetCurrentPosition` | ui | 28 | OnTextChanged |
| `ButtonGetPressedFlag` | ui | 27 | OnEnterPressed, OnFocusLost, OnTextChanged |
| `WindowGetId` | ui | 16 | OnMouseOver, OnTextChanged |
| `WindowAssignFocus` | ui | 9 | OnKeyEnter, OnKeyEscape, OnKeyTab, OnShown |
| `ListBoxSetDisplayOrder` | ui | 7 | OnKeyEnter |
| `WindowGetParent` | ui | 7 | OnMouseOver, OnTextChanged |
| `WindowSetTintColor` | ui | 6 | OnTextChanged |
| `LabelSetText` | ui | 5 | OnTextChanged |
| `TextEditBoxSetTextColor` | ui | 4 | OnTextChanged |
| `LabelSetTextColor` | ui | 3 | OnTextChanged |
| `TextEditBoxGetText:gsub` | ui | 3 | OnTextChanged |
| `TextEditBoxInsertText` | ui | 3 | OnKeyTab |
| `TextEditBoxSelectAll` | ui | 3 | OnKeyTab |
| `WindowSetDimensions` | ui | 3 | OnTextChanged |
| `ButtonSetDisabledFlag` | ui | 2 | OnTextChanged |
| `ComboBoxAddMenuItem` | ui | 2 | OnKeyEnter |
| `ComboBoxSetSelectedMenuItem` | ui | 2 | OnKeyEnter |
| `SystemData.ActiveWindow.name:match` | data | 2 | OnTextChanged |
| `WindowGetShowing` | ui | 2 | OnKeyEscape, OnShown |
| `WindowStartAlphaAnimation` | ui | 2 | OnKeyEnter |
| `WindowStartScaleAnimation` | ui | 2 | OnKeyEnter |
| `ComboBoxGetSelectedText` | ui | 1 | OnKeyEnter |
| `LabelGetText` | ui | 1 | OnKeyEnter |
| `RegisterEventHandler` | event | 1 | OnTextChanged |
| `WindowSetGameActionTrigger` | ui | 1 | OnTextChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnEnterPressed

Confidence: HIGH

### OnFocusLost

Confidence: LOW

### OnKeyEnter

Confidence: LOW

### OnKeyEscape

Confidence: LOW

### OnKeyTab

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
### OnMouseOver

Confidence: HIGH

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

### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnUpdate

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
## Lua Functions Manipulating This Type

- MapPin.local.EditMarker
- SocialWindow.FriendAdd
- mms.SetNewLayout
- SocialWindow.IgnoreAdd
- EA_Window_Macro.Initialize
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- SocialWindow.OnNewMember
- Sequencer.ToggleReset
- Sequencer.Load
- Sequencer.Save
- EA_Window_Macro.OnSave
- BustedGUI.UpdateErrorView
- UiModWindow.UpdateAdvancedSettings
- PP.UpdateDyeFilter
- LPET.RenameProfileOnButtonUp
- Squared.OpenImportExport
- Squared.ExportSettings
- mmw.AutoComplete
- GuildWarden.GetCombosFilters
- MapPin.SetupAccept
- EA_Window_Macro.UpdateDetails
- mmw.FormatRecipient
- Sequencer.SetReturn
- Twister.CloseSettingsWindow
- Twister.OnLoad
- GuildWardenWin.WinSetup
- mmw.LoadRecipient
- mms.SetDefaultLayout
- SocialWindowTabFriends.OnCancelDescription
- XpStatus.InitializeQuotaWindow
- WBStutterLess.OnSave
- LPET.SaveProfileOnButtonUp
- XpStatus.ShowQuotaWindow
- WbLeadHelperMessage.MessageDialogOpen
- SocialWindowTabFriends.OnDescriptionAccept
- mmw.OnRawDeviceInput
- Enemy.IntercomUI_ChooseChannelDialog_OnOkButton
- LPET.AddProfileOnButtonUp
- Squared.ImportSettings
- Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- RandomMountUI.Refresh
- mmo.ValidateSettings
- Twister.OnButtonTextChanged
- Enemy.GroupsUI_EffectFilterDialog_Ok
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- RandomMountUI.OnMinLevelChanged
- Sequencer.local.SetReturn
- WBStutterLess.Initialize
- LibSlash.Initialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- MapPin.OnUpdate
- MapPin.EditMarker
- MapPin.RButtonUp
- XpStatus.OnQuotaTextChanged
- WbLeadHelperMessage.OnOk
- mmw.FixComplete
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- GuildWarden.filter
- XpStatus.OnSetNewQuota
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- MapPin.SendText


## Binding Resolution

- Total handler declarations: 412
- Resolved to Lua functions: 412 (100%)

## Seen In

- AdvancedRenownTrainer
- Aura
- AutoBand
- BuffHead
- Busted
- CastSequence
- CraftingWillard
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- Enemy
- FastFriends
- FozAuction
- GDes
- Ges
- GroupRange
- GuildWarden
- HealGrid
- ItemRack
- KeyBar
- Keyset
- Killer
- LibAddonButton
- LoyalPet
- MapMonster
- MapPin
- Mass Refine
- MegaphonePlusPlus
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- ObjectInspector
- Obsidian
- PartyAd
- PieTracker
- Pocket Palette
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RememberIt
- Res
- SNT_PANEL
- SOR
- Sequencer
- Shinies
- SocialWindow 2.0
- Squared
- TacticSetNames
- TastyButtons
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TurretRange
- Twister
- Vectors
- WARCommander
- WarBoard_FPS
- WarBoard_Loc
- Warbuilder
- XpStatus+G
- bigger_MacroWindow
- nLootLink
- wbLeadHelper
- zMailMod

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox -> EditBox AdvancedRenownTrainingImportNameInputWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox -> EditBox AdvancedRenownTrainingImportWindowLinkInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox -> EditBox AdvancedRenownTrainingImportWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox -> EditBox AdvancedRenownTrainingLinkWindowLinkBox
- AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput -> EditBox AdvancedRenownTrainingPresetsWindowSaveNameInput
- Aura: AuraConfigActivationAlertTextText -> EditBox AuraConfigActivationAlertTextText

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [Color](element_Color.md) (HIGH 100/100) - XML Element Type
- [EA_EditBox_DefaultFrame](../../globals/constants/constant_EA_EditBox_DefaultFrame.md) (HIGH 100/100) - Constant
- [EA_EditBox_DefaultFrame_Multiline](../../globals/constants/constant_EA_EditBox_DefaultFrame_Multiline.md) (HIGH 100/100) - Constant
- [EA_EditBox_NoFrame](../../globals/constants/constant_EA_EditBox_NoFrame.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TextOffset](element_TextOffset.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnKeyEnter](../handlers/handler_OnKeyEnter.md) (HIGH 88/100) - XML Event
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md) (HIGH 88/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnTextChanged](../handlers/handler_OnTextChanged.md) (HIGH 88/100) - XML Event
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [OnEnterPressed](../handlers/handler_OnEnterPressed.md) (HIGH 76/100) - XML Event
- [OnFocusLost](../handlers/handler_OnFocusLost.md) (HIGH 76/100) - XML Event
