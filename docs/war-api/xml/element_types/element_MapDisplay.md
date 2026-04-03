# MapDisplay

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Atlas, CMap, EA_LoadingScreen, Map, MapMonster, Minmap, TurretRange |
| Files seen in | Map/Main.xml, Source/GeneralLoadingScreenTemplates.xml, Source/MapMonster_Calibrate.xml |
| Namespaces detected | MapDisplay |
| Source kinds | xml_frames |
| Example locations | Atlas: AtlasFrameMapContainerMapDisplay, CMap: CMapWindowMapDisplay, CMap: CMapWindowWMap, EA_LoadingScreen: LoadingScreenMapDefDisplay, Map: MapDisplay, MapMonster: MapMonster_CalibrateWindowMapDisplay |
| XML usage count | 8 |
| XML attribute usage count | 8 |
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

MapDisplay is an interactive XML control. It commonly appears under Window. It is typically used to organize structural children such as Anchors, EventHandlers, Size and bind XML events like OnHidden, OnLButtonUp, OnMButtonUp to Lua.

## Common Attributes

- gutterIcon
- handleinput
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

- [OnHidden](../handlers/handler_OnHidden.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md)
- [OnMouseLeave](../handlers/handler_OnMouseLeave.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md)
- [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnShown](../handlers/handler_OnShown.md)
- [OnShutdown](../handlers/handler_OnShutdown.md)
- [OnUpdate](../handlers/handler_OnUpdate.md)

## Common Handler Functions

- Map.OnMouseOverPoint
- AtlasMap.OnMouseOverPoint
- CMapWindow.OnMouseOverPoint
- CMapWindow.WmapOnMouseOverPoint
- Minmap.OnMouseOverPoint
- Tooltips.OnMouseOverMapPoint
- Map.OnClickMap
- CMapWindow.OnClickMap
- CMapWindow.WmapOnClickMap
- Minmap.OnClickMap
- CMapWindow.MWheel
- Minmap.HandleMouseWheel
- Map.OnRClickMap
- AtlasMap.OnMapRBU
- Map.OnMButtonUp
- AtlasMap.OnMouseOver
- CMapWindow.WmapOver
- AtlasMap.OnMouseOverEnd
- CMapWindow.WmapOverEnd
- WindowUtils.OnHidden
- AtlasMap.OnMouseLeaveMapPin
- WindowUtils.OnShown
- Map.Shutdown
- CMapWindow.UpdateCoordinatesWMap


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md) | custom | Map.OnMouseOverPoint, AtlasMap.OnMouseOverPoint, CMapWindow.OnMouseOverPoint, CMapWindow.WmapOnMouseOverPoint, Minmap.OnMouseOverPoint, Tooltips.OnMouseOverMapPoint | `` |  |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Map.OnClickMap, CMapWindow.OnClickMap, CMapWindow.WmapOnClickMap, Minmap.OnClickMap | `flags, x, y` | MEDIUM |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | input | CMapWindow.MWheel, Minmap.HandleMouseWheel | `x, y, delta` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Map.OnRClickMap, AtlasMap.OnMapRBU | `flags, x, y` | MEDIUM |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | Map.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | AtlasMap.OnMouseOver, CMapWindow.WmapOver | `` |  |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | AtlasMap.OnMouseOverEnd, CMapWindow.WmapOverEnd | `` |  |
| [OnHidden](../handlers/handler_OnHidden.md) | lifecycle | WindowUtils.OnHidden | `` |  |
| [OnMouseLeave](../handlers/handler_OnMouseLeave.md) | custom | AtlasMap.OnMouseLeaveMapPin | `` |  |
| [OnShown](../handlers/handler_OnShown.md) | lifecycle | WindowUtils.OnShown | `` |  |
| [OnShutdown](../handlers/handler_OnShutdown.md) | lifecycle | Map.Shutdown | `` |  |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | CMapWindow.UpdateCoordinatesWMap | `elapsed` | MEDIUM |

### Per-Event Lua API Calls

**OnRButtonUp** handlers call: `WindowGetShowing`

**OnMouseOverEnd** handlers call: `WindowSetShowing`

**OnShown** handlers call: `ButtonSetText`, `DoesWindowExist`, `LabelSetText`, `TextEditBoxSetText`, `WindowSetDimensions`, `WindowSetShowing`

**OnShutdown** handlers call: `WindowUnregisterEventHandler`

**OnUpdate** handlers call: `LabelSetText`, `WindowGetScreenPosition`

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 8× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 8× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 8× (HIGH)
- [Size](element_Size.md) — 3× (MEDIUM)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `pinTexture` | **required** | 100% | map_markers01 |
| `layer` | optional | 75% | default, secondary, Overlay |
| `movable` | optional | 62% | false |
| `shape` | optional | 62% | square, circle |
| `iconScale` | optional | 50% | 0.70, 0.7, 0.45 |
| `loadingAnim` | optional | 50% | MapLoadingAnim |
| `sticky` | optional | 50% | false, true |
| `gutterIcon` | optional | 37% | 39 |
| `popable` | optional | 37% | false |
| `handleinput` | optional | 12% | true |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 8 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 8 times as an unnamed child.

### [Size](element_Size.md)

Observed 3 times as an unnamed child.

## Recursive Hierarchy

- Root: [MapDisplay](element_MapDisplay.md)
- [Anchors](element_Anchors.md) (structural, 8×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
      - (cycle)
- [EventHandlers](element_EventHandlers.md) (structural, 8×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
- [Size](element_Size.md) (structural, 3×, MEDIUM)
  - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `TextEditBoxSetText` | ui | 15 | OnShown |
| `LabelSetText` | ui | 14 | OnShown, OnUpdate |
| `ButtonSetText` | ui | 6 | OnShown |
| `WindowSetDimensions` | ui | 3 | OnShown |
| `DoesWindowExist` | ui | 2 | OnShown |
| `WindowSetShowing` | ui | 2 | OnMouseOverEnd, OnShown |
| `WindowUnregisterEventHandler` | ui | 2 | OnShutdown |
| `WindowGetScreenPosition` | ui | 1 | OnUpdate |
| `WindowGetShowing` | ui | 1 | OnRButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHidden

Confidence: HIGH

### OnLButtonUp

Confidence: MEDIUM

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
### OnMouseLeave

Confidence: LOW

### OnMouseOver

Confidence: HIGH

### OnMouseOverEnd

Confidence: HIGH

### OnMouseWheel

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `x` | number | mouse_x |
| 1 | `y` | number | mouse_y |
| 2 | `delta` | number | wheel_delta |
### OnPointMouseOver

Confidence: LOW

### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnShown

Confidence: HIGH

### OnShutdown

Confidence: HIGH

### OnUpdate

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
## Lua Functions Manipulating This Type

- MapPin.OnUpdate
- MapMonster_Calibrate.GetDimensions
- MapMonster_Calibrate.OnMouseOverEnd
- MapPin.UpdateGuide
- Atlas.ShowCoordinatesOnMouseOver
- MapMonster.CreateMarker
- MapMonster.Calibrate_ResizeMap
- MapMonster_Calibrate.OnLMouseButton
- Atlas.local.SetMapTransparency
- Map.Initialize
- MapPin.GetMapPos
- MapMonster.Calibrate_MoveAnchor
- MapMonster_Calibrate.WindowReport
- MapMonster.local.CreateMarker
- MapMonster.local.MoveMapPin
- MapMonster.MoveMapPin
- Atlas.SetMapTransparency
- MapPin.UpdateMapCoordinates
- MapMonster.InitializeMapPins
- SimpleXY.UpdateCoordinates
- MapMonster.local.Calibrate_MoveAnchor
- Map.SetBorderForegroundFRI
- Atlas.local.ShowCoordinatesOnMouseOver


## Binding Resolution

- Total handler declarations: 29
- Resolved to Lua functions: 29 (100%)

## Seen In

- Atlas
- CMap
- EA_LoadingScreen
- Map
- MapMonster
- Minmap
- TurretRange

## Examples

- Atlas: AtlasFrameMapContainerMapDisplay -> MapDisplay AtlasFrameMapContainerMapDisplay
- CMap: CMapWindowMapDisplay -> MapDisplay CMapWindowMapDisplay
- CMap: CMapWindowWMap -> MapDisplay CMapWindowWMap
- EA_LoadingScreen: LoadingScreenMapDefDisplay -> MapDisplay LoadingScreenMapDefDisplay
- Map: MapDisplay -> MapDisplay MapDisplay
- MapMonster: MapMonster_CalibrateWindowMapDisplay -> MapDisplay MapMonster_CalibrateWindowMapDisplay

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 88/100) - XML Event
- [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnShutdown](../handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [OnMouseLeave](../handlers/handler_OnMouseLeave.md) (HIGH 76/100) - XML Event
