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
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:101`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:111`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:121`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1374`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1415`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1456`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1496`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:232` |
| Namespaces detected | Window |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAInstantOnlyHUD, AdvancedPetAssist: APAKitingHUD, AdvancedPetAssist: APAOptions, AdvancedPetAssist: APAOptionsBackground, AdvancedPetAssist: APAOptionsContent |
| XML usage count | 1331 |
| XML attribute usage count | 1331 |
| Lua usage count | 19 |
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

Observed XML element type instantiated by 60 addons.

## Common Attributes

- name
- inherits
- layer
- handleinput
- movable
- savesettings
- popable
- sticky
- scale
- id
- skipinput
- alpha

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnHidden](../handlers/handler_OnHidden.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnShown](../handlers/handler_OnShown.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnInitialize](../handlers/handler_OnInitialize.md)
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnShutdown](../handlers/handler_OnShutdown.md)
- [OnMouseOut](../handlers/handler_OnMouseOut.md)

## Common Handler Functions

- WindowUtils.OnShown
- WindowUtils.OnHidden
- BuffHead.Setup.Layout.OnLayersChanged
- WindowUtils.TrapClick
- BuffHead.Setup.Layout.Properties.OnColorExampleMouseOut
- BuffHead.Setup.Layout.Properties.OnColorExampleMouseOver
- BuffHead.Setup.PriorityEffectsItem.OnTargetTypeLUp
- Killer.HideRowTooltip
- MapPin.RButtonUp
- RoR_SoR.BroadCastOption
- RoR_SoR.OnMouseOverStart
- EA_LabelCheckButton.Initialize


## XML Event Bindings

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnHidden](../handlers/handler_OnHidden.md) | WindowUtils.OnHidden, APAGui.OnFollowTargetHUDHidden, APAGui.OnHidden, APAGui.OnInstantOnlyHUDHidden, APAGui.OnKitingHUDHidden, AdvancedRenownTraining.OnExportHidden | `function()` | MEDIUM |
| [OnInitialize](../handlers/handler_OnInitialize.md) | EA_LabelCheckButton.Initialize, AdvancedRenownTraining.Initialize, AuraConfig.OnInitialize, CMapWindow.Initialize, CMapWindow.RefreshMapPointFilterMenu, ClosetGoblinOptionWindow.OnInitialize | `function()` | MEDIUM |
| [OnKeyEscape](../handlers/handler_OnKeyEscape.md) | DebugWindow.OnKeyEscape, Enemy.CombatLogUI_SnapshotWindow_Hide, Enemy.CombatLogUI_StatsWindow_Hide, Enemy.GroupsUI_EffectFilterDialog_Hide, Enemy.IntercomUI_ChooseChannelDialog_Hide, Enemy.IntercomUI_IntercomDialog_Hide | `function(...)` | LOW |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | WindowUtils.TrapClick, Enemy.UnitFramesUI_Anchor_OnLButtonDown, AdvancedRenownTraining.Select, BuffHead.Setup.AdvancedCompression.OnRowLDown, BuffHead.Setup.AdvancedCompressionItem.OnRowLDown, BuffHead.Setup.AdvancedContainers.OnRowLDown | `function(...)` | LOW |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | BuffHead.Setup.Layout.OnLayersChanged, BuffHead.Setup.PriorityEffectsItem.OnTargetTypeLUp, EA_LabelCheckButton.Toggle, MapPin.LButtonUp, AnywhereTrainer.OnLButtonUp, BustedGUI.ToggleMainWindow | `function(...)` | LOW |
| [OnMButtonDown](../handlers/handler_OnMButtonDown.md) | Enemy.UnitFramesUI_UnitFrame_OnMButtonDown, MoraleCircle.Reset | `function(...)` | LOW |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | Enemy.UnitFramesUI_UnitFrame_OnMButtonUp, TidyRollFrame.OnMButtonUp | `function(...)` | LOW |
| [OnMouseOut](../handlers/handler_OnMouseOut.md) | Killer.HideRowTooltip | `function()` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOver, RoR_SoR.OnMouseOverStart, MapPin.MouseOver, AggroMeter.OnMouseOverStart, Enemy.ConfigurationWindow_ShowTooltip, Enemy.UnitFramesUI_Anchor_OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOut, Enemy.UnitFramesUI_Anchor_OnMouseOverEnd, TexturedButtons.Setup.Fonts.OnColorExampleMouseOut, BuffHead.Setup.AdvancedCompression.OnRowMouseOut, BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut, BuffHead.Setup.AdvancedContainers.OnRowMouseOut | `function(...)` | LOW |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | FrameManager.OnMouseWheel, MoraleCircle.OnMouseWheel, TidyChat.Copy.OnMouseWheel | `function(delta)` | MEDIUM |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | BuffHead.Setup.AdvancedCompression.OnRowRDown, BuffHead.Setup.AdvancedCompressionItem.OnRowRDown, BuffHead.Setup.AdvancedContainers.OnRowRDown, BuffHead.Setup.AdvancedContainersItem.OnRowRDown, BuffHead.Setup.Layout.OnControlFrameRButtonDown, BuffHead.Setup.Layout.OnLayoutWindowRButtonDown | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | MapPin.RButtonUp, RoR_SoR.BroadCastOption, BustedGUI.ClearAlertFlash, MiracleGrow2.onRClick, RoR_SoR.POPOption, AggroMeter.OnTabRBU | `function(...)` | LOW |
| [OnRawDeviceInput](../handlers/handler_OnRawDeviceInput.md) | BuffHead.Setup.Layout.OnRawDeviceInput | `function(...)` | LOW |
| [OnSetMoving](../handlers/handler_OnSetMoving.md) | MapPin.OnMoving | `function(...)` | LOW |
| [OnShown](../handlers/handler_OnShown.md) | WindowUtils.OnShown, APAGui.OnFollowTargetHUDShown, APAGui.OnInstantOnlyHUDShown, APAGui.OnKitingHUDShown, APAGui.OnPetTargetHUDShown, APAGui.OnShown | `function()` | MEDIUM |
| [OnShutdown](../handlers/handler_OnShutdown.md) | CMapWindow.Shutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, DebugWindow.Shutdown, EA_Window_ContextMenu.Shutdown, EA_Window_Macro.Shutdown | `function()` | MEDIUM |
| [OnSizeUpdated](../handlers/handler_OnSizeUpdated.md) | RoR_SoR.OnSizeUpdated | `function(...)` | LOW |
| [OnUpdate](../handlers/handler_OnUpdate.md) | BuffHead.Setup.Layout.OnUpdate, DebugWindow.Update, TidyRoll.OnUpdate | `function(elapsed)` | MEDIUM |

## Common Inherits

- EA_Window_Default
- EA_Window_DefaultFrame
- EA_LabelCheckButton
- EA_Window_DefaultBackgroundFrame
- EA_TitleBar_Default
- EA_Window_DefaultSeparator
- WSCTEvent
- Aura_LabelCheckButton
- EnemyScenarioInfoDialog_CareerStatsTemplate
- EA_Window_DefaultButtonBottomFrame
- MapPinCallTemplateWindow
- WSCTCheckBox

## Common Parent Elements

- [Window](element_Window.md)
- [ScrollWindow](element_ScrollWindow.md)
- [Button](element_Button.md)

## Common Named Child Elements

- [Label](element_Label.md)
- [Window](element_Window.md)
- [Button](element_Button.md)
- [FullResizeImage](element_FullResizeImage.md)
- [DynamicImage](element_DynamicImage.md)
- [ComboBox](element_ComboBox.md)
- [EditBox](element_EditBox.md)
- [ListBox](element_ListBox.md)

## Common Structural Child Elements

- [Sound](element_Sound.md)
- [Sounds](element_Sounds.md)
- [Visual](element_Visual.md)

## Typical XML Structure

```xml
<Window name="..." movable="true" layer="secondary" alpha="50" savesettings="false">
  <Sounds>
    <Sound event="OnShown" script="Sound.Play( Sound.TOME_OPEN )"/>
    <Sound event="OnHidden" script="Sound.Play( Sound.TOME_CLOSE)"/>
  </Sounds>
</Window>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 49% | MiracleGrow2Check, MapMonster_TooltipInfoTemplate, EA_Window_Default, Aggro_Timer_Template, ... |
| `layer` | optional | 25% | default, background, secondary, overlay, ... |
| `handleinput` | optional | 18% | true, false |
| `movable` | optional | 16% | false, true |
| `savesettings` | optional | 10% | false, true |
| `popable` | optional | 5% | false, true |
| `sticky` | optional | 4% | false, true |
| `scale` | optional | 2% | 1.0, 0.7 |
| `id` | optional | 1% | 1, 6, 2, 10, ... |
| `skipinput` | optional | 1% | false, true |
| `alpha` | optional | 1% | 50, 0.85, 0.97, 1, ... |
| `font` | optional | 0% | font_clear_small_bold, font_chat_text |
| `Scale` | optional | 0% | 0.7 |
| `drawchildrenfirst` | optional | 0% | true |
| `textalign` | optional | 0% | center |
| `draganddrop` | optional | 0% | true |
| `parent` | optional | 0% | root, Root |
| `show` | optional | 0% | false |
| `autoresize` | optional | 0% | true |
| `layout` | optional | 0% | background |
| `moveable` | optional | 0% | false, true |
| `ignorealpha` | optional | 0% | false |
| `inhserits` | optional | 0% | EA_Window_Default |
| `localscriptvars` | optional | 0% | true |
| `showing` | optional | 0% | false |
| `texture` | optional | 0% | EA_TintableImage |
| `textureScale` | optional | 0% | 1 |
| `visible` | optional | 0% | false |
| `warnOnTextCropped` | optional | 0% | false |
| `wordwrap` | optional | 0% | false |
## Structural Sub-Elements

### [Sound](element_Sound.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `event` | **required** |  |
| `script` | **required** |  |
### [Sounds](element_Sounds.md)

Observed 1 times as an unnamed child.

### [Visual](element_Visual.md)

Observed 1 times as an unnamed child.

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Call Count | From Events |
| --- | --- | --- |
| `TextEditBoxSetText` | 414 | OnHidden, OnLButtonUp, OnRButtonUp, OnShown |
| `LabelSetText` | 354 | OnInitialize, OnShown |
| `ButtonSetText` | 167 | OnHidden, OnInitialize, OnShown |
| `WindowSetShowing` | 134 | OnHidden, OnInitialize, OnKeyEscape, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp, OnShown |
| `ButtonSetPressedFlag` | 84 | OnLButtonUp, OnRButtonUp, OnShown |
| `DoesWindowExist` | 75 | OnHidden, OnInitialize, OnKeyEscape, OnMouseOver, OnMouseOverEnd, OnShown |
| `WindowSetDimensions` | 73 | OnLButtonDown, OnShown |
| `ButtonGetPressedFlag` | 62 | OnLButtonUp |
| `WindowGetId` | 46 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonDown, OnRButtonUp |
| `WindowSetTintColor` | 39 | OnInitialize, OnLButtonDown, OnMouseOver, OnMouseOverEnd, OnRButtonDown |
| `WindowUtils.OnHidden` | 19 | OnHidden |
| `WindowUtils.OnShown` | 19 | OnShown |
| `WindowGetParent` | 17 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ComboBoxSetSelectedMenuItem` | 16 | OnRButtonUp |
| `AnimatedImageStartAnimation` | 12 | OnLButtonUp |
| `WindowStartAlphaAnimation` | 11 | OnMouseOver, OnMouseOverEnd |
| `RegisterEventHandler` | 9 | OnInitialize |
| `WindowSetAlpha` | 9 | OnInitialize, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `SliderBarSetCurrentPosition` | 8 | OnRButtonUp |
| `WindowGetShowing` | 8 | OnKeyEscape, OnLButtonUp, OnSetMoving, OnShown |
| `WindowSetGameActionData` | 8 | OnInitialize, OnLButtonDown, OnLButtonUp |
| `ButtonGetDisabledFlag` | 6 | OnLButtonUp |
| `ButtonSetDisabledFlag` | 6 | OnHidden, OnInitialize, OnLButtonUp |
| `BroadcastEvent` | 5 | OnHidden, OnLButtonDown, OnShown |
| `LabelSetTextColor` | 5 | OnInitialize, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `WindowSetMovable` | 5 | OnLButtonUp, OnRButtonUp |
| `CreateWindow` | 3 | OnInitialize |
| `DestroyWindow` | 3 | OnKeyEscape |
| `WindowGetScale` | 3 | OnLButtonUp |
| `WindowUtils.AddToOpenList` | 3 | OnShown |
| `WindowUtils.RemoveFromOpenList` | 3 | OnHidden |
| `ComboBoxSetDisabledFlag` | 2 | OnHidden, OnShown |
| `SystemData.MouseOverWindow.name:find` | 2 | OnLButtonDown, OnRButtonDown |
| `SystemData.MouseOverWindow.name:sub` | 2 | OnLButtonDown, OnRButtonDown |
| `WindowAddAnchor` | 2 | OnLButtonDown, OnShown |
| `WindowAssignFocus` | 2 | OnKeyEscape, OnShown |
| `WindowClearAnchors` | 2 | OnLButtonDown, OnShown |
| `WindowGetAlpha` | 2 | OnMouseOver |
| `WindowGetScreenPosition` | 2 | OnLButtonDown, OnMouseOver |
| `WindowRegisterEventHandler` | 2 | OnInitialize |
| `WindowStopAlphaAnimation` | 2 | OnRButtonUp |
| `LabelGetText` | 1 | OnLButtonDown |
| `ListBoxSetDisplayOrder` | 1 | OnUpdate |
| `TextEditBoxGetText` | 1 | OnKeyEscape |
| `UnregisterEventHandler` | 1 | OnHidden |
| `WindowGetDimensions` | 1 | OnMouseOver |
| `WindowGetMovable` | 1 | OnRButtonUp |
| `WindowUtils.ToggleShowing` | 1 | OnLButtonUp |
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
### OnSetMoving

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `bMoving` | unknown | unknown |
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

- MapPin.MapPin.OnMoving
- wbLeadHelper.wbLeadHelperMessagesTab.OnSaveMessages
- WSCT.WSCT.ColorHideMenu
- Swift Assist.WriteLabels
- GetStats.GetStats.OnInitialize
- EA_UiDebugTools.ObjectInspector.DisplayObject
- PotionBar.PotionBarFloating.ReflowButtons
- WhoHealedMe.IsMainWindowVisible
- wbLeadHelper.wbLeadHelperMessagesTab.ListUp
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakOnly
- LoyalPet.LPET.HideMenu
- RandomMount.RandomMountUI.Toggle
- WSCT.WSCT.OnSetCustomColor
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled4
- MapPin.MapPin.UI_ChooseIconDialog_Hide
- wbLeadHelper.wbLeadHelperConfigWindow.Show
- WSCT.WSCT:OpenMenu
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled1
- WSCT.WSCT.HideMenu
- AggroMeter.AggroMeter.Initialize
- Enemy._OnArchetypeChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.Show
- PotionBar.PotionBarSettings.OnUseCheck
- GuardLine.GuardLine.OnLayoutEditorFinished
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSlotIcons
- MoraleCircle.MoraleCircle.ColorChanger3
- PotionBar.PotionBarSettings.OnShown
- DaemonAssist.DaemonAssist.NormalizeWindowLayout
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSlotIcon
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_AddGroup
- JunkDump.JunkDumpOptions.DynamicBagOptions
- PotionBar.PotionBarSettings.RefreshSingleRightLine
- Swift Assist.SetSmartLabel
- WarBoard.WarBoard.Options.OnSlide
- WarBoard.WarBoard.Options.EnableBoard
- Moth.Moth.HideBorders
- RandomMount.RandomMountUI.Show
- EA_UiDebugTools.BustedGUI.UpdateErrorView
- EA_UiDebugTools.BustedGUI.ClearAlertFlash
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Update
- RandomMount.RandomMountUI.OnMinLevelChanged
- EA_UiDebugTools.BustedGUI.NewErrorRecorded
- Effigy.Effigy.RegisterStateInfoForCastbar
- GetStats.GetStats.CloseWindow
- wbLeadHelper.wbLeadHelperMessagesTab.ListEnable
- AdvancedPetAssist.APAGui.ToggleInstantOnlyHUD
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide
- RoR_SoR.RoR_SoR.Restack
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarPageUpProxy
- DAoCBuff.DAoCTooltips.Init
- Enemy.Enemy.MarksUI_MarkConfigDialog_Hide
- PotionBar.PotionBar.LibSlashHandler
- DAoCBuff.DAoCBuffHeadTracker:Create
- MapPin.MapPin.OnUpdate
- DAoCBuff.DAoCTooltips.CreateCondenseTooltip
- PotionBar.UpdateButton
- PotionBar.PotionBarSettings.OnResetButton
- PotionBar.PotionBarSettings.OnAlphaSliderChanged
- GuardLine.GuardLine.update
- wbLeadHelper.WbLeadHelperMessage.OnOk
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloakOptions
- DaemonAssist.DaemonAssist.PopulateBindingCombos
- Aura.Aura:CreateEditWindows
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelm
- TurretRange.TurretRange.OnUpdate
- RoR_SoR.RoR_SoR.SET_BO
- wbLeadHelper.wbLeadHelperMessagesTab.ListSelChanged
- Busted.Busted.Initialize
- AnywhereTrainerAdditions.AnywhereTrainerAdditions.OnLeftClickAuction
- Enemy.Enemy.GroupsUI_EffectFilterDialog_IsOpen
- DAoCBuff.DAoCTooltips.UpdateCondenseTooltip
- PotionBar.PotionBarFloating.ShowSettings
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloak
- wbLeadHelper.wbLeadHelperConfigTab.OnLfgIconsCheckBoxUp
- Enemy.Enemy.UI_ConfigDialog_Hide
- DAoCBuff.DAoCBuffSettings.CreateOptionswindow
- EA_UiDebugTools.ObjectInspector.ShowWindow
- EA_UiDebugTools.BustedGUI.ToggleMainWindow
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOpacity
- LoyalPet.LPET.SaveProfileOnButtonUp
- RoR_SoR.RoR_SoR.CloseSetOffsetWindow
- MapPin.MapPin.SetupCancel
- MoraleCircle.MoraleCircle.OnSetCustomColor
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelmOnly
- EA_UiDebugTools.BustedGUI.Initialize
- Aura.Aura:CreateRuntimeWindows
- wbLeadHelper.WbLeadHelperMessage.MessageDialogHide
- wbLeadHelper.wbLeadHelperConfigTab.OnReset
- DAoCBuff.DAoCBuffSettings.Disable
- MoraleCircle.MoraleCircle.ColorChanger2
- WhoHealedMe.WhoHealedMe.Initialize
- Enemy.Enemy.TimerInitialize
- RandomMount.RandomMountUI.OnDropSlotLButtonUp
- PotionBar.PotionBarSettings.OnAboutShown
- AggroMeter.AggroMeter.SplitText
- JunkDump.JunkDumpOptions.Done
- WSCT.WSCT.OnHidden
- wbLeadHelper.wbLeadHelper.drawWindows
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Open
- JunkDump.JunkDumpOptions.Show
- wbLeadHelper.WbLeadHelperMessage.MessageDialogIsOpen
- Swift Assist.SwiftAssist.OnMacroUpdated
- Enemy.Enemy.GroupsInitialize
- PotionBar.PotionBarSettings.OnAboutButton
- MoraleCircle.MoraleCircle.init
- DAoCBuff.DAoCBuffSettings.OpenOptionswindow
- AdvancedPetAssist.APAGui.Show
- RoR_SoR.RoR_SoR.SET_CITY
- DAoCBuff.DAoCBuffSettings.PopulateSettings
- Aura.AuraAddon.Slash
- GuardLine.GuardLine.Libguard_Toggle
- Killer.Killer.ShowRowTooltip
- EA_UiDebugTools.ObjectInspector.DepthMinus
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Hide
- PotionBar.PotionBarSettings.OnUseCheckUpdateHighlighting
- TurretRange.TurretRange.local.UpdateDisplay
- Enemy.Enemy.CombatLogUI_IDS_Initialize
- WSCT.WSCT.ColorAcceptButtonOnButtonUp
- Swift Assist.Swift Assist.local.SetSmartLabel
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Hide
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.Hide
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_Hide
- WSCT.WSCT.ColorOnButtonUp
- AdvancedRenownTrainer.AdvancedRenownTrainer.local.CreateAbilityWindow
- Enemy.Enemy.MarksUI_MarkConfigDialog_Open
- Enemy.Enemy.CombatLogUI_EpsWindow_Hide
- RandomMount.RandomMountUI.Hide
- Enemy.Enemy._CombatLogUI_IDS_UpdateWindow
- Moth.Moth.Initialize
- BankArkel.BankArkel.CreatePackWin
- MapPin.MapPin.local.OnHyperLinkLButtonUp2
- Swift Assist.SetTexLabel
- wbLeadHelper.wbLeadHelperMessagesTab.ListDown
- Enemy.Enemy.UI_TextEntryDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_IsOpen
- DAoCBuff.DAoCBuffSettings.SetLabels
- QuickWarReport.QuickWarReport.Initialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSetRow
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Open
- AdvancedPetAssist.APAGui.UpdateInstantOnlyHUD
- Aura.AuraTexture.OnLoad
- FastInteract.FastInteract.OptionsShow
- WarBoard.WarBoard.Options.OpenUiModWindow
- Killer.Killer.ShowTopKillersTooltip
- RoR_SoR.RoR_SoR.SET_KEEP
- wbLeadHelper.wbLeadHelper.OnInitialize
- MapPin.MapPin.local.EditMarker
- RoR_SoR.RoR_SoR.Text_Stream_Fetch
- wbLeadHelper.wbLeadHelperConfigTab.OnSave
- AnywhereTrainer.AnywhereTrainer.OnLeftClickRenown
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- WSCT.WSCT:DisplayText
- Effigy.Effigy.RegisterStateInfoForPlayer
- Enemy.Enemy.UI_ChooseIconDialog_Open
- WhoHealedMe.WHMGui.OnOptionsInitialize
- Enemy.Enemy.Timer_Update
- Shinies.Searches_UpdateWindowDisplay
- WhoHealedMe.WhoHealedMe.local.IsMainWindowVisible
- Enemy.Enemy.local.SetStatsRow
- wbLeadHelper.wbLeadHelperMessagesTab.OnResetMessages
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled5
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Update
- FastInteract.FastInteract.OptionsClose
- MapMonster.InitializeFilters
- AggroMeter.AggroMeter.OnTabLBU
- GuardLine.GuardLine.OnLayoutEditorStart
- wbLeadHelper.wbLeadHelperConfigTab.Initialize
- PotionBar.PotionBarSettings.OnDontSplitNameCheck
- PotionBar.PotionBarSettings.OnScaleSliderChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateActionBarSettings
- MapPin.MapPin.UI_ChooseIconDialog_IsOpen
- EA_UiDebugTools.DevPadWindow.CancelRename
- Swift Assist.Swift Assist.local.WriteLabels
- AdvancedRenownTrainer.CreateAbilityWindow
- MapPin.MapPin.SetupAccept
- PotionBar.PotionBarSettings.OnAboutClose
- Enemy.Enemy.UI_ChooseIconDialog_IsOpen
- RoR_SoR.RoR_SoR.DefaultSettings
- Enemy.Enemy._Initialize
- AdvancedPetAssist.APAGui.UpdateKitingHUD
- RoR_SoR.RoR_SoR.ShowPopper
- AdvancedPetAssist.APAGui.OnShown
- AdvancedRenownTrainer.CreateTab
- AdvancedRenownTrainer.AdvancedRenownTrainer.local.CreateTab
- Swift Assist.Swift Assist.local.SetTexLabel
- WarBoard.WarBoard.OnOptionsButton
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_IsOpen
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOffset
- DAoCBuff.DAoCBuffSettings.Change_Setting
- Moth.Moth.UpdateHealthBar
- TidyRoll.TidyRollOptions.Initialize
- WhoHealedMe.WhoHealedMe.Shutdown
- wbLeadHelper.wbLeadHelperMessagesTab.MessageAdd
- wbLeadHelper.wbLeadHelperMessagesTab.MessageClone
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideCloakOptions
- RoR_SoR.RoR_SoR.CloseSetOpacityWindow
- CM_ClosetGoblin.ClosetGoblinZoneWindow.OnInitialize
- DaemonAssist.DaemonAssist.HideWindow
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- Pocket Palette.PP.UpdateDyeFilter
- RoR_SoR.RoR_SoR.OnInitialize
- Miracle Grow Remix.MiracleGrow2.InitConfig
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- BankArkel.BankArkel.PackShow
- BankArkel.BankArkel.PackImg
- MapPin.MapPin.RButtonUp
- GetStats.GetStats.OnChatLogUpdated
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnInitialize
- MapPin.MapPin.SendText
- MoraleCircle.MoraleCircle.ColorChanger1
- AnywhereTrainer.AnywhereTrainer.OnLeftClickAuction
- DaemonAssist.DaemonAssist.ToggleWindow
- Killer.Killer.RefreshPersonalCounter
- LoyalPet.LPET.RenameProfileOnButtonUp
- wbLeadHelper.wbLeadHelper.createWbLeadHelperWindow
- EA_UiDebugTools.DevPadWindow.Hide
- PotionBar.PotionBar.Initialize
- FastInteract.FastInteract.OptionsSetup
- Enemy.Enemy.CombatLogUI_EpsWindow_IsOpen
- WhoHealedMe.WHMCommands.CmdConfig
- LoyalPet.LPET.AddProfileOnButtonUp
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_IsOpen
- GuardLine.GuardLine.OffTarget
- EA_UiDebugTools.ObjectInspector.AddInputHistory
- LibSlash.LibSlash.Initialize
- Busted.BustedGUI.NewErrorRecorded
- Enemy.Enemy.IntercomInitialize
- wbLeadHelper.wbLeadHelperMessagesTab.Hide
- LoyalPet.LPET.OpenMenu
- MapPin.MapPin.UI_ChooseIconDialog_Open
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UseItemRackToggled
- Enemy.Enemy.local._OnKeyModifierChanged
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Open
- JunkDump.JunkDumpOptions.InitSettings
- PotionBar.PotionBar.local.UpdateButton
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloakHeraldry
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideShowHelm
- Enemy.Enemy.UI_ChooseIconDialog_Hide
- QuickWarReport.QuickWarReport.local.PrepareConfirmWindowChrome
- DAoCBuff.DAoCBuffSettings.UC
- MapPin.MapPin.TestTooltip
- EA_UiDebugTools.ObjectInspector.DepthPlus
- Effigy.Effigy.RegisterStateInfoForTargets
- RoR_SoR.RoR_SoR.SetWindowShow
- WhoHealedMe.WHMGui.RefreshConfigurationWindow
- BuffHead.BuffHead.local.RegisterLayoutEditor
- WhoHealedMe.WHMGui.ShowOptionsWindow
- RoR_SoR.RoR_SoR.OnSlideWindowOptionsOpacity
- GuardLine.GuardLine.update2
- Effigy.EffigyBar.Init
- RandomMount.RandomMountUI.Refresh
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Open
- QuickWarReport.PrepareConfirmWindowChrome
- RoR_SoR.RoR_SoR.OnWindowOptionsSetScale
- MapMonster.MapMonster.local.InitializeFilters
- MoraleCircle.MoraleCircle.OnSetCustomColorFill
- GuardLine.GuardLine.GetIDs
- Busted.BustedGUI.ClearAlertFlash
- Moth.Moth.Clear
- AdvancedPetAssist.APAGui.UpdateFollowTargetHUD
- BankArkel.BankArkel.BackPackShow
- wbLeadHelper.wbLeadHelper.chat
- MoraleCircle.MoraleCircle.OnSetCustomColorFull
- Enemy.Enemy.StopwatchEnabledChanged
- WSCT.WSCT.OnLButtonUpColorPicker
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Initialize
- AdvancedPetAssist.ApplyHUDVisibilityFromSettings
- Enemy.Enemy.local._OnArchetypeChanged
- QuickWarReport.QuickWarReport.local.ShowConfirmWindow
- DaemonAssist.DaemonAssist.UpdateWindow
- BankArkel.BankArkel.Init
- Enemy.Enemy.Assist_OnIntercomMessage
- GuardLine.GuardLine.init
- DAoCBuff.DAoCBuffFrame.MouseOverUpdate
- AdvancedRenownTrainer.AdvancedRenownTraining.OnShown
- QuickWarReport.QuickWarReport.local.HideConfirmWindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- PotionBar.PotionBarFloating.Alpha
- PotionBar.PotionBarSettings.Show1Check
- AdvancedRenownTrainer.AdvancedRenownTraining.OnHidden
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateList
- WhoHealedMe.WHMCore.ApplyBackgroundFillColor
- WSCT.WSCT.UpdateAnimationOptions
- Enemy.Enemy.CombatLogUI_EpsWindow_Open
- Enemy.Enemy.Stopwatch_Update
- Enemy.Enemy.IntercomUI_IntercomDialog_Open
- JunkDump.JunkDumpOptions.DestroyOptionsWindow
- PotionBar.PotionBar.Hide
- TurretRange.TurretRange.OnLoadComplete
- AdvancedPetAssist.APAGui.ToggleKitingHUD
- wbLeadHelper.wbLeadHelperConfigTab.Hide
- AdvancedRenownTrainer.AdvancedRenownTraining.Initialize
- Enemy.Enemy.CombatLogUI_StatsWindow_Open
- Enemy.Enemy.CombatLogUI_StatsWindow_Hide
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Hide
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_IsOpen
- Enemy._OnKeyModifierChanged
- Enemy.Enemy.MarksUI_MarkConfigDialog_IsOpen
- PotionBar.PotionBarSettings.OnCancelButton
- Killer.Killer.Initialize
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Enemy.Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- AdvancedPetAssist.APAGui.ApplyPetTargetHUDLayout
- Enemy.Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- AdvancedPetAssist.APAGui.UpdatePetTargetHUD
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Ok
- MapPin.MapPin.ShowIcons
- MoraleCircle.MoraleCircle.OnSetCustomColorEmpty
- QuickWarReport.ShowConfirmWindow
- AggroMeter.AggroMeter.Close
- PotionBar.PotionBarFloating.Scale
- wbLeadHelper.wbLeadHelperConfigTab.OnChanged
- wbLeadHelper.wbLeadHelperMessagesTab.ListDelete
- EA_UiDebugTools.DevPadWindow.HideSaveWindow
- Enemy.Enemy.MarksInitialize
- Enemy.Enemy.UI_ConfigDialog_OnSectionSelChanged
- DAoCBuff.DAoCBuff.CloseMessageWindow
- wbLeadHelper.wbLeadHelper.onZoneMouseOver
- DaemonAssist.DaemonAssist.ShowWindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled3
- Enemy.Enemy.CombatLogUI_IDS_OnSettingsChanged
- EA_UiDebugTools.ObjectInspector.Toggle
- JunkDump.JunkDumpOptions.CreateOptionsWindow
- PotionBar.PotionBarSettings.UpdateLastCheckBasedOnInfoText
- Busted.BustedGUI.UpdateNextPrevButtonStatus
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Update
- AutoMark.AutoMark.local.CreateMarker
- Swift Assist.SwiftAssist.aaLabelColorSet
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnShow
- RandomMount.RandomMountUI.OnAddCustomMount
- MiracleGrowLight.MiracleGrowLight.switchBackground
- Enemy.Enemy.IntercomUI_IntercomDialog_Hide
- JunkDump.CheckAndSetButton
- QuickWarReport.QuickWarReport.local.EnsureConfirmWindow
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- WhoHealedMe.WHMGui.ToggleOptionsWindow
- BankArkel.BankArkel.SetupCombos
- PotionBar.PotionBar.Shutdown
- MoraleCircle.MoraleCircle.ColorChanger4
- wbLeadHelper.wbLeadHelper.showNormalTitle
- AdvancedPetAssist.AnchorInContent
- Enemy.Enemy.MarksUI_EnemyMarksWindow_IsOpen
- JunkDump.JunkDump.local.CheckAndSetButton
- RoR_SoR.RoR_SoR.OnChatLogUpdated
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_IsOpen
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_IsOpen
- RoR_SoR.RoR_SoR.OnCombat
- RoR_SoR.RoR_SoR.SET_CAMPAIGN
- BankArkel.BankArkel.BackPackHide
- DAoCBuff.DAoCBuffSettings.Reactivate
- Enemy.Enemy.CombatLogUI_EpsWindow_Update
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled2
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Hide
- wbLeadHelper.wbLeadHelperConfigTab.ChooseIconMessageStart
- PotionBar.PotionBar.Show
- MiracleGrow.MiracleGrow.Initialize
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- CM_ClosetGoblin.ClosetGoblinZoneWindow.Show
- EA_UiDebugTools.ObjectInspector.InspectObject
- Moth.Moth.Anchor
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged
- Enemy.Enemy.AssistInitialize
- Enemy.SetStatsRow
- AutoMark.CreateMarker
- RandomMount.RandomMountUI.OnInitialize
- QuickWarReport.QuickWarReport.TestConfirmWindow
- EA_UiDebugTools.ObjectInspector.ClearObjects
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide
- WSCT.WSCT.ColorOnInitialize
- Swift Assist.SwiftAssist.Initialize
- RoR_SoR.RoR_SoR.OnSlideWindowOptionsScale
- CM_ClosetGoblin.ClosetGoblinZoneWindow.RefreshOption
- wbLeadHelper.wbLeadHelperConfigTab.locallyStoreFormData
- PotionBar.PotionBarSettings.OnHidden
- BankArkel.BankArkel.PackHide
- MapPin.MapPin.test
- DAoCBuff.DAoCBuffSettings.CloseOptionswindow
- Miracle Grow Remix.MiracleGrow2.ConfigTabChange
- WarBoard.WarBoard.LoadGeneralSettings
- AdvancedPetAssist.AdvancedPetAssist.local.ApplyHUDVisibilityFromSettings
- EA_UiDebugTools.DevPadWindow.HideLoadProject
- QuickWarReport.EnsureConfirmWindow
- QuickWarReport.QuickWarReport.Shutdown
- Aura.AuraAddon.OnLoad
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.Enemy.CombatLogUI_EpsWindow_UpdateLayout
- WhoHealedMe.WHMCore.RunSettingEffects
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarPageDownProxy
- TidyChat.TidyChat.LootRoll.OnRollLinkLButtonUp
- MapPin.OnHyperLinkLButtonUp2
- EA_UiDebugTools.BustedGUI.UpdateNextPrevButtonStatus
- RoR_SoR.RoR_SoR.CloseSetScaleWindow
- LibWBToggler.LibWBTogglerManager.Initialize
- wbLeadHelper.wbLeadHelperConfigTab.ChooseIconMessageEnd
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Open
- AdvancedPetAssist.AdvancedPetAssist.local.AnchorInContent
- Killer.Killer.ShowPersonalStatsTooltip
- wbLeadHelper.wbLeadHelperConfigWindow.Initialize
- PotionBar.PotionBarSettings.AutohideCheck
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowHelm
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_Open
- EA_UiDebugTools.DevPadWindow.HideNewWindow
- Busted.BustedGUI.Initialize
- EA_UiDebugTools.ObjectInspector.WindowInit
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- Enemy.EnemyMarkTemplate:ToggleMark
- wbLeadHelper.wbLeadHelperConfigWindow.Hide
- Enemy.Enemy.UI_ConfigDialog_Open
- Enemy.Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- RoR_SoR.RoR_SoR.OnScenario
- EA_UiDebugTools.ObjectInspector.CloseWindow
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- Busted.BustedGUI.UpdateErrorView
- Enemy.Enemy.CombatLogUI_IDS_Update
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Hide
- wbLeadHelper.wbLeadHelperMessagesTab.Initialize
- EA_UiDebugTools.DevPadWindow.OnKeyEscape
- EA_UiDebugTools.DevPadWindow.HideDeleteWindow
- BankArkel.BankArkel.SetCharInfo
- Enemy.Enemy.CombatLogUI_StatsWindow_IsOpen
- Effigy.EffigyBar:setup
- MapPin.EditMarker
- Enemy.EnemyGroupIcon:Attach
- BuffHead.RegisterLayoutEditor
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_OnOkButton
- Killer.Killer.ApplyFeedLayout
- CM_ClosetGoblin.ClosetGoblinZoneWindow.Hide
- wbLeadHelper.wbLeadHelperMessagesTab.Show
- wbLeadHelper.wbLeadHelperMessagesTab.MessageEdit
- wbLeadHelper.wbLeadHelperConfigTab.Show
- wbLeadHelper.wbLeadHelperConfigTab.OnLoad
- RoR_SoR.RoR_SoR.HidePopper
- WSCT.WSCT:ObjectIDAnimation
- WhoHealedMe.WHMGui.HideOptionsWindow
- PotionBar.PotionBarSettings.QuickActionsSelChanged
- AdvancedRenownTrainer.GeneratePresetByLinkData
- TurretRange.UpdateDisplay
- WarBoard.WarBoard.SlashCommand
- wbLeadHelper.WbLeadHelperMessage.MessageDialogOpen
- MapPin.MapPin.LButtonUp
- Enemy.EnemyUnitFrame:ApplySettings
- DAoCBuff.DAoCBuff.ShowMessageWindow
- Aura.AuraTexture.OnClose
- CM_ClosetGoblin.ClosetGoblinZoneWindow.UpdateHighlightOnRow
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateSessionsList
- Moth.Moth.UpdateTarget
- Enemy.Enemy.AssistUI_Target_Show
- JunkDump.JunkDumpOptions.UpdateBagModeOnOff
- EA_UiDebugTools.DevPadWindow.HideConfirmLoadWindow
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Hide
- AdvancedPetAssist.APAGui.Hide
- QuickWarReport.HideConfirmWindow
- Moth.Moth.HealthBar
- AdvancedPetAssist.APAGui.HidePetTargetHUD
- Pocket Palette.PP.UpdateListRow
- Enemy.Enemy.UnitFramesUI_ConfigDialog_Import
- AdvancedRenownTrainer.AdvancedRenownTraining.ChangeTab
- Moth.Moth.UpdateLevel
- Enemy.Enemy.CombatLogUI_EpsWindow_Initialize

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
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- EA_ThreePartBar
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- Enemy
- FastInteract
- GetStats
- GuardLine
- GuardList
- GuardRange
- JunkDump
- Killer
- LibGroup
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- ObjectInspector
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAFollowTargetHUD -> Window APAFollowTargetHUD
- AdvancedPetAssist: APAInstantOnlyHUD -> Window APAInstantOnlyHUD
- AdvancedPetAssist: APAKitingHUD -> Window APAKitingHUD
- AdvancedPetAssist: APAOptions -> Window APAOptions
- AdvancedPetAssist: APAOptionsBackground -> Window APAOptionsBackground
- AdvancedPetAssist: APAOptionsContent -> Window APAOptionsContent

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [ButtonGetTextColor](../../window_api/functions/window_ButtonGetTextColor.md) (HIGH 100/100) - Window Function
- [ButtonSetCheckButtonFlag](../../window_api/functions/window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ButtonSetTextColor](../../window_api/functions/window_ButtonSetTextColor.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetDisabledFlag](../../window_api/functions/window_ComboBoxGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetDisabledFlag](../../window_api/functions/window_ComboBoxSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [DynamicImageSetRotation](../../window_api/functions/window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](../../window_api/functions/window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap.ShowZone](../../globals/functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
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
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [ScrollWindowSetOffset](../../window_api/functions/window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](../../window_api/functions/window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [StatusBarGetCurrentValue](../../window_api/functions/window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetBackgroundTint](../../window_api/functions/window_StatusBarSetBackgroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetCurrentValue](../../window_api/functions/window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](../../window_api/functions/window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](../../window_api/functions/window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
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
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetPopable](../../window_api/functions/window_WindowSetPopable.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](../../window_api/functions/window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function
- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 90/100) - Window Function
- [ButtonSetTexture](../../window_api/functions/window_ButtonSetTexture.md) (HIGH 90/100) - Window Function
- [ButtonSetTextureSlice](../../window_api/functions/window_ButtonSetTextureSlice.md) (HIGH 90/100) - Window Function
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 90/100) - Window Function
- [WindowRestoreDefaultSettings](../../window_api/functions/window_WindowRestoreDefaultSettings.md) (HIGH 90/100) - Window Function
- [WindowSetMoving](../../window_api/functions/window_WindowSetMoving.md) (HIGH 90/100) - Window Function
- [WindowSetResizing](../../window_api/functions/window_WindowSetResizing.md) (HIGH 90/100) - Window Function
- [WindowStartScaleAnimation](../../window_api/functions/window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [ComboBoxExternalOpenMenu](../../window_api/functions/window_ComboBoxExternalOpenMenu.md) (HIGH 80/100) - Window Function
- [TextEditBoxGetHistory](../../window_api/functions/window_TextEditBoxGetHistory.md) (HIGH 80/100) - Window Function
- [TextEditBoxInsertText](../../window_api/functions/window_TextEditBoxInsertText.md) (HIGH 80/100) - Window Function
- [TextEditBoxSetHistory](../../window_api/functions/window_TextEditBoxSetHistory.md) (HIGH 80/100) - Window Function
- [TextEditBoxSetMaxChars](../../window_api/functions/window_TextEditBoxSetMaxChars.md) (HIGH 80/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 100/100) - XML Handler
- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [OnKeyEscape](../../events/window_events/window_event_OnKeyEscape.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMButtonDown](../../events/window_events/window_event_OnMButtonDown.md) (HIGH 100/100) - Window Event
- [OnMButtonUp](../../events/window_events/window_event_OnMButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../../events/window_events/window_event_OnMouseWheel.md) (HIGH 100/100) - Window Event
- [OnRButtonDown](../../events/window_events/window_event_OnRButtonDown.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnShutdown](../../events/window_events/window_event_OnShutdown.md) (HIGH 100/100) - Window Event
- [OnUpdate](../../events/window_events/window_event_OnUpdate.md) (HIGH 100/100) - Window Event
- [OnMouseOut](../../events/window_events/window_event_OnMouseOut.md) (HIGH 73/100) - Window Event
- [OnSetMoving](../../events/window_events/window_event_OnSetMoving.md) (HIGH 73/100) - Window Event

## Triggered By

- none

## Affects

- none
