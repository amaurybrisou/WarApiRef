# TextEditBoxSetTextColor

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
| Addons seen in | EA_UiDebugTools, Miracle Grow Remix |
| Files seen in | `/workspace_addons/MGRemix/config.lua:17`, `/workspace_addons/MGRemix/config.lua:258`, `/workspace_addons/MGRemix/config.lua:87`, `/workspace_addons/MGRemix/layout.lua:1033`, `/workspace_addons/MGRemix/layout.lua:262`, `/workspace_addons/ea_uidebugtools/Source/objectinsp/ObjectInspector.lua:60` |
| Namespaces detected | TextEditBoxSetTextColor |
| Source kinds | lua_calls |
| Example locations | EA_UiDebugTools: ObjectInspector.WindowInit, Miracle Grow Remix: MiracleGrow2.ConfigCallback, Miracle Grow Remix: MiracleGrow2.ConfigSoundChanged, Miracle Grow Remix: MiracleGrow2.InitConfig, Miracle Grow Remix: MiracleGrow2.InitLayout, Miracle Grow Remix: MiracleGrow2.LayoutPlotArrChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
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
| arg1 | Observed as a function or method reference. | Observed values: ObjectInspector.Window.."ObjectEditBox", mg.sWindowName.."ConfigThrobColorB", mg.sWindowName.."ConfigThrobColorG" |
| arg2 | Observed as a numeric value. | Observed values: 0, 128, 192 |
| arg3 | Observed as a numeric value. | Observed values: 0, 128, 192 |
| arg4 | Observed as a numeric value. | Observed values: 0, 128, 192 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- EA_UiDebugTools
- Miracle Grow Remix

## Examples

- EA_UiDebugTools: ObjectInspector.WindowInit -> TextEditBoxSetTextColor(ObjectInspector.Window.."ObjectEditBox", 225, 225, 225)
- Miracle Grow Remix: MiracleGrow2.ConfigCallback -> TextEditBoxSetTextColor(sWin.."ReserveCount", 255, 255, 255)
- Miracle Grow Remix: MiracleGrow2.ConfigCallback -> TextEditBoxSetTextColor(sWin.."ReserveCount", 128, 128, 128)
- Miracle Grow Remix: MiracleGrow2.ConfigSoundChanged -> TextEditBoxSetTextColor(sName, 0, 255, 0)
- Miracle Grow Remix: MiracleGrow2.ConfigSoundChanged -> TextEditBoxSetTextColor(sName, 255, 255, 0)
- Miracle Grow Remix: MiracleGrow2.ConfigSoundChanged -> TextEditBoxSetTextColor(sName, 255, 0, 0)

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event

## Affects

- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
