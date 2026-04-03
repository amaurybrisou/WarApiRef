# WindowGetTintColor

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
| Addons seen in | InfoScroller, PartyCast |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibConfig.lua:409`, `/workspace/data/raw/InfoScroller/libs/LibConfig.lua:500`, `/workspace/data/raw/PartyCast/PartyCast.lua:655`, `/workspace/data/raw/PartyCast/libs/LibConfig.lua:409`, `/workspace/data/raw/PartyCast/libs/LibConfig.lua:500` |
| Namespaces detected | WindowGetTintColor |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LibConfigMenu:Add, InfoScroller: element.button.OnLButtonUp, PartyCast: LibConfigMenu:Add, PartyCast: PartyCast.Update, PartyCast: element.button.OnLButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
| windowName | Observed as a target window name. | Observed values: "PartyCastWindow"..i.."TimerBar", LibConfig.colorizer.object.name |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- PartyCast

## Examples

- InfoScroller: LibConfigMenu:Add -> WindowGetTintColor(LibConfig.colorizer.object.name)
- InfoScroller: element.button.OnLButtonUp -> WindowGetTintColor(LibConfig.colorizer.object.name)
- PartyCast: LibConfigMenu:Add -> WindowGetTintColor(LibConfig.colorizer.object.name)
- PartyCast: PartyCast.Update -> WindowGetTintColor("PartyCastWindow"..i.."TimerBar")
- PartyCast: element.button.OnLButtonUp -> WindowGetTintColor(LibConfig.colorizer.object.name)

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
