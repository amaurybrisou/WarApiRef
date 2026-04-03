# WindowRegisterEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:713` |
| Namespaces detected | WindowRegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
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
WindowRegisterEventHandler(windowName, eventId, handlerName)
```

## Description

Observed binding SystemData events directly to a named window.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "Moth" |
| eventId | Observed as a SystemData event identifier. | Observed values: SystemData.Events.PLAYER_TARGET_UPDATED |
| handlerName | Observed as a Lua handler reference. | Observed values: "Moth.UpdateTarget" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Moth

## Examples

- Moth: Moth.Initialize -> WindowRegisterEventHandler("Moth", SystemData.Events.PLAYER_TARGET_UPDATED, "Moth.UpdateTarget")

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
