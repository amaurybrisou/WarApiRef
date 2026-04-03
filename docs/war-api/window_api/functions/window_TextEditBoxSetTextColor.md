# TextEditBoxSetTextColor

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
| Addons seen in | EA_UiDebugTools, Miracle Grow Remix, zMailMod |
| Files seen in | Source/objectinsp/ObjectInspector.lua, config.lua, layout.lua, zMailModSend.lua |
| Namespaces detected | TextEditBoxSetTextColor |
| Source kinds | lua_calls |
| Example locations | EA_UiDebugTools: WindowInit, Miracle Grow Remix: ConfigCallback, Miracle Grow Remix: ConfigSoundChanged, Miracle Grow Remix: InitConfig, Miracle Grow Remix: InitLayout, Miracle Grow Remix: LayoutPlotArrChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 25 |
| Global usage count | 25 |
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
TextEditBoxSetTextColor(arg1, arg2, arg3, arg4)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: ObjectInspector.Window.."ObjectEditBox", boxAC, boxMsg |
| arg2 | Observed as a numeric value. | Observed values: 0, 128, 192 |
| arg3 | Observed as a numeric value. | Observed values: 0, 128, 16 |
| arg4 | Observed as a numeric value. | Observed values: 0, 128, 192 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- EA_UiDebugTools
- Miracle Grow Remix
- zMailMod

## Examples

- EA_UiDebugTools: WindowInit -> TextEditBoxSetTextColor(ObjectInspector.Window.."ObjectEditBox", 225, 225, 225)
- Miracle Grow Remix: ConfigCallback -> TextEditBoxSetTextColor(sWin.."ReserveCount", 255, 255, 255)
- Miracle Grow Remix: ConfigCallback -> TextEditBoxSetTextColor(sWin.."ReserveCount", 128, 128, 128)
- Miracle Grow Remix: ConfigSoundChanged -> TextEditBoxSetTextColor(sName, 0, 255, 0)
- Miracle Grow Remix: ConfigSoundChanged -> TextEditBoxSetTextColor(sName, 255, 255, 0)
- Miracle Grow Remix: ConfigSoundChanged -> TextEditBoxSetTextColor(sName, 255, 0, 0)

## Notes

- none
