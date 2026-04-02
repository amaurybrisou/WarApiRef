# WindowGetMovable

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
| Addons seen in | PotionBar |
| Files seen in | `/workspace_addons/PotionBar/source/Floating.lua:100`, `/workspace_addons/PotionBar/source/Floating.lua:90` |
| Namespaces detected | WindowGetMovable |
| Source kinds | lua_calls |
| Example locations | PotionBar: PotionBarFloating.ActivatorRButtonUp, PotionBar: PotionBarFloating.SetMovable |
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
WindowGetMovable(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: EA_Window_ContextMenu.activeWindow |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- PotionBar

## Examples

- PotionBar: PotionBarFloating.ActivatorRButtonUp -> WindowGetMovable(EA_Window_ContextMenu.activeWindow)
- PotionBar: PotionBarFloating.SetMovable -> WindowGetMovable(EA_Window_ContextMenu.activeWindow)

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
