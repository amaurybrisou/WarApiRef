# ColorPicker

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
| Addons seen in | WSCT |
| Files seen in | wsct_options/wsct_options.xml |
| Namespaces detected | ColorPicker |
| Source kinds | xml_frames |
| Example locations | WSCT: WSCTOptionsColorPickerWindowColorPicker |
| XML usage count | 1 |
| XML attribute usage count | 1 |
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

ColorPicker is an interactive XML control. It commonly appears under Window. It is typically used to organize structural children such as Anchors, ColorSize, ColorSpacing and bind XML events like OnLButtonUp to Lua.

## Common Attributes

- columnsPerRow
- name
- texture

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)

## Common Handler Functions

- WSCT.OnLButtonUpColorPicker


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | WSCT.OnLButtonUpColorPicker | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `LabelSetTextColor`, `SliderBarSetCurrentPosition`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowSetShowing`, `WindowSetTintColor`

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 1× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 1× (HIGH)
- [ColorSize](element_ColorSize.md) — 1× (HIGH)
- [ColorSpacing](element_ColorSpacing.md) — 1× (HIGH)
- [ColorTexCoords](element_ColorTexCoords.md) — 1× (HIGH)
- [ColorTexDims](element_ColorTexDims.md) — 1× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 1× (HIGH)
- [Size](element_Size.md) — 1× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `columnsPerRow` | **required** | 100% | 4 |
| `texture` | **required** | 100% | EA_HUD_01 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 1 times as an unnamed child.

### [ColorSize](element_ColorSize.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 28 |
| `y` | **required** | 28 |
### [ColorSpacing](element_ColorSpacing.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5 |
| `y` | **required** | 5 |
### [ColorTexCoords](element_ColorTexCoords.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 304 |
| `y` | **required** | 450 |
### [ColorTexDims](element_ColorTexDims.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 24 |
| `y` | **required** | 24 |
### [EventHandlers](element_EventHandlers.md)

Observed 1 times as an unnamed child.

### [Size](element_Size.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [ColorPicker](element_ColorPicker.md)
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

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `SliderBarSetCurrentPosition` | ui | 3 | OnLButtonUp |
| `LabelSetTextColor` | ui | 2 | OnLButtonUp |
| `WindowAddAnchor` | ui | 1 | OnLButtonUp |
| `WindowClearAnchors` | ui | 1 | OnLButtonUp |
| `WindowSetShowing` | ui | 1 | OnLButtonUp |
| `WindowSetTintColor` | ui | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- WSCT.OnLButtonUpColorPicker


## Binding Resolution

- Total handler declarations: 1
- Resolved to Lua functions: 1 (100%)

## Seen In

- WSCT

## Examples

- WSCT: WSCTOptionsColorPickerWindowColorPicker -> ColorPicker WSCTOptionsColorPickerWindowColorPicker

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [ColorSize](element_ColorSize.md) (MEDIUM 45/100) - XML Element Type
- [ColorSpacing](element_ColorSpacing.md) (MEDIUM 45/100) - XML Element Type
- [ColorTexCoords](element_ColorTexCoords.md) (MEDIUM 45/100) - XML Element Type
- [ColorTexDims](element_ColorTexDims.md) (MEDIUM 45/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
