# SystemData.MousePosition.y

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | MapMonster, MapPin, Miracle Grow Remix, RoR_SoR |
| Files seen in | `/workspace_addons/MGRemix/layout.lua:558`, `/workspace_addons/MGRemix/layout.lua:581`, `/workspace_addons/MGRemix/layout.lua:617`, `/workspace_addons/MapMonster/Source/MapMonster_Calibrate.lua:162`, `/workspace_addons/MapMonster/Source/MapMonster_Calibrate.lua:255`, `/workspace_addons/MapPin/source/MapPin.lua:3110`, `/workspace_addons/MapPin/source/MapPin.lua:3136`, `/workspace_addons/RoR_SoR/RoR_SoR.lua:2820` |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | MapMonster_Calibrate.OnLMouseButton, MapMonster_Calibrate.OnMouseOverEnd, MapPin.GetMapPos, MapPin.UpdateMapCoordinates, MiracleGrow2.LayoutEndDrag, MiracleGrow2.LayoutStartDrag |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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

Observed SystemData field used by 4 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- MapMonster
- MapPin
- Miracle Grow Remix
- RoR_SoR

## Related APIs

- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function

## Used With

- [FullResizeImage](../../xml/element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [SystemData.MousePosition.x](systemdata_SystemData.MousePosition.x.md) (HIGH 100/100) - SystemData Field
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: MapMonster_Calibrate.OnLMouseButton, MapMonster_Calibrate.OnMouseOverEnd, MapPin.GetMapPos, MapPin.UpdateMapCoordinates, MiracleGrow2.LayoutEndDrag, MiracleGrow2.LayoutStartDrag
