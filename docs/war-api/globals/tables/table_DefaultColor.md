# DefaultColor

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CM_ClosetGoblin, DAoCBuff, MapMonster, Moth, TidyRoll |
| Files seen in | `/workspace/BuffHead/Setup/LayoutControlFrame.lua:55`, `/workspace/BuffHead/Setup/LayoutFrame.lua:47`, `/workspace/TidyRoll/CustomAutoRoll.lua:145` |
| Namespaces detected | DefaultColor |
| Source kinds | globals, lua_calls |
| Example locations | BuffHead: BuffHead.Setup.LayoutControlFrame:Create, BuffHead: BuffHead.Setup.LayoutFrame:UpdateFrameColor, TidyRoll: TidyRoll.CustomAutoRoll.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 2 |
| Local definition count | 5 |
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

Observed shared global table or namespace surfaced in 6 addons.

## Functions

- DefaultColor.SetLabelColor
- DefaultColor.SetWindowTint

## Observed Members

- none

## Seen In

- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- MapMonster
- Moth
- TidyRoll

## Examples

- BuffHead: BuffHead.Setup.LayoutControlFrame:Create -> DefaultColor.SetWindowTint(windowName, DefaultColor.YELLOW)
- BuffHead: BuffHead.Setup.LayoutFrame:UpdateFrameColor -> DefaultColor.SetWindowTint(self:GetName().."Frame", frameColor)
- BuffHead: BuffHead.Setup.LayoutFrame:UpdateFrameColor -> DefaultColor.SetWindowTint(self:GetName().."Background", backgroundColor)
- BuffHead: BuffHead.Setup.LayoutFrame:UpdateFrameColor -> DefaultColor.SetLabelColor(self:GetName().."Name", nameColor)
- TidyRoll: TidyRoll.CustomAutoRoll.Initialize -> DefaultColor.SetWindowTint(rowName.."Background", color)

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
