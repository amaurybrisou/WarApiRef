# Size

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 98/100

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Clock |
| Namespaces detected | Size |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: CG_ItemRackEQShow1, CM_ClosetGoblin: CG_ItemRackEQShow1Equipment, CM_ClosetGoblin: ClosetGoblinActionBarPageSelector, CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPage, CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPageText, CM_ClosetGoblin: ClosetGoblinCharacterWindow |
| XML usage count | 39 |
| XML attribute usage count | 39 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

Size is a structural XML sub-element. It commonly appears under Button and DynamicImage. It is typically used to organize structural children such as AbsPoint.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 15× (HIGH)
- [Window](element_Window.md) — 13× (HIGH)
- [Button](element_Button.md) — 5× (MEDIUM)
- [DynamicImage](element_DynamicImage.md) — 3× (MEDIUM)
- [FullResizeImage](element_FullResizeImage.md) — 2× (MEDIUM)
- [ListBox](element_ListBox.md) — 1× (LOW)

## Common Structural Child Elements

- [AbsPoint](element_AbsPoint.md) — 39× (HIGH)

## Structural Sub-Elements

### [AbsPoint](element_AbsPoint.md)

Observed 39 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 75, 65, 355, 340 |
| `y` | **required** | 75, 20, 26, 0 |
## Recursive Hierarchy

- Root: [Size](element_Size.md)
- [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)

## Seen In

- CM_ClosetGoblin
- Clock

## Examples

- CM_ClosetGoblin: CG_ItemRackEQShow1 -> Size in Window CG_ItemRackEQShow1
- CM_ClosetGoblin: CG_ItemRackEQShow1Equipment -> Size in Window CG_ItemRackEQShow1Equipment
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelector -> Size in Window ClosetGoblinActionBarPageSelector
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPage -> Size in DynamicImage ClosetGoblinActionBarPageSelectorCurrentPage
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPageText -> Size in Label ClosetGoblinActionBarPageSelectorCurrentPageText
- CM_ClosetGoblin: ClosetGoblinCharacterWindow -> Size in Window ClosetGoblinCharacterWindow

## Related APIs

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 90/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 90/100) - XML Element Type
- [ListBox](element_ListBox.md) (HIGH 90/100) - XML Element Type
- [AbsPoint](element_AbsPoint.md) (MEDIUM 35/100) - XML Element Type

## Used With

- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
