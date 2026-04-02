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
| `inherits` | optional | 60% | MapPinCallTemplateWindow, Aura_LabelCheckButton, EA_Window_Default, EA_Window_DefaultButtonBottomFrame, ... |
| `layer` | optional | 31% | default, secondary, background, overlay, ... |
| `handleinput` | optional | 22% | true, false |
| `movable` | optional | 20% | false, true |
| `savesettings` | optional | 13% | false, true |
| `popable` | optional | 6% | false, true |
| `sticky` | optional | 5% | false, true |
| `scale` | optional | 2% | 1.0, 0.7 |
| `id` | optional | 2% | 4, 2, 9, 1, ... |
| `skipinput` | optional | 1% | true, false |
| `alpha` | optional | 1% | 0.6, 0.97, 0.0, 0.85, ... |
| `font` | optional | 1% | font_chat_text, font_clear_small_bold |
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

| API Function | Call Count | From Events |
| --- | --- | --- |
| `TextEditBoxSetText` | 414 | OnHidden, OnLButtonUp, OnRButtonUp, OnShown |
| `LabelSetText` | 354 | OnInitialize, OnShown |
| `ButtonSetText` | 167 | OnHidden, OnInitialize, OnShown |
| `WindowSetShowing` | 134 | OnHidden, OnInitialize, OnKeyEscape, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp, OnShown |
| `WindowSetDimensions` | 73 | OnLButtonDown, OnShown |
| `WindowGetId` | 46 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonDown, OnRButtonUp |
| `WindowSetTintColor` | 39 | OnInitialize, OnLButtonDown, OnMouseOver, OnMouseOverEnd, OnRButtonDown |
| `RegisterEventHandler` | 9 | OnInitialize |
| `WindowSetAlpha` | 9 | OnInitialize, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `WindowGetShowing` | 8 | OnKeyEscape, OnLButtonUp, OnSetMoving, OnShown |
| `BroadcastEvent` | 5 | OnHidden, OnLButtonDown, OnShown |
| `LabelSetTextColor` | 5 | OnInitialize, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `CreateWindow` | 3 | OnInitialize |
| `DestroyWindow` | 3 | OnKeyEscape |
| `SystemData.MouseOverWindow.name:find` | 2 | OnLButtonDown, OnRButtonDown |
| `SystemData.MouseOverWindow.name:sub` | 2 | OnLButtonDown, OnRButtonDown |
| `WindowGetAlpha` | 2 | OnMouseOver |
| `WindowRegisterEventHandler` | 2 | OnInitialize |
| `LabelGetText` | 1 | OnLButtonDown |
| `TextEditBoxGetText` | 1 | OnKeyEscape |
| `UnregisterEventHandler` | 1 | OnHidden |
| `WindowGetDimensions` | 1 | OnMouseOver |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

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
### OnRButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnLButtonUp

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
### OnRButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnRButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnRButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnRButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

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
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

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
### OnRButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnUpdate

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

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
### OnRButtonUp

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
### OnLButtonUp

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
### OnRButtonUp

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
### OnRButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnRButtonUp

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
### OnRButtonUp

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
### OnRButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonDown

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
### OnRButtonDown

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonDown

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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnUpdate

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
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
### OnLButtonUp

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
### OnLButtonUp

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
### OnLButtonUp

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
### OnUpdate

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
### OnLButtonDown

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
### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

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
### OnRButtonDown

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
### OnLButtonUp

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
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- Enemy.Enemy.UI_ConfigDialog_Hide
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelm
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Hide
- LoyalPet.LPET.OpenMenu
- QuickWarReport.QuickWarReport.local.ShowConfirmWindow
- DAoCBuff.DAoCBuffSettings.CloseOptionswindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.Hide
- Enemy.Enemy.MarksUI_MarkConfigDialog_Hide
- Enemy.Enemy.MarksUI_MarkConfigDialog_Open
- Enemy.Enemy.CombatLogUI_EpsWindow_Initialize
- WSCT.WSCT.HideMenu
- DaemonAssist.DaemonAssist.ToggleWindow
- BankArkel.BankArkel.CreatePackWin
- QuickWarReport.QuickWarReport.local.EnsureConfirmWindow
- AdvancedPetAssist.APAGui.HidePetTargetHUD
- Swift Assist.SwiftAssist.Initialize
- Swift Assist.SetTexLabel
- Swift Assist.SwiftAssist.OnMacroUpdated
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Hide
- Enemy.Enemy.MarksInitialize
- PotionBar.PotionBar.Hide
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Hide
- DAoCBuff.DAoCTooltips.Init
- wbLeadHelper.wbLeadHelper.drawWindows
- RoR_SoR.RoR_SoR.OnCombat
- Enemy.SetStatsRow
- MapPin.MapPin.ShowIcons
- BankArkel.BankArkel.PackImg
- RoR_SoR.RoR_SoR.SET_KEEP
- Swift Assist.SetSmartLabel
- TidyRoll.TidyRollOptions.Initialize
- BuffHead.RegisterLayoutEditor
- Busted.BustedGUI.ClearAlertFlash
- Enemy.Enemy.UI_ConfigDialog_OnSectionSelChanged
- WSCT.WSCT.OnLButtonUpColorPicker
- RoR_SoR.RoR_SoR.SET_CAMPAIGN
- Enemy.EnemyMarkTemplate:ToggleMark
- wbLeadHelper.wbLeadHelperConfigTab.Hide
- Effigy.EffigyBar:setup
- GuardLine.GuardLine.Libguard_Toggle
- QuickWarReport.ShowConfirmWindow
- AggroMeter.AggroMeter.Initialize
- Moth.Moth.UpdateHealthBar
- AdvancedRenownTrainer.AdvancedRenownTraining.OnShown
- DAoCBuff.DAoCBuffSettings.PopulateSettings
- Effigy.EffigyBar.Init
- Enemy.Enemy.UI_TextEntryDialog_Open
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.Show
- Enemy.Enemy.IntercomUI_IntercomDialog_Hide
- MiracleGrow.MiracleGrow.Initialize
- LibSlash.LibSlash.Initialize
- GuardLine.GuardLine.OnLayoutEditorFinished
- MapPin.MapPin.SetupAccept
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Open
- QuickWarReport.QuickWarReport.TestConfirmWindow
- JunkDump.JunkDumpOptions.CreateOptionsWindow
- Enemy.Enemy.GroupsInitialize
- RoR_SoR.RoR_SoR.CloseSetOffsetWindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakOnly
- WSCT.WSCT:OpenMenu
- GuardLine.GuardLine.GetIDs
- Enemy.Enemy._Initialize
- PotionBar.PotionBar.Shutdown
- BankArkel.BankArkel.BackPackShow
- MapPin.MapPin.UI_ChooseIconDialog_Hide
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Initialize
- PotionBar.PotionBarFloating.ShowSettings
- WhoHealedMe.WHMGui.HideOptionsWindow
- Enemy.Enemy.TimerInitialize
- BankArkel.BankArkel.BackPackHide
- BuffHead.BuffHead.local.RegisterLayoutEditor
- Enemy.Enemy.UI_ChooseIconDialog_Hide
- wbLeadHelper.WbLeadHelperMessage.MessageDialogOpen
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelmOnly
- Enemy.Enemy.StopwatchEnabledChanged
- PotionBar.PotionBarFloating.Alpha
- PotionBar.PotionBarSettings.OnHidden
- EA_UiDebugTools.BustedGUI.ToggleMainWindow
- Enemy.Enemy.CombatLogUI_StatsWindow_Open
- Enemy.Enemy.AssistUI_Target_Show
- GuardLine.GuardLine.init
- DAoCBuff.DAoCBuff.ShowMessageWindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSlotIcons
- Enemy.EnemyGroupIcon:Attach
- TidyChat.TidyChat.LootRoll.OnRollLinkLButtonUp
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_Open
- Enemy.Enemy.IntercomUI_IntercomDialog_Open
- JunkDump.JunkDumpOptions.DynamicBagOptions
- MapPin.EditMarker
- BankArkel.BankArkel.PackShow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideCloakOptions
- wbLeadHelper.wbLeadHelperConfigTab.Show
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Open
- Enemy.Enemy.local.SetStatsRow
- AdvancedPetAssist.APAGui.ToggleInstantOnlyHUD
- GuardLine.GuardLine.update2
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Aura.AuraTexture.OnClose
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide
- Moth.Moth.Initialize
- AdvancedRenownTrainer.AdvancedRenownTraining.ChangeTab
- FastInteract.FastInteract.OptionsShow
- DAoCBuff.DAoCTooltips.UpdateCondenseTooltip
- Miracle Grow Remix.MiracleGrow2.ConfigTabChange
- WSCT.WSCT.ColorOnInitialize
- RandomMount.RandomMountUI.Show
- Aura.AuraAddon.OnLoad
- AdvancedPetAssist.APAGui.OnShown
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOpacity
- AutoMark.CreateMarker
- EA_UiDebugTools.BustedGUI.ClearAlertFlash
- WhoHealedMe.WHMGui.ShowOptionsWindow
- CM_ClosetGoblin.ClosetGoblinZoneWindow.OnInitialize
- RandomMount.RandomMountUI.OnInitialize
- PotionBar.PotionBar.local.UpdateButton
- PotionBar.PotionBarSettings.OnAboutClose
- GuardLine.GuardLine.update
- Enemy.EnemyUnitFrame:ApplySettings
- Enemy.Enemy.CombatLogUI_EpsWindow_UpdateLayout
- MapPin.MapPin.UI_ChooseIconDialog_Open
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- Effigy.Effigy.RegisterStateInfoForCastbar
- LoyalPet.LPET.HideMenu
- PotionBar.UpdateButton
- AdvancedPetAssist.ApplyHUDVisibilityFromSettings
- CM_ClosetGoblin.ClosetGoblinZoneWindow.Hide
- Killer.Killer.ShowPersonalStatsTooltip
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideShowHelm
- EA_UiDebugTools.ObjectInspector.CloseWindow
- WSCT.WSCT.ColorOnButtonUp
- MapPin.MapPin.RButtonUp
- AdvancedRenownTrainer.AdvancedRenownTrainer.local.CreateAbilityWindow
- Enemy.Enemy.CombatLogUI_IDS_OnSettingsChanged
- wbLeadHelper.wbLeadHelperConfigWindow.Hide
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.Enemy.CombatLogUI_IDS_Initialize
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide
- AdvancedRenownTrainer.AdvancedRenownTrainer.local.CreateTab
- BankArkel.BankArkel.Init
- Killer.Killer.Initialize
- QuickWarReport.QuickWarReport.Initialize
- EA_UiDebugTools.BustedGUI.Initialize
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Open
- WhoHealedMe.WHMCore.RunSettingEffects
- QuickWarReport.EnsureConfirmWindow
- PotionBar.PotionBarFloating.ReflowButtons
- Busted.BustedGUI.Initialize
- DaemonAssist.DaemonAssist.ShowWindow
- RoR_SoR.RoR_SoR.OnInitialize
- MapPin.MapPin.local.EditMarker
- Moth.Moth.HideBorders
- Enemy.Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- Effigy.Effigy.RegisterStateInfoForPlayer
- AggroMeter.AggroMeter.Close
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged
- JunkDump.JunkDumpOptions.Done
- Aura.AuraAddon.Slash
- wbLeadHelper.wbLeadHelperMessagesTab.Hide
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateList
- AdvancedRenownTrainer.CreateAbilityWindow
- AdvancedPetAssist.AdvancedPetAssist.local.ApplyHUDVisibilityFromSettings
- Swift Assist.Swift Assist.local.SetTexLabel
- Aura.Aura:CreateRuntimeWindows
- Enemy.Enemy.CombatLogUI_EpsWindow_Open
- EA_UiDebugTools.ObjectInspector.WindowInit
- Enemy.Enemy.UI_ConfigDialog_Open
- PotionBar.PotionBarSettings.OnResetButton
- AdvancedRenownTrainer.CreateTab
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Hide
- AggroMeter.AggroMeter.SplitText
- RoR_SoR.RoR_SoR.ShowPopper
- RoR_SoR.RoR_SoR.SetWindowShow
- GetStats.GetStats.CloseWindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloakOptions
- AdvancedRenownTrainer.AdvancedRenownTraining.OnHidden
- Enemy.Enemy.CombatLogUI_StatsWindow_Hide
- Killer.Killer.ShowRowTooltip
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Open
- PotionBar.PotionBar.LibSlashHandler
- CM_ClosetGoblin.ClosetGoblinZoneWindow.UpdateHighlightOnRow
- DAoCBuff.DAoCBuffSettings.CreateOptionswindow
- PotionBar.PotionBarSettings.OnShown
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Hide
- TurretRange.TurretRange.OnLoadComplete
- MoraleCircle.MoraleCircle.init
- MapPin.OnHyperLinkLButtonUp2
- GetStats.GetStats.OnInitialize
- wbLeadHelper.WbLeadHelperMessage.MessageDialogHide
- RoR_SoR.RoR_SoR.CloseSetOpacityWindow
- GuardLine.GuardLine.OnLayoutEditorStart
- wbLeadHelper.wbLeadHelper.createWbLeadHelperWindow
- DAoCBuff.DAoCBuffFrame.MouseOverUpdate
- Swift Assist.Swift Assist.local.SetSmartLabel
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Hide
- WarBoard.WarBoard.SlashCommand
- PotionBar.PotionBarSettings.OnCancelButton
- DaemonAssist.DaemonAssist.HideWindow
- Moth.Moth.Clear
- RoR_SoR.RoR_SoR.CloseSetScaleWindow
- WarBoard.WarBoard.LoadGeneralSettings
- Busted.Busted.Initialize
- WSCT.WSCT.ColorHideMenu
- AdvancedPetAssist.APAGui.ToggleKitingHUD
- RoR_SoR.RoR_SoR.OnScenario
- AdvancedRenownTrainer.AdvancedRenownTraining.Initialize
- DAoCBuff.DAoCBuffSettings.OpenOptionswindow
- RoR_SoR.RoR_SoR.SET_CITY
- AutoMark.AutoMark.local.CreateMarker
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnInitialize
- Enemy.Enemy.IntercomInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- Miracle Grow Remix.MiracleGrow2.InitConfig
- DAoCBuff.DAoCBuff.CloseMessageWindow
- Effigy.Effigy.RegisterStateInfoForTargets
- Moth.Moth.UpdateTarget
- Moth.Moth.HealthBar
- MiracleGrowLight.MiracleGrowLight.switchBackground
- AdvancedPetAssist.APAGui.Show
- WarBoard.WarBoard.Options.EnableBoard
- DAoCBuff.DAoCBuffHeadTracker:Create
- RandomMount.RandomMountUI.Hide
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Open
- AdvancedPetAssist.APAGui.Hide
- PotionBar.PotionBarSettings.OnAboutButton
- PotionBar.PotionBar.Show
- GetStats.GetStats.OnChatLogUpdated
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- EA_UiDebugTools.ObjectInspector.ShowWindow
- RoR_SoR.RoR_SoR.Text_Stream_Fetch
- Enemy.Enemy.MarksUI_EnemyMarksWindow_Open
- MapPin.MapPin.test
- BankArkel.BankArkel.SetCharInfo
- QuickWarReport.HideConfirmWindow
- Enemy.Enemy.CombatLogUI_EpsWindow_Hide
- DAoCBuff.DAoCBuffSettings.UC
- JunkDump.JunkDumpOptions.DestroyOptionsWindow
- BankArkel.BankArkel.PackHide
- WSCT.WSCT.ColorAcceptButtonOnButtonUp
- Enemy.Enemy.CombatLogUI_IDS_Update
- Enemy.Enemy.AssistInitialize
- MapPin.MapPin.SetupCancel
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- CM_ClosetGoblin.ClosetGoblinZoneWindow.Show
- GuardLine.GuardLine.OffTarget
- Aura.Aura:CreateEditWindows
- wbLeadHelper.wbLeadHelperMessagesTab.Show
- wbLeadHelper.wbLeadHelperConfigWindow.Show
- Enemy.Enemy.UI_ChooseIconDialog_Open
- Killer.Killer.ShowTopKillersTooltip
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOffset
- PotionBar.PotionBarSettings.QuickActionsSelChanged
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_Hide
- WSCT.WSCT.UpdateAnimationOptions
- JunkDump.JunkDumpOptions.Show
- Shinies.Searches_UpdateWindowDisplay
- MapPin.MapPin.local.OnHyperLinkLButtonUp2
- QuickWarReport.QuickWarReport.Shutdown
- QuickWarReport.QuickWarReport.local.HideConfirmWindow
- AdvancedPetAssist.APAGui.UpdatePetTargetHUD
- FastInteract.FastInteract.OptionsClose
- RoR_SoR.RoR_SoR.OnWindowOptionsSetScale
- WarBoard.WarBoard.OnOptionsButton
- PotionBar.PotionBar.Initialize

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
