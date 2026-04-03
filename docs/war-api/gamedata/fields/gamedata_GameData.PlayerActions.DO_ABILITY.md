# GameData.PlayerActions.DO_ABILITY

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, HealGrid, HotbarMorale, MarkBuff, NerfedButtons, Pure, ResHelp, Twister |
| Files seen in | Code/UnitFrames/ClickCasting.lua, Core.lua, HotbarMorale.lua, Modules/HealGridMouseClick.lua, NerfedEngine.lua, ResHelp.lua, Source/PurePlayerPet.lua, Twister.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | CursorSwap, GetAbilityCooldown, Initialize, LoadUnitFrame, OnLoad, OnPlayerSingleAbilityUpdated |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
| Local definition count | 0 |
| Documentation references | 1 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

GameData.GameData.PlayerActions.DO_ABILITY field accessed by 8 addons; commonly found in CursorSwap and GetAbilityCooldown, Initialize, LoadUnitFrame, OnLoad, OnPlayerSingleAbilityUpdated, OnUpdate, Proceed, ProcessAbility, SetActiveAura, changeHotbar, lua_call contexts.

## Seen In

- Enemy
- HealGrid
- HotbarMorale
- MarkBuff
- NerfedButtons
- Pure
- ResHelp
- Twister

## Related APIs

- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [WindowGetGameActionData](../../window_api/functions/window_WindowGetGameActionData.md) (HIGH 100/100) - Window Function
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: CursorSwap, GetAbilityCooldown, Initialize, LoadUnitFrame, OnLoad, OnPlayerSingleAbilityUpdated
