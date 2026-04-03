# ButtonSetTexture

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | NAMBLA, TexturedButtons |
| Files seen in | NAMBLA.lua, TexturedButtons.lua |
| Namespaces detected | ButtonSetTexture |
| Source kinds | lua_calls |
| Example locations | NAMBLA: HideButtonTextures, TexturedButtons: SetButtonTexture |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
ButtonSetTexture(arg1, arg2, arg3, arg4, arg5)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: btn, buttonName |
| arg2 | Observed as a function or method reference. | Observed values: Button.ButtonState.HIGHLIGHTED, Button.ButtonState.NORMAL, Button.ButtonState.PRESSED |
| arg3 | Observed as a text or wstring payload. | Observed values: "", buttonTexture or "" |
| arg4 | Observed as a numeric value. | Observed values: 0 |
| arg5 | Observed as a numeric value. | Observed values: 0 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- NAMBLA
- TexturedButtons

## Examples

- NAMBLA: HideButtonTextures -> ButtonSetTexture(btn, Button.ButtonState.NORMAL, "", 0, 0)
- NAMBLA: HideButtonTextures -> ButtonSetTexture(btn, Button.ButtonState.HIGHLIGHTED, "", 0, 0)
- NAMBLA: HideButtonTextures -> ButtonSetTexture(btn, Button.ButtonState.PRESSED, "", 0, 0)
- NAMBLA: HideButtonTextures -> ButtonSetTexture(btn, Button.ButtonState.PRESSED_HIGHLIGHTED, "", 0, 0)
- TexturedButtons: SetButtonTexture -> ButtonSetTexture(buttonName, buttonState, buttonTexture or "", 0, 0)

## Notes

- none
