# OnMButtonUp

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, GroupSpotter, HealGrid, Map, Minmap, NoOverheal, Statdoll Light, Statdoll Remix |
| Namespaces detected | OnMButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | Enemy: .OnMButtonUp, GroupSpotter: .OnMButtonUp, HealGrid: .OnMButtonUp, Map: .OnMButtonUp, Minmap: .OnMButtonUp, NoOverheal: .OnMButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | yes |
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

## Description

XML handler event observed across 10 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- Enemy
- GroupSpotter
- HealGrid
- Map
- Minmap
- NoOverheal
- Statdoll Light
- Statdoll Remix
- TidyRoll
- followTheLeader

## Examples

- Enemy: .OnMButtonUp -> Enemy.UnitFramesUI_UnitFrame_OnMButtonUp
- GroupSpotter: .OnMButtonUp -> GroupSpotter.Settings.ToggleWindow
- HealGrid: .OnMButtonUp -> HealGridMouseClick.UnitMButtonUp
- Map: .OnMButtonUp -> Map.OnMButtonUp
- Minmap: .OnMButtonUp -> Minmap.AnswerCall
- NoOverheal: .OnMButtonUp -> NoOverheal.OnMButtonUp

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Window_ContextMenu.CreateDefaultContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateDefaultContextMenu.md) (HIGH 100/100) - Global Function
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Affects

- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.BUTTON_CLICK](../../gamedata/fields/gamedata_GameData.Sound.BUTTON_CLICK.md) (HIGH 100/100) - GameData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
