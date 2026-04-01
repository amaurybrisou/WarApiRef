# WindowGetTintColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | GCDsaver, LibWBToggler, WarTriage, WoH-Reticle |
| Files seen in | `/workspace/GCDsaver/libs/LibConfig.lua:419`, `/workspace/GCDsaver/libs/LibConfig.lua:507`, `/workspace/LibWarBoardToggler/libs/LibConfig.lua:415`, `/workspace/LibWarBoardToggler/libs/LibConfig.lua:503`, `/workspace/WarTriage/libs/LibConfig.lua:419`, `/workspace/WarTriage/libs/LibConfig.lua:507`, `/workspace/WoH-Reticle/libs/LibConfig.lua:409`, `/workspace/WoH-Reticle/libs/LibConfig.lua:500` |
| Namespaces detected | WindowGetTintColor |
| Source kinds | lua_calls |
| Example locations | GCDsaver: LibConfigMenu:Add, GCDsaver: element.button.OnLButtonUp, LibWBToggler: LibConfigMenu:Add, LibWBToggler: element.button.OnLButtonUp, WarTriage: LibConfigMenu:Add, WarTriage: element.button.OnLButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
WindowGetTintColor(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: LibConfig.colorizer.object.name |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- GCDsaver
- LibWBToggler
- WarTriage
- WoH-Reticle

## Examples

- GCDsaver: LibConfigMenu:Add -> WindowGetTintColor(LibConfig.colorizer.object.name)
- GCDsaver: element.button.OnLButtonUp -> WindowGetTintColor(LibConfig.colorizer.object.name)
- LibWBToggler: LibConfigMenu:Add -> WindowGetTintColor(LibConfig.colorizer.object.name)
- LibWBToggler: element.button.OnLButtonUp -> WindowGetTintColor(LibConfig.colorizer.object.name)
- WarTriage: LibConfigMenu:Add -> WindowGetTintColor(LibConfig.colorizer.object.name)
- WarTriage: element.button.OnLButtonUp -> WindowGetTintColor(LibConfig.colorizer.object.name)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
