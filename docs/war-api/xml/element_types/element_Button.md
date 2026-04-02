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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, BankArkel, BuffHead, Busted |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:114`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:131`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:145`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:159`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:173`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:187`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:201`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:215` |
| Namespaces detected | Button |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAOptionsClose, AdvancedPetAssist: APAOptionsTabsAutoRecall, AdvancedPetAssist: APAOptionsTabsControls, AdvancedPetAssist: APAOptionsTabsFollowTarget, AdvancedPetAssist: APAOptionsTabsGeneral, AdvancedPetAssist: APAOptionsTabsHUD |
| XML usage count | 917 |
| XML attribute usage count | 917 |
| Lua usage count | 10 |
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

Observed XML element type instantiated by 51 addons.

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
- texturescale
- alpha

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md)
- [OnInitialize](../handlers/handler_OnInitialize.md)
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md)
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md)

## Common Handler Functions

- APAGui.OnTabButtonUp
- BankArkel.PackTab
- BankArkel.PackTabMover
- Enemy.CombatLogUI_StatsWindow_SortColumnClick
- TortallDPSDetail.GenericSort
- Enemy.CombatLogUI_StatsWindow_SortColumnRClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick
- TortallDPSDetail.ShowColumnMenu
- AdvancedRenownTraining.OnLButtonUpTab
- AuraConfig.OnConfigTabSelected
- TortallDPSMeter.ToggleDetail
- AggroMeter.OnTabLBU


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_GenericCheckButton.Initialize, IraConfig.HelpBtnInit | `function()` | MEDIUM |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | ClosetGoblinCharacterWindow.EquipmentLButtonDown, MiracleGrow2.onHClick, QuickWarReport.OnConfirmNoop, BuffHead.Setup.Layout.BeginResize, CMapWindow.OnResizeBeginLO, CMapWindow.OnResizeBeginLU | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | APAGui.OnTabButtonUp, BankArkel.PackTab, Enemy.CombatLogUI_StatsWindow_SortColumnClick, TortallDPSDetail.GenericSort, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick, AdvancedRenownTraining.OnLButtonUpTab | `flags, x, y` | MEDIUM |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | followTheLeader.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseDrag](../handlers/handler_OnMouseDrag.md) | input | EA_Window_Macro.IconMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag, Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | BankArkel.PackTabMover, MapMonster.PinTypeEditor.MouseOverDescription, Enemy.ConfigurationWindow_ShowTooltip, MiracleGrow2.ContextHover, MiracleGrow2.onHHover, WarBoard.OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | ClosetGoblinCharacterWindow.HideCloakOptions, ClosetGoblinCharacterWindow.HideShowHelm, WarBoard.OnMouseOverEnd, WarBoard.OnMouseOverEndBottom, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd, MiracleGrow.onHHoverEnd | `function(...)` | LOW |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | input | CMapWindow.MWheelWholeZoom | `function(delta)` | MEDIUM |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | ClosetGoblinCharacterWindow.EquipmentRButtonDown, CG_ItemRack.EquipmentRButtonDown, EA_Window_Macro.DetailIconRButtonDown, EA_Window_Macro.IconRButtonDown, EA_Window_Macro.SelectionIconRButtonDown | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.CombatLogUI_StatsWindow_SortColumnRClick, TortallDPSDetail.ShowColumnMenu, MiracleGrow2.onRClick, CMapWindow.OnScenarioQueueRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnRClick, MapMonster.OnMouseRightClickFilter | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonDown** handlers call: `WindowGetId`, `WindowGetParent`, `WindowSetShowing`, `WindowUtils.BeginResize`

**OnLButtonUp** handlers call: `BroadcastEvent`, `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonGetText`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ComboBoxAddMenuItem`, `ComboBoxClearMenuItems`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `CreateWindow`, `CreateWindowFromTemplate`, `Cursor.Clear`, `Cursor.IconOnCursor`, `DestroyWindow`, `DoesWindowExist`, `DynamicImageSetTexture`, `GameData.Player.GetRenownRefundCost`, `LabelGetText`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `ListBoxSetDisplayOrder`, `Player.GetMoney`, `Player.GetRenownRefundCost`, `RegisterEventHandler`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `SystemData.ActiveWindow.name:find`, `SystemData.ActiveWindow.name:match`, `TextEditBoxGetText`, `TextEditBoxGetText:len`, `TextEditBoxSelectAll`, `TextEditBoxSetText`, `WindowAddAnchor`, `WindowAssignFocus`, `WindowClearAnchors`, `WindowGetAnchorCount`, `WindowGetDimensions`, `WindowGetHandleInput`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowGetShowing`, `WindowSetHandleInput`, `WindowSetId`, `WindowSetShowing`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnMouseDrag** handlers call: `Cursor.IconOnCursor`, `Cursor.PickUp`

**OnMouseOver** handlers call: `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetShowing`

**OnMouseOverEnd** handlers call: `WindowSetShowing`

**OnRButtonUp** handlers call: `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetId`, `WindowGetShowing`, `WindowSetLayer`, `WindowSetShowing`

## Common Inherits

- EA_Button_DefaultResizeable
- EA_Button_DefaultWindowClose
- EA_Button_DefaultText
- EA_Button_DefaultCheckBox
- EA_Window_MacroIconButton
- MacroIconSelectionWindowIconButton
- EA_Button_Default
- EA_Button_Tab
- ClosetGoblinEquipmentButton
- DefaultButton
- EA_Button_DefaultListSort
- CG_ItemRackEquipmentButton

## Common Parent Elements

- [Window](element_Window.md) — 241× (HIGH)
- [Button](element_Button.md) — 2× (LOW)
- [LogDisplay](element_LogDisplay.md) — 1× (LOW)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md) — 29× (HIGH)
- [Label](element_Label.md) — 28× (HIGH)
- [Window](element_Window.md) — 10× (HIGH)
- [CircleImage](element_CircleImage.md) — 5× (MEDIUM)
- [FullResizeImage](element_FullResizeImage.md) — 5× (MEDIUM)
- [AnimatedImage](element_AnimatedImage.md) — 2× (LOW)
- [Button](element_Button.md) — 2× (LOW)

## Common Structural Child Elements

- [Normal](element_Normal.md) — 84× (HIGH)
- [Pressed](element_Pressed.md) — 77× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 65× (HIGH)
- [Disabled](element_Disabled.md) — 60× (HIGH)
- [PressedHighlit](element_PressedHighlit.md) — 53× (HIGH)
- [TexSlices](element_TexSlices.md) — 44× (HIGH)
- [TextOffset](element_TextOffset.md) — 30× (HIGH)
- [TextColors](element_TextColors.md) — 23× (HIGH)
- [ResizeImages](element_ResizeImages.md) — 21× (HIGH)
- [DisabledPressed](element_DisabledPressed.md) — 14× (HIGH)
- [OverlayOffset](element_OverlayOffset.md) — 14× (HIGH)
- [OverlaySize](element_OverlaySize.md) — 14× (HIGH)
- [OverlayTexCoords](element_OverlayTexCoords.md) — 14× (HIGH)
- [Sound](element_Sound.md) — 7× (MEDIUM)
- [Sounds](element_Sounds.md) — 7× (MEDIUM)
- [Text](element_Text.md) — 3× (MEDIUM)
- [AnimatedImages](element_AnimatedImages.md) — 1× (LOW)
- [Eventhandlers](element_Eventhandlers.md) — 1× (LOW)

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
- CoreWindowResizeButton
- CoreWindowResizeButtonBottomRight
- DefaultButton
- DefaultIconButton
- EA_Button_BottomTab
- EA_Button_Default
- EA_Button_DefaultBigLeftArrow
- EA_Button_DefaultBigRightArrow
- EA_Button_DefaultChatScrollToBottom
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultIconFrame
- EA_Button_DefaultIconFrame_Large
- EA_Button_DefaultIconFrame_Small
- EA_Button_DefaultListRow
- EA_Button_DefaultListSort
- EA_Button_DefaultMenuButton
- EA_Button_DefaultMinus
- EA_Button_DefaultPlus
- EA_Button_DefaultResizableComboBoxSelected
- EA_Button_DefaultResizeable
- EA_Button_DefaultSmallSquare
- EA_Button_DefaultText
- EA_Button_DefaultToggleCircle
- EA_Button_DefaultUp
- EA_Button_DefaultWindowClose
- EA_Button_ListSort
- EA_Button_ResizeIconFrame_NoNormalImage
- EA_Button_Tab
- EA_CheckButtonButton
- EA_FilterMenuButtonTemplate
- EA_Templates_Color_Picker_Button
- EA_Window_MacroDetailIconButton
- EA_Window_MacroIconButton
- EnemyChooseIconDialogList_IconButton
- IraConfigHelpBtn
- ItemWindowSlotButton
- LPETTabButton
- MacroIconSelectionWindowIconButton
- MapFilterContextMenuChoice
- MapMonsterEditorWindowButtonDefault
- MapMonsterPinTypeEditorWindowButtonDefault
- MapPinButton
- MapPinButton_NoBG
- MapPinChooseIconDialogList_IconButton
- MapPinNumberButtonTemplate
- MiracleGrow2Button
- MiracleGrow2HarvestRptButton
- MiracleGrow2Repeat
- MiracleGrow2SoundButton
- MiracleGrowButton
- MiracleGrowLightButton
- ModWindowSortButton
- PotionBarWindowButton
- QuickTacticSwitchWindowSortButton
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
- wbLeadHelperChooseIconDialogList_IconButton


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<Button name="..." backgroundtexture="shared_01" handleinput="false" highlighttexture="shared_01" font="font_default_text" textalign="left" overlaytexture="shared_01" overlayhighlighttexture="shared_01">
  <TextColors>
    <Normal r="255" g="255" b="255" a="255"/>
    <NormalHighlit r="250" g="213" b="63" a="255"/>
    <Pressed r="250" g="213" b="63" a="255"/>
    <PressedHighlit r="250" g="213" b="63" a="255"/>
    <Disabled r="92" g="92" b="92" a="255"/>
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
| `inherits` | optional | 74% | MacroIconSelectionWindowIconButton, BuffHeadLayoutCornerResizeButton, EA_Button_Default, EA_Button_DefaultResizeable, ... |
| `id` | optional | 21% | 10, 1, 18, 5, ... |
| `textalign` | optional | 13% | center, left, leftcenter, right, ... |
| `font` | optional | 11% | font_clear_small_bold, font_chat_text, font_default_text, font_clear_medium, ... |
| `layer` | optional | 7% | overlay, popup, background, default, ... |
| `handleinput` | optional | 6% | false, true |
| `backgroundtexture` | optional | 6% | shared_01, EA_HUD_01, chKtxt, EA_Cultivating01_d5, ... |
| `highlighttexture` | optional | 5% | shared_01, EA_HUD_01, EA_Cultivating01_d5, EA_Guild, ... |
| `drawchildrenfirst` | optional | 2% | true, false |
| `texturescale` | optional | 2% | 2.0, 0.37, 1.171, 0.75, ... |
| `alpha` | optional | 2% | 0, 1 |
| `textureScale` | optional | 1% | 0.625, 0.75, 1, 0.62, ... |
| `overlayhighlighttexture` | optional | 1% | EA_HUD_01, shared_01 |
| `overlaytexture` | optional | 1% | EA_HUD_01, shared_01 |
| `warnOnTextCropped` | optional | 1% | false |
| `sticky` | optional | 0% | false |
| `gameactionbutton` | optional | 0% | left, right |
| `savesettings` | optional | 0% | false, true |
| `autoresize` | optional | 0% | false |
| `draganddrop` | optional | 0% | true, false |
| `textAutoFitMinScale` | optional | 0% | 0.5 |
| `movable` | optional | 0% | false |
| `popable` | optional | 0% | false |
| `wordwrap` | optional | 0% | true |
| `Layer` | optional | 0% | popup |
| `scale` | optional | 0% | .5 |
| `text` | optional | 0% | Save |
| `w` | optional | 0% | false |
## Structural Sub-Elements

### [Normal](element_Normal.md)

Observed 84 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional |  |
| `x` | optional |  |
| `y` | optional |  |
| `b` | optional |  |
| `g` | optional |  |
| `r` | optional |  |
| `a` | optional |  |
| `def` | optional |  |
| `texture` | optional |  |
### [Pressed](element_Pressed.md)

Observed 77 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional |  |
| `x` | optional |  |
| `y` | optional |  |
| `b` | optional |  |
| `g` | optional |  |
| `r` | optional |  |
| `a` | optional |  |
| `def` | optional |  |
| `texture` | optional |  |
### [NormalHighlit](element_NormalHighlit.md)

Observed 65 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional |  |
| `y` | optional |  |
| `id` | optional |  |
| `b` | optional |  |
| `g` | optional |  |
| `r` | optional |  |
| `a` | optional |  |
| `def` | optional |  |
| `texture` | optional |  |
### [Disabled](element_Disabled.md)

Observed 60 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional |  |
| `y` | optional |  |
| `b` | optional |  |
| `g` | optional |  |
| `id` | optional |  |
| `r` | optional |  |
| `a` | optional |  |
| `def` | optional |  |
| `texture` | optional |  |
### [PressedHighlit](element_PressedHighlit.md)

Observed 53 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional |  |
| `b` | optional |  |
| `g` | optional |  |
| `r` | optional |  |
| `a` | optional |  |
| `def` | optional |  |
| `x` | optional |  |
| `y` | optional |  |
### [TexSlices](element_TexSlices.md)

Observed 44 times as an unnamed child.

### [TextOffset](element_TextOffset.md)

Observed 30 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [TextColors](element_TextColors.md)

Observed 23 times as an unnamed child.

### [ResizeImages](element_ResizeImages.md)

Observed 21 times as an unnamed child.

### [DisabledPressed](element_DisabledPressed.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional |  |
| `a` | optional |  |
| `b` | optional |  |
| `g` | optional |  |
| `r` | optional |  |
### [OverlayOffset](element_OverlayOffset.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [OverlaySize](element_OverlaySize.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [OverlayTexCoords](element_OverlayTexCoords.md)

Observed 14 times as an unnamed child.

### [Sound](element_Sound.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `event` | **required** |  |
| `script` | **required** |  |
### [Sounds](element_Sounds.md)

Observed 7 times as an unnamed child.

### [Text](element_Text.md)

Observed 3 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `text` | optional |  |
| `alignment` | optional |  |
### [AnimatedImages](element_AnimatedImages.md)

Observed 1 times as an unnamed child.

### [Eventhandlers](element_Eventhandlers.md)

Observed 1 times as an unnamed child.

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowSetShowing` | ui | 168 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `WindowGetId` | ui | 79 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `TextEditBoxGetText` | ui | 59 | OnLButtonUp |
| `ButtonGetDisabledFlag` | ui | 45 | OnLButtonUp |
| `ButtonSetPressedFlag` | ui | 45 | OnLButtonUp |
| `ComboBoxGetSelectedMenuItem` | ui | 35 | OnLButtonUp |
| `WindowGetShowing` | ui | 32 | OnLButtonUp, OnRButtonUp |
| `TextEditBoxSetText` | ui | 28 | OnLButtonUp |
| `WindowGetParent` | ui | 24 | OnLButtonDown, OnLButtonUp, OnMouseOver |
| `LabelSetTextColor` | ui | 22 | OnLButtonUp, OnMouseOver |
| `SystemData.ActiveWindow.name:match` | data | 22 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `LabelSetText` | ui | 18 | OnLButtonUp, OnMouseOver |
| `WindowAssignFocus` | ui | 18 | OnLButtonUp |
| `DoesWindowExist` | ui | 16 | OnLButtonUp |
| `ButtonGetPressedFlag` | ui | 15 | OnLButtonUp |
| `ButtonSetDisabledFlag` | ui | 15 | OnLButtonUp |
| `DestroyWindow` | ui | 10 | OnLButtonUp |
| `WindowAddAnchor` | ui | 8 | OnLButtonUp |
| `ComboBoxSetSelectedMenuItem` | ui | 7 | OnLButtonUp |
| `WindowSetLayer` | ui | 7 | OnRButtonUp |
| `ComboBoxGetSelectedText` | ui | 6 | OnLButtonUp |
| `LabelGetText` | ui | 6 | OnLButtonUp |
| `WindowClearAnchors` | ui | 6 | OnLButtonUp |
| `ListBoxGetDataIndex` | ui | 5 | OnLButtonUp, OnMouseOver |
| `SliderBarSetCurrentPosition` | ui | 5 | OnLButtonUp |
| `SliderBarGetCurrentPosition` | ui | 4 | OnLButtonUp |
| `SystemData.ActiveWindow.name:find` | data | 4 | OnLButtonUp |
| `CreateWindow` | ui | 3 | OnLButtonUp |
| `SystemData.MouseOverWindow.name:match` | data | 3 | OnMouseOver, OnRButtonUp |
| `WindowGetDimensions` | ui | 3 | OnLButtonUp, OnMouseOver |
| `WindowGetScreenPosition` | ui | 3 | OnLButtonUp, OnMouseOver |
| `WindowUtils.BeginResize` | ui | 3 | OnLButtonDown |
| `ComboBoxAddMenuItem` | ui | 2 | OnLButtonUp |
| `ComboBoxClearMenuItems` | ui | 2 | OnLButtonUp |
| `CreateWindowFromTemplate` | ui | 2 | OnLButtonUp |
| `Cursor.Clear` | ui | 2 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 2 | OnLButtonUp, OnMouseDrag |
| `DynamicImageSetTexture` | ui | 2 | OnLButtonUp |
| `TextEditBoxGetText:len` | ui | 2 | OnLButtonUp |
| `WindowSetId` | ui | 2 | OnLButtonUp |
| `BroadcastEvent` | event | 1 | OnLButtonUp |
| `ButtonGetText` | ui | 1 | OnLButtonUp |
| `Cursor.PickUp` | data | 1 | OnMouseDrag |
| `GameData.Player.GetRenownRefundCost` | data | 1 | OnLButtonUp |
| `ListBoxSetDisplayOrder` | ui | 1 | OnLButtonUp |
| `Player.GetMoney` | data | 1 | OnLButtonUp |
| `Player.GetRenownRefundCost` | data | 1 | OnLButtonUp |
| `RegisterEventHandler` | event | 1 | OnLButtonUp |
| `TextEditBoxSelectAll` | ui | 1 | OnLButtonUp |
| `WindowGetAnchorCount` | ui | 1 | OnLButtonUp |
| `WindowGetHandleInput` | ui | 1 | OnLButtonUp |
| `WindowSetHandleInput` | ui | 1 | OnLButtonUp |
| `WindowSetTintColor` | ui | 1 | OnLButtonUp |
| `WindowStartAlphaAnimation` | ui | 1 | OnLButtonUp |
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

### OnMouseWheel

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `x` | number | mouse_x |
| 1 | `y` | number | mouse_y |
| 2 | `delta` | number | wheel_delta |
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

- DaemonAssist.DaemonAssist.UpdateWindow
- wbLeadHelper.WbLeadHelperMessage.MessageDialogOpen
- BankArkel.BankArkel.CreatePackWin
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- PotionBar.PotionBar.local.UpdateButton
- Killer.Killer.Initialize
- RoR_SoR.RoR_SoR.Restack
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloak
- MapPin.EditMarker
- MapMonster.MapMonster.PinTypeEditor.Initialize
- DAoCBuff.DAoCBuff.ShowMessageWindow
- Enemy.Enemy.ScenarioInfoInitialize
- MapMonster.FilterButtonState
- AdvancedPetAssist.APAGui.OnShown
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.Enemy.IntercomUI_IntercomDialog_Open
- MapMonster.MapMonster.local.CreateFilterMenuEntry
- CM_ClosetGoblin.ClosetGoblinZoneWindow.OnInitialize
- MapMonster.MapMonster.local.OpenTypeViewerWindow
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- QuickWarReport.PrepareConfirmWindowChrome
- MapMonster.MapMonster.local.FilterButtonState
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSetRow
- DAoCBuff.DAoCBuffSettings.SetLabels
- LibWBToggler.LibWBTogglerManager.Initialize
- Miracle Grow Remix.MiracleGrow2.InitConfig
- Enemy.Enemy.UI_ConfigDialog_OnSectionSelChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateActionBarSettings
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloakHeraldry
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnShow
- Enemy.Enemy.CombatLogUI_StatsWindow_Open
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- RandomMount.RandomMountUI.OnAddCustomMount
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled3
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled5
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- MapMonster.CleanEditorWindow
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowCloakOptions
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideCloakOptions
- TidyRoll.TidyRollOptions.Initialize
- Enemy.Enemy.IntercomInitialize
- PotionBar.PotionBarSettings.OnAboutShown
- RandomMount.RandomMountUI.OnInitialize
- WSCT.WSCT.ColorOnInitialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HideShowHelm
- GetStats.GetStats.OnInitialize
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Open
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowCloakOnly
- MapMonster.MapMonster.InitializeMapPins
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled4
- EA_UiDebugTools.BustedGUI.UpdateNextPrevButtonStatus
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- JunkDump.JunkDump.Initialize
- Enemy.Enemy.MarksUI_MarkConfigDialog_Open
- QuickWarReport.QuickWarReport.local.PrepareConfirmWindowChrome
- Enemy.Enemy.UI_ConfigDialog_Open
- Enemy.Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelm
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowShowHelmOnly
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- MapPin.MapPin.local.EditMarker
- nRarity.nRarity.Border:new
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled2
- AggroMeter.AggroMeter.OnTabLBU
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UseItemRackToggled
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Open
- AggroMeter.AggroMeter.Initialize
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.ShowHelm
- MapPin.MapPin.RButtonUp
- RandomMount.RandomMountUI.OnDropSlotLButtonUp
- PotionBar.UpdateButton
- MapMonster.MapMonster.ShutdownPins
- MapMonster.CreateFilterMenuEntry
- MapMonster.OpenTypeViewerWindow
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Open
- Busted.BustedGUI.UpdateNextPrevButtonStatus
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.HotbarChangeToggled1


## Binding Resolution

- Total handler declarations: 741
- Resolved to Lua functions: 734 (99%)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- FastInteract
- GetStats
- GuardList
- GuardRange
- JunkDump
- Killer
- LibGroup
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
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
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- followTheLeader
- nRarity
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAOptionsClose -> Button APAOptionsClose
- AdvancedPetAssist: APAOptionsTabsAutoRecall -> Button APAOptionsTabsAutoRecall
- AdvancedPetAssist: APAOptionsTabsControls -> Button APAOptionsTabsControls
- AdvancedPetAssist: APAOptionsTabsFollowTarget -> Button APAOptionsTabsFollowTarget
- AdvancedPetAssist: APAOptionsTabsGeneral -> Button APAOptionsTabsGeneral
- AdvancedPetAssist: APAOptionsTabsHUD -> Button APAOptionsTabsHUD

## Related APIs

- [AlertTextWindow.AddLine](../../globals/functions/global_AlertTextWindow.AddLine.md) (HIGH 100/100) - Global Function
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DebugWindow.ClearTextLog](../../globals/functions/global_DebugWindow.ClearTextLog.md) (HIGH 100/100) - Global Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [InterfaceCore.ReloadUI](../../globals/functions/global_InterfaceCore.ReloadUI.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Used With

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseDrag](../../events/window_events/window_event_OnMouseDrag.md) (HIGH 100/100) - Window Event
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 100/100) - XML Handler
- [OnRButtonDown](../../events/window_events/window_event_OnRButtonDown.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Triggered By

- none

## Affects

- none
