# ButtonSetTextureSlice

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
| Addons seen in | TexturedButtons |
| Files seen in | `/workspace/data/raw/TexturedButtons/TexturedButtons.lua:214` |
| Namespaces detected | ButtonSetTextureSlice |
| Source kinds | lua_calls |
| Example locations | TexturedButtons: SetButtonTexture, TexturedButtons: TexturedButtons.local.SetButtonTexture |
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
ButtonSetTextureSlice(arg1, arg2, arg3, arg4)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: buttonName |
| arg2 | Observed as a runtime window or control identifier. | Observed values: buttonState |
| arg3 | Observed as a runtime window or control identifier. | Observed values: buttonTexture |
| arg4 | Observed as a runtime window or control identifier. | Observed values: sliceTexture |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TexturedButtons

## Examples

- TexturedButtons: SetButtonTexture -> ButtonSetTextureSlice(buttonName, buttonState, buttonTexture, sliceTexture)
- TexturedButtons: TexturedButtons.local.SetButtonTexture -> ButtonSetTextureSlice(buttonName, buttonState, buttonTexture, sliceTexture)

## Related APIs

- none

## Used With

- [ButtonSetTexture](window_ButtonSetTexture.md) (HIGH 90/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
