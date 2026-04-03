# DialogManager.MakeThreeButtonDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 80/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 80/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TastyButtons |
| Files seen in | TastyButtonsOptions.lua |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | TastyButtons: Button_Set_Texture |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

## Signature (inferred)

```lua
DialogManager.MakeThreeButtonDialog(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: L "Changes will take effext after UI-Reload" |
| arg2 | Observed as a text or wstring payload. | Observed values: L "Reload Now" |
| arg3 | Observed as a function or method reference. | Observed values: function()TastyButtonsOptions.Set_Texture()InterfaceCore.ReloadUI()end |
| arg4 | Observed as a text or wstring payload. | Observed values: L "Reload Later" |
| arg5 | Observed as a function or method reference. | Observed values: function()TastyButtonsOptions.Set_Texture()end |
| arg6 | Observed as a text or wstring payload. | Observed values: L "Cancel" |
| arg7 | Observed as a runtime window or control identifier. | Observed values: nil |
| arg8 | Observed as a runtime window or control identifier. | Observed values: nil |
| arg9 | Observed as a runtime window or control identifier. | Observed values: nil |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- TastyButtons

## Examples

- TastyButtons: Button_Set_Texture -> DialogManager.MakeThreeButtonDialog(L "Changes will take effext after UI-Reload", L "Reload Now", function()TastyButtonsOptions.Set_Texture()InterfaceCore.ReloadUI()end, L "Reload Later", function()TastyButtonsOptions.Set_Texture()end, L "Cancel", nil, nil, nil)

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
