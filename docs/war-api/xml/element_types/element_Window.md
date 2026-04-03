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

- name
- inherits
- layer
- movable
- handleinput
- savesettings
- popable
- sticky
- scale
- skipinput
- drawchildrenfirst
- alpha

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnHidden](../handlers/handler_OnHidden.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnShown](../handlers/handler_OnShown.md)
- [OnInitialize](../handlers/handler_OnInitialize.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md)
- [OnMouseOut](../handlers/handler_OnMouseOut.md)
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md)

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

## Common Inherits

- EA_Window_Default
- EA_LabelCheckButton
- EA_Window_DefaultBackgroundFrame
- EA_TitleBar_Default
- EA_Window_DefaultFrame
- WSCTEvent
- Aura_LabelCheckButton
- EnemyScenarioInfoDialog_CareerStatsTemplate
- EA_Window_DefaultSeparator
- WSCTCheckBox
- EA_Window_DefaultButtonBottomFrame
- TChatCheckboxTemplate

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

- [Size](element_Size.md) — 624× (HIGH)
- [Windows](element_Windows.md) — 476× (HIGH)
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

### [Size](element_Size.md)

Observed 624 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 476 times as an unnamed child.

### [Visual](element_Visual.md)

Observed 1 times as an unnamed child.

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

- Enemy.UI_TextEntryDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- Enemy.local.SetStatsRow
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Open
- AdvancedRenownTraining.OnShown
- AggroMeter.Close
- ClosetGoblinCharacterWindow.ShowCloakOptions
- ClosetGoblinCharacterWindow.HotbarPageDownProxy
- MoraleCircle.OnSetCustomColorEmpty
- TidyChat.LootRoll.OnRollLinkLButtonUp
- WhoHealedMe.local.IsMainWindowVisible
- ClosetGoblinCharacterWindow.OnShow
- PotionBarSettings.OnUseCheckUpdateHighlighting
- PotionBarSettings.Show1Check
- AnywhereTrainer.OnLeftClickRenown
- Enemy.UI_ChooseIconDialog_Hide
- DAoCBuffSettings.OpenOptionswindow
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Open
- Enemy.CombatLogUI_EpsWindow_Update
- MoraleCircle.OnSetCustomColor
- PotionBarFloating.ReflowButtons
- Aura.Aura:CreateRuntimeWindows
- Enemy.IntercomUI_ChooseChannelDialog_OnOkButton
- Enemy.Stopwatch_Update
- Enemy.CombatLogUI_StatsWindow_Open
- ClosetGoblinCharacterWindow.UseItemRackToggled
- Enemy._OnArchetypeChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_IsOpen
- RoR_SoR.OnWindowOptionsSetOpacity
- Killer.ShowPersonalStatsTooltip
- APAGui.UpdateFollowTargetHUD
- AdvancedPetAssist.local.ApplyHUDVisibilityFromSettings
- Aura.Aura:CreateEditWindows
- BankArkel.PackImg
- WSCT.OnHidden
- WhoHealedMe.IsMainWindowVisible
- Enemy.local._OnKeyModifierChanged
- Killer.ShowTopKillersTooltip
- MoraleCircle.OnSetCustomColorFull
- Swift Assist.SetSmartLabel
- Enemy.CombatLogUI_StatsWindow_UpdateSessionsList
- Enemy.CombatLogUI_EpsWindow_Initialize
- MiracleGrowLight.switchBackground
- RoR_SoR.OnWindowOptionsSetScale
- BankArkel.BackPackHide
- Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged
- Enemy.MarksUI_EnemyMarksWindow_Update
- Enemy.EnemyUnitFrame:ApplySettings
- WSCT.WSCT:DisplayText
- RoR_SoR.OnInitialize
- WHMGui.HideOptionsWindow
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.CombatLogUI_TargetDefenseWindow_IsOpen
- PotionBarSettings.UpdateLastCheckBasedOnInfoText
- BankArkel.SetupCombos
- SwiftAssist.OnMacroUpdated
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- ClosetGoblinCharacterWindow.UpdateSetRow
- ClosetGoblinCharacterWindow.UpdateActionBarSettings
- Enemy.IntercomUI_ChooseChannelDialog_Hide
- Enemy.MarksInitialize
- Enemy.MarksUI_MarkConfigDialog_IsOpen
- Enemy.EnemyGroupIcon:Attach
- GuardLine.update2
- AdvancedPetAssist.AnchorInContent
- ClosetGoblinZoneWindow.Show
- Enemy.UnitFramesUI_ConfigDialog_Import
- Enemy.IntercomUI_IntercomJoinDialog_Open
- MoraleCircle.ColorChanger2
- PotionBarSettings.OnResetButton
- RoR_SoR.OnSlideWindowOptionsOpacity
- Enemy.IntercomUI_ChooseChannelDialog_Open
- Enemy.local._OnArchetypeChanged
- MoraleCircle.init
- Swift Assist.WriteLabels
- AuraTexture.OnLoad
- ClosetGoblinCharacterWindow.UpdateSlotIcons
- ClosetGoblinCharacterWindow.HotbarPageUpProxy
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- WHMGui.ShowOptionsWindow
- AdvancedPetAssist.local.AnchorInContent
- ClosetGoblinCharacterWindow.OnInitialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- AdvancedRenownTrainer.local.CreateTab
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- RoR_SoR.SET_KEEP
- APAGui.HidePetTargetHUD
- ClosetGoblinCharacterWindow.UpdateSlotIcon
- Enemy.UI_ConfigDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_Hide
- DAoCBuffSettings.CloseOptionswindow
- Enemy.IntercomUI_IntercomDialog_Hide
- AuraTexture.OnClose
- AutoMark.local.CreateMarker
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- AnywhereTrainerAdditions.OnLeftClickAuction
- WSCT.ColorOnButtonUp
- WHMCore.RunSettingEffects
- BankArkel.SetCharInfo
- DAoCTooltips.Init
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- PotionBarSettings.AutohideCheck
- Enemy.GroupsUI_EffectFilterDialog_Hide
- RoR_SoR.CloseSetScaleWindow
- Swift Assist.local.WriteLabels
- ClosetGoblinZoneWindow.RefreshOption
- Enemy.CombatLogUI_TargetDefenseWindow_Initialize
- GuardLine.update
- WarBoard.Options.OnSlide
- AdvancedRenownTraining.Initialize
- DAoCBuff.ShowMessageWindow
- Enemy.GroupsUI_EffectFilterDialog_IsOpen
- Enemy._CombatLogUI_IDS_UpdateWindow
- WSCT.UpdateAnimationOptions
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Hide
- PotionBarSettings.OnDontSplitNameCheck
- Swift Assist.local.SetTexLabel
- Enemy.MarksUI_MarkConfigDialog_Hide
- Enemy.CombatLogUI_TargetDefenseWindow_Open
- RoR_SoR.ShowPopper
- APAGui.UpdateInstantOnlyHUD
- ClosetGoblinCharacterWindow.Show
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- DAoCBuffSettings.Reactivate
- PotionBar.UpdateButton
- WSCT.WSCT:ObjectIDAnimation
- APAGui.ToggleKitingHUD
- Enemy.IntercomUI_IntercomJoinDialog_AddGroup
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- PotionBar.local.UpdateButton
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Update
- PP.UpdateListRow
- RoR_SoR.OnWindowOptionsSetOffset
- APAGui.Hide
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- Enemy.UnitFramesUI_EffectsIndicatorDialog_IsOpen
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.CombatLogUI_StatsWindow_Hide
- TurretRange.UpdateDisplay
- ClosetGoblinCharacterWindow.Hide
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- ClosetGoblinZoneWindow.Hide
- Enemy.IntercomUI_IntercomDialog_Open
- APAGui.OnShown
- DAoCBuffFrame.MouseOverUpdate
- WarBoard.Options.EnableBoard
- Enemy.CombatLogUI_IDS_OnSettingsChanged
- RoR_SoR.OnScenario
- Shinies.Searches_UpdateWindowDisplay
- WSCT.ColorAcceptButtonOnButtonUp
- AdvancedRenownTraining.ChangeTab
- Enemy.EnemyMarkTemplate:ToggleMark
- RoR_SoR.CloseSetOpacityWindow
- TurretRange.local.UpdateDisplay
- Enemy.CombatLogUI_EpsWindow_IsOpen
- PotionBarFloating.ShowSettings
- ClosetGoblinZoneWindow.UpdateHighlightOnRow
- Enemy.UI_ChooseIconDialog_IsOpen
- Enemy.IntercomInitialize
- Enemy.CombatLogUI_EpsWindow_Open
- WSCT.ColorHideMenu
- BankArkel.CreatePackWin
- Enemy.UnitFramesUI_UnitFramePartDialog_IsOpen
- PotionBar.Shutdown
- RoR_SoR.OnSlideWindowOptionsScale
- AdvancedRenownTrainer.local.CreateAbilityWindow
- Enemy.AssistUI_Target_Show
- MoraleCircle.ColorChanger4
- PotionBarSettings.RefreshSingleRightLine
- AutoMark.CreateMarker
- ClosetGoblinZoneWindow.OnInitialize
- DAoCBuffSettings.SetLabels
- Enemy.MarksUI_EnemyMarksWindow_Open
- PotionBarSettings.OnAboutShown
- WarBoard.SlashCommand
- AggroMeter.Initialize
- ClosetGoblinCharacterWindow.ShowShowHelm
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_IsOpen
- Enemy.CombatLogUI_EpsWindow_UpdateLayout
- WarBoard.LoadGeneralSettings
- WHMGui.RefreshConfigurationWindow
- Enemy.MarksUI_EnemyMarksWindow_IsOpen
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- TurretRange.OnLoadComplete
- PotionBarSettings.OnShown
- SwiftAssist.aaLabelColorSet
- WSCT.ColorOnInitialize
- WSCT.WSCT:OpenMenu
- DAoCBuff.CloseMessageWindow
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- Enemy.CombatLogUI_TargetDefenseWindow_Update
- Killer.ApplyFeedLayout
- AdvancedRenownTraining.OnHidden
- Enemy.GroupsUI_EffectFilterDialog_Ok
- Killer.RefreshPersonalCounter
- TidyRollOptions.Initialize
- WHMCore.ApplyBackgroundFillColor
- ClosetGoblinCharacterWindow.ShowCloak
- DAoCBuffSettings.CreateOptionswindow
- Enemy.MarksUI_MarkConfigDialog_Open
- PotionBarSettings.QuickActionsSelChanged
- PotionBarSettings.OnAlphaSliderChanged
- RoR_SoR.SET_BO
- Swift Assist.SetTexLabel
- WarBoard.OnOptionsButton
- APAGui.UpdatePetTargetHUD
- DAoCBuffSettings.Disable
- Enemy.CombatLogUI_EpsWindow_Hide
- Killer.ShowRowTooltip
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.TimerInitialize
- Enemy.CombatLogUI_StatsWindow_UpdateList
- WSCT.OnSetCustomColor
- PotionBarFloating.Alpha
- AnywhereTrainer.OnLeftClickAuction
- APAGui.UpdateKitingHUD
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide
- Enemy.Timer_Update
- GuardLine.OnLayoutEditorFinished
- ClosetGoblinCharacterWindow.HideCloakOptions
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- RoR_SoR.Text_Stream_Fetch
- RoR_SoR.OnCombat
- PotionBarSettings.OnCancelButton
- SwiftAssist.Initialize
- AuraAddon.OnLoad
- BankArkel.Init
- BuffHead.local.RegisterLayoutEditor
- DAoCTooltips.CreateCondenseTooltip
- WHMGui.OnOptionsInitialize
- APAGui.Show
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- MoraleCircle.ColorChanger3
- PotionBarFloating.Scale
- GuardLine.Libguard_Toggle
- GuardLine.OffTarget
- RoR_SoR.DefaultSettings
- RoR_SoR.Restack
- Enemy.CombatLogUI_TargetDefenseWindow_Hide
- Killer.Initialize
- BankArkel.PackShow
- BankArkel.PackHide
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- Enemy.CombatLogUI_IDS_Initialize
- AdvancedRenownTrainer.GeneratePresetByLinkData
- Enemy.GroupsInitialize
- Enemy.UI_ConfigDialog_Hide
- Enemy.AssistInitialize
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- Enemy.CombatLogUI_IDS_Update
- GuardLine.GetIDs
- BuffHead.RegisterLayoutEditor
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- DAoCBuffSettings.PopulateSettings
- RoR_SoR.SET_CITY
- RoR_SoR.OnChatLogUpdated
- BankArkel.BackPackShow
- Enemy.Assist_OnIntercomMessage
- WHMGui.ToggleOptionsWindow
- AdvancedPetAssist.ApplyHUDVisibilityFromSettings
- AdvancedRenownTrainer.CreateAbilityWindow
- AdvancedRenownTrainer.CreateTab
- AggroMeter.SplitText
- WHMCommands.CmdConfig
- ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- PotionBarSettings.OnAboutButton
- RoR_SoR.SET_CAMPAIGN
- TurretRange.OnUpdate
- RoR_SoR.CloseSetOffsetWindow
- ClosetGoblinCharacterWindow.ShowHelm
- DAoCTooltips.UpdateCondenseTooltip
- PP.UpdateDyeFilter
- PotionBar.Hide
- Swift Assist.local.SetSmartLabel
- Enemy.MarksUI_EnemyMarksWindow_Hide
- Enemy.SetStatsRow
- PotionBar.Show
- PotionBarSettings.OnAboutClose
- APAGui.ToggleInstantOnlyHUD
- PotionBar.LibSlashHandler
- MoraleCircle.OnSetCustomColorFill
- WSCT.HideMenu
- AuraAddon.Slash
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.IntercomUI_IntercomJoinDialog_Hide
- Enemy.StopwatchEnabledChanged
- RoR_SoR.SetWindowShow
- WhoHealedMe.Shutdown
- DAoCBuffSettings.UC
- DAoCBuffSettings.Change_Setting
- Enemy.CombatLogUI_TargetDefenseTotalWindow_IsOpen
- GuardLine.OnLayoutEditorStart
- GuardLine.init
- DAoCBuff.DAoCBuffHeadTracker:Create
- Enemy._Initialize
- Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy._OnKeyModifierChanged
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide
- PotionBar.Initialize
- APAGui.ApplyPetTargetHUDLayout
- AggroMeter.OnTabLBU
- ClosetGoblinCharacterWindow.HideShowHelm
- Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- PotionBarSettings.OnScaleSliderChanged
- WhoHealedMe.Initialize
- Enemy.UI_ChooseIconDialog_Open
- Enemy.CombatLogUI_StatsWindow_IsOpen
- MoraleCircle.ColorChanger1
- PotionBarSettings.OnUseCheck
- PotionBarSettings.OnHidden
- RoR_SoR.HidePopper
- WSCT.OnLButtonUpColorPicker
- LibWBTogglerManager.Initialize


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
- [EA_Window_ContextMenu.AddCascadingMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
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
- [Visual](element_Visual.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
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
