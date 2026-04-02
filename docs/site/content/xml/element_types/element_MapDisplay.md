# MapDisplay

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 186

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CMap, MapMonster, TurretRange |
| Files seen in | `/workspace_addons/MapMonster/Source/MapMonster_Calibrate.xml:83`, `/workspace_addons/TurrentRange/Display.xml:150`, `/workspace_addons/cmap/CMap.xml:66`, `/workspace_addons/cmap/CMap.xml:86` |
| Namespaces detected | MapDisplay |
| Source kinds | xml_frames, xml_handlers |
| Example locations | CMap: CMapWindowMapDisplay, CMap: CMapWindowWMap, MapMonster: MapMonster_CalibrateWindowMapDisplay, TurretRange: TurretMapDisplay |
| XML usage count | 4 |
| XML attribute usage count | 4 |
| Lua usage count | 10 |
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

Observed XML element type instantiated by 3 addons.

## Common Attributes

- name
- pinTexture
- layer
- movable
- shape
- iconScale
- loadingAnim
- sticky
- gutterIcon
- popable

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md)
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md)
- [OnHidden](../handlers/handler_OnHidden.md)
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnShown](../handlers/handler_OnShown.md)
- [OnUpdate](../handlers/handler_OnUpdate.md)

## Common Handler Functions

- CMapWindow.MWheel
- CMapWindow.OnClickMap
- CMapWindow.OnMouseOverPoint
- CMapWindow.UpdateCoordinatesWMap
- CMapWindow.WmapOnClickMap
- CMapWindow.WmapOnMouseOverPoint
- CMapWindow.WmapOver
- CMapWindow.WmapOverEnd
- Map.OnClickMap
- Map.OnMButtonUp
- Map.OnMouseOverPoint
- Map.OnRClickMap


## XML Event Bindings

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnHidden](../handlers/handler_OnHidden.md) | WindowUtils.OnHidden | `function()` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | CMapWindow.OnClickMap, CMapWindow.WmapOnClickMap, Map.OnClickMap | `function(...)` | LOW |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | Map.OnMButtonUp | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | CMapWindow.WmapOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | CMapWindow.WmapOverEnd | `function(...)` | LOW |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | CMapWindow.MWheel | `function(delta)` | MEDIUM |
| [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md) | CMapWindow.OnMouseOverPoint, CMapWindow.WmapOnMouseOverPoint, Map.OnMouseOverPoint | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | Map.OnRClickMap | `function(...)` | LOW |
| [OnShown](../handlers/handler_OnShown.md) | WindowUtils.OnShown | `function()` | MEDIUM |
| [OnUpdate](../handlers/handler_OnUpdate.md) | CMapWindow.UpdateCoordinatesWMap | `function(elapsed)` | MEDIUM |

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `pinTexture` | optional | 57% | map_markers01 |
| `layer` | optional | 42% | Overlay, secondary, default |
| `movable` | optional | 42% | false |
| `shape` | optional | 42% | circle, square |
| `iconScale` | optional | 28% | 0.45, 0.70 |
| `loadingAnim` | optional | 28% | MapLoadingAnim |
| `sticky` | optional | 28% | true, false |
| `gutterIcon` | optional | 14% | 39 |
| `popable` | optional | 14% | false |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Call Count | From Events |
| --- | --- | --- |
| `TextEditBoxSetText` | 15 | OnShown |
| `LabelSetText` | 13 | OnShown |
| `ButtonSetText` | 6 | OnShown |
| `WindowSetDimensions` | 3 | OnShown |
| `DoesWindowExist` | 2 | OnShown |
| `WindowSetShowing` | 1 | OnShown |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHidden

Confidence: HIGH

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
### OnMouseOver

Confidence: LOW

### OnMouseOverEnd

Confidence: LOW

### OnMouseWheel

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `x` | number | mouse_x |
| 1 | `y` | number | mouse_y |
| 2 | `delta` | number | wheel_delta |
### OnPointMouseOver

Confidence: LOW

### OnRButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnShown

Confidence: HIGH

### OnUpdate

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
## Seen In

- CMap
- MapMonster
- TurretRange

## Examples

- CMap: CMapWindowMapDisplay -> MapDisplay CMapWindowMapDisplay
- CMap: CMapWindowWMap -> MapDisplay CMapWindowWMap
- MapMonster: MapMonster_CalibrateWindowMapDisplay -> MapDisplay MapMonster_CalibrateWindowMapDisplay
- TurretRange: TurretMapDisplay -> MapDisplay TurretMapDisplay

## Related APIs

- none

## Used With

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseWheel](../../events/window_events/window_event_OnMouseWheel.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Handler
- [OnPointMouseOver](../../events/window_events/window_event_OnPointMouseOver.md) (HIGH 100/100) - Window Event
- [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
