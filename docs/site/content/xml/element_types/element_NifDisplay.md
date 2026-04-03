# NifDisplay

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
| Addons seen in | RVMOD_3DPortrait |
| Namespaces detected | NifDisplay |
| Source kinds | xml_frames |
| Example locations | RVMOD_3DPortrait: 3DPortraitSceneTemplate, RVMOD_3DPortrait: 3DPortrait_ARCHMAGE_FEMALE, RVMOD_3DPortrait: 3DPortrait_ARCHMAGE_MALE, RVMOD_3DPortrait: 3DPortrait_BLACKGUARD_FEMALE, RVMOD_3DPortrait: 3DPortrait_BLACKGUARD_MALE, RVMOD_3DPortrait: 3DPortrait_BLACK_ORC_FEMALE |
| XML usage count | 50 |
| XML attribute usage count | 50 |
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

NifDisplay is an interactive XML control. It commonly appears under Windows. It is typically used to organize structural children such as Anchors, EventHandlers, Rotation and bind XML events like OnLButtonUp to Lua.

## Common Attributes

- autoprops
- draganddrop
- fov
- handleinput
- inherits
- layer
- name
- nif
- ortho
- rotateaxis
- scale

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)

## Common Handler Functions

- CharacterWindow.PaperDollMouseUp


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | CharacterWindow.PaperDollMouseUp | `flags, x, y` | LOW |

## Common Inherits

- 3DPortraitSceneTemplate

## Common Parent Elements

- [Windows](element_Windows.md) — 50× (HIGH)

## Common Structural Child Elements

- [Translation](element_Translation.md) — 50× (HIGH)
- [Rotation](element_Rotation.md) — 4× (MEDIUM)
- [Anchors](element_Anchors.md) — 1× (LOW)
- [EventHandlers](element_EventHandlers.md) — 1× (LOW)
- [Size](element_Size.md) — 1× (LOW)

## Common Template Bases

- 3DPortraitSceneTemplate


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 96% | 3DPortraitSceneTemplate |
| `autoprops` | optional | 4% | false |
| `draganddrop` | optional | 4% | true, false |
| `fov` | optional | 4% | 0.58 |
| `handleinput` | optional | 4% | true |
| `nif` | optional | 4% | PaperDollNif |
| `ortho` | optional | 4% | false |
| `rotateaxis` | optional | 4% | 3 |
| `scale` | optional | 4% | 1.00 |
| `layer` | optional | 2% | background |
## Structural Sub-Elements

### [Translation](element_Translation.md)

Observed 50 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0.0, -2.0, -5.0, -1.0 |
| `y` | **required** | -170.0, -20.0, -30.0, -23.0 |
| `z` | **required** | 40.0, 65.0, 44.0, 43.0 |
### [Rotation](element_Rotation.md)

Observed 4 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0.0, 20.0 |
| `y` | **required** | 180.0 |
| `z` | **required** | 180.0, 155.0, 170.0 |
### [Anchors](element_Anchors.md)

Observed 1 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 1 times as an unnamed child.

### [Size](element_Size.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [NifDisplay](element_NifDisplay.md)
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

## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- RVMOD_3DPortrait.Update3DPortraitScenes


## Binding Resolution

- Total handler declarations: 1
- Resolved to Lua functions: 0 (0%)

## Seen In

- RVMOD_3DPortrait

## Examples

- RVMOD_3DPortrait: 3DPortraitSceneTemplate -> NifDisplay 3DPortraitSceneTemplate
- RVMOD_3DPortrait: 3DPortrait_ARCHMAGE_FEMALE -> NifDisplay 3DPortrait_ARCHMAGE_FEMALE
- RVMOD_3DPortrait: 3DPortrait_ARCHMAGE_MALE -> NifDisplay 3DPortrait_ARCHMAGE_MALE
- RVMOD_3DPortrait: 3DPortrait_BLACKGUARD_FEMALE -> NifDisplay 3DPortrait_BLACKGUARD_FEMALE
- RVMOD_3DPortrait: 3DPortrait_BLACKGUARD_MALE -> NifDisplay 3DPortrait_BLACKGUARD_MALE
- RVMOD_3DPortrait: 3DPortrait_BLACK_ORC_FEMALE -> NifDisplay 3DPortrait_BLACK_ORC_FEMALE

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Rotation](element_Rotation.md) (MEDIUM 45/100) - XML Element Type
- [Translation](element_Translation.md) (MEDIUM 45/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
