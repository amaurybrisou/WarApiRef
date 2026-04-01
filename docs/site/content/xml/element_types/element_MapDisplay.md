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
| Files seen in | `/workspace/MapMonster/Source/MapMonster_Calibrate.xml:83`, `/workspace/TurrentRange/Display.xml:150`, `/workspace/cmap/CMap.xml:66`, `/workspace/cmap/CMap.xml:86` |
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

- OnLButtonUp
- OnPointMouseOver
- OnMouseWheel
- OnHidden
- OnMButtonUp
- OnMouseOver
- OnMouseOverEnd
- OnRButtonUp
- OnShown
- OnUpdate

## Common Inherits

- none

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

- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 100/100) - XML Handler
- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMButtonUp](../../events/window_events/window_event_OnMButtonUp.md) (HIGH 100/100) - Window Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler
- [OnMouseWheel](../../events/window_events/window_event_OnMouseWheel.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Handler
- [OnPointMouseOver](../../events/window_events/window_event_OnPointMouseOver.md) (HIGH 100/100) - Window Event
- [OnPointMouseOver](../handlers/handler_OnPointMouseOver.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnUpdate](../../events/window_events/window_event_OnUpdate.md) (HIGH 100/100) - Window Event
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
