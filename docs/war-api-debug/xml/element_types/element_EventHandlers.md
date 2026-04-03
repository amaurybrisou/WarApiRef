# EventHandlers

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 45/100

## Confidence Assessment

- Level: MEDIUM

- Score: 45/100

- Rationale: Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Namespaces detected | EventHandlers |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: CG_ItemRackEquipmentButton, CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorDown, CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorUp, CM_ClosetGoblin: ClosetGoblinCharacterWindow, CM_ClosetGoblin: ClosetGoblinCharacterWindowClose, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1 |
| XML usage count | 34 |
| XML attribute usage count | 34 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | yes |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

EventHandlers is a structural XML sub-element. It commonly appears under Button and DynamicImage. It is typically used to organize structural children such as EventHandler.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 25× (HIGH)
- [Label](element_Label.md) — 4× (MEDIUM)
- [Window](element_Window.md) — 3× (MEDIUM)
- [DynamicImage](element_DynamicImage.md) — 2× (MEDIUM)

## Common Structural Child Elements

- [EventHandler](element_EventHandler.md) — 65× (HIGH)

## Structural Sub-Elements

### [EventHandler](element_EventHandler.md)

Observed 65 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `event` | **required** | OnLButtonUp, OnLButtonDown, OnRButtonUp, OnMouseOver |
| `function` | **required** | ClosetGoblinCharacterWindow.EquipmentLButtonUp, ClosetGoblinCharacterWindow.EquipmentLButtonDown, ClosetGoblinCharacterWindow.EquipmentRButtonUp, ClosetGoblinCharacterWindow.EquipmentMouseOver |
## Recursive Hierarchy

- Root: [EventHandlers](element_EventHandlers.md)
- [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: CG_ItemRackEquipmentButton -> EventHandlers in Button CG_ItemRackEquipmentButton
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorDown -> EventHandlers in Button ClosetGoblinActionBarPageSelectorDown
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorUp -> EventHandlers in Button ClosetGoblinActionBarPageSelectorUp
- CM_ClosetGoblin: ClosetGoblinCharacterWindow -> EventHandlers in Window ClosetGoblinCharacterWindow
- CM_ClosetGoblin: ClosetGoblinCharacterWindowClose -> EventHandlers in Button ClosetGoblinCharacterWindowClose
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1 -> EventHandlers in Button ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1

## Related APIs

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 90/100) - XML Element Type
- [EventHandler](element_EventHandler.md) (LOW 20/100) - XML Element Type
