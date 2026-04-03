# EA_Window_ContextMenu.Finalize

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 58 addons

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
| Addons seen in | ActionFraction, AdjustTheTip, AggroMeter, AutoBand, BuffHead, CM_ClosetGoblin, CastSequence, Crusher |
| Files seen in | AdjustTheTip.lua, AggroMeter.lua, AutoBand.lua, Button.lua, ClosetGoblinCharacterWindow.lua, ClosetGoblinOptionWindow.lua, ClosetGoblinZoneWindow.lua, Code/CombatLog/CombatLogEpsWindow.lua |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | ActionFraction: RightClick, ActionFraction: SetFontSelectionMenu, ActionFraction: SetPresetLocation, AdjustTheTip: ContextMenu_Finalize, AdjustTheTip: CreateContextMenu, AdjustTheTip: SpawnAnchorContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 160 |
| Global usage count | 160 |
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
EA_Window_ContextMenu.Finalize()
```

## Description

Observed as a global function across 58 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- ActionFraction
- AdjustTheTip
- AggroMeter
- AutoBand
- BuffHead
- CM_ClosetGoblin
- CastSequence
- Crusher
- Dascore
- DuffTimer
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraft
- EZCraftX
- Effigy
- Enemy
- FozAuction
- HealGrid
- Killer
- LibAddonButton
- MapMonster
- MapPin
- MarkBuff
- Miracle Grow Remix
- MiracleGrow
- MoraleCircle
- Motion
- NerfedButtons
- Obsidian
- PeaceOut
- PlayEffectsOn
- PotionBar
- Pure
- QuickNameActions+
- RandomMount
- Refer
- RoR_SoR
- SOR
- Shinies
- ShowMeTheBubbles
- SocialWindow 2.0
- TacticSetNames
- TastyButtons
- TaxPayer
- TexturedButtons
- TidyChat
- TidyQueue
- Tortall_DPS
- TurretRange
- Vectors
- WARCommander
- WarBoard
- WarBoard_Loc
- WarBoard_Menu
- WarBoard_WarWhisperer
- XpStatus+G
- scenarioInfo

## Examples

- ActionFraction: RightClick -> EA_Window_ContextMenu.Finalize(EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: SetFontSelectionMenu -> EA_Window_ContextMenu.Finalize(EA_Window_ContextMenu.CONTEXT_MENU_3)
- ActionFraction: SetPresetLocation -> EA_Window_ContextMenu.Finalize(EA_Window_ContextMenu.CONTEXT_MENU_2)
- AdjustTheTip: ContextMenu_Finalize -> EA_Window_ContextMenu.Finalize(contextMenuNumber, anchor)
- AdjustTheTip: CreateContextMenu -> EA_Window_ContextMenu.Finalize()
- AdjustTheTip: SpawnAnchorContextMenu -> EA_Window_ContextMenu.Finalize(EA_Window_ContextMenu.CONTEXT_MENU_2)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event

## Used With

- [EA_Window_ContextMenu.AddCascadingMenuItem](global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
