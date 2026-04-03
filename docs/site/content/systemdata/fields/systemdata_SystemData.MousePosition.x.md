# SystemData.MousePosition.x

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Group Icons, MapMonster, MapPin, Miracle Grow Remix, MouseHint, RVMOD_SquaredDistances, RoR_SoR |
| Files seen in | GroupIcons.lua, RVMOD_SquaredDistances.lua, RoR_SoR.lua, Source/MapMonster_Calibrate.lua, layout.lua, mouse-hint.lua, source/MapPin.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | GetMapPos, LayoutEndDrag, LayoutStartDrag, LayoutUpdate, OnLMouseButton, OnMouseOverEnd |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 1 |
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

SystemData.SystemData.MousePosition.x field accessed by 7 addons; commonly found in GetMapPos and LayoutEndDrag, LayoutStartDrag, LayoutUpdate, OnLMouseButton, OnMouseOverEnd, OnUpdate, OnWindowOptionsSetScale, SetWindowAlphaByMousePosition, UpdateMapCoordinates, lua_call contexts.

## Seen In

- Group Icons
- MapMonster
- MapPin
- Miracle Grow Remix
- MouseHint
- RVMOD_SquaredDistances
- RoR_SoR

## Related APIs

- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [SystemData.MousePosition.y](systemdata_SystemData.MousePosition.y.md) (HIGH 100/100) - SystemData Field
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function

## Notes

- Observed in contexts: GetMapPos, LayoutEndDrag, LayoutStartDrag, LayoutUpdate, OnLMouseButton, OnMouseOverEnd
