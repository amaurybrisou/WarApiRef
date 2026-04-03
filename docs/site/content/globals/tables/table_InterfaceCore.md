# InterfaceCore

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, PartyCast |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:14`, `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:236`, `/workspace/data/raw/PartyCast/PartyCast.lua:1015`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:236` |
| Namespaces detected | InterfaceCore |
| Source kinds | globals, lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: InfoScroller.ResetSettings, InfoScroller: LIBGUI_ELEMENT:Scale, PartyCast: LIBGUI_ELEMENT:Scale, PartyCast: PartyCast.Default |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 1 |
| Local definition count | 2 |
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

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- InterfaceCore.GetScale

## Observed Members

- none

## Seen In

- InfoScroller
- PartyCast

## Examples

- InfoScroller: InfoScroller.CreateCard -> InterfaceCore.GetScale()
- InfoScroller: InfoScroller.ResetSettings -> InterfaceCore.GetScale()
- InfoScroller: LIBGUI_ELEMENT:Scale -> InterfaceCore.GetScale()
- PartyCast: LIBGUI_ELEMENT:Scale -> InterfaceCore.GetScale()
- PartyCast: PartyCast.Default -> InterfaceCore.GetScale()

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
