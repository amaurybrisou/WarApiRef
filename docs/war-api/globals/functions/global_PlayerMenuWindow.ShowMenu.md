# PlayerMenuWindow.ShowMenu

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 88/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because referenced by generated docs or reference files, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, Enemy |
| Files seen in | `/workspace/Effigy/Effigy.lua:474`, `/workspace/Enemy/Code/UnitFrames/UnitFrame.lua:413` |
| Namespaces detected | PlayerMenuWindow |
| Source kinds | globals, lua_calls |
| Example locations | Effigy: Effigy.local.FriendlyRightClickMenu, Effigy: FriendlyRightClickMenu, Enemy: Enemy.UnitFramesUI_UnitFrame_OnRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
PlayerMenuWindow.ShowMenu(arg1, arg2, arg3)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: TargetInfo:UnitName("selffriendlytarget"), player.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, TargetInfo:UnitEntityId("selffriendlytarget") |
| arg3 | Observed as a runtime window or control identifier. | Observed values: custom_menu_items, nil |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Effigy
- Enemy

## Examples

- Effigy: Effigy.local.FriendlyRightClickMenu -> PlayerMenuWindow.ShowMenu(TargetInfo:UnitName("selffriendlytarget"), TargetInfo:UnitEntityId("selffriendlytarget"), nil)
- Effigy: FriendlyRightClickMenu -> PlayerMenuWindow.ShowMenu(TargetInfo:UnitName("selffriendlytarget"), TargetInfo:UnitEntityId("selffriendlytarget"), nil)
- Enemy: Enemy.UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.ShowMenu(player.name, 0, custom_menu_items)

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
