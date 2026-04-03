# DefaultColor.SetWindowTint

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, EA_ScenarioGroupWindow, TidyRoll, alertMod |
| Files seen in | CustomAutoRoll.lua, Setup/LayoutControlFrame.lua, Setup/LayoutFrame.lua, Source/ScenarioGroupWindow.lua, alertMod.lua |
| Namespaces detected | DefaultColor |
| Source kinds | lua_calls |
| Example locations | BuffHead: Create, BuffHead: UpdateFrameColor, EA_ScenarioGroupWindow: SetListRowTints, TidyRoll: Initialize, alertMod: SetLabels |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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
DefaultColor.SetWindowTint(arg1, arg2)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "alertModDisplayTimeRowBackgroundName", "alertModFadeInRowBackgroundName", "alertModFadeOutRowBackgroundName" |
| arg2 | Observed as a function or method reference. | Observed values: DefaultColor.GetRowColor(1), DefaultColor.GetRowColor(2), DefaultColor.GetRowColor(row) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- BuffHead
- EA_ScenarioGroupWindow
- TidyRoll
- alertMod

## Examples

- BuffHead: Create -> DefaultColor.SetWindowTint(windowName, DefaultColor.YELLOW)
- BuffHead: UpdateFrameColor -> DefaultColor.SetWindowTint(self:GetName().."Frame", frameColor)
- BuffHead: UpdateFrameColor -> DefaultColor.SetWindowTint(self:GetName().."Background", backgroundColor)
- EA_ScenarioGroupWindow: SetListRowTints -> DefaultColor.SetWindowTint(targetRowWindow, DefaultColor.GetRowColor(row))
- TidyRoll: Initialize -> DefaultColor.SetWindowTint(rowName.."Background", color)
- alertMod: SetLabels -> DefaultColor.SetWindowTint("alertModHeightShiftRowBackgroundName", DefaultColor.GetRowColor(2))

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [DefaultColor.GetRowColor](global_DefaultColor.GetRowColor.md) (HIGH 96/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
