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

- name
- inherits
- id
- textalign
- font
- layer
- handleinput
- backgroundtexture
- highlighttexture
- drawchildrenfirst
- warnOnTextCropped
- texturescale

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md)
- [OnInitialize](../handlers/handler_OnInitialize.md)
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md)

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
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | APAGui.OnTabButtonUp, BankArkel.PackTab, Enemy.CombatLogUI_StatsWindow_SortColumnClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick, AdvancedRenownTraining.OnLButtonUpTab, AuraConfig.OnConfigTabSelected | `flags, x, y` | MEDIUM |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | followTheLeader.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseDrag](../handlers/handler_OnMouseDrag.md) | input | EA_Window_Macro.IconMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag, Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | BankArkel.PackTabMover, Enemy.ConfigurationWindow_ShowTooltip, WarBoard.OnMouseOver, WarBoard.OnMouseOverBottom, ActionBars.OnMouseoverHotbarPageDown, ActionBars.OnMouseoverHotbarPageUp | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | ClosetGoblinCharacterWindow.HideCloakOptions, ClosetGoblinCharacterWindow.HideShowHelm, WarBoard.OnMouseOverEnd, WarBoard.OnMouseOverEndBottom, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd, MiracleGrowLight.harvestEnd | `function(...)` | LOW |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | ClosetGoblinCharacterWindow.EquipmentRButtonDown, CG_ItemRack.EquipmentRButtonDown, EA_Window_Macro.DetailIconRButtonDown, EA_Window_Macro.IconRButtonDown, EA_Window_Macro.SelectionIconRButtonDown | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.CombatLogUI_StatsWindow_SortColumnRClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnRClick, MiracleGrowLight.switchMode, ClosetGoblinCharacterWindow.EquipmentRButtonUp, PP.ItemSlotRMouse, PotionBarFloating.RButtonUp | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonDown** handlers call: `WindowGetId`

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonGetText`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxClearMenuItems`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `CreateWindow`, `DestroyWindow`, `DoesWindowExist`, `GameData.Player.GetRenownRefundCost`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `ListBoxSetDisplayOrder`, `RegisterEventHandler`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `SystemData.ActiveWindow.name:find`, `SystemData.ActiveWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxGetText:len`, `TextEditBoxSelectAll`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowAssignFocus`, `WindowClearAnchors`, `WindowGetAnchorCount`, `WindowGetDimensions`, `WindowGetHandleInput`, `WindowGetId`, `WindowGetParent`, `WindowGetScale`, `WindowGetScreenPosition`, `WindowGetShowing`, `WindowSetHandleInput`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnMouseDrag** handlers call: `Cursor.IconOnCursor`, `Cursor.PickUp`

**OnMouseOver** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetShowing`

**OnMouseOverEnd** handlers call: `WindowSetShowing`

**OnRButtonUp** handlers call: `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetId`, `WindowSetLayer`

## Common Inherits

- EA_Button_DefaultResizeable
- EA_Button_DefaultWindowClose
- EA_Button_DefaultCheckBox
- EA_Window_MacroIconButton
- EA_Button_DefaultText
- MacroIconSelectionWindowIconButton
- ClosetGoblinEquipmentButton
- EA_Button_Tab
- EA_Button_Default
- EA_Button_DefaultListSort
- DefaultButton
- CG_ItemRackEquipmentButton

## Common Parent Elements

- [Windows](element_Windows.md) — 664× (HIGH)
- [Window](element_Window.md) — 8× (MEDIUM)

## Common Structural Child Elements

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
### [AnimatedImages](element_AnimatedImages.md)

Observed 1 times as an unnamed child.

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

- ClosetGoblinCharacterWindow.OnInitialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- DAoCBuff.ShowMessageWindow
- Enemy.IntercomUI_ChooseChannelDialog_Open
- APAGui.OnShown
- AggroMeter.Initialize
- Enemy.UI_ConfigDialog_Open
- Enemy.ScenarioInfoInitialize
- ClosetGoblinCharacterWindow.UseItemRackToggled
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- ClosetGoblinCharacterWindow.HideCloakOptions
- DAoCBuffSettings.SetLabels
- PotionBar.UpdateButton
- ClosetGoblinCharacterWindow.UpdateActionBarSettings
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- WSCT.ColorOnInitialize
- BankArkel.CreatePackWin
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- AggroMeter.OnTabLBU
- ClosetGoblinCharacterWindow.ShowCloakOptions
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- ClosetGoblinCharacterWindow.ShowShowHelm
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- Killer.Initialize
- PotionBar.local.UpdateButton
- RoR_SoR.Restack
- ClosetGoblinZoneWindow.OnInitialize
- Enemy.MarksUI_MarkConfigDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- ClosetGoblinCharacterWindow.UpdateSetRow
- ClosetGoblinCharacterWindow.OnShow
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- Enemy.IntercomUI_IntercomDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.CombatLogUI_StatsWindow_Open
- LibWBTogglerManager.Initialize
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- ClosetGoblinCharacterWindow.HideShowHelm
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.IntercomInitialize
- TidyRollOptions.Initialize
- ClosetGoblinCharacterWindow.ShowCloak
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- ClosetGoblinCharacterWindow.ShowHelm
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- PotionBarSettings.OnAboutShown


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
- [AnimatedImages](element_AnimatedImages.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
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
