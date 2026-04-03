# ButtonSetTextureSlice

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

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
| Addons seen in | Minmap, Pure, TexturedButtons |
| Files seen in | Source/PureTarget_UnitFrame.lua, TexturedButtons.lua, core.lua |
| Namespaces detected | ButtonSetTextureSlice |
| Source kinds | lua_calls |
| Example locations | Minmap: ActivateRallyCall, Minmap: RallyCallUpdate, Pure: SetSigil, TexturedButtons: SetButtonTexture |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
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
ButtonSetTextureSlice(arg1, arg2, arg3, arg4)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "MinmapMapScenarioQueue", buttonName |
| arg2 | Observed as a function or method reference. | Observed values: Button.ButtonState.HIGHLIGHTED, Button.ButtonState.NORMAL, Button.ButtonState.PRESSED |
| arg3 | Observed as a text or wstring payload. | Observed values: "EA_HUD_01", "EA_Texture_Menu01", buttonTexture |
| arg4 | Observed as a text or wstring payload. | Observed values: "PaperDoll-Button", "RvR-Flag", normalSlice |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Minmap
- Pure
- TexturedButtons

## Examples

- Minmap: ActivateRallyCall -> ButtonSetTextureSlice("MinmapMapScenarioQueue", Button.ButtonState.NORMAL, "EA_Texture_Menu01", "PaperDoll-Button")
- Minmap: ActivateRallyCall -> ButtonSetTextureSlice("MinmapMapScenarioQueue", Button.ButtonState.HIGHLIGHTED, "EA_Texture_Menu01", "PaperDoll-Button")
- Minmap: ActivateRallyCall -> ButtonSetTextureSlice("MinmapMapScenarioQueue", Button.ButtonState.PRESSED, "EA_Texture_Menu01", "PaperDoll-Button")
- Minmap: ActivateRallyCall -> ButtonSetTextureSlice("MinmapMapScenarioQueue", Button.ButtonState.PRESSED_HIGHLIGHTED, "EA_Texture_Menu01", "PaperDoll-Button")
- Minmap: RallyCallUpdate -> ButtonSetTextureSlice("MinmapMapScenarioQueue", Button.ButtonState.NORMAL, "EA_HUD_01", "RvR-Flag")
- Minmap: RallyCallUpdate -> ButtonSetTextureSlice("MinmapMapScenarioQueue", Button.ButtonState.HIGHLIGHTED, "EA_HUD_01", "RvR-Flag")

## Used With

- [WindowRegisterCoreEventHandler](window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [AlertTextWindow.AddAlert](../../globals/functions/global_AlertTextWindow.AddAlert.md) (HIGH 75/100) - Global Function

## Notes

- none
