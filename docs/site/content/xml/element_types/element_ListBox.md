# ListBox

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
| Addons seen in | AggroMeter, Aura, BlackBook, BuffHead, CDown, CM_ClosetGoblin, CMap, CastSequence |
| Files seen in | Code/CombatLog/CombatLogSnapshotWindow.xml, Code/CombatLog/CombatLogStatsWindow.xml, Code/Core/ChooseIconDialog.xml, Code/Core/Debug.xml, Code/KillSpam/KillSpam.xml, Code/ScenarioInfo/ScenarioInfo.xml, Code/UnitFrames/EffectsIndicatorDialog.xml, Code/UnitFrames/UnitFramesConfiguration.xml |
| Namespaces detected | ListBox |
| Source kinds | xml_frames |
| Example locations | AggroMeter: AggroMeterGrayListBox, Aura: AuraSettingsAuraList, Aura: AuraSharesAuraList, Aura: AuraTextureIconsIcons, BlackBook: BlackBookWindowList, BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList |
| XML usage count | 113 |
| XML attribute usage count | 113 |
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

ListBox is a container-style XML element for row-based UI lists. It commonly works with structural children such as ListData and ListColumns and binds events to Lua callbacks.

## Common Attributes

- autoHideScrollBar
- color
- draganddrop
- inherits
- layer
- name
- rowcount
- rowdef
- rowspacing
- scrollbar
- scrollbarPosition
- visiblerows

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnUpdate](../handlers/handler_OnUpdate.md)

## Common Handler Functions

- Deathblow.OnTabLBU
- Enemy.UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp
- TidyRoll.CustomAutoRoll.OnListLbuttonUp
- wbLeadHelperMessagesTab.OnListLButtonUp
- EA_Window_OpenPartyManage.OnMouseOverWarbandListBox
- Enemy.KillSpamUI_KillSpamDialog_OnMouseOver
- EA_Window_OpenPartyManage.OnMouseOverEndWarbandListBox
- Enemy.KillSpamUI_KillSpamDialog_OnMouseOverEnd
- AggroMeter.PickedListMenu
- LootAlert.UpdateColor


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | , Deathblow.OnTabLBU, Enemy.UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp, TidyRoll.CustomAutoRoll.OnListLbuttonUp, wbLeadHelperMessagesTab.OnListLButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | EA_Window_OpenPartyManage.OnMouseOverWarbandListBox, Enemy.KillSpamUI_KillSpamDialog_OnMouseOver | `` |  |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | EA_Window_OpenPartyManage.OnMouseOverEndWarbandListBox, Enemy.KillSpamUI_KillSpamDialog_OnMouseOverEnd | `` |  |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | AggroMeter.PickedListMenu | `flags, x, y` | MEDIUM |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | LootAlert.UpdateColor | `elapsed` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `ButtonSetPressedFlag`, `Cursor.Clear`, `Cursor.IconOnCursor`, `ListBoxGetDataIndex`, `WindowGetId`

**OnRButtonUp** handlers call: `ListBoxGetDataIndex`, `WindowGetId`

**OnUpdate** handlers call: `ListBoxSetVisibleRowCount`

## Common Inherits

- EA_Template_WarbandManageListBox
- NBSBChecksList
- NBSBParamList

## Common Parent Elements

- [Windows](element_Windows.md) — 113× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 110× (HIGH)
- [ListData](element_ListData.md) — 110× (HIGH)
- [Size](element_Size.md) — 76× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 16× (HIGH)
- [Windows](element_Windows.md) — 1× (LOW)

## Common Template Bases

- EA_Template_WarbandManageListBox
- NBSBChecksList
- NBSBParamList


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `rowdef` | **required** | 94% | AggroMeterGrayTemplateListboxRow, AuraSharesRow, AuraIconRow, AuraWindowRow, ... |
| `rowspacing` | **required** | 94% | 0, 1, 6, 5, ... |
| `visiblerows` | **required** | 94% | 13, 5, 4, 7, ... |
| `scrollbar` | **required** | 82% | EA_ScrollBar_DefaultVerticalChain, DPSMeterVerticalScrollbar, ScenarioSummaryVertScroll, UngroupedPlayerVertScroll, ... |
| `rowcount` | optional | 46% | 120, 150, 250, 100, ... |
| `layer` | optional | 7% | secondary, popup |
| `inherits` | optional | 5% | EA_Template_WarbandManageListBox, NBSBChecksList, NBSBParamList |
| `color` | optional | 4% | 155, 255, 155, 0, 255, 255, 255, 0 |
| `autoHideScrollBar` | optional | 0% | true |
| `draganddrop` | optional | 0% | true |
| `scrollbarPosition` | optional | 0% | left |
| `useDoubleAnchoringForRows` | optional | 0% | true |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 110 times as an unnamed child.

### [ListData](element_ListData.md)

Observed 110 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `table` | **required** | AggroMeter.Listdata, AuraShares.listDisplayData, AuraTexture.listIconDisplayData, AuraSettings.listDisplayData |
| `populationfunction` | **required** | AuraShares.PopulateDisplay, AuraTexture.PopulateIconsListDisplay, AuraSettings.PopulateDisplay, BlackBookWindow.Populate |
### [Size](element_Size.md)

Observed 76 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 16 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [ListBox](element_ListBox.md)
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
| `WindowGetId` | ui | 8 | OnLButtonUp, OnRButtonUp |
| `ButtonSetPressedFlag` | ui | 6 | OnLButtonUp |
| `ListBoxGetDataIndex` | ui | 6 | OnLButtonUp, OnRButtonUp |
| `ListBoxSetVisibleRowCount` | ui | 2 | OnUpdate |
| `Cursor.Clear` | ui | 1 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseOver

Confidence: LOW

### OnMouseOverEnd

Confidence: LOW

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
## Lua Functions Manipulating This Type

- RealmStatus.Loading
- mml.OnItemMouseOver
- ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- DAoCBuffSettings.SetLabels
- DAoCBuffSettings.Change_Setting
- DAoCBuffSettings.CreateOptionswindow
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- TTitan.UI.DisplayRow
- ClosetGoblinZoneWindow.UpdateHighlightOnRow


## Binding Resolution

- Total handler declarations: 18
- Resolved to Lua functions: 14 (77%)

## Seen In

- AggroMeter
- Aura
- BlackBook
- BuffHead
- CDown
- CM_ClosetGoblin
- CMap
- CastSequence
- Cheeseboard
- Crusher
- DAoCBuff
- DPSMeter
- Dascore
- Deathblow
- Deathblow2
- DetauntHelper
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiModWindow
- Emojii
- Enemy
- FozAuction
- HealGrid
- Hopper
- Kwestor
- LibAddonButton
- LootAlert
- MapMonster
- MapPin
- Minmap
- Motion
- NerfedButtons
- Obsidian
- PieTracker
- Pocket Palette
- Pure
- QuickTacticSwitch
- RealmStatus
- SessionRPs
- Shinies
- SocialWindow 2.0
- TastyButtons
- TaxPayer
- TexturedButtons
- TidyChat
- TidyRoll
- Tome Titan
- TomeTracker
- Tortall_DPS
- TurretRange
- WTes
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- nLootLink
- wbLeadHelper
- zMailMod

## Examples

- AggroMeter: AggroMeterGrayListBox -> ListBox AggroMeterGrayListBox
- Aura: AuraSettingsAuraList -> ListBox AuraSettingsAuraList
- Aura: AuraSharesAuraList -> ListBox AuraSharesAuraList
- Aura: AuraTextureIconsIcons -> ListBox AuraTextureIconsIcons
- BlackBook: BlackBookWindowList -> ListBox BlackBookWindowList
- BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList -> ListBox BuffHeadSetupAdvancedCompressionItemWindowList

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [ListData](element_ListData.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_Template_WarbandManageListBox](../../globals/constants/constant_EA_Template_WarbandManageListBox.md) (HIGH 90/100) - Constant

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
