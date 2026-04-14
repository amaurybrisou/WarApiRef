# WindowSetResizing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead |
| Files seen in | Setup/LayoutControlFrame.lua |
| Namespaces detected | WindowSetResizing |
| Source kinds | lua_calls |
| Example locations | BuffHead: BeginResize, BuffHead: EndResize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
WindowSetResizing(windowName, arg2, arg3, arg4)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: self:GetName() |
| arg2 | Observed as a boolean toggle. | Observed values: false, true |
| arg3 | Observed as a text or wstring payload. | Observed values: "", resizeData.sizePoint |
| arg4 | Observed as a boolean toggle. | Observed values: false, resizeData.lockAspect |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- BuffHead

## Examples

- BuffHead: BeginResize -> WindowSetResizing(self:GetName(), true, resizeData.sizePoint, resizeData.lockAspect)
- BuffHead: EndResize -> WindowSetResizing(self:GetName(), false, "", false)

## Notes

- Only one addon surfaced this symbol in the current corpus.
