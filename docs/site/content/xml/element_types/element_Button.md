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

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnInitialize](../handlers/handler_OnInitialize.md) | EA_GenericCheckButton.Initialize, IraConfig.HelpBtnInit | `function()` | MEDIUM |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | ClosetGoblinCharacterWindow.EquipmentLButtonDown, MiracleGrow2.onHClick, QuickWarReport.OnConfirmNoop, BuffHead.Setup.Layout.BeginResize, CMapWindow.OnResizeBeginLO, CMapWindow.OnResizeBeginLU | `function(...)` | LOW |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | APAGui.OnTabButtonUp, BankArkel.PackTab, Enemy.CombatLogUI_StatsWindow_SortColumnClick, TortallDPSDetail.GenericSort, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnClick, AdvancedRenownTraining.OnLButtonUpTab | `function(...)` | LOW |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | followTheLeader.OnMButtonUp | `function(...)` | LOW |
| [OnMouseDrag](../handlers/handler_OnMouseDrag.md) | EA_Window_Macro.IconMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag, Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag, Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | BankArkel.PackTabMover, MapMonster.PinTypeEditor.MouseOverDescription, Enemy.ConfigurationWindow_ShowTooltip, MiracleGrow2.ContextHover, MiracleGrow2.onHHover, WarBoard.OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | ClosetGoblinCharacterWindow.HideCloakOptions, ClosetGoblinCharacterWindow.HideShowHelm, WarBoard.OnMouseOverEnd, WarBoard.OnMouseOverEndBottom, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd, MiracleGrow.onHHoverEnd | `function(...)` | LOW |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | CMapWindow.MWheelWholeZoom | `function(delta)` | MEDIUM |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | ClosetGoblinCharacterWindow.EquipmentRButtonDown, CG_ItemRack.EquipmentRButtonDown, EA_Window_Macro.DetailIconRButtonDown, EA_Window_Macro.IconRButtonDown, EA_Window_Macro.SelectionIconRButtonDown | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | Enemy.CombatLogUI_StatsWindow_SortColumnRClick, TortallDPSDetail.ShowColumnMenu, MiracleGrow2.onRClick, CMapWindow.OnScenarioQueueRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnRClick, MapMonster.OnMouseRightClickFilter | `function(...)` | LOW |

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

- [Window](element_Window.md)
- [Button](element_Button.md)
- [LogDisplay](element_LogDisplay.md)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md)
- [Label](element_Label.md)
- [Window](element_Window.md)
- [CircleImage](element_CircleImage.md)
- [FullResizeImage](element_FullResizeImage.md)
- [AnimatedImage](element_AnimatedImage.md)
- [Button](element_Button.md)


## Structural Sub-Elements

### [Normal](element_Normal.md)

- Observed in 84 parent frames
- Attributes: `a`, `b`, `def`, `g`, `id`, `r`, `texture`, `x`, `y`
  - `a`: `255`
  - `b`: `255`, `45`, `73`, `102`
  - `def`: `EA_HorizontalResizeImage_DefaultComboBox`, `EA_AnimatedImage_DefaultChoiceOverlay`, `BuffHeadLayoutHorizontalButtonNormal`, `BuffHeadLayoutVerticalButtonNormal`, `EA_HorizontalResizeImage_DefaultComboBox2`, `EA_HorizontalResizeImage_ResizeableButtonNormal`, `EA_Button_ListSortNormal`
  - `g`: `255`, `216`, `175`, `204`
  - `id`: `morale-yellow`, `RightTabFrame`, `LayoutCorner-BottomLeft`, `LayoutCorner-BottomRight`, `LayoutCorner-TopLeft`, `LayoutCorner-TopRight`, `Map-Filters-Button`, `Map-Scenario-Button-Blue`
  - `r`: `255`, `226`
  - `texture`: `bpKtxt`, `EA_SquareFrame`, `PinBG`, `ShiniesIconBorderNormal`, `TidyRoll_SquareFrame`, `nRarityBorderNormal`
  - `x`: `0`, `92`, `172`, `494`
  - `y`: `28`, `44`, `0`, `341`, `434`, `420`

### [Pressed](element_Pressed.md)

- Observed in 77 parent frames
- Attributes: `a`, `b`, `def`, `g`, `id`, `r`, `texture`, `x`, `y`
  - `a`: `255`
  - `b`: `63`, `0`, `45`, `36`
  - `def`: `EA_HorizontalResizeImage_DefaultComboBox`, `BuffHeadLayoutHorizontalButtonPressed`, `BuffHeadLayoutVerticalButtonPressed`, `EA_HorizontalResizeImage_DefaultComboBox2`, `EA_HorizontalResizeImage_ResizeableButtonPressed`, `EA_Button_ListSortPressed`
  - `g`: `213`, `255`, `216`, `57`, `85`
  - `id`: `morale-white`, `RightTabFrame-Rollover`, `LayoutCorner-BottomLeft-ROLLOVER`, `LayoutCorner-BottomRight-ROLLOVER`, `LayoutCorner-TopLeft-ROLLOVER`, `LayoutCorner-TopRight-ROLLOVER`, `Map-Filters-Button-Depressed`, `Map-Scenario-Button-Blue-Queue-Frame-1`
  - `r`: `250`, `255`, `226`, `95`
  - `texture`: `EA_SquareFrame_Pressed`, `PinBG`, `ShiniesIconBorderHighlight`, `TidyRoll_SquareFrame`, `nRarityBorderHighlight`
  - `x`: `0`, `120`, `172`, `494`
  - `y`: `56`, `44`, `0`, `370`, `434`, `420`

### [NormalHighlit](element_NormalHighlit.md)

- Observed in 65 parent frames
- Attributes: `a`, `b`, `def`, `g`, `id`, `r`, `texture`, `x`, `y`
  - `a`: `255`
  - `b`: `63`, `0`, `45`
  - `def`: `EA_HorizontalResizeImage_DefaultComboBox`, `EA_AnimatedImage_DefaultChoiceOverlay`, `EA_FullResizeImage_RedTransparent`, `BuffHeadLayoutHorizontalButtonHighlight`, `BuffHeadLayoutVerticalButtonHighlight`, `EA_HorizontalResizeImage_DefaultComboBox2`, `EA_HorizontalResizeImage_ResizeableButtonNormalHighlit`, `EA_Button_ListSortHighlight`
  - `g`: `213`, `255`, `216`, `85`
  - `id`: `morale-white`, `RightTabFrame-Rollover`, `LayoutCorner-BottomLeft-ROLLOVER`, `LayoutCorner-BottomRight-ROLLOVER`, `LayoutCorner-TopLeft-ROLLOVER`, `LayoutCorner-TopRight-ROLLOVER`, `Map-Filters-Button-ROLLOVER`, `Map-Scenario-Button-ROLLOVER`
  - `r`: `250`, `255`, `226`
  - `texture`: `EA_SquareFrame_Highlight`, `PinBG`, `ShiniesIconBorderHighlight`, `TidyRoll_SquareFrame_Highlight`, `nRarityBorderHighlight`
  - `x`: `27`, `105`, `0`, `201`, `475`
  - `y`: `28`, `44`, `0`, `341`, `434`, `420`

### [Disabled](element_Disabled.md)

- Observed in 60 parent frames
- Attributes: `a`, `b`, `def`, `g`, `id`, `r`, `texture`, `x`, `y`
  - `a`: `255`
  - `b`: `92`, `36`, `102`
  - `def`: `EA_HorizontalResizeImage_DefaultComboBox`, `EA_HorizontalResizeImage_DefaultComboBox2`, `EA_HorizontalResizeImage_ResizeableButtonDisabled`, `EA_Button_ListSortNormal`
  - `g`: `92`, `57`, `204`
  - `id`: `morale-white`, `RightTabFrame`, `Map-Plus-Button-Disabled`, `Map-Minus-Button-Disabled`, `IconFrame-1`, `adminbutton-base`
  - `r`: `92`, `95`, `255`
  - `texture`: `EA_SquareFrame`, `PinBG`, `ShiniesIconBorderNormal`, `TidyRoll_SquareFrame`, `nRarityBorderNormal`
  - `x`: `27`, `92`, `0`, `230`
  - `y`: `56`, `44`, `0`, `341`

### [PressedHighlit](element_PressedHighlit.md)

- Observed in 53 parent frames
- Attributes: `a`, `b`, `def`, `g`, `id`, `r`, `x`, `y`
  - `a`: `255`
  - `b`: `63`, `0`, `45`, `36`
  - `def`: `EA_HorizontalResizeImage_DefaultComboBox`, `BuffHeadLayoutHorizontalButtonPressed`, `BuffHeadLayoutVerticalButtonPressed`, `EA_FullResizeImage_RedTransparent`, `EA_HorizontalResizeImage_DefaultComboBox2`
  - `g`: `213`, `255`, `216`, `57`, `85`
  - `id`: `morale-white`, `RightTabFrame-Rollover`, `LayoutCorner-BottomLeft-ROLLOVER`, `LayoutCorner-BottomRight-ROLLOVER`, `LayoutCorner-TopLeft-ROLLOVER`, `LayoutCorner-TopRight-ROLLOVER`, `Map-Filters-Button-ROLLOVER`, `Map-Scenario-Button-ROLLOVER`
  - `r`: `250`, `255`, `226`, `95`
  - `x`: `0`, `120`, `172`, `475`
  - `y`: `56`, `44`, `370`, `434`, `420`

### [TexSlices](element_TexSlices.md)

- Observed in 44 parent frames

### [TextOffset](element_TextOffset.md)

- Observed in 30 parent frames
- Attributes: `x`, `y`
  - `x`: `5`, `0`, `10`
  - `y`: `5`, `7`, `0`

### [TextColors](element_TextColors.md)

- Observed in 23 parent frames

### [ResizeImages](element_ResizeImages.md)

- Observed in 21 parent frames

### [DisabledPressed](element_DisabledPressed.md)

- Observed in 14 parent frames
- Attributes: `a`, `b`, `g`, `id`, `r`
  - `a`: `255`
  - `b`: `36`
  - `g`: `57`
  - `id`: `morale-white`, `RightTabFrame`, `Map-Plus-Button-Disabled`, `Map-Minus-Button-Disabled`
  - `r`: `95`

### [OverlayOffset](element_OverlayOffset.md)

- Observed in 14 parent frames
- Attributes: `x`, `y`
  - `x`: `93`, `193`, `48`, `148`, `223`, `183`, `5`, `73`
  - `y`: `0`, `9`, `7`

### [OverlaySize](element_OverlaySize.md)

- Observed in 14 parent frames
- Attributes: `x`, `y`
  - `x`: `27`, `18`
  - `y`: `28`, `13`

### [OverlayTexCoords](element_OverlayTexCoords.md)

- Observed in 14 parent frames

### [Sound](element_Sound.md)

- Observed in 7 parent frames
- Attributes: `event`, `script`
  - `event`: `OnLButtonDown`, `OnMouseOver`, `OnLButtonUp`
  - `script`: `Sound.Play( Sound.BUTTON_OVER )`, `Sound.Play(Sound.CULTIVATING_HARVEST_CROP)`, `Sound.Play( Sound.CULTIVATING_HARVEST_CROP )`

### [Sounds](element_Sounds.md)

- Observed in 7 parent frames

### [Text](element_Text.md)

- Observed in 3 parent frames

### [AnimatedImages](element_AnimatedImages.md)

- Observed in 1 parent frames

### [Eventhandlers](element_Eventhandlers.md)

- Observed in 1 parent frames

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

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `APAOptionsClose`, `APAOptionsTabsAutoRecall`, `APAOptionsTabsControls`, `APAOptionsTabsFollowTarget`, … | 917 |
| `inherits` | frame-ref | `EA_Button_DefaultWindowClose`, `EA_Button_Tab`, `EA_Button_DefaultResizeable`, `EA_Button_DefaultCheckBox`, … | 868 |
| `id` | number | `3`, `2`, `1`, `4`, … | 249 |
| `textalign` | string | `left`, `center`, `leftcenter`, `bottomright`, … | 160 |
| `font` | string | `font_default_text`, `font_chat_text`, `font_clear_small_bold`, `font_clear_medium`, … | 139 |
| `layer` | string | `overlay`, `popup`, `default`, `secondary`, … | 90 |
| `handleinput` | boolean | `false`, `true` | 78 |
| `backgroundtexture` | string | `shared_01`, `EA_Training_Specialization`, `EA_Abilities01_d5`, `bpKtxt`, … | 74 |
| `highlighttexture` | string | `shared_01`, `EA_Training_Specialization`, `EA_Abilities01_d5`, `EA_HUD_01`, … | 67 |
| `drawchildrenfirst` | boolean | `true`, `false` | 32 |
| `texturescale` | number | `2.0`, `0.74`, `1.171`, `0.68`, … | 32 |
| `alpha` | number | `1`, `0` | 24 |
| `textureScale` | number | `0.62`, `0.75`, `1`, `0.85`, … | 15 |
| `overlayhighlighttexture` | string | `shared_01`, `EA_HUD_01` | 14 |
| `overlaytexture` | string | `shared_01`, `EA_HUD_01` | 14 |

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
