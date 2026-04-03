# WindowGetAnchorCount

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
| Addons seen in | AnywhereTrainer, Pocket Palette, Shinies |
| Files seen in | `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.lua:160`, `/workspace/data/raw/PocketPalette/PocketPalette.lua:338`, `/workspace/data/raw/Shinies/Modules/Aggregator/Shinies-Aggregator-Tooltip.lua:452` |
| Namespaces detected | WindowGetAnchorCount |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainer: AnywhereTrainer.ReadjustWindowAnchors, Pocket Palette: PP.ToggleWindow, Shinies: Tooltips.AddExtraWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
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
WindowGetAnchorCount(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: b, windowName, windowToAnchorTo |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AnywhereTrainer
- Pocket Palette
- Shinies

## Examples

- AnywhereTrainer: AnywhereTrainer.ReadjustWindowAnchors -> WindowGetAnchorCount(windowName)
- Pocket Palette: PP.ToggleWindow -> WindowGetAnchorCount(b)
- Shinies: Tooltips.AddExtraWindow -> WindowGetAnchorCount(windowToAnchorTo)

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
