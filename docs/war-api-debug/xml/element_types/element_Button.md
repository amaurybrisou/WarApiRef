# Button

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 90/100

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Namespaces detected | Button |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot1, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot10, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot2, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot3, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot4, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot5 |
| XML usage count | 57 |
| XML attribute usage count | 57 |
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

Button is an interactive XML control used for click and mouse-driven callbacks. It usually binds input events to Lua handler functions.

## Common Attributes

- draganddrop
- drawchildrenfirst
- font
- gameactionbutton
- id
- inherits
- layer
- name
- textalign
- texturescale

## Common Handlers

- OnLButtonDown
- OnLButtonUp
- OnMouseOver
- OnMouseOverEnd
- OnRButtonDown
- OnRButtonUp

## Common Handler Functions

- ClosetGoblinCharacterWindow.EquipmentLButtonUp
- ClosetGoblinCharacterWindow.Hide
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- ClosetGoblinCharacterWindow.HotbarPageDownProxy
- ClosetGoblinCharacterWindow.HotbarPageUpProxy
- ClosetGoblinCharacterWindow.OnClickDeleteSetButton
- ClosetGoblinCharacterWindow.OnClickNewSetButton
- ClosetGoblinCharacterWindow.OnClickSortNameButton
- ClosetGoblinCharacterWindow.OnClickSortSetButton
- ClosetGoblinCharacterWindow.OnClickSortTacticsButton
- ClosetGoblinCharacterWindow.OnClickZoneConfigButton
- ClosetGoblinCharacterWindow.ShowCloak
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- ClosetGoblinCharacterWindow.ShowHelm
- ClosetGoblinCharacterWindow.UseItemRackToggled
- ClosetGoblinZoneWindow.Hide
- ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton
- ClosetGoblinZoneWindow.OnClickSortZoneButton
- ActionBars.OnMouseoverHotbarPageDown
- ActionBars.OnMouseoverHotbarPageUp
- CG_ItemRack.EquipmentMouseOver
- ClosetGoblinCharacterWindow.EquipmentMouseOver
- ClosetGoblinCharacterWindow.ShowCloakOptions
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- ClosetGoblinCharacterWindow.ShowShowHelm
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- ClosetGoblinCharacterWindow.HideCloakOptions
- ClosetGoblinCharacterWindow.HideShowHelm
- ClosetGoblinCharacterWindow.EquipmentMouseOverEnd
- ClosetGoblinCharacterWindow.EquipmentLButtonDown
- ClosetGoblinCharacterWindow.EquipmentRButtonDown
- CG_ItemRack.EquipmentRButtonDown
- ClosetGoblinCharacterWindow.EquipmentRButtonUp


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| OnLButtonUp | input | ClosetGoblinCharacterWindow.EquipmentLButtonUp, ClosetGoblinCharacterWindow.Hide, ClosetGoblinCharacterWindow.HotbarChangeToggled1, ClosetGoblinCharacterWindow.HotbarChangeToggled2, ClosetGoblinCharacterWindow.HotbarChangeToggled3, ClosetGoblinCharacterWindow.HotbarChangeToggled4, ClosetGoblinCharacterWindow.HotbarChangeToggled5, ClosetGoblinCharacterWindow.HotbarPageDownProxy, ClosetGoblinCharacterWindow.HotbarPageUpProxy, ClosetGoblinCharacterWindow.OnClickDeleteSetButton, ClosetGoblinCharacterWindow.OnClickNewSetButton, ClosetGoblinCharacterWindow.OnClickSortNameButton, ClosetGoblinCharacterWindow.OnClickSortSetButton, ClosetGoblinCharacterWindow.OnClickSortTacticsButton, ClosetGoblinCharacterWindow.OnClickZoneConfigButton, ClosetGoblinCharacterWindow.ShowCloak, ClosetGoblinCharacterWindow.ShowCloakHeraldry, ClosetGoblinCharacterWindow.ShowHelm, ClosetGoblinCharacterWindow.UseItemRackToggled, ClosetGoblinZoneWindow.Hide, ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton, ClosetGoblinZoneWindow.OnClickSortZoneButton | `flags, x, y` | MEDIUM |
| OnMouseOver | input | ActionBars.OnMouseoverHotbarPageDown, ActionBars.OnMouseoverHotbarPageUp, CG_ItemRack.EquipmentMouseOver, ClosetGoblinCharacterWindow.EquipmentMouseOver, ClosetGoblinCharacterWindow.ShowCloakOptions, ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly, ClosetGoblinCharacterWindow.ShowShowCloakOnly, ClosetGoblinCharacterWindow.ShowShowHelm, ClosetGoblinCharacterWindow.ShowShowHelmOnly | `` |  |
| OnMouseOverEnd | input | ClosetGoblinCharacterWindow.HideCloakOptions, ClosetGoblinCharacterWindow.HideShowHelm, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd | `` |  |
| OnLButtonDown | input | ClosetGoblinCharacterWindow.EquipmentLButtonDown | `flags, x, y` | MEDIUM |
| OnRButtonDown | input | ClosetGoblinCharacterWindow.EquipmentRButtonDown, CG_ItemRack.EquipmentRButtonDown | `flags, x, y` | LOW |
| OnRButtonUp | input | ClosetGoblinCharacterWindow.EquipmentRButtonUp | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `LabelSetText`, `WindowGetId`

**OnMouseOver** handlers call: `WindowGetId`, `WindowSetShowing`

**OnMouseOverEnd** handlers call: `WindowSetShowing`

**OnRButtonUp** handlers call: `WindowGetId`

## Common Inherits

- CG_ItemRackEquipmentButton
- CharacterWindowDefaultButton
- ClosetGoblinDefaultButton
- ClosetGoblinEquipmentButton
- EA_Button_Default
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultResizeable
- EA_Button_DefaultUp
- EA_Button_DefaultWindowClose
- EA_Button_ListSort

## Common Parent Elements

- [Windows](element_Windows.md) — 57× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 53× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 25× (HIGH)
- [Size](element_Size.md) — 5× (MEDIUM)
- [TexCoords](element_TexCoords.md) — 1× (LOW)
- [TextOffset](element_TextOffset.md) — 1× (LOW)
- [Windows](element_Windows.md) — 1× (LOW)

## Common Template Bases

- CG_ItemRackEquipmentButton
- CharacterWindowDefaultButton
- ClosetGoblinDefaultButton
- ClosetGoblinEquipmentButton
- EA_Button_Default
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultResizeable
- EA_Button_DefaultUp
- EA_Button_DefaultWindowClose
- EA_Button_ListSort

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 100% | EA_Button_Default, ClosetGoblinDefaultButton, EA_Button_DefaultUp, EA_Button_DefaultDown, ... |
| `id` | optional | 64% | 1, 2, 3, 12, ... |
| `layer` | optional | 15% | overlay, popup |
| `font` | optional | 8% | font_default_text, font_chat_text |
| `textalign` | optional | 8% | bottomright, center |
| `draganddrop` | optional | 5% | true |
| `gameactionbutton` | optional | 5% | right |
| `drawchildrenfirst` | optional | 1% | true |
| `texturescale` | optional | 1% | 1.171 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 53 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 25 times as an unnamed child.

### [Size](element_Size.md)

Observed 5 times as an unnamed child.

### [TexCoords](element_TexCoords.md)

Observed 1 times as an unnamed child.

### [TextOffset](element_TextOffset.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5 |
### [Windows](element_Windows.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [Button](element_Button.md)
- [Anchors](element_Anchors.md) (structural, 53×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
- [EventHandlers](element_EventHandlers.md) (structural, 25×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
- [Size](element_Size.md) (structural, 5×, MEDIUM)
  - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
  - [BottomCenter](element_BottomCenter.md) (structural, 1×, HIGH)
  - [BottomLeft](element_BottomLeft.md) (structural, 1×, HIGH)
  - [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 1×, HIGH)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 1×, HIGH)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 1×, HIGH)
  - [MiddleRight](element_MiddleRight.md) (structural, 1×, HIGH)
  - [Normal](element_Normal.md) (structural, 1×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 1×, HIGH)
  - [TopCenter](element_TopCenter.md) (structural, 1×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
  - [TopRight](element_TopRight.md) (structural, 1×, HIGH)
- [TextOffset](element_TextOffset.md) (structural, 1×, LOW)
- [Windows](element_Windows.md) (structural, 1×, LOW)
  - [Button](element_Button.md) (named, 57×, HIGH)
    - (cycle)
  - [DynamicImage](element_DynamicImage.md) (named, 13×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 12×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 2×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
    - [Size](element_Size.md) (structural, 3×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [FullResizeImage](element_FullResizeImage.md) (named, 7×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 6×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [Size](element_Size.md) (structural, 2×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 1×, LOW)
      - [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
      - [Middle](element_Middle.md) (structural, 1×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 1×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 1×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 1×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 1×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 1×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 1×, HIGH)
      - [Normal](element_Normal.md) (structural, 1×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 1×, HIGH)
      - [TopCenter](element_TopCenter.md) (structural, 1×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 1×, HIGH)
  - [Label](element_Label.md) (named, 61×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 60×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [Color](element_Color.md) (structural, 4×, MEDIUM)
    - [EventHandlers](element_EventHandlers.md) (structural, 4×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
    - [Size](element_Size.md) (structural, 15×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
  - [ListBox](element_ListBox.md) (named, 2×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 2×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [ListData](element_ListData.md) (structural, 2×, HIGH)
      - [ListColumns](element_ListColumns.md) (structural, 2×, HIGH)
        - [ListColumn](element_ListColumn.md) (structural, 4×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
  - [Window](element_Window.md) (named, 25×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 20×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 3×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
    - [Size](element_Size.md) (structural, 13×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
    - [Windows](element_Windows.md) (structural, 13×, HIGH)
      - (cycle)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowSetShowing` | ui | 14 | OnMouseOver, OnMouseOverEnd |
| `ButtonSetPressedFlag` | ui | 9 | OnLButtonUp |
| `WindowGetId` | ui | 6 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ButtonSetDisabledFlag` | ui | 2 | OnLButtonUp |
| `LabelSetText` | ui | 2 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

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

- ClosetGoblinCharacterWindow.UseItemRackToggled
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- ClosetGoblinCharacterWindow.HideShowHelm
- ClosetGoblinCharacterWindow.ShowHelm
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- ClosetGoblinCharacterWindow.ShowShowHelm
- ClosetGoblinCharacterWindow.UpdateActionBarSettings
- ClosetGoblinCharacterWindow.HideCloakOptions
- ClosetGoblinCharacterWindow.ShowCloakOptions
- ClosetGoblinCharacterWindow.OnShow
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- ClosetGoblinCharacterWindow.ShowCloak
- ClosetGoblinZoneWindow.OnInitialize
- ClosetGoblinCharacterWindow.UpdateSetRow
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- ClosetGoblinCharacterWindow.OnInitialize


## Binding Resolution

- Total handler declarations: 46
- Resolved to Lua functions: 43 (93%)

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot1 -> Button CG_ItemRackEQShow1EquipmentSlot1
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot10 -> Button CG_ItemRackEQShow1EquipmentSlot10
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot2 -> Button CG_ItemRackEQShow1EquipmentSlot2
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot3 -> Button CG_ItemRackEQShow1EquipmentSlot3
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot4 -> Button CG_ItemRackEQShow1EquipmentSlot4
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot5 -> Button CG_ItemRackEQShow1EquipmentSlot5

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
- [Size](element_Size.md) (HIGH 98/100) - XML Element Type
- [EA_Button_Default](../../globals/constants/constant_EA_Button_Default.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultCheckBox](../../globals/constants/constant_EA_Button_DefaultCheckBox.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultDown](../../globals/constants/constant_EA_Button_DefaultDown.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultResizeable](../../globals/constants/constant_EA_Button_DefaultResizeable.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultUp](../../globals/constants/constant_EA_Button_DefaultUp.md) (HIGH 90/100) - Constant
- [EA_Button_DefaultWindowClose](../../globals/constants/constant_EA_Button_DefaultWindowClose.md) (HIGH 90/100) - Constant
- [EA_Button_ListSort](../../globals/constants/constant_EA_Button_ListSort.md) (HIGH 90/100) - Constant
- [EventHandlers](element_EventHandlers.md) (MEDIUM 45/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (MEDIUM 45/100) - XML Element Type
- [TextOffset](element_TextOffset.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 90/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 90/100) - XML Element Type
