# EA_Window_ContextMenu.AddMenuDivider

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, BuffHead, Enemy, RoR_SoR |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.lua:251`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:378`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.lua:195`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogEpsWindow.lua:216`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogTargetDefenseWindow.lua:318`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogTargetDefenseWindow.lua:446`, `/workspace/data/raw/Enemy/Code/Marks/Marks.lua:357`, `/workspace/data/raw/Enemy/Code/Marks/Marks.lua:442` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnTabRBU, AggroMeter: AggroMeter.PickedListMenu, BuffHead: BuffHead.Setup.EffectCache.CreateContextMenu, Enemy: Enemy.CombatLogUI_EpsWindow_OnRButtonUp, Enemy: Enemy.CombatLogUI_TargetDefenseTotalWindow_OnRButtonUp, Enemy: Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
EA_Window_ContextMenu.AddMenuDivider(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: EA_Window_ContextMenu.CONTEXT_MENU_1, EA_Window_ContextMenu.CONTEXT_MENU_2 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AggroMeter
- BuffHead
- Enemy
- RoR_SoR

## Examples

- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuDivider(EA_Window_ContextMenu.CONTEXT_MENU_1)
- AggroMeter: AggroMeter.PickedListMenu -> EA_Window_ContextMenu.AddMenuDivider(EA_Window_ContextMenu.CONTEXT_MENU_1)
- BuffHead: BuffHead.Setup.EffectCache.CreateContextMenu -> EA_Window_ContextMenu.AddMenuDivider()
- Enemy: Enemy.CombatLogUI_EpsWindow_OnRButtonUp -> EA_Window_ContextMenu.AddMenuDivider()
- Enemy: Enemy.CombatLogUI_TargetDefenseTotalWindow_OnRButtonUp -> EA_Window_ContextMenu.AddMenuDivider()
- Enemy: Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp -> EA_Window_ContextMenu.AddMenuDivider()

## Related APIs

- none

## Used With

- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Triggered By

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
