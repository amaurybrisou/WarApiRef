# DefaultColor.SetWindowTint

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 113

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, TidyRoll |
| Files seen in | `/workspace/BuffHead/Setup/LayoutControlFrame.lua:55`, `/workspace/BuffHead/Setup/LayoutFrame.lua:47`, `/workspace/TidyRoll/CustomAutoRoll.lua:145` |
| Namespaces detected | DefaultColor |
| Source kinds | globals, lua_calls |
| Example locations | BuffHead: BuffHead.Setup.LayoutControlFrame:Create, BuffHead: BuffHead.Setup.LayoutFrame:UpdateFrameColor, TidyRoll: TidyRoll.CustomAutoRoll.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
DefaultColor.SetWindowTint(arg1, arg2)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: rowName.."Background", self:GetName().."Background", self:GetName().."Frame" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: DefaultColor.YELLOW, backgroundColor, color |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- BuffHead
- TidyRoll

## Examples

- BuffHead: BuffHead.Setup.LayoutControlFrame:Create -> DefaultColor.SetWindowTint(windowName, DefaultColor.YELLOW)
- BuffHead: BuffHead.Setup.LayoutFrame:UpdateFrameColor -> DefaultColor.SetWindowTint(self:GetName().."Frame", frameColor)
- BuffHead: BuffHead.Setup.LayoutFrame:UpdateFrameColor -> DefaultColor.SetWindowTint(self:GetName().."Background", backgroundColor)
- TidyRoll: TidyRoll.CustomAutoRoll.Initialize -> DefaultColor.SetWindowTint(rowName.."Background", color)

## Related APIs

- none

## Used With

- [DefaultColor.SetLabelColor](global_DefaultColor.SetLabelColor.md) (HIGH 70/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
