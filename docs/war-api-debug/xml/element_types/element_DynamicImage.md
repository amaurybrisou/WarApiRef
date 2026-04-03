# DynamicImage

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
| Namespaces detected | DynamicImage |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPage, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortDownArrow, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortTacticsDownArrow, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortTacticsUpArrow, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortUpArrow, CM_ClosetGoblin: ClosetGoblinCharacterWindowCornerImage |
| XML usage count | 13 |
| XML attribute usage count | 13 |
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

DynamicImage is an interactive XML control. It commonly appears under Button and Window. It is typically used to organize structural children such as Anchors, EventHandlers, Size and bind XML events like OnLButtonUp, OnRButtonUp to Lua.

## Common Attributes

- handleinput
- inherits
- layer
- movable
- name
- popable
- savesettings
- slice
- sticky
- texture
- textureScale
- texturescale

## Common Handlers

- OnLButtonUp
- OnRButtonUp

## Common Handler Functions

- ClosetGoblinOptionWindow.OnLButtonUp
- ClosetGoblinOptionWindow.OnRButtonUp


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| OnLButtonUp | input | ClosetGoblinOptionWindow.OnLButtonUp | `flags, x, y` | MEDIUM |
| OnRButtonUp | input | ClosetGoblinOptionWindow.OnRButtonUp | `flags, x, y` | MEDIUM |

## Common Inherits

- EA_Default_CharacterImage
- EA_ListSortDownArrow
- EA_ListSortUpArrow

## Common Parent Elements

- [Windows](element_Windows.md) — 13× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 12× (HIGH)
- [Size](element_Size.md) — 3× (MEDIUM)
- [EventHandlers](element_EventHandlers.md) — 2× (MEDIUM)
- [Windows](element_Windows.md) — 1× (LOW)

## Common Template Bases

- EA_Default_CharacterImage
- EA_ListSortDownArrow
- EA_ListSortUpArrow

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 69% | EA_Default_CharacterImage, EA_ListSortUpArrow, EA_ListSortDownArrow |
| `layer` | optional | 30% | background, overlay, default, popup |
| `texture` | optional | 23% | shared_01, ClosetGoblinIcon, WHLCG |
| `handleinput` | optional | 15% | false |
| `movable` | optional | 7% | false |
| `popable` | optional | 7% | true |
| `savesettings` | optional | 7% | false |
| `slice` | optional | 7% | Radio-Button |
| `sticky` | optional | 7% | false |
| `textureScale` | optional | 7% | 1.171 |
| `texturescale` | optional | 7% | 1.00 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 12 times as an unnamed child.

### [Size](element_Size.md)

Observed 3 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 2 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [DynamicImage](element_DynamicImage.md)
- [Anchors](element_Anchors.md) (structural, 12×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
- [EventHandlers](element_EventHandlers.md) (structural, 2×, MEDIUM)
  - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
- [Size](element_Size.md) (structural, 3×, MEDIUM)
  - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
- [Windows](element_Windows.md) (structural, 1×, LOW)
  - [Button](element_Button.md) (named, 57×, HIGH)
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
      - (cycle)
  - [DynamicImage](element_DynamicImage.md) (named, 13×, HIGH)
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

## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

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

- ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- ClosetGoblinCharacterWindow.OnInitialize
- ClosetGoblinZoneWindow.OnInitialize
- ClosetGoblin.Initialize


## Binding Resolution

- Total handler declarations: 4
- Resolved to Lua functions: 4 (100%)

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPage -> DynamicImage ClosetGoblinActionBarPageSelectorCurrentPage
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortDownArrow -> DynamicImage ClosetGoblinCharacterWindowContentsSortDownArrow
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortTacticsDownArrow -> DynamicImage ClosetGoblinCharacterWindowContentsSortTacticsDownArrow
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortTacticsUpArrow -> DynamicImage ClosetGoblinCharacterWindowContentsSortTacticsUpArrow
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortUpArrow -> DynamicImage ClosetGoblinCharacterWindowContentsSortUpArrow
- CM_ClosetGoblin: ClosetGoblinCharacterWindowCornerImage -> DynamicImage ClosetGoblinCharacterWindowCornerImage

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
- [Size](element_Size.md) (HIGH 98/100) - XML Element Type
- [EA_Default_CharacterImage](../../globals/constants/constant_EA_Default_CharacterImage.md) (HIGH 90/100) - Constant
- [EA_ListSortDownArrow](../../globals/constants/constant_EA_ListSortDownArrow.md) (HIGH 90/100) - Constant
- [EA_ListSortUpArrow](../../globals/constants/constant_EA_ListSortUpArrow.md) (HIGH 90/100) - Constant
- [EventHandlers](element_EventHandlers.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
