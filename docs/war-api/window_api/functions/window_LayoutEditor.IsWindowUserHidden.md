# LayoutEditor.IsWindowUserHidden

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | FixGit |
| Files seen in | Enhancements.lua |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | FixGit: ApplyUnitFrameSettings, FixGit: EnhancePetUnitFrame, FixGit: LayoutEditorDone |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | yes |
| Consistent returns | yes |
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
LayoutEditor.IsWindowUserHidden(arg1)
```

## Description

Observed as a window function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "PetHealthWindow" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- FixGit

## Examples

- FixGit: ApplyUnitFrameSettings -> LayoutEditor.IsWindowUserHidden("PetHealthWindow")
- FixGit: EnhancePetUnitFrame -> LayoutEditor.IsWindowUserHidden("PetHealthWindow")
- FixGit: LayoutEditorDone -> LayoutEditor.IsWindowUserHidden("PetHealthWindow")

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Advanced return analysis: No strong return evidence observed
