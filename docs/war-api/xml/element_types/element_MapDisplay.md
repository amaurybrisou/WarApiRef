# MapDisplay

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TurretRange |
| Files seen in | `/workspace/data/raw/TurrentRange/Display.xml:0` |
| Namespaces detected | MapDisplay |
| Source kinds | xml_frames, xml_handlers |
| Example locations | TurretRange: TurretMapDisplay |
| XML usage count | 1 |
| XML attribute usage count | 1 |
| Lua usage count | 4 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed XML element type instantiated by 1 addons.

## Common Attributes

- iconScale
- layer
- loadingAnim
- movable
- name
- pinTexture
- popable
- shape
- sticky

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md)
- [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

## Common Handler Functions

- Map.OnClickMap
- Map.OnMButtonUp
- Map.OnMouseOverPoint
- Map.OnRClickMap


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Map.OnClickMap | `function(...)` | LOW |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | Map.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md) | custom | Map.OnMouseOverPoint | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Map.OnRClickMap | `function(...)` | LOW |

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 1× (HIGH)

## Common Structural Child Elements

- [Size](element_Size.md) — 1× (HIGH)

## Typical XML Structure

```xml
<MapDisplay iconScale="0.45" layer="Overlay" loadingAnim="MapLoadingAnim" movable="false" name="..." pinTexture="map_markers01" popable="false" shape="circle" sticky="true">
  <Size/>
</MapDisplay>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `iconScale` | **required** | 100% | 0.45 |
| `layer` | **required** | 100% | Overlay |
| `loadingAnim` | **required** | 100% | MapLoadingAnim |
| `movable` | **required** | 100% | false |
| `pinTexture` | **required** | 100% | map_markers01 |
| `popable` | **required** | 100% | false |
| `shape` | **required** | 100% | circle |
| `sticky` | **required** | 100% | true |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 1 times as an unnamed child.

## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnPointMouseOver

Confidence: LOW

### OnRButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |

## Binding Resolution

- Total handler declarations: 4
- Resolved to Lua functions: 1 (25%)

## Seen In

- TurretRange

## Examples

- TurretRange: TurretMapDisplay -> MapDisplay TurretMapDisplay

## Related APIs

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md) (HIGH 73/100) - XML Event

## Affects

- none
