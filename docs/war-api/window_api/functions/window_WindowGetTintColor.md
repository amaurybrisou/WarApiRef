# WindowGetTintColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | Amethyst, CMap, EZGuard, GCDsaver, InfoScroller, LibWBToggler, PartyCast, TargetRing |
| Files seen in | Amethyst.lua, LibConfig.lua, PartyCast.lua, libs/LibConfig.lua |
| Namespaces detected | WindowGetTintColor |
| Source kinds | lua_calls |
| Example locations | Amethyst: ApplySettings, Amethyst: CreateGUI, Amethyst: OnLButtonUp, CMap: Add, CMap: OnLButtonUp, EZGuard: Add |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 39 |
| Global usage count | 39 |
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
| windowName | Observed as a target window name. | Observed values: "PartyCastWindow"..i.."TimerBar", GUI.Colorizer.Object.name, I.BarColor.Overlay.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Amethyst
- CMap
- EZGuard
- GCDsaver
- InfoScroller
- LibWBToggler
- PartyCast
- TargetRing
- WarTriage
- WoH-Reticle

## Examples

- Amethyst: ApplySettings -> WindowGetTintColor(I.BgColor.Overlay.name)
- Amethyst: ApplySettings -> WindowGetTintColor(I.BarColor.Overlay.name)
- Amethyst: ApplySettings -> WindowGetTintColor(I.SuccessColor.Overlay.name)
- Amethyst: ApplySettings -> WindowGetTintColor(I.CancelColor.Overlay.name)
- Amethyst: CreateGUI -> WindowGetTintColor(GUI.Colorizer.Object.name)
- Amethyst: OnLButtonUp -> WindowGetTintColor(GUI.Colorizer.Object.name)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Used With

- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Notes

- none
