# WindowGetAnchorCount

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | AnywhereTrainer, CaVES, Countdown, DammazKron, KillTracker, Pocket Palette, Shinies, Vectors |
| Files seen in | Core.lua, Core/DK_Core.lua, Modules/Aggregator/Shinies-Aggregator-Tooltip.lua, PocketPalette.lua, Source/KillTracker.lua, Vis.lua, countdown.lua, emotesOBJECT.lua |
| Namespaces detected | WindowGetAnchorCount |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainer: ReadjustWindowAnchors, CaVES: ReadjustWindowAnchors, Countdown: OnInitialize, DammazKron: AnchorHTS, KillTracker: CreateTrackerWindow, Pocket Palette: ToggleWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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
| windowName | Observed as a target window name. | Observed values: "DammazKronHTS", b, name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AnywhereTrainer
- CaVES
- Countdown
- DammazKron
- KillTracker
- Pocket Palette
- Shinies
- Vectors
- emotes

## Examples

- AnywhereTrainer: ReadjustWindowAnchors -> WindowGetAnchorCount(windowName)
- CaVES: ReadjustWindowAnchors -> WindowGetAnchorCount(wndName)
- Countdown: OnInitialize -> WindowGetAnchorCount(window)
- DammazKron: AnchorHTS -> WindowGetAnchorCount("DammazKronHTS")
- KillTracker: CreateTrackerWindow -> WindowGetAnchorCount(windowName)
- Pocket Palette: ToggleWindow -> WindowGetAnchorCount(b)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed
