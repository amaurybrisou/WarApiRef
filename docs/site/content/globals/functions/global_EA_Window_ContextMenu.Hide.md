# EA_Window_ContextMenu.Hide

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

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
| Addons seen in | ActionFraction, AdjustTheTip, BuffHead, EZCraftX, MapMonster, Miracle Grow Remix, Obsidian, ShowMeTheBubbles |
| Files seen in | AdjustTheTip.lua, Setup/SelectFont.lua, Setup/SetupDisplay.lua, Source/EZCraftX.lua, Source/MapMonster_Pins.lua, context.lua, showmethebubbles.lua, source/ActionFraction.lua |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | ActionFraction: AutoHideOnMouseOver, ActionFraction: SetFontSelectionMenu, ActionFraction: SetPresetLocation, AdjustTheTip: OnAnchor, AdjustTheTip: OnDelayedFadingEnabled, AdjustTheTip: SpawnAnchorContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
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
EA_Window_ContextMenu.Hide(arg1)
```

## Description

Observed as a global function across 11 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: 1, 2, EA_Window_ContextMenu.CONTEXT_MENU_2 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- ActionFraction
- AdjustTheTip
- BuffHead
- EZCraftX
- MapMonster
- Miracle Grow Remix
- Obsidian
- ShowMeTheBubbles
- TexturedButtons
- TurretRange
- XpStatus+G

## Examples

- ActionFraction: AutoHideOnMouseOver -> EA_Window_ContextMenu.Hide(EA_Window_ContextMenu.CONTEXT_MENU_3)
- ActionFraction: SetFontSelectionMenu -> EA_Window_ContextMenu.Hide(EA_Window_ContextMenu.CONTEXT_MENU_2)
- ActionFraction: SetPresetLocation -> EA_Window_ContextMenu.Hide(EA_Window_ContextMenu.CONTEXT_MENU_3)
- AdjustTheTip: OnAnchor -> EA_Window_ContextMenu.Hide(EA_Window_ContextMenu.CONTEXT_MENU_2)
- AdjustTheTip: OnDelayedFadingEnabled -> EA_Window_ContextMenu.Hide(EA_Window_ContextMenu.CONTEXT_MENU_2)
- AdjustTheTip: SpawnAnchorContextMenu -> EA_Window_ContextMenu.Hide(EA_Window_ContextMenu.CONTEXT_MENU_3)

## Used With

- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
