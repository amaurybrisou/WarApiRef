# Label

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
| Namespaces detected | Label |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPageText, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabel, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB1, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB2, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB3, CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB4 |
| XML usage count | 61 |
| XML attribute usage count | 61 |
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

Label is a display-focused XML element used to present text content in UI layouts. It commonly appears inside container elements such as Window.

## Common Attributes

- autoresize
- autoresizewidth
- font
- handleinput
- id
- inherits
- layer
- maxchars
- name
- popable
- textalign

## Common Handlers

- OnLButtonUp
- OnRButtonUp

## Common Handler Functions

- ClosetGoblinCharacterWindow.OnClickSetRow
- ClosetGoblinZoneWindow.OnClickZoneRow
- ClosetGoblinCharacterWindow.OnSetRowContextMenu
- ClosetGoblinZoneWindow.OnSetZoneRowContextMenu


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| OnLButtonUp | input | ClosetGoblinCharacterWindow.OnClickSetRow, ClosetGoblinZoneWindow.OnClickZoneRow | `flags, x, y` | MEDIUM |
| OnRButtonUp | input | ClosetGoblinCharacterWindow.OnSetRowContextMenu, ClosetGoblinZoneWindow.OnSetZoneRowContextMenu | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `WindowGetId`

**OnRButtonUp** handlers call: `WindowGetId`

## Common Inherits

- ClosetGoblinTalismanLabel

## Common Parent Elements

- [Windows](element_Windows.md) — 61× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 60× (HIGH)
- [Size](element_Size.md) — 15× (HIGH)
- [Color](element_Color.md) — 4× (MEDIUM)
- [EventHandlers](element_EventHandlers.md) — 4× (MEDIUM)

## Common Template Bases

- ClosetGoblinTalismanLabel

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 75% | ClosetGoblinTalismanLabel |
| `font` | optional | 24% | font_clear_medium_bold, font_default_text, font_clear_tiny, font_chat_text, ... |
| `textalign` | optional | 21% | center, left |
| `handleinput` | optional | 18% | false |
| `layer` | optional | 14% | overlay, popup |
| `autoresize` | optional | 3% | true |
| `autoresizewidth` | optional | 1% | true |
| `id` | optional | 1% | 1 |
| `maxchars` | optional | 1% | 10 |
| `popable` | optional | 1% | false |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 60 times as an unnamed child.

### [Size](element_Size.md)

Observed 15 times as an unnamed child.

### [Color](element_Color.md)

Observed 4 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `a` | **required** | 255 |
| `b` | **required** | 255 |
| `g` | **required** | 255 |
| `r` | **required** | 255 |
### [EventHandlers](element_EventHandlers.md)

Observed 4 times as an unnamed child.

## Recursive Hierarchy

- Root: [Label](element_Label.md)
- [Anchors](element_Anchors.md) (structural, 60×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
- [Color](element_Color.md) (structural, 4×, MEDIUM)
- [EventHandlers](element_EventHandlers.md) (structural, 4×, MEDIUM)
  - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
- [Size](element_Size.md) (structural, 15×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetId` | ui | 8 | OnLButtonUp, OnRButtonUp |
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

- ClosetGoblinCharacterWindow.OnInitialize
- ClosetGoblinZoneWindow.RefreshOption
- Clock.OnUpdate


## Binding Resolution

- Total handler declarations: 8
- Resolved to Lua functions: 8 (100%)

## Seen In

- CM_ClosetGoblin
- Clock

## Examples

- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPageText -> Label ClosetGoblinActionBarPageSelectorCurrentPageText
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabel -> Label ClosetGoblinCharacterWindowContentsActionBarSettingsLabel
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB1 -> Label ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB1
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB2 -> Label ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB2
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB3 -> Label ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB3
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB4 -> Label ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB4

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
- [Color](element_Color.md) (HIGH 98/100) - XML Element Type
- [Size](element_Size.md) (HIGH 98/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 90/100) - XML Element Type
