# Sounds

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CMap, FastInteract, Miracle Grow Remix, MiracleGrow |
| Files seen in | `/workspace_addons/FastInteract/FastInteract.xml:3`, `/workspace_addons/MGRemix/MGRemix.xml:169`, `/workspace_addons/MGRemix/MGRemix.xml:183`, `/workspace_addons/MiracleGrow/window.xml:63`, `/workspace_addons/cmap/CMap.xml:148`, `/workspace_addons/cmap/CMap.xml:183`, `/workspace_addons/cmap/CMap.xml:271`, `/workspace_addons/cmap/CMap.xml:304` |
| Namespaces detected | Sounds |
| Source kinds | xml_frames |
| Example locations | CMap: CMapWindowMapRallyCall, CMap: CMapWindowMapRankedLeaderboard, CMap: CMapWindowMapScenarioQueue, CMap: CMapWindowMapWorldMapButton, FastInteract: FastInteractWindow, Miracle Grow Remix: MiracleGrow2LineHarvest |
| XML usage count | 8 |
| XML attribute usage count | 8 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

Observed XML element type instantiated by 4 addons.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 7× (HIGH)
- [Window](element_Window.md) — 1× (LOW)

## Seen In

- CMap
- FastInteract
- Miracle Grow Remix
- MiracleGrow

## Examples

- CMap: CMapWindowMapRallyCall -> Sounds in Button CMapWindowMapRallyCall
- CMap: CMapWindowMapRankedLeaderboard -> Sounds in Button CMapWindowMapRankedLeaderboard
- CMap: CMapWindowMapScenarioQueue -> Sounds in Button CMapWindowMapScenarioQueue
- CMap: CMapWindowMapWorldMapButton -> Sounds in Button CMapWindowMapWorldMapButton
- FastInteract: FastInteractWindow -> Sounds in Window FastInteractWindow
- Miracle Grow Remix: MiracleGrow2LineHarvest -> Sounds in Button MiracleGrow2LineHarvest

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
