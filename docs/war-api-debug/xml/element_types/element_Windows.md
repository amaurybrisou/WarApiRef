# Windows

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Clock |
| Namespaces detected | Windows |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: CG_ItemRackEQShow1, CM_ClosetGoblin: CG_ItemRackEQShow1Equipment, CM_ClosetGoblin: ClosetGoblinActionBarPageSelector, CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPage, CM_ClosetGoblin: ClosetGoblinCharacterWindow, CM_ClosetGoblin: ClosetGoblinCharacterWindowContents |
| XML usage count | 15 |
| XML attribute usage count | 15 |
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

Windows is a container-style XML element. It commonly appears under Button and DynamicImage. It is typically used to host named child elements such as Button, DynamicImage, FullResizeImage.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md) — 13× (HIGH)
- [Button](element_Button.md) — 1× (LOW)
- [DynamicImage](element_DynamicImage.md) — 1× (LOW)

## Common Named Child Elements

- [Label](element_Label.md) — 61× (HIGH)
- [Button](element_Button.md) — 57× (HIGH)
- [Window](element_Window.md) — 25× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 13× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 7× (HIGH)
- [ListBox](element_ListBox.md) — 2× (MEDIUM)

## Recursive Hierarchy

- Root: [Windows](element_Windows.md)
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

## Seen In

- CM_ClosetGoblin
- Clock

## Examples

- CM_ClosetGoblin: CG_ItemRackEQShow1 -> Windows in Window CG_ItemRackEQShow1
- CM_ClosetGoblin: CG_ItemRackEQShow1Equipment -> Windows in Window CG_ItemRackEQShow1Equipment
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelector -> Windows in Window ClosetGoblinActionBarPageSelector
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPage -> Windows in DynamicImage ClosetGoblinActionBarPageSelectorCurrentPage
- CM_ClosetGoblin: ClosetGoblinCharacterWindow -> Windows in Window ClosetGoblinCharacterWindow
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContents -> Windows in Window ClosetGoblinCharacterWindowContents

## Related APIs

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 90/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 90/100) - XML Element Type
- [ListBox](element_ListBox.md) (HIGH 90/100) - XML Element Type
