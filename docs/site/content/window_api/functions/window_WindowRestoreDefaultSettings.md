# WindowRestoreDefaultSettings

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
| Addons seen in | GuardLine |
| Files seen in | `/workspace/GuardLine/GuardLine.lua:897` |
| Namespaces detected | WindowRestoreDefaultSettings |
| Source kinds | lua_calls |
| Example locations | GuardLine: GuardLine.ResetSettings |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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
WindowRestoreDefaultSettings(arg1)
```

## Description

Observed as a window function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "GuardLineLine", "GuardLineOffGuardLine", "GuardLineOffGuardSelfWindow" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- GuardLine

## Examples

- GuardLine: GuardLine.ResetSettings -> WindowRestoreDefaultSettings("GuardLineLine")
- GuardLine: GuardLine.ResetSettings -> WindowRestoreDefaultSettings("GuardLineSelfWindow")
- GuardLine: GuardLine.ResetSettings -> WindowRestoreDefaultSettings("GuardLineTargetWindow")
- GuardLine: GuardLine.ResetSettings -> WindowRestoreDefaultSettings("GuardLineOffGuardSelfWindow")
- GuardLine: GuardLine.ResetSettings -> WindowRestoreDefaultSettings("GuardLineOffGuardTargetWindow")
- GuardLine: GuardLine.ResetSettings -> WindowRestoreDefaultSettings("GuardLineOffGuardLine")

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
