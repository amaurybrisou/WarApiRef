# EA_Window_ContextMenu.HideAll

- Category: Global Function
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
| Addons seen in | FozAuction, LibAddonButton, MapMonster, Motion, Vectors |
| Files seen in | Menu.lua, Settings.lua, Source/MapMonster_Pins.lua, Source/Motion.lua, Source/auctionwindowsearchcontrols.lua |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | FozAuction: OnLButtonDownContextMenuItem, LibAddonButton: OnLButtonUp, MapMonster: HideFilterMenus, Motion: OnLButtonUp_ContextMenu_Item, Motion: OnLButtonUp_ContextMenu_Recipe, Vectors: OnUpdateInteractiveSelection |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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
EA_Window_ContextMenu.HideAll()
```

## Description

Observed as a global function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- FozAuction
- LibAddonButton
- MapMonster
- Motion
- Vectors

## Examples

- FozAuction: OnLButtonDownContextMenuItem -> EA_Window_ContextMenu.HideAll()
- LibAddonButton: OnLButtonUp -> EA_Window_ContextMenu.HideAll()
- MapMonster: HideFilterMenus -> EA_Window_ContextMenu.HideAll()
- Motion: OnLButtonUp_ContextMenu_Item -> EA_Window_ContextMenu.HideAll()
- Motion: OnLButtonUp_ContextMenu_Recipe -> EA_Window_ContextMenu.HideAll()
- Vectors: OnUpdateInteractiveSelection -> EA_Window_ContextMenu.HideAll()

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Used With

- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
