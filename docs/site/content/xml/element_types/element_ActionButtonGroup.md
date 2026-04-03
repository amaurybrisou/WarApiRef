# ActionButtonGroup

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
| Addons seen in | BankWindowFix, EA_UiModWindow, RVMOD_Manager |
| Files seen in | Source/BankWindowFix.xml, Source/UiModInfoTemplate.xml |
| Namespaces detected | ActionButtonGroup |
| Source kinds | xml_frames |
| Example locations | BankWindowFix: BankWindowSlotsFixed, EA_UiModWindow: EA_ScrollWindow_ModInfoTemplateScrollChildCareers, RVMOD_Manager: RVMOD_ManagerModInfoTemplateScrollChildCareers |
| XML usage count | 3 |
| XML attribute usage count | 3 |
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

ActionButtonGroup is an interactive XML control. It commonly appears under Window. It is typically used to organize structural children such as Anchors, EventHandlers and bind XML events like OnActionButtonLButtonDown, OnActionButtonLButtonUp, OnActionButtonMouseOver to Lua.

## Common Attributes

- draganddrop
- gameactionbutton
- hideButtonWhenIconBlank
- inherits
- name

## Common Handlers

- OnActionButtonLButtonDown
- OnActionButtonLButtonUp
- OnActionButtonMouseOver
- OnActionButtonRButtonDown

## Common Handler Functions

- BankWindow.EquipmentLButtonDown
- BankWindow.EquipmentLButtonUp
- BankWindow.EquipmentMouseOver
- BankWindow.EquipmentRButtonDown


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| OnActionButtonLButtonDown | custom | BankWindow.EquipmentLButtonDown | `buttonIndex, flags` | LOW |
| OnActionButtonLButtonUp | custom | BankWindow.EquipmentLButtonUp | `flags` | LOW |
| OnActionButtonMouseOver | custom | BankWindow.EquipmentMouseOver | `` |  |
| OnActionButtonRButtonDown | custom | BankWindow.EquipmentRButtonDown | `` |  |

### Per-Event Lua API Calls

**OnActionButtonLButtonDown** handlers call: `Cursor.Clear`, `Cursor.IconOnCursor`, `Cursor.PickUp`

**OnActionButtonLButtonUp** handlers call: `WindowGetId`

**OnActionButtonMouseOver** handlers call: `WindowGetId`

## Common Inherits

- EA_ActionButtonGroup_CareerIconsWithTooltip
- EA_ActionButtonGroup_DefaultSmall

## Common Parent Elements

- [Windows](element_Windows.md) — 3× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 3× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 1× (LOW)

## Common Template Bases

- EA_ActionButtonGroup_CareerIconsWithTooltip
- EA_ActionButtonGroup_DefaultSmall


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 100% | EA_ActionButtonGroup_DefaultSmall, EA_ActionButtonGroup_CareerIconsWithTooltip |
| `hideButtonWhenIconBlank` | optional | 66% | true |
| `draganddrop` | optional | 33% | true |
| `gameactionbutton` | optional | 33% | right |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 3 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [ActionButtonGroup](element_ActionButtonGroup.md)
- [Anchors](element_Anchors.md) (structural, 3×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
      - (cycle)
- [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
  - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetId` | ui | 2 | OnActionButtonLButtonUp, OnActionButtonMouseOver |
| `Cursor.Clear` | ui | 1 | OnActionButtonLButtonDown |
| `Cursor.IconOnCursor` | data | 1 | OnActionButtonLButtonDown |
| `Cursor.PickUp` | data | 1 | OnActionButtonLButtonDown |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnActionButtonLButtonDown

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `buttonIndex` | number | number |
| 1 | `flags` | boolean | modifier_flags |
### OnActionButtonLButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | boolean | modifier_flags |
### OnActionButtonMouseOver

Confidence: LOW

### OnActionButtonRButtonDown

Confidence: LOW

## Lua Functions Manipulating This Type

- BankWindowFix.Initialize


## Binding Resolution

- Total handler declarations: 4
- Resolved to Lua functions: 3 (75%)

## Seen In

- BankWindowFix
- EA_UiModWindow
- RVMOD_Manager

## Examples

- BankWindowFix: BankWindowSlotsFixed -> ActionButtonGroup BankWindowSlotsFixed
- EA_UiModWindow: EA_ScrollWindow_ModInfoTemplateScrollChildCareers -> ActionButtonGroup EA_ScrollWindow_ModInfoTemplateScrollChildCareers
- RVMOD_Manager: RVMOD_ManagerModInfoTemplateScrollChildCareers -> ActionButtonGroup RVMOD_ManagerModInfoTemplateScrollChildCareers

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EA_ActionButtonGroup_CareerIconsWithTooltip](../../globals/constants/constant_EA_ActionButtonGroup_CareerIconsWithTooltip.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_ActionButtonGroup_DefaultSmall](../../globals/constants/constant_EA_ActionButtonGroup_DefaultSmall.md) (HIGH 90/100) - Constant
