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

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnHidden](../handlers/handler_OnHidden.md) | lifecycle | WindowUtils.OnHidden, APAGui.OnFollowTargetHUDHidden, APAGui.OnHidden, APAGui.OnInstantOnlyHUDHidden, APAGui.OnKitingHUDHidden, AdvancedRenownTraining.OnExportHidden | `function()` | MEDIUM |
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_LabelCheckButton.Initialize, AdvancedRenownTraining.Initialize, AuraConfig.OnInitialize, CMapWindow.Initialize, CMapWindow.RefreshMapPointFilterMenu, ClosetGoblinOptionWindow.OnInitialize | `function()` | MEDIUM |
| [OnKeyEscape](../handlers/handler_OnKeyEscape.md) | custom | DebugWindow.OnKeyEscape, Enemy.CombatLogUI_SnapshotWindow_Hide, Enemy.CombatLogUI_StatsWindow_Hide, Enemy.GroupsUI_EffectFilterDialog_Hide, Enemy.IntercomUI_ChooseChannelDialog_Hide, Enemy.IntercomUI_IntercomDialog_Hide | `function(...)` | LOW |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | WindowUtils.TrapClick, Enemy.UnitFramesUI_Anchor_OnLButtonDown, AdvancedRenownTraining.Select, BuffHead.Setup.AdvancedCompression.OnRowLDown, BuffHead.Setup.AdvancedCompressionItem.OnRowLDown, BuffHead.Setup.AdvancedContainers.OnRowLDown | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | BuffHead.Setup.Layout.OnLayersChanged, BuffHead.Setup.PriorityEffectsItem.OnTargetTypeLUp, EA_LabelCheckButton.Toggle, MapPin.LButtonUp, AnywhereTrainer.OnLButtonUp, BustedGUI.ToggleMainWindow | `flags, x, y` | MEDIUM |
| [OnMButtonDown](../handlers/handler_OnMButtonDown.md) | input | Enemy.UnitFramesUI_UnitFrame_OnMButtonDown, MoraleCircle.Reset | `flags, x, y` | MEDIUM |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | Enemy.UnitFramesUI_UnitFrame_OnMButtonUp, TidyRollFrame.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOut](../handlers/handler_OnMouseOut.md) | input | Killer.HideRowTooltip | `function()` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOver, RoR_SoR.OnMouseOverStart, MapPin.MouseOver, AggroMeter.OnMouseOverStart, Enemy.ConfigurationWindow_ShowTooltip, Enemy.UnitFramesUI_Anchor_OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | BuffHead.Setup.Layout.Properties.OnColorExampleMouseOut, Enemy.UnitFramesUI_Anchor_OnMouseOverEnd, TexturedButtons.Setup.Fonts.OnColorExampleMouseOut, BuffHead.Setup.AdvancedCompression.OnRowMouseOut, BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut, BuffHead.Setup.AdvancedContainers.OnRowMouseOut | `function(...)` | LOW |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | input | FrameManager.OnMouseWheel, MoraleCircle.OnMouseWheel, TidyChat.Copy.OnMouseWheel | `function(delta)` | MEDIUM |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | BuffHead.Setup.AdvancedCompression.OnRowRDown, BuffHead.Setup.AdvancedCompressionItem.OnRowRDown, BuffHead.Setup.AdvancedContainers.OnRowRDown, BuffHead.Setup.AdvancedContainersItem.OnRowRDown, BuffHead.Setup.Layout.OnControlFrameRButtonDown, BuffHead.Setup.Layout.OnLayoutWindowRButtonDown | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | MapPin.RButtonUp, RoR_SoR.BroadCastOption, BustedGUI.ClearAlertFlash, MiracleGrow2.onRClick, RoR_SoR.POPOption, AggroMeter.OnTabRBU | `flags, x, y` | MEDIUM |
| [OnRawDeviceInput](../handlers/handler_OnRawDeviceInput.md) | custom | BuffHead.Setup.Layout.OnRawDeviceInput | `function(...)` | LOW |
| [OnSetMoving](../handlers/handler_OnSetMoving.md) | custom | MapPin.OnMoving | `function(...)` | LOW |
| [OnShown](../handlers/handler_OnShown.md) | lifecycle | WindowUtils.OnShown, APAGui.OnFollowTargetHUDShown, APAGui.OnInstantOnlyHUDShown, APAGui.OnKitingHUDShown, APAGui.OnPetTargetHUDShown, APAGui.OnShown | `function()` | MEDIUM |
| [OnShutdown](../handlers/handler_OnShutdown.md) | lifecycle | CMapWindow.Shutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, DebugWindow.Shutdown, EA_Window_ContextMenu.Shutdown, EA_Window_Macro.Shutdown | `function()` | MEDIUM |
| [OnSizeUpdated](../handlers/handler_OnSizeUpdated.md) | custom | RoR_SoR.OnSizeUpdated | `function(...)` | LOW |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | BuffHead.Setup.Layout.OnUpdate, DebugWindow.Update, TidyRoll.OnUpdate | `function(elapsed)` | MEDIUM |

### Per-Event Lua API Calls

**OnHidden** handlers call: `BroadcastEvent`, `ButtonSetDisabledFlag`, `ButtonSetText`, `ComboBoxSetDisabledFlag`, `DoesWindowExist`, `TextEditBoxSetText`, `UnregisterEventHandler`, `WindowSetShowing`, `WindowUtils.OnHidden`, `WindowUtils.RemoveFromOpenList`

**OnInitialize** handlers call: `ButtonSetDisabledFlag`, `ButtonSetText`, `CreateWindow`, `DoesWindowExist`, `LabelSetText`, `LabelSetTextColor`, `RegisterEventHandler`, `WindowRegisterEventHandler`, `WindowSetAlpha`, `WindowSetGameActionData`, `WindowSetShowing`, `WindowSetTintColor`

**OnKeyEscape** handlers call: `DestroyWindow`, `DoesWindowExist`, `TextEditBoxGetText`, `WindowAssignFocus`, `WindowGetShowing`, `WindowSetShowing`

**OnLButtonDown** handlers call: `BroadcastEvent`, `LabelGetText`, `SystemData.MouseOverWindow.name:find`, `SystemData.MouseOverWindow.name:sub`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetDimensions`, `WindowSetGameActionData`, `WindowSetTintColor`

**OnLButtonUp** handlers call: `AnimatedImageStartAnimation`, `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `TextEditBoxSetText`, `WindowGetId`, `WindowGetParent`, `WindowGetScale`, `WindowGetShowing`, `WindowSetGameActionData`, `WindowSetMovable`, `WindowSetShowing`, `WindowUtils.ToggleShowing`

**OnMouseOver** handlers call: `DoesWindowExist`, `LabelSetTextColor`, `WindowGetAlpha`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetAlpha`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnMouseOverEnd** handlers call: `DoesWindowExist`, `LabelSetTextColor`, `WindowSetAlpha`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnRButtonDown** handlers call: `SystemData.MouseOverWindow.name:find`, `SystemData.MouseOverWindow.name:sub`, `WindowGetId`, `WindowSetTintColor`

**OnRButtonUp** handlers call: `ButtonSetPressedFlag`, `ComboBoxSetSelectedMenuItem`, `LabelSetTextColor`, `SliderBarSetCurrentPosition`, `TextEditBoxSetText`, `WindowGetId`, `WindowGetMovable`, `WindowGetParent`, `WindowSetAlpha`, `WindowSetMovable`, `WindowSetShowing`, `WindowStopAlphaAnimation`

**OnSetMoving** handlers call: `WindowGetShowing`

**OnShown** handlers call: `BroadcastEvent`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxSetDisabledFlag`, `DoesWindowExist`, `LabelSetText`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowAssignFocus`, `WindowClearAnchors`, `WindowGetShowing`, `WindowSetDimensions`, `WindowSetShowing`, `WindowUtils.AddToOpenList`, `WindowUtils.OnShown`

**OnUpdate** handlers call: `ListBoxSetDisplayOrder`

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

- [Window](element_Window.md) — 262× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 26× (HIGH)
- [Button](element_Button.md) — 10× (HIGH)

## Common Named Child Elements

- [Label](element_Label.md) — 379× (HIGH)
- [Window](element_Window.md) — 262× (HIGH)
- [Button](element_Button.md) — 241× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 153× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 132× (HIGH)
- [ComboBox](element_ComboBox.md) — 97× (HIGH)
- [EditBox](element_EditBox.md) — 88× (HIGH)
- [ListBox](element_ListBox.md) — 47× (HIGH)
- [SliderBar](element_SliderBar.md) — 41× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 25× (HIGH)
- [CircleImage](element_CircleImage.md) — 13× (HIGH)
- [AnimatedImage](element_AnimatedImage.md) — 12× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 12× (HIGH)
- [StatusBar](element_StatusBar.md) — 5× (MEDIUM)
- [MapDisplay](element_MapDisplay.md) — 3× (MEDIUM)
- [ActionButtonGroup](element_ActionButtonGroup.md) — 2× (LOW)
- [LogDisplay](element_LogDisplay.md) — 2× (LOW)
- [VerticalScrollbar](element_VerticalScrollbar.md) — 2× (LOW)
- [ColorPicker](element_ColorPicker.md) — 1× (LOW)
- [CooldownDisplay](element_CooldownDisplay.md) — 1× (LOW)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 1× (LOW)

## Common Structural Child Elements

- [Sound](element_Sound.md) — 1× (LOW)
- [Sounds](element_Sounds.md) — 1× (LOW)
- [Visual](element_Visual.md) — 1× (LOW)

## Common Template Bases

- Aggro_Timer_Template
- AnywhereTrainerTabTemplate
- Aura_LabelCheckButton
- Aura_LargeLabelCheckButton
- BuffHeadColorExample
- BuffHeadSetupSelectTextureRowTemplate
- ClosetGoblinActionBarPageSelector
- Cmap_Template_OverheadMapZoomSlider
- DAoCBuffCondenseTooltipItem
- DAoCBuff_FakeSettingsRow
- EA_LabelCheckButton
- EA_LabelCheckButtonSmallCopy
- EA_TitleBar_Default
- EA_Window_AdvancedMainUIPathToggle
- EA_Window_CityRating
- EA_Window_ComboBoxMenuBackground
- EA_Window_Default
- EA_Window_DefaultBackgroundFrame
- EA_Window_DefaultButtonBottomFrame
- EA_Window_DefaultContextMenuFrame
- EA_Window_DefaultFrame
- EA_Window_DefaultFrameStatusBar
- EA_Window_DefaultLabelToggleCircle
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
- GuardList_Window0
- GuardRange_Window0
- IraConfigCheckBox
- LogFilterButton
- MapMonster_TooltipDefault
- MapMonster_TooltipInfoTemplate
- MapPinCallTemplateWindow
- MiracleGrow2BarTint
- MiracleGrow2Check
- MiracleGrow2Checker
- MiracleGrow2ConfigTemplate
- MiracleGrow2LayoutPlotsTemplate
- MiracleGrow2LayoutProgTemplate
- MiracleGrow2LayoutSettingsTemplate
- MiracleGrow2LayoutTemplate
- MiracleGrow2Line
- MiracleGrow2PresetTemplate
- MiracleGrow2SoundLine
- MiracleGrowLightLine
- MiracleGrowLine
- MiracleGrowStartAllLine
- MoneyFrame
- OptionsTemplate
- RVAPI_ColorDialogColorBoxTemplate
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
- wbLeadHelperConfigTab
- wbLeadHelperMessagesTab


> **Note**: This element type commonly acts as a template base.

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
| `inherits` | optional | 49% | EA_Window_DefaultSeparator, EA_Window_DefaultFrame, EA_Window_DefaultBackgroundFrame, EA_LabelCheckButton, ... |
| `layer` | optional | 25% | popup, background, secondary, default, ... |
| `handleinput` | optional | 18% | false, true |
| `movable` | optional | 16% | true, false |
| `savesettings` | optional | 10% | false, true |
| `popable` | optional | 5% | false, true |
| `sticky` | optional | 4% | false, true |
| `scale` | optional | 2% | 1.0, 0.7 |
| `id` | optional | 1% | 1, 2, 9, 10, ... |
| `skipinput` | optional | 1% | false, true |
| `alpha` | optional | 1% | 0.97, 0.0, 50, 1, ... |
| `font` | optional | 0% | font_clear_small_bold, font_chat_text |
| `Scale` | optional | 0% | 0.7 |
| `drawchildrenfirst` | optional | 0% | true |
| `textalign` | optional | 0% | center |
| `draganddrop` | optional | 0% | true |
| `parent` | optional | 0% | Root, root |
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

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `TextEditBoxSetText` | ui | 414 | OnHidden, OnLButtonUp, OnRButtonUp, OnShown |
| `LabelSetText` | ui | 354 | OnInitialize, OnShown |
| `ButtonSetText` | ui | 167 | OnHidden, OnInitialize, OnShown |
| `WindowSetShowing` | ui | 134 | OnHidden, OnInitialize, OnKeyEscape, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp, OnShown |
| `ButtonSetPressedFlag` | ui | 84 | OnLButtonUp, OnRButtonUp, OnShown |
| `DoesWindowExist` | ui | 75 | OnHidden, OnInitialize, OnKeyEscape, OnMouseOver, OnMouseOverEnd, OnShown |
| `WindowSetDimensions` | ui | 73 | OnLButtonDown, OnShown |
| `ButtonGetPressedFlag` | ui | 62 | OnLButtonUp |
| `WindowGetId` | ui | 46 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonDown, OnRButtonUp |
| `WindowSetTintColor` | ui | 39 | OnInitialize, OnLButtonDown, OnMouseOver, OnMouseOverEnd, OnRButtonDown |
| `WindowUtils.OnHidden` | ui | 19 | OnHidden |
| `WindowUtils.OnShown` | ui | 19 | OnShown |
| `WindowGetParent` | ui | 17 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ComboBoxSetSelectedMenuItem` | ui | 16 | OnRButtonUp |
| `AnimatedImageStartAnimation` | ui | 12 | OnLButtonUp |
| `WindowStartAlphaAnimation` | ui | 11 | OnMouseOver, OnMouseOverEnd |
| `RegisterEventHandler` | event | 9 | OnInitialize |
| `WindowSetAlpha` | ui | 9 | OnInitialize, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `SliderBarSetCurrentPosition` | ui | 8 | OnRButtonUp |
| `WindowGetShowing` | ui | 8 | OnKeyEscape, OnLButtonUp, OnSetMoving, OnShown |
| `WindowSetGameActionData` | ui | 8 | OnInitialize, OnLButtonDown, OnLButtonUp |
| `ButtonGetDisabledFlag` | ui | 6 | OnLButtonUp |
| `ButtonSetDisabledFlag` | ui | 6 | OnHidden, OnInitialize, OnLButtonUp |
| `BroadcastEvent` | event | 5 | OnHidden, OnLButtonDown, OnShown |
| `LabelSetTextColor` | ui | 5 | OnInitialize, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `WindowSetMovable` | ui | 5 | OnLButtonUp, OnRButtonUp |
| `CreateWindow` | ui | 3 | OnInitialize |
| `DestroyWindow` | ui | 3 | OnKeyEscape |
| `WindowGetScale` | ui | 3 | OnLButtonUp |
| `WindowUtils.AddToOpenList` | ui | 3 | OnShown |
| `WindowUtils.RemoveFromOpenList` | ui | 3 | OnHidden |
| `ComboBoxSetDisabledFlag` | ui | 2 | OnHidden, OnShown |
| `SystemData.MouseOverWindow.name:find` | data | 2 | OnLButtonDown, OnRButtonDown |
| `SystemData.MouseOverWindow.name:sub` | data | 2 | OnLButtonDown, OnRButtonDown |
| `WindowAddAnchor` | ui | 2 | OnLButtonDown, OnShown |
| `WindowAssignFocus` | ui | 2 | OnKeyEscape, OnShown |
| `WindowClearAnchors` | ui | 2 | OnLButtonDown, OnShown |
| `WindowGetAlpha` | ui | 2 | OnMouseOver |
| `WindowGetScreenPosition` | ui | 2 | OnLButtonDown, OnMouseOver |
| `WindowRegisterEventHandler` | event | 2 | OnInitialize |
| `WindowStopAlphaAnimation` | ui | 2 | OnRButtonUp |
| `LabelGetText` | ui | 1 | OnLButtonDown |
| `ListBoxSetDisplayOrder` | ui | 1 | OnUpdate |
| `TextEditBoxGetText` | ui | 1 | OnKeyEscape |
| `UnregisterEventHandler` | event | 1 | OnHidden |
| `WindowGetDimensions` | ui | 1 | OnMouseOver |
| `WindowGetMovable` | ui | 1 | OnRButtonUp |
| `WindowUtils.ToggleShowing` | ui | 1 | OnLButtonUp |
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

- TurretRange.TurretRange.local.UpdateDisplay
- Busted.BustedGUI.NewErrorRecorded
- PotionBar.PotionBarSettings.UpdateLastCheckBasedOnInfoText
- wbLeadHelper.wbLeadHelperMessagesTab.ListDelete
- EA_UiDebugTools.ObjectInspector.CloseWindow
- WhoHealedMe.WHMCore.RunSettingEffects
- Busted.BustedGUI.UpdateNextPrevButtonStatus
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloakOptions
- Swift Assist.SwiftAssist.aaLabelColorSet
- RoR_SoR.RoR_SoR.OnSlideWindowOptionsScale
- Enemy.Enemy.CombatLogUI_EpsWindow_Update
- RoR_SoR.RoR_SoR.OnScenario
- Aura.Aura:CreateEditWindows
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- JunkDump.JunkDumpOptions.DestroyOptionsWindow
- RandomMount.RandomMountUI.Show
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOffset
- PotionBar.PotionBarSettings.OnDontSplitNameCheck
- RandomMount.RandomMountUI.OnMinLevelChanged
- DaemonAssist.DaemonAssist.NormalizeWindowLayout
- RandomMount.RandomMountUI.OnInitialize
- PotionBar.PotionBarSettings.OnUseCheckUpdateHighlighting
- wbLeadHelper.WbLeadHelperMessage.MessageDialogHide
- MapPin.MapPin.LButtonUp
- Enemy.Enemy.CombatLogUI_IDS_Initialize
- Enemy.Enemy.IntercomUI_IntercomDialog_Hide
- wbLeadHelper.wbLeadHelperConfigWindow.Hide
- Enemy.Enemy.CombatLogUI_StatsWindow_Open
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- PotionBar.PotionBarFloating.ReflowButtons
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowHelm
- PotionBar.PotionBarSettings.OnHidden
- AdvancedPetAssist.APAGui.ApplyPetTargetHUDLayout
- PotionBar.PotionBarSettings.OnCancelButton
- GuardLine.GuardLine.update
- PotionBar.PotionBarFloating.Scale
- AutoMark.AutoMark.local.CreateMarker
- Effigy.Effigy.RegisterStateInfoForTargets
- WarBoard.WarBoard.LoadGeneralSettings
- RoR_SoR.RoR_SoR.DefaultSettings
- wbLeadHelper.wbLeadHelperConfigTab.Show
- EA_UiDebugTools.DevPadWindow.Hide
- Busted.Busted.Initialize
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- DaemonAssist.DaemonAssist.HideWindow
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- RoR_SoR.RoR_SoR.HidePopper
- LoyalPet.LPET.HideMenu
- AdvancedRenownTrainer.AdvancedRenownTraining.Initialize
- RoR_SoR.RoR_SoR.OnInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled4
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakOnly
- Enemy.Enemy.CombatLogUI_StatsWindow_Hide
- Swift Assist.Swift Assist.local.SetTexLabel
- wbLeadHelper.wbLeadHelperMessagesTab.OnResetMessages
- BankArkel.BankArkel.Init
- BankArkel.BankArkel.PackHide
- DAoCBuff.DAoCTooltips.CreateCondenseTooltip
- WSCT.WSCT.HideMenu
- Enemy.Enemy.TimerInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- WarBoard.WarBoard.Options.OnSlide
- RoR_SoR.RoR_SoR.SET_CITY
- JunkDump.JunkDump.local.CheckAndSetButton
- Enemy.SetStatsRow
- Enemy.Enemy.local._OnKeyModifierChanged
- Enemy.Enemy.IntercomInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSlotIcons
- MiracleGrow.MiracleGrow.Initialize
- BankArkel.BankArkel.BackPackShow
- AdvancedPetAssist.APAGui.UpdateInstantOnlyHUD
- QuickWarReport.QuickWarReport.local.EnsureConfirmWindow
- MoraleCircle.MoraleCircle.ColorChanger3
- wbLeadHelper.wbLeadHelperMessagesTab.Initialize
- AdvancedRenownTrainer.AdvancedRenownTraining.OnHidden
- QuickWarReport.QuickWarReport.Shutdown
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnShow
- Enemy.Enemy.CombatLogUI_EpsWindow_Hide
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelm
- DAoCBuff.DAoCBuff.ShowMessageWindow
- PotionBar.PotionBar.local.UpdateButton
- WhoHealedMe.WhoHealedMe.Initialize
- Enemy._OnArchetypeChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled3
- PotionBar.PotionBarSettings.OnAboutButton
- Enemy.Enemy.IntercomUI_IntercomDialog_Open
- DAoCBuff.DAoCBuffSettings.Disable
- BankArkel.BankArkel.PackShow
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- PotionBar.PotionBarSettings.AutohideCheck
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateList
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Open
- AdvancedPetAssist.AdvancedPetAssist.local.ApplyHUDVisibilityFromSettings
- Shinies.Searches_UpdateWindowDisplay
- AdvancedRenownTrainer.AdvancedRenownTraining.OnShown
- Effigy.Effigy.RegisterStateInfoForPlayer
- Enemy.Enemy.CombatLogUI_EpsWindow_UpdateLayout
- RoR_SoR.RoR_SoR.OnSlideWindowOptionsOpacity
- Moth.Moth.Anchor
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.Hide
- DAoCBuff.DAoCBuffSettings.UC
- Killer.Killer.Initialize
- AggroMeter.AggroMeter.Close
- MapPin.MapPin.OnUpdate
- Enemy.Enemy.StopwatchEnabledChanged
- Enemy.Enemy.GroupsUI_EffectFilterDialog_IsOpen
- QuickWarReport.PrepareConfirmWindowChrome
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideCloakOptions
- MapPin.MapPin.RButtonUp
- wbLeadHelper.wbLeadHelperMessagesTab.ListSelChanged
- Enemy.Enemy.UI_ConfigDialog_OnSectionSelChanged
- WhoHealedMe.WHMGui.ShowOptionsWindow
- WhoHealedMe.WHMGui.ToggleOptionsWindow
- MoraleCircle.MoraleCircle.ColorChanger1
- EA_UiDebugTools.DevPadWindow.HideSaveWindow
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Hide
- MiracleGrowLight.MiracleGrowLight.switchBackground
- RoR_SoR.RoR_SoR.SET_KEEP
- WSCT.WSCT.ColorHideMenu
- Enemy.Enemy.UI_ChooseIconDialog_Open
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_Open
- wbLeadHelper.wbLeadHelperConfigTab.Initialize
- GuardLine.GuardLine.update2
- Moth.Moth.UpdateLevel
- WSCT.WSCT:DisplayText
- MapPin.MapPin.OnMoving
- Enemy.Enemy.CombatLogUI_IDS_Update
- GetStats.GetStats.OnInitialize
- Enemy.Enemy._Initialize
- RandomMount.RandomMountUI.Hide
- EA_UiDebugTools.DevPadWindow.HideLoadProject
- AdvancedPetAssist.APAGui.HidePetTargetHUD
- Moth.Moth.UpdateHealthBar
- wbLeadHelper.wbLeadHelper.createWbLeadHelperWindow
- wbLeadHelper.wbLeadHelperConfigWindow.Show
- PotionBar.PotionBarSettings.OnAboutClose
- MoraleCircle.MoraleCircle.OnSetCustomColor
- Enemy.Enemy.Timer_Update
- RoR_SoR.RoR_SoR.CloseSetScaleWindow
- wbLeadHelper.wbLeadHelperMessagesTab.ListUp
- EA_UiDebugTools.ObjectInspector.ShowWindow
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Update
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelmOnly
- Swift Assist.SetTexLabel
- AdvancedPetAssist.ApplyHUDVisibilityFromSettings
- AdvancedPetAssist.APAGui.Hide
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_IsOpen
- GuardLine.GuardLine.init
- wbLeadHelper.wbLeadHelper.OnInitialize
- RoR_SoR.RoR_SoR.SET_BO
- WhoHealedMe.WHMCore.ApplyBackgroundFillColor
- WhoHealedMe.IsMainWindowVisible
- Enemy.Enemy.CombatLogUI_IDS_OnSettingsChanged
- Enemy.Enemy.UI_ChooseIconDialog_IsOpen
- Enemy.Enemy.CombatLogUI_EpsWindow_IsOpen
- RoR_SoR.RoR_SoR.SET_CAMPAIGN
- Effigy.EffigyBar.Init
- Swift Assist.SwiftAssist.Initialize
- PotionBar.PotionBar.Initialize
- JunkDump.CheckAndSetButton
- EA_UiDebugTools.ObjectInspector.AddInputHistory
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_IsOpen
- WSCT.WSCT:OpenMenu
- RandomMount.RandomMountUI.OnDropSlotLButtonUp
- CM_ClosetGoblin.ClosetGoblinZoneWindow.UpdateHighlightOnRow
- RoR_SoR.RoR_SoR.Restack
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled2
- Effigy.EffigyBar:setup
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Hide
- Enemy.Enemy.CombatLogUI_StatsWindow_IsOpen
- DAoCBuff.DAoCBuffSettings.Reactivate
- AdvancedRenownTrainer.AdvancedRenownTrainer.local.CreateTab
- Aura.AuraTexture.OnClose
- MapPin.OnHyperLinkLButtonUp2
- EA_UiDebugTools.BustedGUI.ClearAlertFlash
- FastInteract.FastInteract.OptionsClose
- MoraleCircle.MoraleCircle.OnSetCustomColorFill
- QuickWarReport.QuickWarReport.local.PrepareConfirmWindowChrome
- wbLeadHelper.wbLeadHelperConfigTab.OnLoad
- Killer.Killer.ShowRowTooltip
- Swift Assist.Swift Assist.local.SetSmartLabel
- AdvancedPetAssist.APAGui.UpdatePetTargetHUD
- MapPin.MapPin.ShowIcons
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_OnOkButton
- DAoCBuff.DAoCBuffSettings.PopulateSettings
- PotionBar.PotionBar.Shutdown
- EA_UiDebugTools.BustedGUI.ToggleMainWindow
- AnywhereTrainer.AnywhereTrainer.OnLeftClickRenown
- DAoCBuff.DAoCTooltips.Init
- MapPin.MapPin.SetupAccept
- WSCT.WSCT.OnSetCustomColor
- Miracle Grow Remix.MiracleGrow2.ConfigTabChange
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_IsOpen
- FastInteract.FastInteract.OptionsSetup
- WhoHealedMe.WhoHealedMe.Shutdown
- Enemy.Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- JunkDump.JunkDumpOptions.UpdateBagModeOnOff
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSetRow
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Hide
- LibSlash.LibSlash.Initialize
- Moth.Moth.HealthBar
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Hide
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide
- DaemonAssist.DaemonAssist.ToggleWindow
- RoR_SoR.RoR_SoR.CloseSetOffsetWindow
- AutoMark.CreateMarker
- MapPin.MapPin.TestTooltip
- MapPin.MapPin.SetupCancel
- WarBoard.WarBoard.SlashCommand
- PotionBar.PotionBarSettings.OnScaleSliderChanged
- Moth.Moth.HideBorders
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_IsOpen
- Swift Assist.WriteLabels
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateActionBarSettings
- Moth.Moth.Initialize
- DAoCBuff.DAoCBuffSettings.OpenOptionswindow
- CM_ClosetGoblin.ClosetGoblinZoneWindow.RefreshOption
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Open
- Killer.Killer.RefreshPersonalCounter
- Enemy.Enemy.UnitFramesUI_ConfigDialog_Import
- Enemy.Enemy.Assist_OnIntercomMessage
- BankArkel.BankArkel.BackPackHide
- JunkDump.JunkDumpOptions.CreateOptionsWindow
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- TurretRange.TurretRange.OnLoadComplete
- wbLeadHelper.wbLeadHelperConfigTab.Hide
- WSCT.WSCT.ColorAcceptButtonOnButtonUp
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.Show
- PotionBar.PotionBarSettings.OnAlphaSliderChanged
- MoraleCircle.MoraleCircle.ColorChanger4
- wbLeadHelper.wbLeadHelperConfigTab.OnLfgIconsCheckBoxUp
- Enemy.Enemy.UI_ConfigDialog_Open
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- AdvancedRenownTrainer.GeneratePresetByLinkData
- LoyalPet.LPET.SaveProfileOnButtonUp
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_IsOpen
- GuardLine.GuardLine.GetIDs
- Pocket Palette.PP.UpdateListRow
- wbLeadHelper.wbLeadHelperConfigTab.OnChanged
- AdvancedRenownTrainer.AdvancedRenownTrainer.local.CreateAbilityWindow
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.Enemy.MarksInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled5
- CM_ClosetGoblin.ClosetGoblinZoneWindow.Hide
- Enemy.Enemy.CombatLogUI_EpsWindow_Open
- EA_UiDebugTools.ObjectInspector.Toggle
- Swift Assist.SetSmartLabel
- DAoCBuff.DAoCBuff.CloseMessageWindow
- PotionBar.PotionBarSettings.OnResetButton
- MapMonster.MapMonster.local.InitializeFilters
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged
- PotionBar.PotionBar.LibSlashHandler
- Enemy.Enemy.AssistInitialize
- MapPin.MapPin.local.OnHyperLinkLButtonUp2
- BuffHead.RegisterLayoutEditor
- PotionBar.PotionBarSettings.OnShown
- Enemy.Enemy.CombatLogUI_EpsWindow_Initialize
- AdvancedPetAssist.APAGui.ToggleInstantOnlyHUD
- RandomMount.RandomMountUI.OnAddCustomMount
- EA_UiDebugTools.BustedGUI.UpdateErrorView
- DAoCBuff.DAoCBuffHeadTracker:Create
- QuickWarReport.QuickWarReport.local.HideConfirmWindow
- RoR_SoR.RoR_SoR.OnCombat
- GetStats.GetStats.CloseWindow
- wbLeadHelper.wbLeadHelper.drawWindows
- GuardLine.GuardLine.OnLayoutEditorStart
- Enemy.Enemy.local.SetStatsRow
- LibWBToggler.LibWBTogglerManager.Initialize
- PotionBar.PotionBar.Show
- JunkDump.JunkDumpOptions.Show
- CM_ClosetGoblin.ClosetGoblinZoneWindow.OnInitialize
- TidyRoll.TidyRollOptions.Initialize
- DAoCBuff.DAoCBuffSettings.CreateOptionswindow
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- AggroMeter.AggroMeter.SplitText
- BuffHead.BuffHead.local.RegisterLayoutEditor
- GuardLine.GuardLine.OffTarget
- Enemy.Enemy.GroupsInitialize
- WSCT.WSCT.OnHidden
- WarBoard.WarBoard.OnOptionsButton
- JunkDump.JunkDumpOptions.InitSettings
- AdvancedPetAssist.AnchorInContent
- PotionBar.PotionBarFloating.Alpha
- wbLeadHelper.wbLeadHelperMessagesTab.MessageAdd
- wbLeadHelper.WbLeadHelperMessage.MessageDialogIsOpen
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarPageUpProxy
- WhoHealedMe.WHMCommands.CmdConfig
- WarBoard.WarBoard.Options.OpenUiModWindow
- JunkDump.JunkDumpOptions.DynamicBagOptions
- AdvancedPetAssist.AdvancedPetAssist.local.AnchorInContent
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Open
- MoraleCircle.MoraleCircle.init
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled1
- AdvancedPetAssist.APAGui.ToggleKitingHUD
- Pocket Palette.PP.UpdateDyeFilter
- wbLeadHelper.wbLeadHelperConfigTab.ChooseIconMessageEnd
- AdvancedRenownTrainer.CreateAbilityWindow
- TidyChat.TidyChat.LootRoll.OnRollLinkLButtonUp
- MapPin.MapPin.UI_ChooseIconDialog_IsOpen
- Enemy.EnemyUnitFrame:ApplySettings
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Update
- QuickWarReport.QuickWarReport.TestConfirmWindow
- RoR_SoR.RoR_SoR.Text_Stream_Fetch
- wbLeadHelper.wbLeadHelperConfigWindow.Initialize
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Ok
- MapMonster.InitializeFilters
- wbLeadHelper.wbLeadHelperConfigTab.OnSave
- EA_UiDebugTools.DevPadWindow.CancelRename
- Enemy._OnKeyModifierChanged
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- AdvancedPetAssist.APAGui.UpdateFollowTargetHUD
- AdvancedPetAssist.APAGui.Show
- Aura.AuraTexture.OnLoad
- wbLeadHelper.wbLeadHelper.onZoneMouseOver
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- RoR_SoR.RoR_SoR.OnWindowOptionsSetScale
- Busted.BustedGUI.UpdateErrorView
- Moth.Moth.UpdateTarget
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_IsOpen
- RoR_SoR.RoR_SoR.OnChatLogUpdated
- GuardLine.GuardLine.Libguard_Toggle
- MoraleCircle.MoraleCircle.OnSetCustomColorFull
- RandomMount.RandomMountUI.Refresh
- wbLeadHelper.wbLeadHelperMessagesTab.ListDown
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Hide
- WSCT.WSCT:ObjectIDAnimation
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- EA_UiDebugTools.ObjectInspector.DepthPlus
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Open
- DAoCBuff.DAoCBuffFrame.MouseOverUpdate
- QuickWarReport.QuickWarReport.local.ShowConfirmWindow
- EA_UiDebugTools.BustedGUI.NewErrorRecorded
- TurretRange.UpdateDisplay
- LoyalPet.LPET.OpenMenu
- WSCT.WSCT.OnLButtonUpColorPicker
- Enemy.EnemyGroupIcon:Attach
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSlotIcon
- WhoHealedMe.WHMGui.RefreshConfigurationWindow
- WSCT.WSCT.UpdateAnimationOptions
- EA_UiDebugTools.DevPadWindow.HideConfirmLoadWindow
- wbLeadHelper.WbLeadHelperMessage.MessageDialogOpen
- Killer.Killer.ShowPersonalStatsTooltip
- TurretRange.TurretRange.OnUpdate
- PotionBar.PotionBar.Hide
- EA_UiDebugTools.ObjectInspector.DepthMinus
- Enemy.Enemy.MarksUI_EnemyMarksWindow_IsOpen
- RoR_SoR.RoR_SoR.CloseSetOpacityWindow
- wbLeadHelper.wbLeadHelperMessagesTab.MessageEdit
- QuickWarReport.QuickWarReport.Initialize
- Aura.AuraAddon.OnLoad
- wbLeadHelper.wbLeadHelperConfigTab.OnReset
- Enemy.Enemy.Stopwatch_Update
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Initialize
- DAoCBuff.DAoCBuffSettings.SetLabels
- EA_UiDebugTools.ObjectInspector.InspectObject
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarPageDownProxy
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- EA_UiDebugTools.ObjectInspector.DisplayObject
- RandomMount.RandomMountUI.Toggle
- EA_UiDebugTools.BustedGUI.UpdateNextPrevButtonStatus
- Enemy.Enemy.local._OnArchetypeChanged
- RoR_SoR.RoR_SoR.SetWindowShow
- FastInteract.FastInteract.OptionsShow
- EA_UiDebugTools.ObjectInspector.ClearObjects
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide
- BankArkel.BankArkel.CreatePackWin
- Moth.Moth.Clear
- AnywhereTrainer.AnywhereTrainer.OnLeftClickAuction
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOpacity
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloakHeraldry
- Enemy.Enemy.MarksUI_MarkConfigDialog_Hide
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- BankArkel.BankArkel.SetCharInfo
- DAoCBuff.DAoCTooltips.UpdateCondenseTooltip
- DAoCBuff.DAoCBuffSettings.Change_Setting
- AggroMeter.AggroMeter.OnTabLBU
- Enemy.Enemy._CombatLogUI_IDS_UpdateWindow
- wbLeadHelper.wbLeadHelperMessagesTab.OnSaveMessages
- Killer.Killer.ApplyFeedLayout
- Busted.BustedGUI.Initialize
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Update
- AggroMeter.AggroMeter.Initialize
- AdvancedPetAssist.APAGui.OnShown
- MoraleCircle.MoraleCircle.ColorChanger2
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Open
- QuickWarReport.ShowConfirmWindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideShowHelm
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloak
- EA_UiDebugTools.DevPadWindow.OnKeyEscape
- WSCT.WSCT.ColorOnButtonUp
- QuickWarReport.EnsureConfirmWindow
- EA_UiDebugTools.BustedGUI.Initialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- AdvancedPetAssist.APAGui.UpdateKitingHUD
- wbLeadHelper.wbLeadHelperMessagesTab.MessageClone
- WhoHealedMe.WHMGui.HideOptionsWindow
- Enemy.Enemy.UI_ChooseIconDialog_Hide
- AnywhereTrainerAdditions.AnywhereTrainerAdditions.OnLeftClickAuction
- Swift Assist.Swift Assist.local.WriteLabels
- WSCT.WSCT.ColorOnInitialize
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_Hide
- RoR_SoR.RoR_SoR.ShowPopper
- MapPin.MapPin.UI_ChooseIconDialog_Hide
- PotionBar.PotionBarFloating.ShowSettings
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- EA_UiDebugTools.DevPadWindow.HideNewWindow
- PotionBar.UpdateButton
- CM_ClosetGoblin.ClosetGoblinZoneWindow.Show
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UseItemRackToggled
- MoraleCircle.MoraleCircle.OnSetCustomColorEmpty
- PotionBar.PotionBarSettings.OnUseCheck
- PotionBar.PotionBarSettings.Show1Check
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_AddGroup
- QuickWarReport.HideConfirmWindow
- DaemonAssist.DaemonAssist.UpdateWindow
- Miracle Grow Remix.MiracleGrow2.InitConfig
- Enemy.Enemy.UI_TextEntryDialog_Open
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Hide
- MapPin.MapPin.local.EditMarker
- EA_UiDebugTools.ObjectInspector.WindowInit
- BankArkel.BankArkel.SetupCombos
- Busted.BustedGUI.ClearAlertFlash
- BankArkel.BankArkel.PackImg
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Open
- Aura.Aura:CreateRuntimeWindows
- wbLeadHelper.wbLeadHelperConfigTab.locallyStoreFormData
- AdvancedRenownTrainer.CreateTab
- Enemy.Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- Enemy.Enemy.AssistUI_Target_Show
- Effigy.Effigy.RegisterStateInfoForCastbar
- Enemy.Enemy.MarksUI_MarkConfigDialog_Open
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide
- Enemy.EnemyMarkTemplate:ToggleMark
- MapPin.EditMarker
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- WarBoard.WarBoard.Options.EnableBoard
- wbLeadHelper.WbLeadHelperMessage.OnOk
- AdvancedRenownTrainer.AdvancedRenownTraining.ChangeTab
- Enemy.Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- wbLeadHelper.wbLeadHelperMessagesTab.Hide
- JunkDump.JunkDumpOptions.Done
- PotionBar.PotionBarSettings.OnAboutShown
- wbLeadHelper.wbLeadHelperMessagesTab.Show
- DAoCBuff.DAoCBuffSettings.CloseOptionswindow
- MapPin.MapPin.test
- LoyalPet.LPET.AddProfileOnButtonUp
- WhoHealedMe.WhoHealedMe.local.IsMainWindowVisible
- MapPin.MapPin.SendText
- Swift Assist.SwiftAssist.OnMacroUpdated
- Enemy.Enemy.UI_ConfigDialog_Hide
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateSessionsList
- GuardLine.GuardLine.OnLayoutEditorFinished
- wbLeadHelper.wbLeadHelper.showNormalTitle
- LoyalPet.LPET.RenameProfileOnButtonUp
- wbLeadHelper.wbLeadHelper.chat
- PotionBar.PotionBarSettings.RefreshSingleRightLine
- GetStats.GetStats.OnChatLogUpdated
- DaemonAssist.DaemonAssist.PopulateBindingCombos
- Aura.AuraAddon.Slash
- wbLeadHelper.wbLeadHelperConfigTab.ChooseIconMessageStart
- EA_UiDebugTools.DevPadWindow.HideDeleteWindow
- Enemy.Enemy.MarksUI_MarkConfigDialog_IsOpen
- DaemonAssist.DaemonAssist.ShowWindow
- wbLeadHelper.wbLeadHelperMessagesTab.ListEnable
- WhoHealedMe.WHMGui.OnOptionsInitialize
- MapPin.MapPin.UI_ChooseIconDialog_Open
- Killer.Killer.ShowTopKillersTooltip
- PotionBar.PotionBarSettings.QuickActionsSelChanged


## Binding Resolution

- Total handler declarations: 609
- Resolved to Lua functions: 609 (100%)

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
