# ActiveWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 113

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, TexturedButtons, TurretRange |
| Files seen in | `/workspace_addons/BuffHead/Setup/Setup.lua:141`, `/workspace_addons/TexturedButtons/Setup/Setup.lua:123`, `/workspace_addons/TurrentRange/Setup/Setup.lua:31` |
| Namespaces detected | ActiveWindow |
| Source kinds | lua_calls |
| Example locations | BuffHead: BuffHead.Setup.SetActiveWindow, TexturedButtons: TexturedButtons.Setup.SetActiveWindow, TurretRange: TurretRange.Setup.SetActiveWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 2 |
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

Observed shared global table or namespace surfaced in 3 addons.

## Functions

- ActiveWindow.Hide
- ActiveWindow.Show

## Observed Members

- none

## Seen In

- BuffHead
- TexturedButtons
- TurretRange

## Examples

- BuffHead: BuffHead.Setup.SetActiveWindow -> ActiveWindow.Hide()
- BuffHead: BuffHead.Setup.SetActiveWindow -> ActiveWindow.Show()
- TexturedButtons: TexturedButtons.Setup.SetActiveWindow -> ActiveWindow.Hide()
- TexturedButtons: TexturedButtons.Setup.SetActiveWindow -> ActiveWindow.Show()
- TurretRange: TurretRange.Setup.SetActiveWindow -> ActiveWindow.Hide()
- TurretRange: TurretRange.Setup.SetActiveWindow -> ActiveWindow.Show()

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
