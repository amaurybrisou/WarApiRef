# ColorPicker

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 108

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | WSCT |
| Files seen in | `/workspace_addons/wsct/wsct_options/wsct_options.xml:1115` |
| Namespaces detected | ColorPicker |
| Source kinds | xml_frames, xml_handlers |
| Example locations | WSCT: WSCTOptionsColorPickerWindowColorPicker |
| XML usage count | 1 |
| XML attribute usage count | 1 |
| Lua usage count | 1 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed XML element type instantiated by 1 addons.

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

- [Window](element_Window.md) — 1× (HIGH)

## Common Structural Child Elements

- [ColorSize](element_ColorSize.md) — 1× (HIGH)
- [ColorSpacing](element_ColorSpacing.md) — 1× (HIGH)
- [ColorTexCoords](element_ColorTexCoords.md) — 1× (HIGH)
- [ColorTexDims](element_ColorTexDims.md) — 1× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `columnsPerRow` | optional | 50% | 4 |
| `texture` | optional | 50% | EA_HUD_01 |
## Structural Sub-Elements

### [ColorSize](element_ColorSize.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [ColorSpacing](element_ColorSpacing.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [ColorTexCoords](element_ColorTexCoords.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [ColorTexDims](element_ColorTexDims.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
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

- WSCT.WSCT.OnLButtonUpColorPicker


## Binding Resolution

- Total handler declarations: 1
- Resolved to Lua functions: 1 (100%)

## Seen In

- WSCT

## Examples

- WSCT: WSCTOptionsColorPickerWindowColorPicker -> ColorPicker WSCTOptionsColorPickerWindowColorPicker

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
