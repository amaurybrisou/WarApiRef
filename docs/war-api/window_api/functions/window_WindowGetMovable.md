# WindowGetMovable

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | ActionFraction, LibAddonButton, NoOverheal, PotionBar, XpStatus+G |
| Files seen in | Button.lua, NoOverheal.lua, source/ActionFraction.lua, source/Floating.lua, source/XpStatus.lua |
| Namespaces detected | WindowGetMovable |
| Source kinds | lua_calls |
| Example locations | ActionFraction: RightClick, LibAddonButton: OnRButtonUp, LibAddonButton: ToggleLock, NoOverheal: OnLButtonUp, PotionBar: ActivatorRButtonUp, PotionBar: SetMovable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
WindowGetMovable(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "NoOverheal_Window", EA_Window_ContextMenu.activeWindow, self:GetName() |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- ActionFraction
- LibAddonButton
- NoOverheal
- PotionBar
- XpStatus+G

## Examples

- ActionFraction: RightClick -> WindowGetMovable(windowName)
- LibAddonButton: OnRButtonUp -> WindowGetMovable(self:GetName())
- LibAddonButton: ToggleLock -> WindowGetMovable(self:GetName())
- NoOverheal: OnLButtonUp -> WindowGetMovable("NoOverheal_Window")
- PotionBar: ActivatorRButtonUp -> WindowGetMovable(EA_Window_ContextMenu.activeWindow)
- PotionBar: SetMovable -> WindowGetMovable(EA_Window_ContextMenu.activeWindow)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- none
