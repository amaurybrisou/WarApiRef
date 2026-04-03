# DefaultColor.SetLabelColor

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 88/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, EA_UiModWindow, Tome Titan |
| Files seen in | Setup/LayoutFrame.lua, TTitan_UI.lua, source/uimodwindow.lua |
| Namespaces detected | DefaultColor |
| Source kinds | lua_calls |
| Example locations | BuffHead: UpdateFrameColor, EA_UiModWindow: UpdateModDetailsData, EA_UiModWindow: UpdateModRowByIndex, Tome Titan: MouseOverRow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 0 |
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
DefaultColor.SetLabelColor(arg1, arg2)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: rowName.."Status", self:GetName().."Name", targetRowWindow.."Name" |
| arg2 | Observed as a function or method reference. | Observed values: DefaultColor.MAGENTA, loadStatusData.color, nameColor |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- BuffHead
- EA_UiModWindow
- Tome Titan

## Examples

- BuffHead: UpdateFrameColor -> DefaultColor.SetLabelColor(self:GetName().."Name", nameColor)
- EA_UiModWindow: UpdateModDetailsData -> DefaultColor.SetLabelColor(windowName.."ScrollChildStatusText", loadStatusData.color)
- EA_UiModWindow: UpdateModRowByIndex -> DefaultColor.SetLabelColor(rowName.."Status", loadStatusData.color)
- Tome Titan: MouseOverRow -> DefaultColor.SetLabelColor(targetRowWindow.."Name", DefaultColor.MAGENTA)

## Used With

- [DefaultColor.SetWindowTint](global_DefaultColor.SetWindowTint.md) (HIGH 100/100) - Global Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
